package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/webvillain/vikashbank14/controllers"
	"github.com/webvillain/vikashbank14/routes"
	"github.com/webvillain/vikashbank14/touchup.go"
)

func main() {
	touchup.WelcomeIntro()
	controllers.NewDatabase()
	touchup.IntroAfterDatabaseConnection()
	r := mux.NewRouter()
	r.HandleFunc("/bank", routes.ListUsers).Methods("GET")
	r.HandleFunc("/bank/{Id}", routes.UserById).Methods("GET")
	r.HandleFunc("/bank/{Name}/{Email}/{UserName}", routes.CreateUser).Methods("POST")
	r.HandleFunc("/bank/{Id}", routes.DeleteUser).Methods("DELETE")
	r.HandleFunc("/bank/{Id}/{Name}/{Email}/{UserName}", routes.UpdateUser)
	log.Fatal(http.ListenAndServe(":8000", r))

}
