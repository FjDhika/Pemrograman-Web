package main

import(
	"database/sql"
	"log"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl=template.Must(template.ParseGlob("view/*"))

type Biodata struct{
	Nama string
	Npm string
	Phone string
	Email string
	Alamat string
	General string
	Org string
	Org_desc string
}

func checkErr(err error){
	if err != nil{
		panic(err.Error())
	}
}

func New(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "New",nil)
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "bio"
	db, err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	
	checkErr(err)

	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db:=dbConn()
	selDB,err := db.Query("SELECT NPM,NAMA,EMAIL,PHONE,ALAMAT,ORG FROM biodata")
	checkErr(err)

	bio:=Biodata{}
	res:=[]Biodata{}

	for selDB.Next() {
		var nama,npm,phone,email string
		var alamat, org,org_desc string

		err := selDB.Scan(&npm,&nama,&email,&phone,&alamat,&org)
		checkErr(err)
		bio.Npm = npm
		bio.Nama = nama
		bio.Email = email
		bio.Phone = phone
		bio.Alamat = alamat
		bio.Org = org
		bio.Org_desc = org_desc
		res = append(res, bio)
	}

	tmpl.ExecuteTemplate(w,"Index",res)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		
		Npm:= r.FormValue("Npm")
		Nama:= r.FormValue("Nama")
		Email:= r.FormValue("Email")
		Phone:= r.FormValue("Telp")
		Alamat:= r.FormValue("Alamat")
		Org:= r.FormValue("Org")
		Org_desc:= r.FormValue("Org-desc")
		General:= r.FormValue("General")

		InsData, err := db.Prepare("INSERT INTO biodata(NPM,NAMA,EMAIL,PHONE,ALAMAT,ORG,ORG_DESC,GENERAL) VALUES (?,?,?,?,?,?,?,?)")

		checkErr(err)

		log.Println("Inserting "+ Npm +" Biodata")
		InsData.Exec(Npm,Nama,Email,Phone,Alamat,Org,Org_desc,General)
		defer db.Close()
		
		http.Redirect(w,r,"/",301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db:=dbConn()
	Npm:= r.URL.Query().Get("Npm")
	
	DelData, err := db.Prepare("DELETE FROM biodata WHERE NPM=?")
	checkErr(err)

	log.Println("Delete "+ Npm +" From Biodata")
	DelData.Exec(Npm)
	defer db.Close()

	http.Redirect(w,r,"/",302)
}

func View(w http.ResponseWriter, r *http.Request) {
	db:=dbConn()
	Npm:= r.URL.Query().Get("Npm")
	viewData, err := db.Query("SELECT NPM,NAMA,EMAIL,PHONE,ALAMAT,ORG,ORG_DESC,GENERAL FROM biodata WHERE NPM=?",Npm)
	checkErr(err)

	bio:=Biodata{}
	for viewData.Next() {
		var nama,npm,phone,email string
		var alamat, general, org, org_desc string

		err:= viewData.Scan(&npm,&nama,&email,&phone,&alamat,&org,&org_desc,&general)
		checkErr(err)

		bio.Npm = npm
		bio.Nama = nama
		bio.Email = email
		bio.Phone = phone
		bio.Alamat = alamat
		bio.Org = org
		bio.Org_desc = org_desc
		bio.General = general
	}

	tmpl.ExecuteTemplate(w, "Show", bio)
	defer db.Close()
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db:=dbConn()
	Npm:= r.URL.Query().Get("Npm")
	viewData, err := db.Query("SELECT NPM,NAMA,EMAIL,PHONE,ALAMAT,ORG,ORG_DESC,GENERAL FROM biodata WHERE NPM=?",Npm)
	checkErr(err)

	bio:=Biodata{}
	for viewData.Next() {
		var nama,npm,phone,email string
		var alamat, general, org, org_desc string

		err:= viewData.Scan(&npm,&nama,&email,&phone,&alamat,&org,&org_desc,&general)
		checkErr(err)

		bio.Npm = npm
		bio.Nama = nama
		bio.Email = email
		bio.Phone = phone
		bio.Alamat = alamat
		bio.Org = org
		bio.Org_desc = org_desc
		bio.General = general
	}

	tmpl.ExecuteTemplate(w, "Edit", bio)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		
		Npm:= r.URL.Query().Get("Npm")
		Nama:= r.FormValue("Nama")
		Email:= r.FormValue("Email")
		Phone:= r.FormValue("Telp")
		Alamat:= r.FormValue("Alamat")
		Org:= r.FormValue("Org")
		Org_desc:= r.FormValue("Org-desc")
		General:= r.FormValue("General")

		InsData, err := db.Prepare("UPDATE biodata SET NAMA=?,EMAIL=?,PHONE=?,ALAMAT=?,ORG=?,ORG_DESC=?,GENERAL=? WHERE NPM=?")

		checkErr(err)

		log.Println("Update "+ Npm +" Biodata")
		InsData.Exec(Nama,Email,Phone,Alamat,Org,Org_desc,General,Npm)
		defer db.Close()
		
		http.Redirect(w,r,"/",303)
	}
}

func main() {
	log.Println("Server started on localhost:8080 ")
	
	http.HandleFunc("/",Index)
	http.HandleFunc("/New",New)
	http.HandleFunc("/Insert",Insert)

	http.HandleFunc("/Show",View)

	http.HandleFunc("/Edit",Edit)
	http.HandleFunc("/Update",Update)
	
	http.HandleFunc("/Delete",Delete)
	
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	
	http.ListenAndServe(":8080",nil)
}