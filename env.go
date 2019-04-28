package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, req *http.Request) {
	if env, exist := os.LookupEnv("HOSTNAME"); exist {
		fmt.Fprintf(w, "%+v: %+v\n", "HOSTNAME", env)
	}
	fmt.Fprintf(w, "%+v: %+v\n", "PORT", os.Getenv("PORT"))

	if _, exist := os.LookupEnv("CF_INSTANCE_GUID"); exist {
		fmt.Fprintf(w, "%+v: %+v\n", "CF_INSTANCE_GUID", os.Getenv("CF_INSTANCE_GUID"))
		fmt.Fprintf(w, "%+v: %+v\n", "CF_INSTANCE_ADDR", os.Getenv("CF_INSTANCE_ADDR"))
		fmt.Fprintf(w, "%+v: %+v\n", "CF_INSTANCE_INDEX", os.Getenv("CF_INSTANCE_INDEX"))
		fmt.Fprintf(w, "%+v: %+v\n", "VCAP_SERVICES", os.Getenv("VCAP_SERVICES"))
	}

	fmt.Fprintln(w, "")

	fmt.Printf("%+v\n", req)

	fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))

	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	addr := ":" + os.Getenv("PORT")

	fmt.Printf("Listening on %v\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
