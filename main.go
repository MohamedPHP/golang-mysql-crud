package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User Struct
type User struct {
	ID            int
	Name          string
	Email         string
	Password      string
	RememberToken string
	CreatedAt     string
	UpdatedAt     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

// Index users
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	users, err := db.Query("SELECT * FROM users ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}

	user := User{}

	all := []User{}
	for users.Next() {
		var id int
		var name string
		var email string
		var password string
		var rememberToken string
		var createdAt string
		var updatedAt string

		err = users.Scan(&id, &name, &email, &password, &rememberToken, &createdAt, &updatedAt)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.RememberToken = rememberToken
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt

		all = append(all, user)
	}

	tmpl.ExecuteTemplate(w, "Index", all)

	defer db.Close()
}

// Show user
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nID := r.URL.Query().Get("id")

	users, err := db.Query("SELECT * FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	user := User{}

	for users.Next() {
		var id int
		var name string
		var email string
		var password string
		var rememberToken string
		var createdAt string
		var updatedAt string

		err = users.Scan(&id, &name, &email, &password, &rememberToken, &createdAt, &updatedAt)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.RememberToken = rememberToken
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}
	tmpl.ExecuteTemplate(w, "Show", user)
	defer db.Close()
}

// New User
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit User
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	nID := r.URL.Query().Get("id")

	users, err := db.Query("SELECT * FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	user := User{}

	for users.Next() {
		var id int
		var name string
		var email string
		var password string
		var rememberToken string
		var createdAt string
		var updatedAt string

		err = users.Scan(&id, &name, &email, &password, &rememberToken, &createdAt, &updatedAt)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.RememberToken = rememberToken
		user.CreatedAt = createdAt
		user.UpdatedAt = updatedAt
	}

	tmpl.ExecuteTemplate(w, "Edit", user)
	defer db.Close()
}

// Insert User
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		// check if email unique
		emailUnique, err := db.Query("SELECT * FROM users WHERE email=?", r.FormValue("email"))

		if err != nil {
			panic(err.Error())
		}

		if !emailUnique.Next() {
			// create new user
			insForm, err := db.Prepare("INSERT INTO users(name, email, password, createdAt, updatedAt) VALUES(?,?,?,?,?)")

			if err != nil {
				panic(err.Error())
			}

			currentTime := time.Now()

			insForm.Exec(r.FormValue("name"), r.FormValue("email"), r.FormValue("password"), currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

			log.Println("INSERT: Name: " + r.FormValue("name"))
		}
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

// Update User
func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		insForm, err := db.Prepare("UPDATE Employee SET name=?, email=?, password=?, updatedAt=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		currentTime := time.Now()

		insForm.Exec(r.FormValue("name"), r.FormValue("email"), r.FormValue("password"), currentTime.Format("2006-01-02 15:04:05"), r.FormValue("uid"))

		log.Println("Update Record #" + r.FormValue("uid") + " : Name: " + r.FormValue("name"))
	}

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

// Delete User
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	ID := r.URL.Query().Get("id")

	delForm, err := db.Prepare("DELETE FROM users WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(ID)

	log.Println("DELETE")

	defer db.Close()

	http.Redirect(w, r, "/", 301)
}

func main() {
	fmt.Println("HTTP server started at http://127.0.0.1:8000")

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":8000", nil)
}
