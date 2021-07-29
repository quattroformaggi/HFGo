package main

/*Exercise 1: A Simple Web App

Here’s the start of a web app that simulates rolling a six-sided die.
When a browser makes a request for the /roll path, the app will respond with a random number from 1 to 6.
Each refresh of the page will generate a new random number.

Within rollHandler:
    Call rollDie.
    Convert the return value to a string. You may want to consult the documentation for the "strconv" package’s Itoa function.
    Convert the string to a slice of bytes, and store it in a variable named body.
Within main, set up the rollHandler function to handle all requests with a path of "/roll".*/

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// rollDie simulates rolling a six-sided die.
func rollDie() int {
	return rand.Intn(6) + 1
}

func rollHandler(writer http.ResponseWriter, request *http.Request) {
	// YOUR CODE HERE: Call rollDie, convert the return
	// value to a string, and convert the string to a
	// slice of bytes. Store the result in a "body"
	// variable.
	diestring := strconv.Itoa(rollDie())
	body := []byte(diestring)

	_, err := writer.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	// YOUR CODE HERE: Have all requests for a URL with a
	// path of "/roll" go to the rollHandler function.
	http.HandleFunc("/roll", rollHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
