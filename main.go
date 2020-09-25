package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, this is multibranches app v3")
	fmt.Printf("Result of call Min(2, 3) is: %d\n", Min(2, 3))
	fmt.Println("Start web server")
	http.HandleFunc("/ping", Ping)
	_ = http.ListenAndServe(":8000", nil)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s\n", "Pong v3")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
