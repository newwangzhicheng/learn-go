package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "wzc")
	//cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet

	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
}

//MyGreetHandler for http
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

//Greet print greet
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
