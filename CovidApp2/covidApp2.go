package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type AllInfo struct {
	All India `json:"All"`
}
type India struct {
	Country   string `json:"country"`
	Confirmed int    `json:"confirmed"`
	Recovered int    `json:"recovered"`
	Deaths    int    `json:"Deaths"`
}

func home(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	//dataString := string(data)
	//fmt.Println(dataString)

	if err != nil {
		log.Fatal(err)
	}
	var responseObject AllInfo
	json.Unmarshal(data, &responseObject)
	out := India{Country: responseObject.All.Country, Confirmed: responseObject.All.Confirmed, Recovered: responseObject.All.Recovered, Deaths: responseObject.All.Deaths}
	p1, err := template.ParseFiles("templates/test.html")
	if err != nil {
		panic(err)
	}
	p1.Execute(w, out)

}
func handleReq() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8066", nil))
}
func main() {
	handleReq()
}
