package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/upgrade")
}

func handlerUpgrade()
