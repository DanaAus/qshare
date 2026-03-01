package handlers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"qshare/internal/network"
	"qshare/internal/server"

	"github.com/mdp/qrterminal/v3"
)

// StartSendServer initializes the ephemeral server and handles file sending.
func StartSendServer(targetPath string, secure bool) error {
	// 1. Check if path exists and is a directory
	info, err := os.Stat(targetPath)
	if err != nil {
		return fmt.Errorf("target path not found: %w", err)
	}

	// 2. Discover Network Interface and Port
	ip, err := network.GetActiveIPv4Interface()
	if err != nil {
		fmt.Printf("[Warning] Could not auto-detect primary IP. Using 127.0.0.1. Error: %v\n", err)
		ip = "127.0.0.1"
	}
	port, err := network.GetAvailablePort()
	if err != nil {
		return fmt.Errorf("could not find open port: %w", err)
	}

	// 3. Generate secure download URL
	hash, err := network.GenerateRandomString(6) // 12-char hex
	if err != nil {
		return err
	}

	downloadPath := fmt.Sprintf("/d/%s", hash)
	downloadURL := fmt.Sprintf("http://%s:%d%s", ip, port, downloadPath)

	// Output Info
	fmt.Printf("[Network] Using active interface: %s\n", ip)
	fmt.Printf("[Server]  Started on port %d\n", port)
	fmt.Printf("[URL]     %s\n", downloadURL)

	// Print QR
	fmt.Println("[QR]")
	qrterminal.GenerateHalfBlock(downloadURL, qrterminal.L, os.Stdout)

	// 4. Setup Server
	srv := server.NewEphemeralServer(port)

	srv.Handle(downloadPath, func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Security: Prevent path traversal (handled implicitly out of box but good to log)

		fmt.Printf("\n[Server] Connection established from %s\n", r.RemoteAddr)

		if info.IsDir() {
			// Stream folder as ZIP
			w.Header().Set("Content-Type", "application/zip")
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, info.Name()))

			zipWriter := zip.NewWriter(w)
			defer zipWriter.Close()

			err := filepath.Walk(targetPath, func(path string, fileInfo os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if fileInfo.IsDir() {
					return nil
				}

				relPath, err := filepath.Rel(targetPath, path)
				if err != nil {
					return err
				}

				zipFile, err := zipWriter.Create(filepath.ToSlash(relPath))
				if err != nil {
					return err
				}

				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(zipFile, file)
				return err
			})

			if err != nil {
				fmt.Printf("[Error] Failed to stream ZIP: %v\n", err)
			}
		} else {
			// Serve a single file
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, info.Name()))
			http.ServeFile(w, r, targetPath)
		}

		// Shutdown after transfer
		go func() {
			time.Sleep(1 * time.Second) // allow buffer to flush
			srv.TriggerShutdown()
		}()
	})

	fmt.Println("\nStatus: Waiting for connection... (Server will close after 1 download, timeout 5m)")

	// 5. Start Server with 5-minute timeout
	return srv.Start(5 * time.Minute)
}

// ServeFileWithProgress serves a single file with a progress bar.
func ServeFileWithProgress(w http.ResponseWriter, r *http.Request, filePath string) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
