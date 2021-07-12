package yaml_export

import "fmt"

type YamlExport struct{}

func (y *YamlExport) PrintYaml() {
	fmt.Println("Hello YAML world!")
}
