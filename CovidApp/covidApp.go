package main

import (
	"CovidApp/util"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "homepage")
	url := util.ReadUrl()
	response, err := http.Get(url)

	//response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
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
	var Updated string
	p2, err := template.ParseFiles("html/headers.html")
	if err != nil {
		panic(err)
	}
	p2.Execute(w, "HI")
	//flag := 0
	for _, i := range keys {
		state := responseObject[i].(map[string]interface{})
		for key, value := range state {
			if key == "confirmed" && value != nil {
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

		//s := []float64{Confirmed, Recovered, Deaths}
		//var s []float64 =abc(Confirmed,Recovered,Deaths)
		//p1.Execute(w, s)
		p1, err := template.ParseFiles("html/states.html")
		//data1 := Data{i, Confirmed, Recovered, Deaths}
		data1 := Data{i, s, t, u, Updated}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7050", nil))
}
func main() {
	handleRequests()
}

/*func main() {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=Mauritius")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	if err != nil {
		log.Fatal(err)
	}
	//data1:=(string(data))
	var responseObject Result
	err = json.Unmarshal(data, &responseObject)
	if err != nil {
		panic(err)
	}

	//for key, value := range data {
	// Each value is an interface{} type, that is type asserted as a string
	//	fmt.Println(key, value.(string))
	//}
	//handleRequests()
}*/

//out := Andaman1{Confirmed: responseObject.Andaman.Confirmed, Recovered: responseObject.Andaman.Recovered, Deaths: responseObject.Andaman.Deaths}
//out1 := Andhra1{Confirmed: responseObject.Andhra.Confirmed, Recovered: responseObject.Andhra.Recovered, Deaths: responseObject.Andhra.Deaths}
//out2 := Arunachal1{Confirmed: responseObject.Arunachal.Confirmed, Recovered: responseObject.Arunachal.Recovered, Deaths: responseObject.Arunachal.Deaths}
//p1, err := template.ParseFiles("html/covid.html")
//if err != nil {
//	panic(err)
//}
//p1.Execute(w)
//p1.Execute(w, out)
///p1.Execute(w, out1)
//p1.Execute(w, out2)
