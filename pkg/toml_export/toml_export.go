package yaml_export

import "fmt"

type TomlExport struct{}

func (t *TomlExport) PrintToml() {
	fmt.Println("Hello TOML world!")
}
