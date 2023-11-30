package main

import (
	"fmt"
	"log"
	"make-api-script/variables"
	"os"
)

func main() {
	fmt.Println("sempre precisamos começar por algum lugar")

	for _, folderName := range variables.FoldersName {
		if err := os.Mkdir(folderName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	for _, fileContent := range variables.ContentFiles {
		file, err := os.Create(fileContent["fileName"])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println("File is created successfully.")

		_, err = file.WriteString(fileContent["content"])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Content file add successfully.")
	}
}
