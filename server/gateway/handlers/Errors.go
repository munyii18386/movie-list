package handlers

import(
	"fmt"
)

// HandleError is used to log errors
func HandleError (err error){
	if err != nil{
		fmt.Printf("Error Generated is: %v\n", err)
	}
}


// ExitTransaction is used for error that need to abort immediately
func ExitTransaction (err error){
	if err != nil {
		panic(err)
	}
}