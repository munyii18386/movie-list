package handlers

import(
    // "fmt"
	// "log"
	// "os"
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
)

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}


// SignUp handles 
// func (p *ReverseProxy) SignUp (w http.ResponseWriter, r *http.Request) {

// 	f, err := os.Create("/tmp/input")
// 	if err != nil {
//         panic(err)
// 	}
// 	defer f.Close()

// 	if r.Method == http.MethodPost{
// 		var u NewUser
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Println(string(body)) 

// 		err = json.Unmarshal(body, &u)

// 		if err != nil {
// 			panic(err)
// 		}

// 		f.WriteString(u.Email)
		
// 	}

	

// }