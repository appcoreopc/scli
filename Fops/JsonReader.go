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

type IJsonReader interface {
	GetCommandJson(path string) interface{}
}
