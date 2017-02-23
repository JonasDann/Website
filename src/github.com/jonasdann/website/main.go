package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type CV struct {
	activities  []Activity
	Bio         Bio
	educations  []Education
	experiences []Experience
}

type Activity struct {
	time        string
	title       string
	description string
}

type Bio struct {
	FirstName     string
	SecondName    string
	LastName      string
	birthDate     string
	birthLocation string
	street        string
	houseNumber   string
	postalCode    string
	city          string
	country       string
	phone         string
	email         string
	languages     map[string]string
	skills        map[string]map[string]string
	links         map[string][]string
}

type Education struct {
	from               string
	until              string
	expectedGraduation string
	name               string
	location           string
	degree             string
	grade              string
	description        string
}

type Experience struct {
	from        string
	until       string
	name        string
	location    string
	position    string
	department  string
	description string
	appliedTech []string
}

type Page struct {
	title string
	body  template.HTML
}

func cv(w http.ResponseWriter, r *http.Request) {
	cv := *new(CV)
	cv.Bio = *new(Bio)
	cv.Bio.FirstName = "Jonas"
	cv.Bio.SecondName = "Christian"
	cv.Bio.LastName = "Dann"
	fmt.Println(cv)
	tmpl, err1 := template.New("cv").ParseFiles("templates/cv.go.html")
	fmt.Println(err1)
	err2 := tmpl.Execute(w, cv)
	fmt.Println(err2)
}

func main() {
	http.HandleFunc("/", cv)
	http.ListenAndServe(":8000", nil)
}
