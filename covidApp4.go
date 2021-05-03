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
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India") //Get json data from api
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
	p2, err := template.ParseFiles("html/headers_State.html") //to print header of table
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
		}
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64) //convert float to string
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)
		//fmt.Printf("%T, %v\n", s, s)
		//fmt.Printf("%T, %v\n", t, t)
		//fmt.Printf("%T, %v\n", u, u)
		data := []string{i, s, t, u} //pass converted string data to a variable
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp\\data\\state.csv")
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
		p1, err := template.ParseFiles("html/states.html") //pass data to html table
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
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases") //Get json data from api
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body)) //read json data

	if err != nil {
		log.Fatal(err)
	}
	responseObject := map[string]interface{}{} //create a map to keep all types of values using interface
	json.Unmarshal(data, &responseObject)      //converts json to an object type
	stringdata := string(data)                 //convert all values to string
	//fmt.Println(stringdata)
	var p fastjson.Parser
	v, err := p.Parse(stringdata) //pass the converted string data to get all keys
	if err != nil {
		log.Fatal(err)
	}
	var keys []string                                              //create a slice of string
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) { //without creating multiple structs get all keys from json object.
		keys = append(keys, string(key)) //append keys to a variable of keys
	})
	//fmt.Println(keys)
	var Capital_city string //create variables of json type to compare and get data
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	p1, err := template.ParseFiles("html/headers_Country.html")
	if err != nil {
		log.Fatalln(err)
	}
	p1.Execute(w, "Hi")
	for _, i := range keys { //iterate through keys
		state := responseObject[i].(map[string]interface{}) //create a map to get all types of data
		for _, value := range state {
			all := value.(map[string]interface{}) //create another map to go inside and get values
			//country=i
			for key, value := range all { //to go inside each country and get data
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
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64) //convert all types to string
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)

		data := []string{i, Capital_city, s, t, u} //pass converted string data to varaible
		file := ("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp3\\data\\data.csv")
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //create a csv file
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		m := csv.NewWriter(f)
		m.Write(data) //write data to csv file
		m.Flush()

		err = m.Error()
		if err != nil {
			log.Fatalln(err)
		}
		p1, err := template.ParseFiles("html/country.html") //pass the filtered data to html table
		data1 := Country{i, Capital_city, s, t, u}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homePage, err := template.ParseFiles("html/home.html")
	if err != nil {
		panic(err)
	}
	homePage.Execute(w, "Home")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/2", stateCovid)
	http.HandleFunc("/1", countryCovid)
	//http.HandleFunc("/3", VaccinatedDetails)
	log.Fatal(http.ListenAndServe(":6014", nil))
}
func main() {
	handleRequests()
}
