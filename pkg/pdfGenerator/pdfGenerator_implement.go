package pdfGenerator

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkHTMLToPDF struct {
	rootPath string
}

func NewWkHTMLToPDF(rootPath string) PDFGeneratorInterface {
	return &wkHTMLToPDF{rootPath: rootPath}
}

func (w *wkHTMLToPDF) Create(htmlFile string) (string, error) {
	file, err := os.Open(htmlFile)

	if err != nil {
		return "", err
	}

	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return "", err
	}

	pdfGenerator.AddPage(wkhtmltopdf.NewPageReader(file))

	if err = pdfGenerator.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err = pdfGenerator.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
