package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content	  string
}

func main() {
	fileName := flag.String("file", "first-post.txt", "The name of the file to convert to HTML")
	flag.Parse()
	TextFileName := strings.TrimSuffix(*fileName, ".txt")

page := Page{
	TextFilePath: fmt.Sprintf("%s.txt", TextFileName),
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

}
