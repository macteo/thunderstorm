package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/RobotsAndPencils/buford/certificate"
	"github.com/RobotsAndPencils/buford/payload"
	"github.com/RobotsAndPencils/buford/payload/badge"
	"github.com/RobotsAndPencils/buford/push"
	"github.com/codegangsta/cli"
)

func send(token string, headers *push.Headers, payload interface{}, service *push.Service) {
	if token != "" {
		println("Sending push: ", token)

		id, err := service.Push(token, headers, payload)
		if err != nil {
			log.Fatal(err, id)
		} else {
			println("Push sent successfully ", id)
		}

	}
}

func main() {
	var filename string
	var alert string
	var badgeInt int
	var sound string
	var category string
	var environmentString string
	var passphrase string
	var contentAvailable bool
	var payloadString string
	var payloadFile string
	var tokensFile string

	app := cli.NewApp()
	app.Name = "thunderstorm"
	// app.EnableBashCompletion = true
	app.Usage = "push TOKEN [...]"
	app.Version = "0.2.0"
	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "Sends an Apple Push Notification to specified devices",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "certificate, c",
					Value:       "",
					Usage:       "Path to certificate (.p12) file",
					Destination: &filename,
				},
				cli.StringFlag{
					Name:        "alert, m",
					Value:       "",
					Usage:       "Body of the alert to send in the push notification",
					Destination: &alert,
				},
				cli.IntFlag{
					Name:        "badge, b",
					Usage:       "Badge number to set with the push notification",
					Destination: &badgeInt,
				},
				cli.StringFlag{
					Name:        "sound, s",
					Usage:       "Sound to play with the notification",
					Destination: &sound,
				},
				cli.StringFlag{
					Name:        "category, y",
					Usage:       "Category of notification",
					Destination: &category,
				},
				cli.StringFlag{
					Name:        "environment, e",
					Usage:       "Environment to send push notification (production or development (default))",
					Destination: &environmentString,
				},
				cli.StringFlag{
					Name:        "passphrase, p",
					Usage:       "Provides the certificate passphrase",
					Destination: &passphrase,
				},
				cli.BoolFlag{
					Name:        "content-available, n",
					Usage:       "Indicates content available",
					Destination: &contentAvailable,
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
					Name:        "tokens-path, t",
					Usage:       "Path of JSON file containing the tokens",
					Destination: &tokensFile,
				},
			},

			// TODO: support other parameters https://github.com/nomad/houston/blob/master/bin/apn

			Action: func(c *cli.Context) {
				deviceToken := c.Args().First()

				cert, privateKey, err := certificate.Load(filename, passphrase)
				if err != nil {
					log.Fatal(err)
				}

				tls := certificate.TLS(cert, privateKey)

				commonName := tls.Leaf.Subject.CommonName
				bundle := strings.Replace(commonName, "Apple Push Services: ", "", 1)

				var environment = push.Development

				if environmentString == "production" {
					environment = push.Production
				}

				client, err := push.NewClient(tls)
				if err != nil {
					log.Fatal(err)
				}

				service := push.Service{
					Client: client,
					Host:   environment,
				}

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
						return
					}
				}

				if p == nil {
					bUint := uint(badgeInt)

					// TODO: use Alert title, body and action or let it be with paylod
					pay := payload.APS{
						Alert:            payload.Alert{Body: alert},
						Badge:            badge.New(bUint),
						Sound:            sound,
						Category:         category,
						ContentAvailable: contentAvailable,
					}
					p = pay.Map()
				}

				headers := &push.Headers{
					Topic: bundle,
				}

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
							return
						}
						for _, token := range tokens {
							send(token, headers, p, &service)
						}
					}
				} else {
					send(deviceToken, headers, p, &service)
				}
			},
		},
	}

	app.Run(os.Args)
}
