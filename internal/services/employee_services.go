package services

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Street string
	City   string
}

type Employee struct {
	//personal info
	ID            uuid.UUID
	Firstname     string
	Lastname      string
	DOB           time.Time
	Age           int64
	MaritalStatus string //Married, Single, Divorced, Widowed
	Parent        string //Yes, No
	Address       Address
	//company info
	JobTitle         Position   //from list of positions
	Department       Department //from list
	Division         Division   //from list
	Location         string     //from list
	HireDate         time.Time
	EmploymentStatus string //permanent, contract
	YearsOfService   int64
	Sick_Bal         int64
	Casual_Bal       int64
	Vacation_Bal     int64
	RateOfPay        string //weekly, monthly
	Salary           int64
}
