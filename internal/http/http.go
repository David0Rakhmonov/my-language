package http

import (
	"fmt"
	"net/http"
)

// Listen запускает HTTP-сервер на указанном порту
func Listen(port int, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc("/", handler)
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
