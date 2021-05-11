package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	Name    string `json:"person_name" xml:"xml-name"`
	Address string `json:"person_adress" xml:"xml-address"`
	Pincode string `json:"person_pincode" xml:"xml-pincode"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func Customer(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{Name: "Mahi", Address: "DSr", Pincode: "531011"},
		{Name: "Nani", Address: "Npt", Pincode: "531012"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}
