package main

import (
    "net/http"
    "path"
    "log"
    "os"
)

func main() {
    http.HandleFunc("/css/", serveStatic("", "text/css"))
    http.HandleFunc("/js/", serveStatic("", "application/javascript"))
    http.HandleFunc("/", serveStatic("/index.html", "text/html"))
    log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func serveStatic(filename, contentType string) func(http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        var filepath string
        if filename == "" {
            filepath = path.Join("client/dist/", req.URL.Path)
        } else {
            filepath = path.Join("client/dist/", filename)
        }

        res.Header().Set("Content-Type", contentType)
        http.ServeFile(res, req, filepath)
    }
}