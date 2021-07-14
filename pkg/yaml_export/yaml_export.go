package yaml_export

import (
	"fmt"
	"io/ioutil"
)

type YamlExport struct{}

func (y *YamlExport) SaveFile(b []byte, f string) {
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
