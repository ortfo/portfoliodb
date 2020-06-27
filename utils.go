package main

import (
	"io/ioutil"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

// ReadFile reads the content of ``filepath`` and returns the contents as a string
func ReadFile(filepath string) string {
    file, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    b, err := ioutil.ReadAll(file)
    return string(b)
}

// ValidateWithJSONSchema checks if the JSON document at ``documentFilepath`` conforms to the JSON schema at ``schemaFilepath``
func ValidateWithJSONSchema(documentFilepath string, schemaFilepath string) (bool, []gojsonschema.ResultError) {
	schema := gojsonschema.NewReferenceLoader("file://$schemaFilepath")
	document := gojsonschema.NewReferenceLoader("file://$documentFilepath")
	result, err := gojsonschema.Validate(schema, document)
	if err != nil {
		panic(err.Error())
	}
	
	if result.Valid() {
		return true, nil
	}
	var errorMessages []gojsonschema.ResultError
	for _, desc := range result.Errors() {
		errorMessages = append(errorMessages, desc)
	}
	return false, errorMessages
}
