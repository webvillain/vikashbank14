package controllers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/webvillain/vikashbank14/model"
)

var DB *sql.DB

func NewDatabase() {
	db, err := sql.Open("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS USERS (Id INTEGER PRIMARY KEY AUTOINCREMENT , Name TEXT NOT NULL, Email TEXT , UserName TEXT);")
	if err != nil {
		fmt.Println(err)
	}
	stmt.Exec()
	DB = db
}

// list all users
func AllUsers() ([]model.User, error) {

	var users []model.User
	rows, err := DB.Query("SELECT * FROM USERS")
	if err != nil {
		fmt.Println(err)
	}
	var Id int64
	var Name string
	var Email string
	var UserName string
	for rows.Next() {

		rows.Scan(&Id, &Name, &Email, &UserName)
		users = append(users, model.User{Id: Id, Name: Name, Email: Email, UserName: UserName})
	}

	// /defer rows.Close()
	//defer DB.Close()
	return users, nil
}

//user by id
func UserById(Id int64) (model.User, error) {
	var user model.User
	// var Name string
	// var Email string
	// var UserName string
	listusers := `SELECT * FROM USERS WHERE Id = ?`
	rows, err := DB.Query(listusers, Id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {

		rows.Scan(&user.Id, &user.Name, &user.Email, &user.UserName)
	}
	// user.Id = Id
	// user.Name = Name
	// user.Email = Email
	// user.UserName = UserName
	//defer DB.Close()
	//defer rows.Close()

	return user, nil
}

// create new user
func CreateUser(Name string, Email string, UserName string) error {

	stmt, err := DB.Prepare("INSERT INTO USERS(Name , Email , UserName) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(Name, Email, UserName)
	if err != nil {
		fmt.Println(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("Rows Affected : ", n)
	//defer DB.Close()
	return nil
}

//update a user
func UpdateUser(Id int64, Name string, Email string, UserName string) error {

	stmt, err := DB.Prepare("UPDATE USERS SET Name = ?, Email = ?, UserName = ? WHERE Id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(Name, Email, UserName, Id)
	if err != nil {
		fmt.Println(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("Rows Affected : ", n)
	//defer DB.Close()
	return nil
}

//delete a user
func DeleteUser(Id int64) error {

	stmt, err := DB.Prepare("DELETE FROM USERS WHERE Id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(Id)
	if err != nil {
		panic(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("Rows Affected : ", n)
	//defer DB.Close()
	return nil
}
