package main

import (
	"fmt"
	"net/http"

	"github.com/loginpage/controllers"
)

func main() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/admin", controllers.Admin)
	http.HandleFunc("/test", controllers.Test)
	http.HandleFunc("/logout", controllers.Logout)

	fmt.Println("Server starting at 4000")

	http.ListenAndServe(":4000", nil)

}
