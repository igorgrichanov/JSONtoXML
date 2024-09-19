package ex00

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

const (
	xmlExtension = iota
	jsonExtension
)

type DBReader interface {
	ReadDBFromFile() (*Recipes, error)
}

type JSONDBReader struct {
	Reader *json.Decoder
}

type XMLDBReader struct {
	Reader *xml.Decoder
}

func (j *JSONDBReader) ReadDBFromFile() (*Recipes, error) {
	recipes := new(Recipes)
	err := j.Reader.Decode(recipes)
	return recipes, err
}

func (x *XMLDBReader) ReadDBFromFile() (*Recipes, error) {
	recipes := new(Recipes)
	err := x.Reader.Decode(&recipes)
	return recipes, err
}

type IngredientItem struct {
	ItemName  string `xml:"itemname"           json:"ingredient_name"`
	ItemCount string `xml:"itemcount"          json:"ingredient_count"`
	ItemUnit  string `xml:"itemunit,omitempty" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	Name        string           `xml:"name"             json:"name"`
	StoveTime   string           `xml:"stovetime"        json:"time"`
	Ingredients []IngredientItem `xml:"ingredients>item" json:"ingredients"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []Cake   `xml:"cake"    json:"cake"`
}

// ParseDB parses .xml and .json files to Recipes
// It returns the code of extension (see const declaration at the top of the file) and a pointer
// to the filled Recipes struct
func ParseDB(filepath *string) (int, *Recipes) {
	file, err := os.Open(*filepath)
	for err != nil {
		log.Fatal("Incorrect filepath, try again")
	}
	defer file.Close()

	var reader DBReader
	var extension int

	if strings.HasSuffix(*filepath, ".xml") {
		extension = xmlExtension
		reader = &XMLDBReader{Reader: xml.NewDecoder(file)}
	} else if strings.HasSuffix(*filepath, ".json") {
		extension = jsonExtension
		reader = &JSONDBReader{Reader: json.NewDecoder(file)}
	} else {
		log.Fatal("Unsupported file type, only can work with .json and .xml")
	}

	recipes, err := reader.ReadDBFromFile()

	if err != nil {
		log.Fatal("Error occurred when reading recipes: ", err)
	}
	return extension, recipes
}

func MarshalDB(recipes *Recipes, extension int) []byte {
	var db []byte
	var err error

	if extension == jsonExtension {
		db, err = xml.MarshalIndent(recipes, "", "    ")
	} else {
		db, err = json.MarshalIndent(recipes, "", "    ")
	}
	if err != nil {
		log.Fatal("Error occurred when marshalling recipes: ", err)
	}
	if extension == jsonExtension {
		dbWithHeader := make([]byte, 0, len(db)+len(xml.Header))
		dbWithHeader = append(dbWithHeader, xml.Header...)
		dbWithHeader = append(dbWithHeader, db...)
		db = dbWithHeader
	}
	return db
}
