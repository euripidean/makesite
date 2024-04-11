package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content	  string
}

var totalFileSize int

func main() {
	startTime := time.Now()
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
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("\033[1;32mSuccess!\033[0m Generated \033[1m%d\033[0m pages (%.1fKB total) in %.2f seconds.\n", len(fileNames), float64(totalFileSize)/1024, duration.Seconds())
	
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

originalText := string(fileContents)
translatedText, err := gtranslate.TranslateWithParams(originalText, gtranslate.TranslationParams{
	From: "en",
	To: "fr",
})
if err != nil {
	panic(err)
} else {
	fmt.Printf("\033[1;33mTranslated %s\033[0m\n", page.TextFilePath)
}

page.Content = translatedText

t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

newFile, err := os.Create(page.HTMLPagePath)
if err != nil {
	panic(err)
}

t.Execute(newFile, page)

// get the size of the new file
fileInfo, err := newFile.Stat()
if err != nil {
	panic(err)
}

if fileInfo.Size() == 0 {
	panic("The file was not created")
} else {
	totalFileSize += int(fileInfo.Size())
	fmt.Printf("\033[1;33mGenerated %s\033[0m\n", page.HTMLPagePath)
}
}
