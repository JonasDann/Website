package main

type CV struct {
	activities []Activity
	bio Bio
	educations []Education
	experiences []Experience
}

type Activity struct {
	time        string
	title       string
	description string
}

type Bio struct {
	firstName string
	secondName string
	lastName string
	birthDate string
	birthLocation string
	street string
	houseNumber string
	postalCode string
	city string
	country string
	phone string
	email string
	languages map[string]string
	skills map[string]map[string]string
	links map[string][]string
}

type Education struct {
	from string
	until string
	expectedGraduation string
	name string
	location string
	degree string
	grade string
	description string
}

type Experience struct {
	from string
	until string
	name string
	location string
	position string
	department string
	description string
	appliedTech []string
}