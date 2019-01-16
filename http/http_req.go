// make_http_request.go
package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    // Make HTTP GET request
    // response, err := http.Get("https://weather.com/es-ES/tiempo/hoy/l/ae60f7c43022aa56e26e3498ec7ff0293103aa8e67c59886ae7c33b3ce0355dc")
    response, err := http.Get("http://127.0.0.1:8889/") //to test in local server
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Copy data from the response to standard output
    n, err := io.Copy(os.Stdout, response.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Number of bytes copied to STDOUT:", n)
}