package main

import (
	"fmt"

	"log"
	"os"

	"github.com/Caixetadev/pdf-generator/pkg/htmlParser"
	"github.com/Caixetadev/pdf-generator/pkg/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {
	h := htmlParser.New("pkg/templates")

	wk := pdfGenerator.NewWkHTMLToPDF("pkg/tmp")

	dataHTML := Data{Name: "Rafael"}

	htmlGenerated, err := h.Create("pkg/templates/example.html", dataHTML)

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
