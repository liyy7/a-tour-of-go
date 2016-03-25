package main

import (
	"fmt"
	"github.com/voxelbrain/goptions"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	options := struct {
		Servers  []string      `goptions:"-s, --server, obligatory, description='Servers to connect to'"`
		Password string        `goptions:"-p, --password, description='Don\\'t prompt for password'"`
		Timeout  time.Duration `goptions:"-t, --timeout, description='Connection timeout in seconds'"`
		Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs

		Execute struct {
			Command string   `goptions:"--command, mutexgroup='input', description='Command to execute', obligatory"`
			Script  *os.File `goptions:"--script, mutexgroup='input', description='Script to execute', rdonly"`
		} `goptions:"execute"`

		Delete struct {
			Path  string `goptions:"-n, --name, obligatory, description='Name of the entity to be deleted'"`
			Force bool   `goptions:"-f, --force, description='Force removal'"`
		} `goptions:"delete"`
	}{
		Timeout: 10 * time.Second,
	}

	goptions.ParseAndFail(&options)

	var printStructure func(st interface{}, level int)
	printStructure = func(st interface{}, level int) {
		for i := 0; i < reflect.TypeOf(st).NumField(); i++ {
			typeField := reflect.TypeOf(st).Field(i)
			valueField := reflect.ValueOf(st).Field(i)

			fmt.Printf("%s%s: ", strings.Repeat("\t", level), typeField.Name)

			if valueField.Kind() == reflect.Struct {
				fmt.Println()
				printStructure(valueField.Interface(), level+1)
			} else {
				fmt.Println(valueField)
			}
		}
	}

	printStructure(options, 0)
}
