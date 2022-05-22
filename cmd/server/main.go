package main

import (
    "flag"
    "log"
    "net/http"

)

func main() {
    port := flag.String("port", "8081", "port")
    flag.Parse()
    log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("."))))
}
