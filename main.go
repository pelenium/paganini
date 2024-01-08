package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
*	tn - template name
*	pn - project name
*	tf - template folder
 */

var basicJson string = `{"template-folder": ""}`

func main() {
	var (
		tn = flag.String("tn", "", "Name of project template")
		pn = flag.String("pn", "", "Name of your project")
		tf = flag.String("tf", "", "Set root directory with your templates (absolute path)")
	)
	flag.Lookup("help")
	flag.Lookup("h")

	flag.Parse()

	if *tn != "" {
	rep:
		var file *os.File
		dir, err := os.Executable()
		if err != nil {
			panic(err)
		}
		fmt.Println(dir)
		dir = dir[0:max(strings.LastIndex(dir, `\`), strings.LastIndex(dir, `/`))];		
		path := dir + `\settings.json`
		fmt.Println(path)

		bytes, err := os.ReadFile(path)
		if err != nil {
			file, err = os.Create(path)
			if err != nil {
				printError("cannot create \"settings.json\"")
			}
			file.WriteString(basicJson)
			goto rep
		}
		var data map[string]any
		json.Unmarshal(bytes, &data)

		data["template-folder"] = *tn

		res, _ := json.Marshal(data)
		fmt.Println(string(res))

		if err := os.WriteFile(path, res, 0666); err != nil {
			printError("cannot write into \"settings.json\"")
		}
	}

	fmt.Printf("\ntarget name: %s\nproject name: %s\ntemplate folder: %s\n", *tn, *pn, *tf)
}

func printError(errorData string) {
	fmt.Println("paga error: " + errorData)
	os.Exit(1)
}
