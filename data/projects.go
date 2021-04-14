package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

type (
	File struct {
		Headers []string
		Data    [][]string
	}
)

func LoadProjects(file string) (File, error) {
	recordFile, err := os.Open(file)
	if err != nil {
		color.New(color.FgRed).Printf("Error reading file: %s", err.Error())
		return File{}, err
	}

	reader := csv.NewReader(recordFile)

	headers, err := reader.Read()
	if err != nil {
		color.New(color.FgRed).Printf("Error reading file: %s", err.Error())
		return File{}, err
	}

	response := File{
		Headers: headers,
		Data:    make([][]string, 0),
	}

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return File{}, err
		}

		response.Data = append(response.Data, record)
	}

	return response, nil
}
