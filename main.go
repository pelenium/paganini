package main

import (
	"flag"
	"fmt"
)

func main() {
	tn := flag.String("tn", "template-name", "Name of project template")
	pn := flag.String("pn", "project-name", "Name of your project")
	flag.PrintDefaults()
	flag.Parse()

	fmt.Printf("\ntarget name: %s\nproject name: %s\n", *tn, *pn)
}
