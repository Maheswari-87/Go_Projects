package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fastjson"
)

type Country struct {
	All string `json:"All"`
}

/*type All struct {
	Confirmed float64
	Recovered float64
	Deaths    float64
}*/

type Data struct {
	State        string
	Capital_city string
	Confirmed    string
	Recovered    string
	Deaths       string
}

/*func parseData() (*Data, error) {

}*/

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "homepage")

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
	//fmt.Println(v)
	var keys []string
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) {
		keys = append(keys, string(key))
	})
	fmt.Println(keys)

	//fmt.Println(all)
	//m := make(map[string]interface{})
	var Capital_city string
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	p1, err := template.ParseFiles("html/headers.html")
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
			//header := []string{`State`, `Capital_city`, `Confirmed`, `Recovered`, `Deaths`}
		}
		//}
		//}
		//fmt.Println(v)
		//}

		fmt.Println(i)
		fmt.Println(Capital_city)
		fmt.Println(Confirmed)
		fmt.Println(Recovered)
		fmt.Println(Deaths)
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)
		//fmt.Printf("%T, %v\n", s, s)
		//fmt.Printf("%T, %v\n", t, t)
		//fmt.Printf("%T, %v\n", u, u)

		//data := [][]string{{`State`, `Capital_city`, `Confirmed`, `Recovered`, `Deaths`},
		//	{i, Capital_city, s, t, u}}
		//f, err := os.Create("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp3\\data\\db.csv")
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//fmt.Println(user)
		//defer f.Close()
		//m := csv.NewWriter(f)
		//data1 := [][]string{
		//	header, user1, user,
		//}
		//m.Write(header)
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
		//m.WriteAll(data)

		//	m.WriteAll(data)
		//for _, user := range db.users {
		//ss := user.EncodeAsStrings()
		//w.Write(ss)
		//	}
		err = m.Error()
		if err != nil {
			log.Fatalln(err)
		}

		//s := []float64{Confirmed, Recovered, Deaths}
		//var s []float64 =abc(Confirmed,Recovered,Deaths)
		//p1.Execute(w, s)
		/*header := []string{`State`, `Capital_city`, `Confirmed`, `Recovered`, `Deaths`}
		s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
		t := strconv.FormatFloat(Recovered, 'f', -1, 64)
		u := strconv.FormatFloat(Deaths, 'f', -1, 64)
		fmt.Printf("%T, %v\n", s, s)
		fmt.Printf("%T, %v\n", t, t)
		fmt.Printf("%T, %v\n", u, u)

		data := []string{i, Capital_city, s, t, u}
		f, err := os.Create("C:\\Users\\SRS\\gocode\\src\\workspace\\CovidApp3\\data\\db.csv")
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(user)
		defer f.Close()
		m := csv.NewWriter(f)
		//data1 := [][]string{
		//	header, user1, user,
		//}
		m.Write(header)
		for i := 0; i < len(data); i++ {
			m.Write(data)
		}

		//for _, user := range db.users {
		//ss := user.EncodeAsStrings()
		//w.Write(ss)
		//	}
		m.Flush()
		err = m.Error()
		if err != nil {
			log.Fatalln(err)
		}*/

		//m := csv.NewWriter(os.Stdout)
		p1, err := template.ParseFiles("html/country.html")
		data1 := Data{i, Capital_city, s, t, u}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data1)
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7097", nil))
}
func main() {
	/*db, err := readJSONFile("CovidApp3/db.json")
	if err != nil {
		log.Fatalln(err)
	}
	//header := []string{`State`, `Capital_city`, `Confirmed`, `Recovered`, `Deaths`}
	//user1 := []string{`Afghanistan`, `Kabul`, `58312`, `52348`, `2561`}
	//user := []string{
	//	`Albania`, `Tirana`, `130114`, `103582`, `2364`,
	//}
	f, err := os.Create("CovidApp3/db.csv")
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(user)
	defer f.Close()
	w := csv.NewWriter(f)
	//data1 := [][]string{
	//	header, user1, user,
	//}
	w.Write(types.GetHeader())
	for _, user := range db.users {
		ss := user.EncodeAsStrings()
		w.Write(ss)
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		log.Fatalln(err)
	}*/
	handleRequests()
}

/*func readJSONFile(s string) { //(db *types.UserDb,err error){
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var dec = json.NewDecoder(f)
	db := new(types.User)
	dec.Decode(db)
	return
}*/

/*func main() {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
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
	fmt.Printf("India:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.All.Confirmed, responseObject.All.Recovered, responseObject.All.Confirmed-responseObject.All.Recovered, responseObject.All.Deaths)
	fmt.Printf("\n\nAndaman:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Andaman.Confirmed, responseObject.Andaman.Recovered, responseObject.Andaman.Confirmed-responseObject.Andaman.Recovered, responseObject.Andaman.Deaths)
	fmt.Printf("\n\nAndhra:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Andhra.Confirmed, responseObject.Andhra.Recovered, responseObject.Andhra.Confirmed-responseObject.Andhra.Recovered, responseObject.Andhra.Deaths)
	fmt.Printf("\n\nArunachal pradesh:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Arunachal.Confirmed, responseObject.Arunachal.Recovered, responseObject.Arunachal.Confirmed-responseObject.Arunachal.Recovered, responseObject.Arunachal.Deaths)
	fmt.Printf("\n\nAssam:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Assam.Confirmed, responseObject.Assam.Recovered, responseObject.Assam.Confirmed-responseObject.Assam.Recovered, responseObject.Assam.Deaths)
	fmt.Printf("\n\nBihar:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Bihar.Confirmed, responseObject.Bihar.Recovered, responseObject.Bihar.Confirmed-responseObject.Bihar.Recovered, responseObject.Bihar.Deaths)
	fmt.Printf("\n\nChandigarh:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Chandigarh.Confirmed, responseObject.Chandigarh.Recovered, responseObject.Chandigarh.Confirmed-responseObject.Chandigarh.Recovered, responseObject.Chandigarh.Deaths)

	//fmt.Println(len(responseObject.IndianStates))
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
/*s := strconv.FormatFloat(Confirmed, 'f', -1, 64)
t := strconv.FormatFloat(Confirmed, 'f', -1, 64)
u := strconv.FormatFloat(Confirmed, 'f', -1, 64)
fmt.Printf("%T, %v\n", s, s)
fmt.Printf("%T, %v\n", t, t)
fmt.Printf("%T, %v\n", u, u)

data := []string{"i", "Capital_city", "s", "t", "u"}*/
