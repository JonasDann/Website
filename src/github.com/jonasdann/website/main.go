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
	"encoding/json"
)

type CV struct {
	Activities  []Activity
	Bio         Bio
	Educations  []Education
	Experiences []Experience
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
	From               string
	Until              string
	ExpectedGraduation string
	Name               string
	Location           string
	Degree             string
	Grade              string
	Description        string
}

type Experience struct {
	From        string
	Until       string
	Name        string
	Location    string
	Position    string
	Department  string
	Description string
	AppliedTech []string
}

type Page struct {
	title string
	body  template.HTML
}

func showCv(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("cv").ParseFiles("templates/cv.go.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, curriculumVitae)
	if err != nil {
		fmt.Println(err)
	}
}

var curriculumVitae CV

func main() {
	curriculumVitae = readCv()
	http.HandleFunc("/", showCv)
	http.ListenAndServe(":8000", nil)
}

const activitiesDir = "cv/activity"
const bioDir = "cv/bio"
const educationDir = "cv/education"
const experienceDir = "cv/experience"

func readCv() CV {
	cv := *new(CV)

	// Read Activities
	typ := reflect.TypeOf(Activity{})
	files, _ := ioutil.ReadDir(activitiesDir)
	cv.Activities = make([]Activity, len(files))
	for i, elem := range files {
		activity := fillStruct(activitiesDir, elem, typ).(Activity)
		cv.Activities[i] = activity
	}
	// Read bio
	typ = reflect.TypeOf(Bio{})
	files, _ = ioutil.ReadDir(bioDir)
	cv.Bio = fillStruct(bioDir, files[0], typ).(Bio)
	// Read Education
	typ = reflect.TypeOf(Education{})
	files, _ = ioutil.ReadDir(educationDir)
	cv.Educations = make([]Education, len(files))
	for i, elem := range files {
		education := fillStruct(educationDir, elem, typ).(Education)
		cv.Educations[i] = education
	}
	// Read Experience
	typ = reflect.TypeOf(Experience{})
	files, _ = ioutil.ReadDir(experienceDir)
	cv.Experiences = make([]Experience, len(files))
	for i, elem := range files {
		experience := fillStruct(experienceDir, elem, typ).(Experience)
		cv.Experiences[i] = experience
	}

	fmt.Println(cv)
	return cv
}

func getName(source string) string {
	name := strings.Title(source)
	underscoreFinder, _ := regexp.Compile("_([a-z])")
	characters := underscoreFinder.FindAllStringSubmatch(name, -1)
	for _, character := range characters {
		name = strings.Replace(name, character[0], strings.ToUpper(character[1]), -1)
	}
	return name
}

func fillStruct(dir string, fileInfo os.FileInfo, typ reflect.Type) interface{} {
	data, err := ioutil.ReadFile(dir + "/" + fileInfo.Name())
	text := string(data)
	if err != nil {
		fmt.Println(err)
	}
	struc := reflect.New(typ).Elem()
	// Handle strings
	stringRegex, _ := regexp.Compile("@([a-z_]*)=\"([A-Za-z0-9 .:/äöüß+\\-]*)\"")
	submatches := stringRegex.FindAllStringSubmatch(text, -1)
	for _, submatch := range submatches {
		name := getName(submatch[1])
		field := struc.FieldByName(name)
		field.SetString(submatch[2])
	}
	// Handle JSON
	jsonRegex, _ := regexp.Compile("@([a-z_]*)=({[\\s\\w\":,#{}/\\[\\].]*})")
	byteSubmatches := jsonRegex.FindAllSubmatch(data, -1)
	for _, submatch := range byteSubmatches {
		name := getName(string(submatch[1]))
		fmt.Println(name)
		field := struc.FieldByName(name)
		switch name {
		case "Languages":
			var languages map[string]string
			json.Unmarshal(submatch[2], &languages)
			field.Set(reflect.ValueOf(languages))
		case "Skills":
			var skills map[string]map[string]string
			json.Unmarshal(submatch[2], &skills)
			field.Set(reflect.ValueOf(skills))
		case "Links":
			var links map[string][]string
			json.Unmarshal(submatch[2], &links)
			field.Set(reflect.ValueOf(links))
		case "AppliedTech":
			var tech []string
			json.Unmarshal(submatch[2], &tech)
			field.Set(reflect.ValueOf(tech))
		}
	}
	return struc.Interface()
}