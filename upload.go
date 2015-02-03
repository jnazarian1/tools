package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"gitlab.mitre.org/intervention-engine/hdsfhir"
	"io/ioutil"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "upload"
	app.Usage = "Convert health-data-standards JSON to FHIR JSON and upload it to a FHIR Server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "fhir, f",
			Usage: "URL for the FHIR server",
		},
		cli.StringFlag{
			Name:  "json, j",
			Usage: "Path to the directory of JSON files",
		},
	}
	app.Action = func(c *cli.Context) {
		fhirUrl := c.String("fhir")
		path := c.String("json")
		if fhirUrl == "" || path == "" {
			fmt.Println("You must provide a FHIR URL and path to JSON files")
		} else {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				panic("Couldn't list the directory")
			}
			for _, file := range files {
				patient := &hdsfhir.Patient{}
				jsonBlob, err := ioutil.ReadFile(path + "/" + file.Name())
				if err != nil {
					panic("Couldn't read the JSON file" + err.Error())
				}
				json.Unmarshal(jsonBlob, patient)
				patient.PostToFHIRServer(fhirUrl)
			}

		}
	}

	app.Run(os.Args)
}