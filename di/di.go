package di

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, "+name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// func main() {
// 	http.ListenAndServe(":8080", http.HandlerFunc(MyGreeterHandler))
// }
