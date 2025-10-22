package services

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	//personal info
	ID            uuid.UUID
	Firstname     string
	Lastname      string
	DOB           time.Time
	MaritalStatus string //Married, Single, Divorced, Widowed
	Parent        string //Yes, No
	street        string
	city          string
	//company info
	Position         string //from list of positions
	Department       string //from list
	Division         string //from list
	Location         string //from list
	HireDate         time.Time
	EmploymentStatus string //permanent, contract
	Sick_Bal         int64
	Casual_Bal       int64
	Vacation_Bal     int64
	RateOfPay        string //weekly, monthly
	Salary           int64
}

func NewEmployee(
	firstname string,
	lastname string,
	dob time.Time,
	maritalStatus string,
	parent string,
	street string,
	city string,
	position string,
	department string,
	division string,
	location string,
	hireDate time.Time,
	employmentStatus string,
	tenure int64,
) Employee {
	return Employee{
		ID:               uuid.New(),
		Firstname:        firstname,
		Lastname:         lastname,
		DOB:              dob,
		MaritalStatus:    maritalStatus,
		Parent:           parent,
		street:           street,
		city:             city,
		Position:         position,
		Department:       department,
		Division:         division,
		Location:         location,
		HireDate:         hireDate,
		EmploymentStatus: employmentStatus,
		Tenure:           tenure,
	}
}
