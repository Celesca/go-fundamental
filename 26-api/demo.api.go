package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

const coursePath = "courses"
const basePath = "/api"

type Course struct {
	CourseID   int     `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"imageurl"`
}

func SetupDB() {
	var err error
	Db, err = sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/coursedb")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Db)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

// GET : Get Course List
func getCourseList() ([]Course, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second) // set time to use this function
	defer cancle()

	results, err := Db.QueryContext(ctx,
		`SELECT * FROM courseonline`)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	courses := make([]Course, 0)
	for results.Next() {
		var course Course
		results.Scan(&course.CourseID, &course.CourseName, &course.Price, &course.ImageURL)

		courses = append(courses, course)
	}

	return courses, nil
}

// POST : Insert Product
func insertProduct(course Course) (int, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	result, err := Db.ExecContext(ctx, `INSERT INTO courseonline 
	(courseid, coursename, price, imageurl) VALUES (?, ?, ?, ?)`,
		course.CourseID, course.CourseName, course.Price, course.ImageURL)

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return int(insertID), nil

}

// GET : Get Course (only one)
func getCourse(courseid int) (*Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := Db.QueryRowContext(ctx, `SELECT * FROM courseonline WHERE courseid = ?`, courseid)

	course := &Course{}
	err := row.Scan(&course.CourseID, &course.CourseName, &course.Price, &course.ImageURL)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return course, nil
}

// DELETE : Remove Course
func removeCourse(courseid int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := Db.ExecContext(ctx, `DELETE FROM courseonline WHERE courseid = ?`, courseid)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Handle /course/
func handleCourse(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", coursePath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		course, err := getCourse(courseID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if course == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, err := json.Marshal(course)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodDelete:
		err := removeCourse(courseID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// Handle /course - Get all courses
func handleCourses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		courseList, err := getCourseList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(courseList)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

		// Post
	case http.MethodPost:
		var course Course
		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		CourseID, err := insertProduct(course)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id":%d}`, CourseID)))

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// cors middleware
func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, X-Requested-With")
		handler.ServeHTTP(w, r)
	})

}

// Set เส้นทางของ API
func SetupRoutes(apibasePath string) {

	// handle Course by ID
	courseHandler := http.HandlerFunc(handleCourse)
	http.Handle(fmt.Sprintf("%s/%s/", apibasePath, coursePath), corsMiddleware(courseHandler))

	// handle Courses
	coursesHandler := http.HandlerFunc(handleCourses)
	http.Handle(fmt.Sprintf("%s/%s", apibasePath, coursePath), corsMiddleware(coursesHandler))
}

func main() {

	SetupDB()
	SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
