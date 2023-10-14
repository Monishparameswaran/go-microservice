package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	info "github.com/Monishparameswaran/go-microservice/ContainerInfo"
	"github.com/gorilla/mux"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, http.FileServer(http.Dir("static/")))
// }

func orderDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	response := map[string]string{
		"Book name":  title,
		"timestamp":  time.Now().String(),
		"developerr": "Monish",
	}
	json.NewEncoder(w).Encode(response)
}

func health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"ip":        info.IP().String(),
		"hostname":  info.Hostname(),
		"developer": "Monish",
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}", orderDetails)
	r.HandleFunc("/health", health)
	log.Println("Server has started serving at 8070")
	http.ListenAndServe(":8070", r)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w, "Welcome to my website!");
// }
// func main() {
//     http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

//     })

//     fs := http.FileServer(http.Dir("static/"))
//     http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	log.Println("Server started at localhost:8070");
//     log.Fatal(http.ListenAndServe(":8070", nil));
// }
