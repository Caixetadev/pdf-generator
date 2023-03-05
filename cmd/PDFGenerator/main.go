package main

import (
	"fmt"
	"generatePDF/pkg/htmlParser"
	"generatePDF/pkg/pdfGenerator"
	"log"
	"os"
)

type Data struct {
	Name string
}

func main() {
	h := htmlParser.New("tmp")

	wk := pdfGenerator.NewWkHTMLToPDF("tmp")

	dataHTML := Data{Name: "Cesar Augusto"}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)

	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(htmlGenerated)

	fmt.Println("HTML generated", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF generated", filePDFName)
}
