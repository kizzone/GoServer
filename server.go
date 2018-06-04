package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

//rappresent the acceleration of the lever

type levelPosition struct {
	XCord int
	YCord int
	ZCord int
}

//handler for http request
func handler(w http.ResponseWriter, r *http.Request) {

	currentlevelPs := levelPosition{XCord: 0, YCord: 0, ZCord: 0}

	i := rand.Intn(100) + 1

	//nel 10 percento dei casi
	if i <= 35 && i >= 0 {
		fmt.Println("Genero una nuova accelerazione per la leva perchè il valore di è ")
		fmt.Println(i)
		//genero una nuova accelerazione
		currentlevelPs.XCord = rand.Intn(100)
		fmt.Println("L'accelerazione generata è ")
		fmt.Println(currentlevelPs.XCord)
	}

	//2 random uno che decide se aggiornare la posizione e uno aggiorna

	switch r.Method {

	case "GET":

		jsonMarsh, err := json.Marshal(currentlevelPs)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", string(jsonMarsh))

	case "POST":
		fmt.Fprintf(w, "Not implemented yet")

	default:
		fmt.Fprintf(w, "Sorry only POST and GET method allow")
	}

	//creating an array of levelPosition

}

func main() {

	http.HandleFunc("/accelerometer", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
