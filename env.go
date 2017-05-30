package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "strings"
)

func handler(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "%+v: %+v\n", "STACKATO_APP_NAME", os.Getenv("STACKATO_APP_NAME"))
        fmt.Fprintf(w, "%+v: %+v\n", "CF_INSTANCE_ADDR", os.Getenv("CF_INSTANCE_ADDR"))
        fmt.Fprintf(w, "%+v: %+v\n", "CF_INSTANCE_INDEX", os.Getenv("CF_INSTANCE_INDEX"))
        fmt.Fprintf(w, "%+v: %+v\n", "STACKATO_APP_ENV", os.Getenv("STACKATO_APP_ENV"))
        fmt.Fprintf(w, "%+v: %+v\n", "STACKATO_SERVICES", os.Getenv("STACKATO_SERVICES"))
        fmt.Fprintln(w, "")

        fmt.Printf("%+v\n", req)
        fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
        if req.URL.Path == "/crash" {
                os.Exit(1)
        }
}

func main() {
        http.HandleFunc("/", handler)
        addr := ":" + os.Getenv("PORT")
        fmt.Printf("Listening on %v\n", addr)
        log.Fatal(http.ListenAndServe(addr, nil))
}
