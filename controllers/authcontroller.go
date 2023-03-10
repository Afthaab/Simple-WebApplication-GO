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
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else if session.Values["adminIn"] == true {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			temp, err := template.ParseFiles("views/index.html")
			CheckError(err)
			temp.Execute(w, nil)
		}
	}

}
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		session, _ := config.Store.Get(r, config.SESSION_ID)
		if session.Values["loggedIn"] == true {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else if session.Values["adminIn"] == true {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			temp, err := template.ParseFiles("views/login.html")
			CheckError(err)
			temp.Execute(w, nil)
		}

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		Userinput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		if Userinput.Username == "Admin" && Userinput.Password == "12345" {
			session, _ := config.Store.Get(r, config.SESSION_ID)
			var temp_user entities.Users
			session.Values["adminIn"] = true
			session.Values["username"] = temp_user.Username
			session.Save(r, w)
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			DB := config.DBConn()
			var temp_user entities.Users
			result := DB.First(&temp_user, "username LIKE ? AND password LIKE ?", Userinput.Username, Userinput.Password)
			if result.Error != nil {
				http.Redirect(w, r, "/login", http.StatusFound)

			} else {
				session, _ := config.Store.Get(r, config.SESSION_ID)
				session.Values["loggedIn"] = true
				session.Values["username"] = temp_user.Username
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusFound)
			}
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
		result := DB.First(&temp_user, "username LIKE ? ", Registerinput.Username)
		if result.Error != nil {
			DB.Create(&entities.Users{First_Name: Registerinput.fname, Last_name: Registerinput.lname, Username: Registerinput.Username, Password: Registerinput.Password})
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			http.Redirect(w, r, "/register", http.StatusFound)
		}

	}

}

func Admin(w http.ResponseWriter, r *http.Request) {
	var temp_user []entities.Users
	DB := config.DBConn()
	if r.Method == http.MethodGet {
		result := DB.Find(&temp_user)
		if result.Error != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			session, _ := config.Store.Get(r, config.SESSION_ID)
			if session.Values["adminIn"] == true {
				temp, err := template.ParseFiles("views/admin.html")
				CheckError(err)
				temp.Execute(w, temp_user)
			} else if session.Values["loggedIn"] == true {
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}

		}

	} else if r.Method == http.MethodPost {
		r.ParseForm()
		uid := r.PostForm["id"][0]
		DB.Delete(&entities.Users{}, uid)
		http.Redirect(w, r, "/admin", http.StatusFound)
	}

}
func Test(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	untyped, ok := session.Values["username"]
	if !ok {
		return
	}
	username, ok := untyped.(string)
	if !ok {
		return
	}
	w.Write([]byte(username))

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
