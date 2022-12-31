package controllers

import (
	"html/template"
	"net/http"

	"github.com/loginpage/config"
	"github.com/loginpage/entities"
)

type UserInput struct {
	Username string
	Password string
	fname    string
	lname    string
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")
	CheckError(err)
	temp.Execute(w, nil)

}
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/login.html")
		CheckError(err)
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		Userinput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		DB := config.DBConn()
		var temp_user entities.Users
		result := DB.First(&temp_user, "username LIKE ? AND password LIKE ?", Userinput.Username, Userinput.Password)
		if result.Error != nil {
			http.Redirect(w, r, "/login", http.StatusFound)

		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}

	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/register.html")
		CheckError(err)
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		Registerinput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
			fname:    r.Form.Get("firstname"),
			lname:    r.Form.Get("lastname"),
		}
		DB := config.DBConn()
		var temp_user entities.Users
		result := DB.First(&temp_user, "username LIKE ?", Registerinput.Username)
		if result.Error != nil {
			DB.Create(&entities.Users{First_Name: Registerinput.fname, Last_name: Registerinput.lname, Username: Registerinput.Username, Password: Registerinput.Password})
			http.Redirect(w, r, "/login", http.StatusFound)

		} else {
			http.Redirect(w, r, "/register", http.StatusFound)
		}

	}

}

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/admin.html")
		CheckError(err)
		temp.Execute(w, nil)
	}

}
