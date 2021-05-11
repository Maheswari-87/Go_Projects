package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	router.HandleFunc("/customer", Customer).Methods(http.MethodGet)
	router.HandleFunc("/customer", PostReqCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8012", router))

}
func PostReqCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post request recieved")
}
