package main

import (
	"html/template"
	"os"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content	  string
}

func main() {
	fileContents, err := os.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
}

page := Page{
	TextFilePath: "first-post.txt", 
	TextFileName: "first-post.txt", 
	HTMLPagePath: "first-post.html", 
	Content: string(fileContents),
}

t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))


newFile, err := os.Create(page.HTMLPagePath)
if err != nil {
	panic(err)
}

t.Execute(newFile, page)

}
