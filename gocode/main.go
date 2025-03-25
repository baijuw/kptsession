package main

import (
        "net"
        "net/http"
        "os"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request to:", r.URL.Path)
        hostname, err := os.Hostname()
        if err != nil {
                http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
                return
        }

        clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
                http.Error(w, "Unable to get client IP", http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "text/html") // Set the content type

        fmt.Fprintf(w, `
                <!DOCTYPE html>
                <html>
                <head>
                        <title>Hostname and Client IP</title>
                        <style>
                          body {
                            font-family: sans-serif;
                            background-color: #f0f0f0;
                            display: flex;
                            justify-content: center;
                            align-items: center;
                            min-height: 100vh;
                          }
                          .container {
                            background-color: #fff;
                            padding: 20px;
                            border-radius: 8px;
                            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
                            text-align: center;
                          }
                          h1 {
                            color: #333;
                            margin-bottom: 10px;
                          }
                        </style>
                </head>
                <body>
                        <div class="container">
                                <h1>Hostname: %s</h1>
                                <h1>Client IP: %s</h1>
				<h1> Version: 0.2</h1>
                        </div>
                </body>
                </html>
        `, hostname, clientIP)
}

func main() {
	fmt.Println("Starting server....")
        http.HandleFunc("/", handler)
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                fmt.Println("Error starting server:", err)
        }
}

