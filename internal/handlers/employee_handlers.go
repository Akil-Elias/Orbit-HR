package handlers

import (
	"net/http"
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

	dobStr := r.FormValue("dob")
	dob, err := time.Parse("2006-01-02", dobStr) // Go's date layout for yyyy-mm-dd
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	hdStr := r.FormValue("hdStr")

	position := r.FormValue("position")
	department := r.FormValue("department")
	division := r.FormValue("division")
	location := r.FormValue("location")
}
