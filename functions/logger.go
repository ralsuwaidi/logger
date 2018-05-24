package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var path = "D:/Go-work/src/sql/logger.md"

// MakeLog - Saves a log to DB with only title nad text input
func MakeLog(title, body string) {

	l := Log{
		Title:       title,
		Body:        body,
		Tags:        getTags(body),
		LastUpdated: time.Now(),
	}
	l.CreateLog()
}

// this works on the command line
func writeCommandLineLog() {
	fmt.Print("What is the title?\n")
	title := scanner()
	fmt.Println("Type the log\n ")
	body := scanner()

	MakeLog(title, body)
	//body := scanner()
	//MakeLog(title, "#spank this is a # foshizzle my guchi is #cashizzle")

	CallClear()
	fmt.Printf("The log is saved with title: %s", title)
}

func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err) //okay to call os.exit()
		defer file.Close()
	}
}

func writeFile(title, body string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	_, err = file.WriteString(title)
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return //same as above
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func getPath() string {
	folderpath, _ := os.Getwd()
	return folderpath
}

func getTags(s string) []string {
	tags := []string{}
	if strings.Contains(s, "#") {

		split := strings.SplitAfter(s, "#")
		for i := 1; i < len(split); i++ {

			//take all text, find any slice with #inside
			v := strings.ToLower(split[i])
			tag := strings.Split(v, " ")
			tags = append(tags, tag[0])
		}

		//clean tags
		for i, v := range tags {
			if v == "" {
				tags = append(tags[:i], tags[i+1:]...)
				break
			}
		}

		//return the tag
		return tags
	}

	//return empty tag if no
	return tags
}
