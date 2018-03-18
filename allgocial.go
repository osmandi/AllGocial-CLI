package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

// Check Error
func check(e error) {
	if e != nil {
		fmt.Println("Variables de OAuth no seteadas, genera tus propios keys en apps.twitter.com :P")
		panic(e)
	}
}

// Read variables key
func readVariables() []string {
	fileText, err := ioutil.ReadFile("keys.txt")
	check(err)

	return strings.Split(string(fileText), ",")
}

func main() {

	// Get variables keys
	keys := readVariables()
	anaconda.SetConsumerKey(keys[0])
	anaconda.SetConsumerSecret(keys[1])
	api := anaconda.NewTwitterApi(keys[2], keys[3])

	// Flag insert
	var message *string = flag.String(
		"m",
		"",
		"Message to tweet")

	flag.Parse()

	// Post Tweet
	_, err := api.PostTweet(*message, nil)
	if err != nil {
		fmt.Println("No se pudo publicar el tweet", err)
	} else {
		fmt.Println("Tweet publicado exitosamente")
	}

}
