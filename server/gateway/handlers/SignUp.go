package handlers

import(
    // "fmt"
	// "log"
	// "os"
	// "encoding/json"
	// "io/ioutil"
	"net/http"
)

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}


// SignupHandler handles 
func  SignupHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sign up!"))
	// if r.Method == http.MethodPost{
	// 	var u NewUser
	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(string(body)) 

	// 	err = json.Unmarshal(body, &u)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(u.Email)
	// 	w.Write([]byte("sign up!"))
		
	// }

}