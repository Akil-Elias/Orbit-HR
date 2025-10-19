package services

import "github.com/google/uuid"

type LeavePolicy struct {
	ID          int64
	Name        string
	Sick        int64
	Casual      int64
	Vacation    int64
	Bereavement int64
	Study       int64
	Maternity   int64
	Paternity   int64
}

type TimeOffRequest struct {
	ID            int64
	Policy        LeavePolicy
	EmployeeID    uuid.UUID
	Firstname     string
	LastName      string
	LeaveType     string
	DaysRequested int64
	Reason        string
}
