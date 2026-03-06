package handlers

import (
	"archive/zip"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"magshare/internal/network"
	"magshare/internal/server"
	"magshare/ui"

	"github.com/mdp/qrterminal/v3"
	"github.com/schollz/progressbar/v3"
)

// SendOptions defines configuration for the send server.
type SendOptions struct {
	Port   int
	Secure bool
	PIN    string
	Demo   bool
}

// StartSendServer initializes the ephemeral server and handles file sending.
func StartSendServer(targetPath string, opts SendOptions) error {
	// 1. Check if path exists
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

	port := opts.Port
	if port <= 0 {
		port, err = network.GetAvailablePort()
		if err != nil {
			return fmt.Errorf("could not find open port: %w", err)
		}
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
	if opts.Secure {
		if opts.PIN == "" {
			opts.PIN, _ = server.GeneratePIN()
		}
		fmt.Printf("[Auth]    PIN REQUIRED: %s\n", opts.PIN)
	}
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

		// Check PIN if secure mode is enabled
		if opts.Secure {
			clientPin := r.URL.Query().Get("pin")
			if clientPin != opts.PIN {
				// Serve PIN entry page instead of raw error
				tmpl, err := template.ParseFS(ui.Files, "pin.html")
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				data := struct {
					Error string
				}{}
				if clientPin != "" {
					data.Error = "Invalid PIN. Please try again."
				}

				w.Header().Set("Content-Type", "text/html")
				tmpl.Execute(w, data)
				return
			}
		}

		fmt.Printf("\n[Server] Connection established from %s\n", r.RemoteAddr)

		var serveErr error
		if info.IsDir() {
			serveErr = ServeDirWithProgress(w, r, targetPath)
		} else {
			serveErr = ServeFileWithProgress(w, r, targetPath)
		}

		if serveErr != nil {
			fmt.Printf("[Error] Failed to serve content: %v\n", serveErr)
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

// ServeFileWithProgress serves a single file to the client with a real-time
// progress bar (bytes sent, speed, ETA) shown on the host terminal.
func ServeFileWithProgress(w http.ResponseWriter, r *http.Request, filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return fmt.Errorf("stat %q: %w", filePath, err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return fmt.Errorf("open %q: %w", filePath, err)
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(info.Size(), "Sending file")

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, info.Name()))
	pr := NewProgressReader(r.Context(), file, bar)
	if _, err := io.Copy(w, pr); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	fmt.Println() // newline after bar
	return nil
}

// ServeDirWithProgress streams a directory as a ZIP archive to the client, showing
// an indeterminate spinner progress bar on the host terminal (ZIP size is unknown upfront).
func ServeDirWithProgress(w http.ResponseWriter, r *http.Request, dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil || !info.IsDir() {
		http.Error(w, "Directory not found", http.StatusNotFound)
		return fmt.Errorf("stat dir %q: %w", dirPath, err)
	}

	bar := progressbar.Default(-1, "Streaming ZIP")

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, info.Name()))

	pw := NewProgressWriter(r.Context(), w, bar)
	zipWriter := zip.NewWriter(pw)
	defer func() {
		zipWriter.Close()
		fmt.Println() // newline after bar
	}()

	return filepath.Walk(dirPath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		zipEntry, err := zipWriter.Create(filepath.ToSlash(relPath))
		if err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(zipEntry, f)
		return err
	})
}
