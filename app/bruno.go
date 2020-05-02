package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

func loop() float64 {
	x := 0.0001
	for index := 0; index < 10000000; index++ {
		x = x + math.Sqrt(x)
		fmt.Println(x)
	}
	return x
}

func greeting(text string) string {
	data := fmt.Sprintf("<b>%s</b>", text)
	return data
}

func BrunoFunc(w http.ResponseWriter, r *http.Request) {
	mensagem := greeting("Bruno Leal - Code.education Rocks!")
	//w.Write([]byte(mensagem))
	fmt.Fprintf(w, mensagem)
	loop()
	fmt.Println("Fim!")
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", BrunoFunc)
	http.ListenAndServe(":80", route)
}
