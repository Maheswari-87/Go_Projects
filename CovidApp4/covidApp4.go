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

type Data struct {
	State     string
	Confirmed string
	Recovered string
	Deaths    string
}
type Country struct {
	State        string
	Capital_city string
	Confirmed    string
	Recovered    string
	Deaths       string
}

func stateCovid(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "homepage")

	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	//webPage, err := template.ParseFiles("html/states.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//states := []Data{}
	//details := Data{}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	if err != nil {
		log.Fatal(err)
	}
	responseObject := map[string]interface{}{}
	//var responseObject Result
	json.Unmarshal(data, &responseObject)
	stringdata := string(data)
	//fmt.Println(stringdata)
	var p fastjson.Parser
	v, err := p.Parse(stringdata)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(v)
	var keys []string
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) {
		keys = append(keys, string(key))
	})
	fmt.Println(keys)

	//m := make(map[string]interface{})
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	p2, err := template.ParseFiles("html/headers_State.html")
	if err != nil {
		panic(err)
	}
	p2.Execute(w, "HI")
	//flag := 0
	for _, i := range keys {
		state := responseObject[i].(map[string]interface{})
		for key, value := range state {
			//details:=Data{}
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
		fmt.Println(i)
		fmt.Println(Confirmed)
		fmt.Println(Recovered)
		fmt.Println(Deaths)
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)
		//fmt.Printf("%T, %v\n", s, s)
		//fmt.Printf("%T, %v\n", t, t)
		//fmt.Printf("%T, %v\n", u, u)
		data := []string{i, s, t, u}
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp\\data\\state.csv")
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
		p1, err := template.ParseFiles("html/states.html")
		//data1 := Data{i, Confirmed, Recovered, Deaths}
		data1 := Data{i, s, t, u}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}
	//webPage.Execute(w, states)

}
func countryCovid(w http.ResponseWriter, r *http.Request) {
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
	//var responseObject Result
	json.Unmarshal(data, &responseObject)
	stringdata := string(data)
	//fmt.Println(stringdata)
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
	var Capital_city string
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	p1, err := template.ParseFiles("html/headers_Country.html")
	if err != nil {
		log.Fatalln(err)
	}

	p1.Execute(w, "Hi")

	for _, i := range keys {

		state := responseObject[i].(map[string]interface{})

		for _, value := range state {
			all := value.(map[string]interface{})
			//country=i
			for key, value := range all {
				if key == "capital_city" && value != nil {
					Capital_city = value.(string)
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

		fmt.Println(i)
		fmt.Println(Capital_city)
		fmt.Println(Confirmed)
		fmt.Println(Recovered)
		fmt.Println(Deaths)
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)

		data := []string{i, Capital_city, s, t, u}
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp3\\data\\data.csv")
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

		p1, err := template.ParseFiles("html/country.html")
		data1 := Country{i, Capital_city, s, t, u}
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
	http.HandleFunc("/2", stateCovid)
	http.HandleFunc("/1", countryCovid)
	log.Fatal(http.ListenAndServe(":7079", nil))
}
func main() {
	handleRequests()
}
