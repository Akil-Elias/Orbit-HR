package services

import "github.com/google/uuid"

type Position struct {
	ID          uuid.UUID
	Name        string
	Department  Department
	LeavePolicy string
}

type Department struct {
	ID        uuid.UUIDs
	Name      string
	Division  Division
	Headcount int64
}

type Division struct {
	ID          uuid.UUID
	Name        string
	Description string
	Headcount   int64
}

type Company struct {
	ID        uuid.UUID
	Name      string
	Industry  string
	Headcount int64
	Division  []Division
}
