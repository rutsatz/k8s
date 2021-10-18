// Arquivo para subir um server go simples para ser gerenciado pelo kubernetes
package main

import "net/http"

func main() {
	// Quando acessar o endereço "/", ele executa a função "Hello"
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}
