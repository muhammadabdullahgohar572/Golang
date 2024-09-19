package main

import (
	"encoding/json"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json");
	var empp Employee;
	json.NewDecoder(r.Body).Decode(empp);
	json.NewEncoder(w).Encode(empp);
	//db.Create(&empp);
	Database.Create(&empp)
	//fmt.Fprintf(w, "New Employee: %+v", empp)
	json.NewEncoder(w).Encode(&empp)


}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	
}

func GetEmployeeID(w http.ResponseWriter, r *http.Request) {
	
}