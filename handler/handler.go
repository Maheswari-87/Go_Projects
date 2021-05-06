package handler

import (
	"CovidWebApp/covidDetails"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", covidDetails.Home)
	http.HandleFunc("/IndiaStates", covidDetails.IndiaStateCovid)
	http.HandleFunc("/Worldwide", covidDetails.CountryWiseCovid)
	http.HandleFunc("/states", covidDetails.IndividualStates)
	//http.HandleFunc("/3", VaccinatedDetails)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
