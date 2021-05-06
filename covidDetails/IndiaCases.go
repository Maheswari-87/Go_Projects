package covidDetails

import (
	"CovidWebApp/urlData"
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
	Updated   string
}

func IndiaStateCovid(w http.ResponseWriter, r *http.Request) {
	url := urlData.ReadIndiaUrl()
	response, err := http.Get(url)
	//response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India") //Get json data from api
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body)) //Read all json Data

	if err != nil {
		log.Fatal(err)
	}
	responseObject := map[string]interface{}{} //Take all types of values using map and interface
	json.Unmarshal(data, &responseObject)      //converts json to an object type
	stringdata := string(data)                 //keeping values in string form
	//fmt.Println(stringdata)
	var p fastjson.Parser
	v, err := p.Parse(stringdata)
	if err != nil {
		log.Fatal(err)
	}
	var keys []string //create slice of string to keep data
	//without fastjson we need to create multiple structs to get all keys in api.
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) { //get all keys from object
		keys = append(keys, string(key)) //append each key to a variable
	})
	fmt.Println(keys)
	var Confirmed float64 //create variables of json type to compare and get data
	var Recovered float64
	var Deaths float64
	var Updated string
	p2, err := template.ParseFiles("html/header_india.html") //to print header of table
	if err != nil {
		panic(err)
	}
	p2.Execute(w, "HI")      //execute template
	for _, i := range keys { //range over each key
		state := responseObject[i].(map[string]interface{}) //create an interface to get all types of data
		for key, value := range state {
			if key == "confirmed" && value != nil { //check for key in json
				Confirmed = value.(float64)
			}
			if key == "recovered" && value != nil {
				Recovered = value.(float64)
			}
			if key == "deaths" && value != nil {
				Deaths = value.(float64)
			}
			if key == "updated" && value != nil {
				Updated = value.(string)
			}
		}
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64) //convert float to string
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)
		//fmt.Printf("%T, %v\n", s, s)
		//fmt.Printf("%T, %v\n", t, t)
		//fmt.Printf("%T, %v\n", u, u)
		data := []string{i, s, t, u, Updated} //pass converted string data to a variable
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidWebApp\\csvData\\india.csv")
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //create csv file
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		m := csv.NewWriter(f)

		m.Write(data) //write the string data in csv file
		m.Flush()
		err = m.Error()
		if err != nil {
			log.Fatalln(err)
		}
		p1, err := template.ParseFiles("html/india.html") //pass data to html table
		//data1 := Data{i, Confirmed, Recovered, Deaths}
		data1 := Data{i, s, t, u, Updated}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}
	//webPage.Execute(w, states)

}
