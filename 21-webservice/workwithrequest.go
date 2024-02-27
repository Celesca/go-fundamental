package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Course struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{"id": 1, "name": "Golang", "price": 150, "instructor": "Ardan Labs"},
		{"id": 2, "name": "Python", "price": 150, "instructor": "Udemy"},
		{"id": 3, "name": "Java", "price": 150, "instructor": "Coursera"}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList) // Convert JSON to Go Object (Interface)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}

	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)

	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Convert JSON to Go Object to err
		err = json.Unmarshal(Bodybyte, &newCourse)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// check if user send ID in the request (don't send the ID)
		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()

		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)

}
