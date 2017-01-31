package main

import (
	"flag"
	"net/http"
)

func main() {
	dir := flag.String("dir", "staticwebsite", "Directory of Static Website")
	flag.Parse()
	http.ListenAndServe(":8000", http.FileServer(http.Dir(*dir)))
}
