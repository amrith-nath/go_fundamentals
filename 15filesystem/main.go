package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hellp, %q", html.EscapeString(r.URL.Path))

	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hellp, %q", html.EscapeString(r.URL.Path))

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

	// content := "this is a sample text"

	// file, err := os.Create("./sampletext.text")
	// chechError(err)

	// length, er := io.WriteString(file, content)

	// chechError(er)

	// file.Close()

	// readFile(file.Name())

	// fmt.Println("length is :", length)

}

func readFile(filename string) {

	data, err := os.ReadFile(filename)

	chechError(err)
	fmt.Println(string(data))
}

func chechError(err error) {
	if err != nil {

		panic(err)
	}
}
