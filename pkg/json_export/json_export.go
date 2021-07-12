package yaml_export

import "fmt"

type JsonExport struct{}

func (j *JsonExport) PrintJson() {
	fmt.Println("Hello JSON world!")
}
