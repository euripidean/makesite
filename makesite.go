package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content	  string
}

func main() {
	// fileName := flag.String("file", "first-post.txt", "The name of the file to convert to HTML")
	dir := flag.String("dir", ".", "The directory where the text file is located")
	flag.Parse()
	

	// find all the text files in the directory
	var fileNames []string

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(info.Name(), ".txt") {
			fileNames = append(fileNames, info.Name())
			createHTMLPage(path, strings.TrimSuffix(info.Name(), ".txt"))
		}

		return nil
	})

    if err != nil {
        panic(err)
    }
	fmt.Printf("\033[1;32mSuccess!\033[0m Generated %d pages\n", len(fileNames))
}

func createHTMLPage(TextFilePath string, TextFileName string) {
	page := Page{
	TextFilePath: TextFilePath,
	HTMLPagePath: fmt.Sprintf("%s.html", TextFileName), 
	Content: "",
}

	fileContents, err := os.ReadFile(page.TextFilePath)
	if err != nil {
		panic(err)
} 

page.Content = string(fileContents)

t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

newFile, err := os.Create(page.HTMLPagePath)
if err != nil {
	panic(err)
}

t.Execute(newFile, page)

fmt.Printf("Created %s\n", page.HTMLPagePath)

}
