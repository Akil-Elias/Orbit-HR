package handlers

import (
	"net/http"
	"strconv"
	"time"
)

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	maritalStatus := r.FormValue("marital_status")
	parent := r.FormValue("parent")
	street := r.FormValue("street")
	city := r.FormValue("city")
	position := r.FormValue("position")
	department := r.FormValue("department")
	division := r.FormValue("division")
	location := r.FormValue("location")
	employmentStatus := r.FormValue("employment")
	sick := r.FormValue("sick_leave_bal")
	casual := r.FormValue("casual_leave_bal")
	vacation := r.FormValue("vacation_leave_bal")
	rop := r.FormValue("rate_of_pay")
	salary := r.FormValue("salary")

	dobStr := r.FormValue("dob")
	dob, err := time.Parse("2006-01-02", dobStr) // Go's date layout for yyyy-mm-dd
	if err != nil {
		http.Error(w, "Invalid date format for date of birth", http.StatusBadRequest)
		return
	}

	hdStr := r.FormValue("hire_date")
	hireDate, err := time.Parse("2006-01-02", hdStr) // Go's date layout for yyyy-mm-dd
	if err != nil {
		http.Error(w, "Invalid date format for hire date", http.StatusBadRequest)
		return
	}

	sickLeave, err := strconv.Atoi(sick)
	if err != nil {
		http.Error(w, "Invalid sick leave balance", http.StatusBadRequest)
		return
	}

	casualLeave, err := strconv.Atoi(casual)
	if err != nil {
		http.Error(w, "Invalid sick leave balance", http.StatusBadRequest)
		return
	}

	vacationLeave, err := strconv.Atoi(vacation)
	if err != nil {
		http.Error(w, "Invalid sick leave balance", http.StatusBadRequest)
		return
	}

	rateOfPay, err := strconv.ParseFloat(rop, 64)
	if err != nil {
		http.Error(w, "Invalid rate of pay", http.StatusBadRequest)
		return
	}

}
