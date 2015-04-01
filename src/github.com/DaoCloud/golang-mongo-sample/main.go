package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	log()
	http.HandleFunc("/", hello)
	http.HandleFunc("/env", env)

	fmt.Println("Start listening...")
	go log()
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func log() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello World !" + time.Now().String())
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World, ", GetResult())
}

func env(res http.ResponseWriter, req *http.Request) {
	env := os.Environ()

	fmt.Fprintln(res, "List of Environtment variables : \n")

	for index, value := range env {
		name := strings.Split(value, "=") // split by = sign

		fmt.Fprintf(res, "[%d] %s : %v\n", index, name[0], name[1])
	}
}
