package main

import (
	"net/http"
	"html/template"
	"fmt"
	"io/ioutil"
	"regexp"
	"reflect"
	"strings"
	"os"
)

type CV struct {
	activities  []Activity
	Bio         Bio
	educations  []Education
	experiences []Experience
}

type Activity struct {
	Time        string
	Title       string
	Description string
}

type Bio struct {
	FirstName     string
	SecondName    string
	LastName      string
	BirthDate     string
	BirthLocation string
	Street        string
	HouseNumber   string
	PostalCode    string
	City          string
	Country       string
	Phone         string
	Email         string
	Languages     map[string]string
	Skills        map[string]map[string]string
	Links         map[string][]string
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
	tmpl, err := template.New("cv").ParseFiles("templates/cv.go.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, cv)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	readCv()
	http.HandleFunc("/", cv)
	http.ListenAndServe(":8000", nil)
}

const activitiesDir = "cv/activity"
const bioDir = "cv/bio"

func readCv() {
	cv := *new(CV)

	// Read activities
	typ := reflect.TypeOf(Activity{})
	files, _ := ioutil.ReadDir(activitiesDir)
	cv.activities = make([]Activity, len(files))
	for i, elem := range files {
		activity := fillStruct(activitiesDir, elem, typ).(Activity)
		cv.activities[i] = activity
	}
	// Read bio
	typ = reflect.TypeOf(Bio{})
	files, _ = ioutil.ReadDir(bioDir)
	cv.Bio = fillStruct(bioDir, files[0], typ).(Bio)

	fmt.Println(cv)
}

func fillStruct(dir string, fileInfo os.FileInfo, typ reflect.Type) interface{} {
	data, err := ioutil.ReadFile(dir + "/" + fileInfo.Name())
	text := string(data)
	if err != nil {
		fmt.Println(err)
	}
	r, _ := regexp.Compile("@([a-z_]*)=\"([A-Za-z0-9 .:/äöüß+\\-]*)\"")
	submatches := r.FindAllStringSubmatch(text, -1)
	struc := reflect.New(typ).Elem()
	for _, submatch := range submatches {
		name := strings.Title(submatch[1])
		underscoreFinder, _ := regexp.Compile("_([a-z])")
		characters := underscoreFinder.FindAllStringSubmatch(name, -1)
		for _, character := range characters {
			name = strings.Replace(name, character[0], strings.ToUpper(character[1]), -1)
		}
		fmt.Println(name)
		field := struc.FieldByName(name)
		field.SetString(submatch[2])
	}
	return struc.Interface()
}