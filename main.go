package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Successfully Deleted User")
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
func newUser(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func initialMigration() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	//defer db.Close()
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")
	//initialMigration()
	handleRequests()
}
