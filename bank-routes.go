package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/webvillain/vikashbank14/controllers"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank You For Connecting To V-Banking Application .")
	fmt.Println("We Are Happy To Help.")
	fmt.Println("You Have Selected Option - 1")
	users, err := controllers.AllUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("__________________________________________________")

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\t HTTP Status Code :", http.StatusOK)
	fmt.Println("\t Request Method   :", r.Method)
	fmt.Println("\t Request URL      :", r.URL)

	fmt.Println("__________________________________________________")

	json.NewEncoder(w).Encode(users)

}

func UserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank You For Connecting To V-Banking Application .")
	fmt.Println("We Are Happy To Help.")
	fmt.Println("You Have Selected Option - 2")
	params := mux.Vars(r)
	id := params["Id"]
	newID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error While Converting String Into Int64")
	}

	user, err := controllers.UserById(newID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("__________________________________________________")

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\t HTTP Status Code :", http.StatusOK)
	fmt.Println("\t Request Method   :", r.Method)
	fmt.Println("\t Request URL      :", r.URL)

	fmt.Println("__________________________________________________")

	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank You For Connecting To V-Banking Application .")
	fmt.Println("We Are Happy To Help.")
	fmt.Println("You Have Selected Option - 3")
	params := mux.Vars(r)
	name := params["Name"]
	email := params["Email"]
	username := params["UserName"]
	err := controllers.CreateUser(name, email, username)
	if err != nil {
		fmt.Println("We Got A Problem While Creating A New User Into Database")
		log.Fatal(err)
	}
	fmt.Fprintln(w, "New User Is Created Successfully.")
	fmt.Println("New User Is Created Successfully.")
	fmt.Println("__________________________________________________")

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\t HTTP Status Code :", http.StatusOK)
	fmt.Println("\t Request Method   :", r.Method)
	fmt.Println("\t Request URL      :", r.URL)

	fmt.Println("__________________________________________________")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank You For Connecting To V-Banking Application .")
	fmt.Println("We Are Happy To Help.")
	fmt.Println("You Have Selected Option - 4")
	params := mux.Vars(r)
	id := params["Id"]
	newID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error While Converting String Into Int64")
	}
	err = controllers.DeleteUser(newID)
	if err != nil {
		fmt.Println("We Got A Problem While Deleting User From Database")
		log.Fatal(err)
	}
	fmt.Fprintf(w, "User With Id :%d Is Deleted Successfully .\n", newID)
	fmt.Printf("User With Id :%d Is Deleted Successfully .\n", newID)
	fmt.Println("__________________________________________________")

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\t HTTP Status Code :", http.StatusOK)
	fmt.Println("\t Request Method   :", r.Method)
	fmt.Println("\t Request URL      :", r.URL)

	fmt.Println("__________________________________________________")

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Thank You For Connecting To V-Banking Application .")
	fmt.Println("We Are Happy To Help.")
	fmt.Println("You Have Selected Option - 5")
	params := mux.Vars(r)
	id := params["Id"]
	newID, err := strconv.ParseInt(id, 0, 0)
	name := params["Name"]
	email := params["Email"]
	username := params["UserName"]
	if err != nil {
		fmt.Println("Error While Converting String Into Int64")
	}
	err = controllers.UpdateUser(newID, name, email, username)
	if err != nil {
		fmt.Println("We Got A Problem While Updating User From Database")
		log.Fatal(err)
	}
	fmt.Fprintf(w, "User With Id :%d Is Updated Successfully .\n", newID)
	fmt.Printf("User With Id :%d Is Updated Successfully .\n", newID)
	fmt.Println("__________________________________________________")

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("\t HTTP Status Code :", http.StatusOK)
	fmt.Println("\t Request Method   :", r.Method)
	fmt.Println("\t Request URL      :", r.URL)

	fmt.Println("__________________________________________________")

}
