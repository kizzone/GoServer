package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type levelPosition struct {
	XCord int
	YCord int
	ZCord int
}

func handler(w http.ResponseWriter, r *http.Request) {

	levelPs := levelPosition{XCord: rand.Intn(100), YCord: rand.Intn(100), ZCord: rand.Intn(100)}
	/*fmt.Println("New X position: ", levelPs.XCord)
	fmt.Println("New Y position: ", levelPs.YCord)
	fmt.Println("New Z position: ", levelPs.ZCord)
	*/
	jsonMarsh, err := json.Marshal(&levelPs)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "String rappresentation of Json: %s!", string(jsonMarsh))
}

func main() {
	http.HandleFunc("/accelerometer", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
