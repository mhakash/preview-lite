package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func getAvailablePort(startPort int) (int, error) {
	for port := startPort; port <= 65535; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close()
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available port found")
}

func main() {
	var dir string
	if len(os.Args) < 2 {
		execPath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dir = filepath.Dir(execPath)
		fmt.Println("execPath: ", dir)
	} else {
		dir = os.Args[1]
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		fmt.Println("path: ", path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(dir, "index.html"))
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	port, err := getAvailablePort(8081)
	if err != nil {
		log.Fatalf("Could not find an available port: %v\n", err)
	}

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		url := fmt.Sprintf("http://localhost:%d", port)
		fmt.Printf("Serving %s on HTTP port: %d\n", dir, port)
		fmt.Printf("Open %s in your browser\n", url)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :%d: %v\n", port, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
