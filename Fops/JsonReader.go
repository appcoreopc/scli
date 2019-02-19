package Fops

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/appcoreopc/scli/Model"
)

type JsonReader struct {
}

func (j *JsonReader) GetCommandJson(path string) interface{} {

	var model = Model.CommandCliModel{}
	f, err := os.Open(path)

	if err == nil {
		jsonBytes, _ := ioutil.ReadAll(f)
		json.Unmarshal(jsonBytes, &model)
	}

	defer f.Close()

	return model
}

func (j *JsonReader) WriteJson(path string, content []byte) bool {

	// f, err := os.Create(path)
	// defer f.Close()

	err := ioutil.WriteFile(path, content, 0644)

	if err != nil {
		return false
	}

	// f.Write(content)

	return true
}

func (j *JsonReader) Deserialize(jsonText string) interface{} {

	var model = Model.CommandCliModel{}

	// f, err := os.Open(path)

	// if err == nil {
	// 	jsonBytes, _ := ioutil.ReadAll(f)
	// 	json.Unmarshal(jsonBytes, &model)
	// }

	// defer f.Close()

	return model
}

type IJsonReader interface {
	GetCommandJson(path string) interface{}
}
