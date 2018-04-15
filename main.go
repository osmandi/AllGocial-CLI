package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

// TODO: Agregar funcionalidad de authenticación por OAuth a AllGocial usando url y abriéndose automáticamente.
// TODO: Realizado la autorización, publicar el codigo y dejarlo como portafolio.

// Check Error
func check(e error) {
	if e != nil {
		fmt.Println("Variables de OAuth no seteadas, genera tus propios keys en apps.twitter.com :P")
		panic(e)
	}
}

// Read variables key
func readVariables() []string {
	/*
		fileText, err := ioutil.ReadFile("keys.txt")
		check(err)

		return strings.Split(string(fileText), ",")
	*/

	variables := []string{os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRET"), os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSTOKENSECRET")}
	return variables
}

func main() {
	// Get variables keys
	keys := readVariables()
	anaconda.SetConsumerKey(keys[0])
	anaconda.SetConsumerSecret(keys[1])
	api := anaconda.NewTwitterApi(keys[2], keys[3])

	app := cli.NewApp()

	// Define command your client
	app.Commands = []cli.Command{
		{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "Input message. [Usage]: allgocial message 'I am using allgocial :D'",
			Action: func(c *cli.Context) error {
				if c.NArg() == 1 {
					// Logic
					// Post Tweet
					_, err := api.PostTweet(c.Args()[0], nil)
					if err != nil {
						fmt.Println("No se pudo publicar el tweet", err)
					} else {
						fmt.Println("Tweet success!")
					}

				} else {
					log.Println("Error input message. See -h to see help")
				}
				return nil
			},
		},
	}
	app.Version = "0.4"
	app.Run(os.Args)
}
