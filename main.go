package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Tool struct
type Tool struct {
	Id         int
	FirstName  string
	LastName   string
	PetName    string
	PetSpecies string
	Bloodtype  string
	Phone      int
	PLZ        int
	Street     string
	Country    string
	Notes      string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	// dbServer := os.Getenv("MYSQL_SERVER")
	dbPort := "3306"
	// log.Println("Database host: " + dbServer)
	log.Println("Database port: " + dbPort)

	// This connectionstring is needed, if we wanna push this code to a container.
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+"172.19.0.2"+":"+dbPort+")/"+dbName)

	// This connectionstring is for local development purposes only.
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+"0.0.0.0"+":"+dbPort+")/"+dbName)

	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(:"+dbPort+")/"+dbName)

	// db, err := sql.Open("mysql", "docker:docker@tcp(:3306)/abdb")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM tools ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}
	res := []Tool{}

	for selDB.Next() {
		var id, phone, plz int
		var firstname, lastname, petname, petspecies, bloodtype, street, country, notes string
		err := selDB.Scan(&id, &firstname, &lastname, &petname, &petspecies, &bloodtype, &phone, &plz, &street, &country, &notes)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(rune(id)) + " | firstname " + firstname + " | lastname " + lastname + " | petname " + petname + " | petspecies " + petspecies + " | bloodtype " + bloodtype + " | phone " + string(rune(phone)) + " | plz " + string(rune(plz)) + " | street " + street + " | country " + country + " | notes " + notes)

		tool.Id = id
		tool.FirstName = firstname
		tool.LastName = lastname
		tool.PetName = petname
		tool.PetSpecies = petspecies
		tool.Bloodtype = bloodtype
		tool.Phone = phone
		tool.PLZ = plz
		tool.Street = street
		tool.Country = country
		tool.Notes = notes
		res = append(res, tool)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id, phone, plz int
		var firstname, lastname, petname, petspecies, bloodtype, street, country, notes string
		err := selDB.Scan(&id, &firstname, &lastname, &petname, &petspecies, &bloodtype, &phone, &plz, &street, &country, &notes)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Listing Row: Id " + string(rune(id)) + " | firstname " + firstname + " | lastname " + lastname + " | petname " + petname + " | petspecies " + petspecies + " | bloodtype " + bloodtype + " | phone " + string(rune(phone)) + " | plz " + string(rune(plz)) + " | street " + street + " | country " + country + " | notes " + notes)

		tool.Id = id
		tool.FirstName = firstname
		tool.LastName = lastname
		tool.PetName = petname
		tool.PetSpecies = petspecies
		tool.Bloodtype = bloodtype
		tool.Phone = phone
		tool.PLZ = plz
		tool.Street = street
		tool.Country = country
		tool.Notes = notes
	}
	tmpl.ExecuteTemplate(w, "Show", tool)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func EntriesList(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM tools ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}
	res := []Tool{}

	for selDB.Next() {
		var id, phone, plz int
		var firstname, lastname, petname, petspecies, bloodtype, street, country, notes string
		err := selDB.Scan(&id, &firstname, &lastname, &petname, &petspecies, &bloodtype, &phone, &plz, &street, &country, &notes)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing EntriesList Row: Id " + string(rune(id)) + " | firstname " + firstname + " | lastname " + lastname + " | petname " + petname + " | petspecies " + petspecies + " | bloodtype " + bloodtype + " | phone " + string(rune(phone)) + " | plz " + string(rune(plz)) + " | street " + street + " | country " + country + " | notes " + notes)

		tool.Id = id
		tool.FirstName = firstname
		tool.LastName = lastname
		tool.PetName = petname
		tool.PetSpecies = petspecies
		tool.Bloodtype = bloodtype
		tool.Phone = phone
		tool.PLZ = plz
		tool.Street = street
		tool.Country = country
		tool.Notes = notes
		res = append(res, tool)
	}
	tmpl.ExecuteTemplate(w, "EntriesList", res)
	defer db.Close()
}

func Cover(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Cover", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id, phone, plz int
		var firstname, lastname, petname, petspecies, bloodtype, street, country, notes string
		err := selDB.Scan(&id, &firstname, &lastname, &petname, &petspecies, &bloodtype, &phone, &plz, &street, &country, &notes)
		if err != nil {
			panic(err.Error())
		}

		tool.Id = id
		tool.FirstName = firstname
		tool.LastName = lastname
		tool.PetName = petname
		tool.PetSpecies = petspecies
		tool.Bloodtype = bloodtype
		tool.Phone = phone
		tool.PLZ = plz
		tool.Street = street
		tool.Country = country
		tool.Notes = notes
	}

	tmpl.ExecuteTemplate(w, "Edit", tool)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		petname := r.FormValue("petName")
		petspecies := r.FormValue("petSpecies")
		bloodtype := r.FormValue("bloodType")
		phone := r.FormValue("phone")
		plz := r.FormValue("plz")
		street := r.FormValue("street")
		country := r.FormValue("country")
		notes := r.FormValue("notes")
		insForm, err := db.Prepare("INSERT INTO tools (firstname, lastname, petName, petSpecies, bloodType, phone, plz, street, country, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(firstname, lastname, petname, petspecies, bloodtype, phone, plz, street, country, notes)
		log.Println("Insert Data: firstname " + firstname + " | lastname " + lastname + " | petname " + petname + " | petspecies " + petspecies + " | bloodtype " + bloodtype + " | phone " + phone + " | plz " + plz + " | street " + street + " | country " + country + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		petname := r.FormValue("petName")
		petspecies := r.FormValue("petSpecies")
		bloodtype := r.FormValue("bloodType")
		phone := r.FormValue("phone")
		plz := r.FormValue("plz")
		street := r.FormValue("street")
		country := r.FormValue("country")
		notes := r.FormValue("notes")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE tools SET firstname=?, lastname=?, petName=?, petSpecies=?, bloodType=?, phone=?, plz=?, street=?, country=?, notes=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(firstname, lastname, petname, petspecies, bloodtype, phone, plz, street, country, notes, id)
		log.Println("UPDATE Data: firstname " + firstname + " | lastname " + lastname + " | petname " + petname + " | petspecies " + petspecies + " | bloodtype " + bloodtype + " | phone " + phone + " | plz " + plz + " | street " + street + " | country " + country + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	tool := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM tools WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tool)
	log.Println("DELETE " + tool)
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/entrieslist", EntriesList)
	http.HandleFunc("/cover", Cover)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
