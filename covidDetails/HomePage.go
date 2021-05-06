package covidDetails

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	homePage, err := template.ParseFiles("html/home.html")
	if err != nil {
		panic(err)
	}
	homePage.Execute(w, "Home")
}
