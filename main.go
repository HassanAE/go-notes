package main

import (
	"bufio"
	"encoding/json/v2"
	"example/notes/notes"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

func main() {
	mynotes := notes.Notes{}
	groups := []string{"glob"}

	//filePath := "./notes.json"
	home := os.Getenv("HOME")
	os.Mkdir(path.Join(home, ".notes"), 0777)
	filePath := path.Join("/", home, ".notes", "notes.json")

	myFile, err := os.OpenFile(filePath, os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("notes: can not create or open \"%s\"", filePath)
		os.Exit(1)
	}
	defer myFile.Close()

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json.Unmarshal(file, &mynotes)

	for _, note := range mynotes {
		if !slices.Contains(groups, note.Group) {
			groups = append(groups, note.Group)
		}
	}

	if len(os.Args) > 1 {
		//fmt.Println("You are in command line mode")
		line := strings.Join(os.Args[1:], " ")
		notes.Parse(line, &mynotes, groups, filePath)
	} else {
		fmt.Print("\033[5;40H\033[2J")
		fmt.Print("Notes app to take notes version 1.0")
		fmt.Print("\033[7;2H")

		//fmt.Print("\033[5;2H\033[2k\033[5J")

		scanner := bufio.NewScanner(os.Stdin)

	cli:

		for {
			fmt.Print(">> ")
			for scanner.Scan() {
				line := scanner.Text()
				if strings.TrimSpace(line) == "exit" {
					break cli
				}
				notes.Parse(line, &mynotes, groups, filePath)
				//fmt.Println(line)
				break
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
