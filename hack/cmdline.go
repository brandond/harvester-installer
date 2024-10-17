package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/harvester/harvester-installer/pkg/config"
	"github.com/harvester/harvester-installer/pkg/util"
	"github.com/rancher/mapper"
	"github.com/rancher/mapper/convert"
	"sigs.k8s.io/yaml"
)

var (
	schemas = mapper.NewSchemas().Init(func(s *mapper.Schemas) *mapper.Schemas {
		s.DefaultMappers = func() []mapper.Mapper {
			return []mapper.Mapper{
				config.NewToMap(),
				config.NewToSlice(),
				config.NewToBool(),
				&config.FuzzyNames{},
			}
		}
		return s
	}).MustImport(config.HarvesterConfig{})
	schema = schemas.Schema("harvesterConfig")
)

func main() {
	cmdLine := strings.Join(os.Args[1:], " ")
	fmt.Printf("Parsing command line: %q\n", cmdLine)

	data, err := util.ParseCmdLine(cmdLine, "harvester")
	if err != nil {
		fmt.Print("Failed to parse command line: %v", err)
		os.Exit(1)
	}

	result := config.NewHarvesterConfig()
	schema.Mapper.ToInternal(data)
	if err := convert.ToObj(data, result); err != nil {
		fmt.Printf("Failed to convert args to HarvesterConfig: %v\n", err)
		os.Exit(1)
	}

	b, err := yaml.Marshal(result)
	if err != nil {
		fmt.Printf("Failed to marshal HarvesterConfig: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", string(b))
}
