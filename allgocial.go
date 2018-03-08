package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey("llh63iPofChnGnFnQxSteZozH")
	anaconda.SetConsumerSecret("jcH7IJ0nHzCa5dL1zxIP9mCs09JxHXGi0xhfOO98OhiqsHRVcK")
	api := anaconda.NewTwitterApi("129853182-a9yexVbhKwy0EFo5wVluv0A6qZ46XXbkE1uBhdDt", "dE0d8N9ZyqGRL3QJj3fUTYHoDXfN284rfYouSnVBX6qxv")

	// Post Tweet
	fmt.Print("Ingresa el tweet a Publicar: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()
	_, err := api.PostTweet(message, nil)
	if err != nil {
		fmt.Println("No se pudo publicar el tweet", err)
	} else {
		fmt.Println("Tweet publicado exitosamente")
	}

}
