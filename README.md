# 26 - API Service by Go (API Endpoints)

### Initialize Project
- Run demo.api.go
- the server will start at localhost:5500
- use Postman to test the APIs.

* `GET: /courses`

Get all courses list

* `POST: /courses`

Request Body :
{
    "CourseID": 2,
    "Coursename" : "Complete Arduino For Beginners" ,
    "Price": 19000,
    "Image_URL" : "Hello World"
}

* `GET: /courses/{id}`

Get course by ID

* `DELETE: /courses/{id}`

Delete course by ID


## 21-webservice

* `GET: /course`

to get all courses in the CourseList (Slice) and Course (struct)

*  `POST: /course`

body : { "Name": "Python", "Price": 350, "Instructor": "Adan Wong" }
