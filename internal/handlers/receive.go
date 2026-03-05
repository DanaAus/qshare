package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"magshare/internal/network"
	"magshare/internal/server"
	uiHelper "magshare/internal/ui"
	"magshare/ui"

	"github.com/mdp/qrterminal/v3"
	"github.com/schollz/progressbar/v3"
)

// ReceiveOptions defines configuration for the receive server.
type ReceiveOptions struct {
	Port   int
	Secure bool
	PIN    string
}

// StartReceiveServer initializes the ephemeral server to receive files.
func StartReceiveServer(destDir string, opts ReceiveOptions) error {
	// 1. Discover Network Interface and Port
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

	// 2. Generate secure upload URL and optional PIN
	hash, err := network.GenerateRandomString(6) // 12-char hex
	if err != nil {
		return err
	}

	uploadPath := fmt.Sprintf("/u/%s", hash)
	uploadURL := fmt.Sprintf("http://%s:%d%s", ip, port, uploadPath)

	pin := opts.PIN
	if opts.Secure && pin == "" {
		pin, err = server.GeneratePIN()
		if err != nil {
			return err
		}
	}

	// Output Info
	fmt.Printf("[Network] Using active interface: %s\n", ip)
	fmt.Printf("[Server]  Dropzone started on port %d\n", port)
	if opts.Secure {
		fmt.Printf("[Auth]    PIN REQUIRED: %s\n", pin)
	}
	fmt.Printf("[URL]     %s\n", uploadURL)

	// Print QR
	fmt.Println("[QR]")
	qrterminal.GenerateHalfBlock(uploadURL, qrterminal.L, os.Stdout)

	// 3. Setup Server
	srv := server.NewEphemeralServer(port)

	// Serve the UI Dropzone
	srv.Handle(uploadPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		tmpl, err := template.ParseFS(ui.Files, "dropzone.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Secure bool
		}{
			Secure: opts.Secure,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			fmt.Printf("[Error] Failed to render UI: %v\n", err)
		}
	})

	// Handle the actual file upload
	srv.Handle("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Handle upload with progress bar
		ReceiveFileWithProgress(w, r, destDir, opts.Secure, pin)
	})

	fmt.Println("\nStatus: Ready to receive files... Press Ctrl+C to stop. (timeout 5m)")

	// 5. Start Server with 5-minute timeout
	return srv.Start(5 * time.Minute)
}

// ReceiveFileWithProgress handles file upload with progress bar.
func ReceiveFileWithProgress(w http.ResponseWriter, r *http.Request, destDir string, secure bool, expectedPin string) {
	// 1. Wrap Body for Progress
	contentLength := r.ContentLength
	bar := uiHelper.NewProgressBar(contentLength, "Receiving Upload")
	defer bar.Finish()
	// Use NewProxyReader to wrap body
	// Since r.Body is a ReadCloser, and NewReader returns Reader, we wrap with NopCloser
	proxyReader := progressbar.NewReader(r.Body, bar)
	r.Body = io.NopCloser(&proxyReader)

	// 2. Parse Multipart
	// Maximum 5GB upload size
	// Note: ParseMultipartForm reads the body, which triggers the progress bar updates.
	err := r.ParseMultipartForm(5000 << 20)
	if err != nil {
		http.Error(w, "File too large or malformed", http.StatusBadRequest)
		return
	}

	// 3. Check PIN if secure mode is enabled
	// We check AFTER parsing because the PIN is likely in the form data
	if secure {
		clientPin := r.FormValue("pin")
		if clientPin != expectedPin {
			http.Error(w, "Invalid PIN", http.StatusUnauthorized)
			return
		}
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file format", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("\n[Server] Receiving file '%s' from %s\n", header.Filename, r.RemoteAddr)

	// Sanitize file name
	sanitizedName := filepath.Base(header.Filename)

	// Destination
	if destDir == "" {
		var err error
		destDir, err = os.Getwd()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	dstPath := filepath.Join(destDir, sanitizedName)
	dst, err := os.Create(dstPath)
	if err != nil {
		fmt.Printf("[Error] Failed to create file: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		fmt.Printf("[Error] Failed to save file: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("[Success] Saved to %s\n", dstPath)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File received successfully"))
}
