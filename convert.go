package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/urfave/cli/v2"
)

type ToolData struct {
	// timestamp         string
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Tags             []string `json:"tags"`
	OperatingSystems []string `json:"operating_systems"`
	License          string   `json:"license"`
	Availability     []string `json:"availability"`
	GithubURL        string   `json:"github_url"`
	URL              string   `json:"url"`
}

const (
	sheetName = "Output"
)

func convertXLSXToJSON(c *cli.Context) error {
	inputXLSXFilePath := c.String("input")
	// outputJSONFilePath := c.String("output")
	f, err := excelize.OpenFile(inputXLSXFilePath)
	if err != nil {
		log.Fatal(err)
	}

	rows := f.GetRows(sheetName)

	for _, row := range rows {
		var tool ToolData
		tool.Name = row[1]
		tool.Description = row[2]
		tool.Tags = stringTrimSplit(row[3])
		tool.OperatingSystems = stringTrimSplit(row[4])
		tool.License = row[5]
		tool.Availability = stringTrimSplit(row[6])
		tool.GithubURL = row[7]
		tool.URL = row[8]

		// currently hardcoded the tools folder
		outputJSONFilePath := "tools/" + safeCleanName(row[1])

		jsonData, err := json.Marshal(tool)
		if err != nil {
			log.Println(err)
		}

		err = ioutil.WriteFile(outputJSONFilePath, jsonData, 0644)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Successfully saved the output to:", outputJSONFilePath)
	}
	return nil
}
