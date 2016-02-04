package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RobotsAndPencils/buford/certificate"
	"github.com/RobotsAndPencils/buford/payload/badge"
	"github.com/RobotsAndPencils/buford/push"
	"github.com/codegangsta/cli"
	// TODO: if the pull request to support category is accepted we can go back
	// to the main repo
	"github.com/macteo/buford/payload"
)

func main() {
	var filename string
	var alert string
	var badgeString string
	var sound string
	var category string
	var environmentString string

	app := cli.NewApp()
	app.Name = "tds"
	app.EnableBashCompletion = true
	app.Usage = "tds push TOKEN [...]"
	app.Commands = []cli.Command{

		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "Sends an Apple Push Notification to specified devices",
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
				cli.StringFlag{
					Name:        "badge, b",
					Usage:       "Badge number to set with the push notification",
					Destination: &badgeString,
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
			},

			// TODO: support other parameters https://github.com/nomad/houston/blob/master/bin/apn

			Action: func(c *cli.Context) {
				println("Sending push to: ", c.Args().First())

				// set these variables appropriately
				password := ""
				deviceToken := c.Args().First()

				cert, err := certificate.Load(filename, password)
				if err != nil {
					log.Fatal(err)
				}

				commonName := cert.Leaf.Subject.CommonName
				bundle := strings.Replace(commonName, "Apple Push Services: ", "", 1)

				var environment = push.Development

				if environmentString == "production" {
					environment = push.Production
				}
				service := push.Service{
					Client: push.NewClient(cert),
					Host:   environment,
				}

				bUint64, err := strconv.ParseUint(badgeString, 10, 32)
				bUint := uint(bUint64)

				p := payload.APS{
					Alert:    payload.Alert{Body: alert},
					Badge:    badge.New(bUint),
					Sound:    sound,
					Category: category,
				}

				headers := &push.Headers{
					Topic: bundle,
				}

				id, err := service.Push(deviceToken, headers, p)
				if err != nil {
					log.Fatal(err, id)
				} else {
					println("Push sent successfully")
				}
			},
		},
	}

	app.Run(os.Args)
}
