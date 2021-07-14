package json_export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonExport struct{}

func (j *JsonExport) PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func (j *JsonExport) SaveFile(b []byte, f string) {
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
