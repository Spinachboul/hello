package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloworld)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

}
