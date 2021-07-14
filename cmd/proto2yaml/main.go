package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/emicklei/proto"
	"github.com/fatih/color"
	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"github.com/krzko/proto2yaml/pkg/json_export"
	"github.com/krzko/proto2yaml/pkg/yaml_export"
)

var (
	buildVersion string
	// commit       string
)

type ProtoExport struct {
	Version  string        `json:"version" yaml:"version"`
	Packages []PackageItem `json:"packages" yaml:"packages"`
}
type PackageItem struct {
	Package  string        `json:"package" yaml:"package"`
	Services []ServiceItem `json:"services" yaml:"services"`
}
type ServiceItem struct {
	Service string   `json:"service" yaml:"service"`
	RPCs    []string `json:"rpc" yaml:"rpc"`
}

func main() {
	// Rainbow
	c := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgMagenta, color.FgCyan, color.FgWhite, color.FgHiRed, color.FgHiGreen, color.FgHiYellow, color.FgHiBlue, color.FgHiMagenta, color.FgHiCyan, color.FgHiWhite}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
	c0 := color.New(c[0]).SprintFunc()
	c1 := color.New(c[1]).SprintFunc()
	c2 := color.New(c[2]).SprintFunc()
	c3 := color.New(c[3]).SprintFunc()
	c4 := color.New(c[4]).SprintFunc()
	c5 := color.New(c[5]).SprintFunc()
	c6 := color.New(c[6]).SprintFunc()
	c7 := color.New(c[7]).SprintFunc()
	c8 := color.New(c[8]).SprintFunc()
	c9 := color.New(c[9]).SprintFunc()
	appName := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", c0("p"), c1("r"), c2("o"), c3("t"), c4("o"), c5("2"), c6("y"), c7("a"), c8("m"), c9("l"))

	app := &cli.App{
		Name:      appName,
		Usage:     "A command-line utility to convert Protocol Buffers (proto) files to YAML",
		UsageText: appName + " [global options] command [command options] [arguments...]",
		Version:   buildVersion,
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "pro2yaml: Command not found: %q\n", command)
		},
	}

	app.Commands = cli.Commands{
		{
			Name:  "json",
			Usage: "The outputs are formatted as a JSON",
			Subcommands: []*cli.Command{
				{
					Name:    "export",
					Aliases: []string{"x"},
					Usage:   "Exports the proto defintions to a file",
					Flags: []cli.Flag{
						&cli.StringSliceFlag{
							Name:        "exclude-option",
							Usage:       "Exclude this option to filter on, e.g. --exclude-option 'deprecated=true'",
							Aliases:     []string{"eo"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringFlag{
							Name:        "file",
							Usage:       "The exported file",
							DefaultText: "./foobar_protos.yaml",
							Aliases:     []string{"f"},
							Required:    true,
						},
						&cli.StringSliceFlag{
							Name:        "include-option",
							Usage:       "Include this option to filter on, e.g. --include-option 'go_package=api'",
							Aliases:     []string{"io"},
							DefaultText: "",
							Required:    false,
						},
						&cli.BoolFlag{
							Name:        "pretty",
							Usage:       "Pretty prints the output, export --pretty",
							DefaultText: "",
						},
						&cli.StringFlag{
							Name:        "source",
							Usage:       "The source directory",
							DefaultText: "~/foobar/proto",
							Aliases:     []string{"s"},
							Required:    true,
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("%s %s %s %s %s\n", color.GreenString("==>"), color.HiWhiteString("Using Source:"), color.HiGreenString(c.String("source")), color.HiWhiteString("Destination:"), color.HiGreenString(c.String("file")))

						if len(c.StringSlice("exclude-option")) != 0 && len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s\n", color.HiRedString("==>"), color.HiWhiteString("❌ Please use 'exclude-option' or 'include-option' only"))
							os.Exit(1)
						}

						// Get files
						files, err := getFiles(c.String("source"), ".proto")
						if err != nil {
							fmt.Println(err)
						}

						// Return filtered files
						ff, err := searchFiles(files, "Request) returns")
						if err != nil {
							fmt.Println(err)
						}

						// Parse the protos
						parseFiles(ff)

						// Generate object to export
						var obj *ProtoExport
						if len(c.StringSlice("exclude-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("exclude"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("exclude-option"), "exclude")
							if err != nil {
								fmt.Println(err)
							}
						} else if len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("include"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("include-option"), "include")
							if err != nil {
								fmt.Println(err)
							}
						} else {
							obj, err = generateExport(ff, nil, "")
							if err != nil {
								fmt.Println(err)
							}
						}

						// Forat the obj
						j, _ := json.Marshal(obj)
						je := json_export.JsonExport{}

						fmt.Printf("%s %s %s\n", color.GreenString("==>"), color.HiWhiteString("💾 Saving to:"), color.HiGreenString(c.String("file")))
						if c.Bool("pretty") {
							jpp, _ := je.PrettyPrint([]byte(j))
							je.SaveFile(jpp, c.String("file"))
						} else {
							je.SaveFile(j, c.String("file"))
						}

						return nil
					},
				},
				{
					Name:    "print",
					Aliases: []string{"p"},
					Usage:   "Prints the proto defintions to console",
					Flags: []cli.Flag{
						&cli.StringSliceFlag{
							Name:        "exclude-option",
							Usage:       "Exclude this option to filter on, e.g. --exclude-option 'deprecated=true'",
							Aliases:     []string{"eo"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringSliceFlag{
							Name:        "include-option",
							Usage:       "Include this option to filter on, e.g. --include-option 'go_package=api'",
							Aliases:     []string{"io"},
							DefaultText: "",
							Required:    false,
						},
						&cli.BoolFlag{
							Name:        "pretty",
							Usage:       "Pretty prints the output,  --pretty",
							DefaultText: "",
						},
						&cli.StringFlag{
							Name:        "source",
							Usage:       "The source directory",
							DefaultText: "~/foobar/proto",
							Aliases:     []string{"s"},
							Required:    true,
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("%s %s %s %s %s\n", color.GreenString("==>"), color.HiWhiteString("Using Source:"), color.HiGreenString(c.String("source")), color.HiWhiteString("Destination:"), color.HiGreenString(c.String("file")))

						if len(c.StringSlice("exclude-option")) != 0 && len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s\n", color.HiRedString("==>"), color.HiWhiteString("❌ Please use 'exclude-option' or 'include-option' only"))
							os.Exit(1)
						}

						// Get files
						files, err := getFiles(c.String("source"), ".proto")
						if err != nil {
							fmt.Println(err)
						}

						// Return filtered files
						ff, err := searchFiles(files, "Request) returns")
						if err != nil {
							fmt.Println(err)
						}

						// Parse the protos
						parseFiles(ff)

						// Generate object to export
						var obj *ProtoExport
						if len(c.StringSlice("exclude-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("exclude"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("exclude-option"), "exclude")
							if err != nil {
								fmt.Println(err)
							}
						} else if len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("include"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("include-option"), "include")
							if err != nil {
								fmt.Println(err)
							}
						} else {
							obj, err = generateExport(ff, nil, "")
							if err != nil {
								fmt.Println(err)
							}
						}

						// Forat the obj
						j, _ := json.Marshal(obj)
						je := json_export.JsonExport{}

						fmt.Printf("%s %s\n", color.GreenString("==>"), color.HiWhiteString("🖥️ Printing to console"))
						if c.Bool("pretty") {
							fmt.Println()
							jpp, _ := je.PrettyPrint([]byte(j))
							fmt.Printf("%s", jpp)
						} else {
							fmt.Println()
							fmt.Println(string(j))
						}

						return nil
					},
				},
			},
		},
		// {
		// 	Name:  "toml",
		// 	Usage: "The outputs are formatted as a TOML",
		// 	Subcommands: []*cli.Command{
		// 		{
		// 			Name:    "print",
		// 			Aliases: []string{"p"},
		// 			Usage:   "Prints the proto defintions to console",
		// 			Action: func(c *cli.Context) error {
		// 				// t := pt.TomlPkg{}
		// 				// t.PrintToml()
		// 				return nil
		// 			},
		// 		},
		// 	},
		// },
		{
			Name:  "yaml",
			Usage: "The outputs are formatted as a YAML",
			Subcommands: []*cli.Command{
				{
					Name:    "export",
					Aliases: []string{"x"},
					Usage:   "Exports the proto defintions to a file",
					Flags: []cli.Flag{
						// &cli.BoolFlag{
						// 	Name:        "openslo",
						// 	Usage:       "Exports in OpenSLO format, --file slo",
						// 	DefaultText: "--openslo=false",
						// 	Aliases:     []string{"slo"},
						// },
						&cli.StringSliceFlag{
							Name:        "exclude-option",
							Usage:       "Exclude this option to filter on, e.g. --exclude-option 'deprecated=true'",
							Aliases:     []string{"eo"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringFlag{
							Name:        "file",
							Usage:       "The exported file, e.g. ./foobar_protos.yaml",
							DefaultText: "",
							Aliases:     []string{"f"},
							Required:    true,
						},
						&cli.StringSliceFlag{
							Name:        "include-option",
							Usage:       "Include this option to filter on, e.g. --include-option 'go_package=api'",
							Aliases:     []string{"io"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringFlag{
							Name:        "source",
							Usage:       "The source directory, e.g. ~/foobar/proto",
							DefaultText: "",
							Aliases:     []string{"s"},
							Required:    true,
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("%s %s %s %s %s\n", color.GreenString("==>"), color.HiWhiteString("Using Source:"), color.HiGreenString(c.String("source")), color.HiWhiteString("Destination:"), color.HiGreenString(c.String("file")))

						if len(c.StringSlice("exclude-option")) != 0 && len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s\n", color.HiRedString("==>"), color.HiWhiteString("❌ Please use 'exclude-option' or 'include-option' only"))
							os.Exit(1)
						}

						// Get files
						files, err := getFiles(c.String("source"), ".proto")
						if err != nil {
							fmt.Println(err)
						}

						// Return filtered files
						ff, err := searchFiles(files, "Request) returns")
						if err != nil {
							fmt.Println(err)
						}

						// Parse the protos
						parseFiles(ff)

						// Generate object to export
						var obj *ProtoExport
						if len(c.StringSlice("exclude-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("exclude"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("exclude-option"), "exclude")
							if err != nil {
								fmt.Println(err)
							}
						} else if len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("include"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("include-option"), "include")
							if err != nil {
								fmt.Println(err)
							}
						} else {
							obj, err = generateExport(ff, nil, "")
							if err != nil {
								fmt.Println(err)
							}
						}

						// Print obj
						y, _ := yaml.Marshal(obj)
						ye := yaml_export.YamlExport{}
						fmt.Printf("%s %s %s\n", color.GreenString("==>"), color.HiWhiteString("💾 Saving to:"), color.HiGreenString(c.String("file")))
						ye.SaveFile(y, c.String("file"))

						return nil
					},
				},
				{
					Name:    "print",
					Aliases: []string{"p"},
					Usage:   "Prints the proto defintions to console",
					Flags: []cli.Flag{
						&cli.StringSliceFlag{
							Name:        "exclude-option",
							Usage:       "Exclude this option to filter on, e.g. --exclude-option 'deprecated=true'",
							Aliases:     []string{"eo"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringSliceFlag{
							Name:        "include-option",
							Usage:       "Include this option to filter on, e.g. --include-option 'go_package=api'",
							Aliases:     []string{"io"},
							DefaultText: "",
							Required:    false,
						},
						&cli.StringFlag{
							Name:        "source",
							Usage:       "The source directory, e.g. ~/foobar/proto",
							DefaultText: "",
							Aliases:     []string{"s"},
							Required:    true,
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("%s %s %s %s %s\n", color.GreenString("==>"), color.HiWhiteString("✨ Using Source:"), color.HiGreenString(c.String("source")), color.HiWhiteString("Destination:"), color.HiGreenString("console"))

						if len(c.StringSlice("exclude-option")) != 0 && len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s\n", color.HiRedString("==>"), color.HiWhiteString("❌ Please use 'exclude-option' or 'include-option' only"))
							os.Exit(1)
						}

						// Get files
						files, err := getFiles(c.String("source"), ".proto")
						if err != nil {
							fmt.Println(err)
						}

						// Return filtered files
						ff, err := searchFiles(files, "Request) returns")
						if err != nil {
							fmt.Println(err)
						}

						// Parse the protos
						parseFiles(ff)

						// Generate object to export
						var obj *ProtoExport
						if len(c.StringSlice("exclude-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("exclude"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("exclude-option"), "exclude")
							if err != nil {
								fmt.Println(err)
							}
						} else if len(c.StringSlice("include-option")) != 0 {
							fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Using Filter '"), color.HiGreenString("include"), color.HiWhiteString("'"))
							obj, err = generateExport(ff, c.StringSlice("include-option"), "include")
							if err != nil {
								fmt.Println(err)
							}
						} else {
							obj, err = generateExport(ff, nil, "")
							if err != nil {
								fmt.Println(err)
							}
						}

						// Print obj
						y, _ := yaml.Marshal(obj)
						fmt.Printf("%s %s\n", color.GreenString("==>"), color.HiWhiteString("🖥️ Printing to console"))
						fmt.Println()
						fmt.Println(string(y))

						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// generateExport a filtered object, don't @ me :P
func generateExport(files, filter []string, filterType string) (*ProtoExport, error) {
	pe := &ProtoExport{}
	pe.Version = buildVersion

	for _, f := range files {
		reader, _ := os.Open(f)
		defer reader.Close()

		parser := proto.NewParser(reader)
		definition, _ := parser.Parse()

		// Use filters or not
		if len(filter) != 0 {
			// Validate = assignment
			filSplit := strings.Split(filter[0], "=")
			if len(filSplit) < 2 {
				fmt.Printf("%s %s\n", color.HiRedString("==>"), color.HiWhiteString("❌ Please use '=' for assignment of options"))
				os.Exit(1)
			}

			//TODO(krzko): Dodgy exclude, I don't know why
			add := false
			// Filters
			switch filterType {
			case "exclude":
				proto.Walk(definition,
					proto.WithOption(func(o *proto.Option) {
						// fmt.Printf("%v - %v - %v\n", f, filter, filterType)
						// Split option to get option k,v
						ck := fmt.Sprintf("%v", o.Constant)
						cv := strings.Split(ck, " ")

						// Compare flags to option value
						if o.Name == filSplit[0] && cv[1] == filSplit[1] {
							fmt.Printf("%s %s %v\n", color.BlueString("==>"), color.HiWhiteString("Excluding:"), f)
							add = false
						} else {
							add = true
						}
					}))
				//TODO(krzko): Had to implement this as it iterated over options continually
				if add {
					// only one package per file, keep track to add services to this package
					var index int
					for _, e := range definition.Elements {
						if p, ok := e.(*proto.Package); ok {
							index, ok = findPackage(pe.Packages, p.Name)
							if !ok {
								pe.Packages = append(pe.Packages, PackageItem{
									Package: p.Name,
								})
								index = len(pe.Packages) - 1
							}
							// only one package per file so can break
							break
						}
					}
					proto.Walk(definition,
						proto.WithService(func(s *proto.Service) {
							check := containsService(pe.Packages[index].Services, s.Name)
							if !check {
								pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
									Service: s.Name,
								})
							}
						}),
						proto.WithRPC(func(rpc *proto.RPC) {
							parent, ok := rpc.Parent.(*proto.Service)
							if !ok {
								return
							}

							i, check := findService(pe.Packages[index].Services, parent.Name)
							if check {
								pe.Packages[index].Services[i].RPCs = append(pe.Packages[index].Services[i].RPCs, rpc.Name)
							} else {
								// Add service and rpc
								pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
									Service: parent.Name,
									RPCs:    []string{rpc.Name},
								})
							}
						}))
				}
			case "include":
				proto.Walk(definition,
					proto.WithOption(func(o *proto.Option) {
						// Split option to get option k,v
						ck := fmt.Sprintf("%v", o.Constant)
						cv := strings.Split(ck, " ")

						// Compare flags to option value
						if o.Name == filSplit[0] && cv[1] == filSplit[1] {
							fmt.Printf("%s %s %v\n", color.BlueString("==>"), color.HiWhiteString("Including:"), f)

							// only one package per file, keep track to add services to this package
							var index int
							for _, e := range definition.Elements {
								if p, ok := e.(*proto.Package); ok {
									index, ok = findPackage(pe.Packages, p.Name)
									if !ok {
										pe.Packages = append(pe.Packages, PackageItem{
											Package: p.Name,
										})
										index = len(pe.Packages) - 1
									}
									// only one package per file so can break
									break
								}
							}
							proto.Walk(definition,
								proto.WithService(func(s *proto.Service) {
									check := containsService(pe.Packages[index].Services, s.Name)
									if !check {
										pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
											Service: s.Name,
										})
									}
								}),
								proto.WithRPC(func(rpc *proto.RPC) {
									parent, ok := rpc.Parent.(*proto.Service)
									if !ok {
										return
									}

									i, check := findService(pe.Packages[index].Services, parent.Name)
									if check {
										pe.Packages[index].Services[i].RPCs = append(pe.Packages[index].Services[i].RPCs, rpc.Name)
									} else {
										// Add service and rpc
										pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
											Service: parent.Name,
											RPCs:    []string{rpc.Name},
										})
									}
								}))
						}
					}))
			}
		} else {
			// only one package per file, keep track to add services to this package
			var index int
			for _, e := range definition.Elements {
				if p, ok := e.(*proto.Package); ok {
					index, ok = findPackage(pe.Packages, p.Name)
					if !ok {
						pe.Packages = append(pe.Packages, PackageItem{
							Package: p.Name,
						})
						index = len(pe.Packages) - 1
					}
					// only one package per file so can break
					break
				}
			}
			proto.Walk(definition,
				proto.WithService(func(s *proto.Service) {
					check := containsService(pe.Packages[index].Services, s.Name)
					if !check {
						pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
							Service: s.Name,
						})
					}
				}),
				proto.WithRPC(func(rpc *proto.RPC) {
					parent, ok := rpc.Parent.(*proto.Service)
					if !ok {
						return
					}

					i, check := findService(pe.Packages[index].Services, parent.Name)
					if check {
						pe.Packages[index].Services[i].RPCs = append(pe.Packages[index].Services[i].RPCs, rpc.Name)
					} else {
						// Add service and rpc
						pe.Packages[index].Services = append(pe.Packages[index].Services, ServiceItem{
							Service: parent.Name,
							RPCs:    []string{rpc.Name},
						})
					}
				}))

		}
	}
	return pe, nil
}

func getFiles(root, extension string) ([]string, error) {
	var files []string

	fmt.Printf("%s %s %v\n", color.BlueString("==>"), color.HiWhiteString(("Walking")), root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relpath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if !strings.HasSuffix(relpath, extension) || info.IsDir() {
			// Exclude directories
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s %s %s %s %s\n", color.BlueString("==>"), color.HiWhiteString("Found"), color.HiGreenString(fmt.Sprint(len(files))), color.HiWhiteString(extension), color.HiWhiteString("files"))
	return files, nil
}

// func handleMessage(m *proto.Message) {
// 	fmt.Printf("%v \n", m.Name)
// }

func handlePackage(p *proto.Package) {
	fmt.Printf("%v\n", p.Name)
}

func handleRPC(r *proto.RPC) {
	parent := fmt.Sprintf("%v", r.Parent)
	p := strings.Split(parent, " ")
	fmt.Printf("%s %v | %v\n", color.BlueString("==>"), p[2], r.Name)

	// fmt.Println("data written")
}

func handleService(s *proto.Service) {
	fmt.Printf("%v\n", s.Name)
}

func parseFiles(files []string) {

	for _, f := range files {
		reader, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
		}
		defer reader.Close()

		parser := proto.NewParser(reader)
		definition, _ := parser.Parse()

		fmt.Printf("%s %s\n", color.GreenString("==>"), color.HiWhiteString("📄 "+f))
		// proto.Walk(definition, proto.WithService(handleService), proto.WithMessage(handleMessage))
		fmt.Printf("%s %s ", color.GreenString("==>"), color.HiWhiteString("📦 Package"))
		proto.Walk(definition, proto.WithPackage(handlePackage))

		// fmt.Printf("%s %s\n", color.GreenString("==>"), color.HiWhiteString("🎁 💈 Service/s"))
		// proto.Walk(definition, proto.WithService(handleService))

		fmt.Printf("%s %s\n", color.GreenString("==>"), color.HiWhiteString("🧩 Service | RPC"))
		proto.Walk(definition, proto.WithRPC(handleRPC))

	}

}

func contains(s []string, str string) (result bool) {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func containsPackage(pis []PackageItem, str string) (result bool) {
	for _, pi := range pis {
		if pi.Package == str {
			return true
		}
	}
	return false
}

func containsService(sis []ServiceItem, str string) (result bool) {
	for _, si := range sis {
		if si.Service == str {
			return true
		}
	}
	return false
}

func findPackage(pis []PackageItem, str string) (int, bool) {
	for i, pi := range pis {
		if pi.Package == str {
			return i, true
		}
	}
	return -1, false
}

func findService(sis []ServiceItem, str string) (int, bool) {
	for i, si := range sis {
		if si.Service == str {
			return i, true
		}
	}
	return -1, false
}

func searchFiles(files []string, filter string) ([]string, error) {
	var found []string

	fmt.Printf("%s %s%s%s\n", color.BlueString("==>"), color.HiWhiteString("Filter '"), color.HiGreenString(filter), color.HiWhiteString("'"))

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		f := string(data)
		//TODO(krzko): Better filter for stream/unary RPCs
		if strings.Contains(f, "Request) returns (stream....") {
			// Skip streams
		} else if strings.Contains(f, filter) {
			found = append(found, file)
		}
	}

	fmt.Printf("%s %s %s %s\n", color.BlueString("==>"), color.HiWhiteString("Using"), color.HiGreenString(fmt.Sprint(len(found))), color.HiWhiteString("filtered files"))
	return found, nil
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
