package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type Student struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	NoteA int       `json:"noteA"`
	NoteB int       `json:"noteB"`
}

type allStudents []Student

var students = allStudents{}

// @title           Swagger API-Students
// @version         1.0
// @description     Documentação da API de alunos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Students Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://localhost:8080/students
func main() {
	log.Println("iniciando API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/student", CreateNewStudent).Methods("POST")

	http.ListenAndServe(":8080", router)
	//http.ListenAndServe(":"+port, router)

}

func Home(w http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(w, "Seja bem vindo!\n")
}

func HealthCheck(w http.ResponseWriter, rq *http.Request) {
	log.Println("Acessando o Health-Check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

// ShowAllStudents godoc
// @Summary      Show all students
// @Description  List all students
// @Tags         students
// @Accept       json
// @Produce      json
// @Success      200  {object}  Student
// @Router       /students [get]
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint get all students")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

// CreateNewStudent godoc
// @Summary      New student
// @Description  Creater new student
// @Tags         student
// @Accept       json
// @Produce      json
// @Succes      201  {object}  Student
// @Router       /student [post]
func CreateNewStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("acessando o endpoint create new event")
	var newStudent Student

	newStudent.Id, _ = uuid.NewV4()

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newStudent)

	students = append(students, newStudent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}
