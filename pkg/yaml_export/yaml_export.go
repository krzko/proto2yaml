package yaml_export

import "io/ioutil"

type YamlExport struct{}

func (y *YamlExport) SaveFile(b []byte, f string) {
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		panic(err)
	}
}
