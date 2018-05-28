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

	//creating an array of levelPosition

	var levelPositionArray = []levelPosition{
		levelPosition{XCord: 0, YCord: 0, ZCord: 0},  //rappresenting the quiet state
		levelPosition{XCord: 0, YCord: 20, ZCord: 0}, //rappresenting the pulled state
		levelPosition{XCord: rand.Intn(100), YCord: rand.Intn(100), ZCord: rand.Intn(100)},
	}

	//levelPs := levelPosition{XCord: rand.Intn(100), YCord: rand.Intn(100), ZCord: rrand.Intn(100)}
	/*fmt.Println("New X position: ", levelPs.XCord)
	fmt.Println("New Y position: ", levelPs.YCord)
	fmt.Println("New Z position: ", levelPs.ZCord)
	*/

	//return from 0 to n
	i := rand.Intn(3)

	jsonMarsh, err := json.Marshal(&levelPositionArray[i])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", string(jsonMarsh))
}

func main() {
	http.HandleFunc("/accelerometer", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
