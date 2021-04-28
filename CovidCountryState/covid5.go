package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/valyala/fastjson"
)

type CountryState struct {
	Country   string
	State     string
	Confirmed string
	Recovered string
	Deaths    string
}

func CountryStateCovid(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	if err != nil {
		log.Fatal(err)
	}
	responseObject := map[string]interface{}{}
	json.Unmarshal(data, &responseObject)
	stringdata := string(data)
	var p fastjson.Parser
	v, err := p.Parse(stringdata)
	if err != nil {
		log.Fatal(err)
	}
	var keys []string
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) {
		keys = append(keys, string(key))
	})
	fmt.Println(keys)
	var State string
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	p1, err := template.ParseFiles("html/header.html")
	if err != nil {
		log.Fatalln(err)
	}

	p1.Execute(w, "Hi")

	for _, i := range keys {

		all := responseObject[i].(map[string]interface{})

		for k, v := range all {
			if k == "All" {
				allV := v.(map[string]interface{})
				for key, value := range allV {
					if key == "capital_city" && value != nil {
						State = value.(string)
					}
					if key == "confirmed" && value != nil {
						Confirmed = value.(float64)
					}
					if key == "recovered" && value != nil {
						Recovered = value.(float64)
					}
					if key == "deaths" && value != nil {
						Deaths = value.(float64)
					}
				}
			}
		}
		fmt.Println(State)
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)

		data := []string{i, s, t, u}
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp4\\data\\CountryStatedata.csv")
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		m := csv.NewWriter(f)
		m.Write(data)
		m.Flush()

		err = m.Error()
		if err != nil {
			log.Fatalln(err)
		}

		p1, err := template.ParseFiles("html/state.html")
		data1 := CountryState{i, State, s, t, u}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	homePage, err := template.ParseFiles("html/new.html")
	if err != nil {
		panic(err)
	}
	homePage.Execute(w, "Home")
}
func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/1", CountryStateCovid)
	log.Fatal(http.ListenAndServe(":7067", nil))
}
func main() {
	handleRequests()
}
