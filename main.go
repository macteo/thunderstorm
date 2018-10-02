package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

func send(notification *apns2.Notification, client *apns2.Client) {
	res, err := client.Push(notification)
	if err != nil {
		log.Fatal(err, res)
	} else {
		fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	}
}

func main() {
	var team string
	var key string
	var topic string
	var payloadString string
	var payloadFile string
	var tokensFile string

	app := cli.NewApp()
	app.Name = "thunderstorm"
	// app.EnableBashCompletion = true
	app.Usage = "push TOKEN [...]"
	app.Version = "0.3.0"
	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "Sends an Apple Push Notification to specified devices",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "team, t",
					Value:       "",
					Usage:       "Team ID",
					Destination: &team,
				}, cli.StringFlag{
					Name:        "bundle, b",
					Value:       "",
					Usage:       "Topic (bundle-id)",
					Destination: &topic,
				}, cli.StringFlag{
					Name:        "key, k",
					Value:       "",
					Usage:       "Key ID",
					Destination: &key,
				},
				cli.StringFlag{
					Name:        "payload, P",
					Usage:       "JSON payload for notifications",
					Destination: &payloadString,
				},
				cli.StringFlag{
					Name:        "payload-file, f",
					Usage:       "Path of JSON file containing the payload to be sent",
					Destination: &payloadFile,
				},
				cli.StringFlag{
					Name:        "tokens-path, tp",
					Usage:       "Path of JSON file containing the tokens",
					Destination: &tokensFile,
				},
			},
			Action: func(c *cli.Context) error {
				deviceToken := c.Args().First()

				authKey, err := token.AuthKeyFromFile("AuthKey_" + key + ".p8")
				if err != nil {
					log.Fatal("token error:", err)
				}

				token := &token.Token{
					AuthKey: authKey,
					// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
					KeyID: key,
					// TeamID from developer account (View Account -> Membership)
					TeamID: team,
				}
				client := apns2.NewTokenClient(token)

				if payloadFile != "" {
					buffer, err := ioutil.ReadFile(payloadFile)
					if err == nil {
						payloadString = string(buffer)
					}
				}

				var p map[string]interface{}
				if payloadString != "" {
					error := json.Unmarshal([]byte(payloadString), &p)
					if error != nil {
						log.Fatal("Cannot parse payload")
						return err
					}
				}
				notification := &apns2.Notification{}
				notification.Topic = topic
				notification.Payload = payloadString

				tokensString := ""

				if tokensFile != "" {
					buffer, err := ioutil.ReadFile(tokensFile)
					if err == nil {
						tokensString = string(buffer)
					}

					var tokens []string
					if tokensString != "" {
						error := json.Unmarshal([]byte(tokensString), &tokens)
						if error != nil {
							log.Fatal("Cannot parse tokens array", error)
							return error
						}
						for _, token := range tokens {
							notification.DeviceToken = token
							send(notification, client)
						}
					}
				} else {
					notification.DeviceToken = deviceToken
					send(notification, client)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
