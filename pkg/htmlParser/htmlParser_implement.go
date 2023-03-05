package htmlParser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type HTMLStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &HTMLStruct{rootPath: rootPath}
}

func (h *HTMLStruct) Create(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)

	if err != nil {
		return "", err
	}

	fileName := h.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)

	if err != nil {
		return "", err
	}

	if err = templateGenerator.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
