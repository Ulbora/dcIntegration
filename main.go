package main

import (
	"fmt"
	d "github.com/Ulbora/dcIntegration/delegates"
	"os"
)

func main() {
	var supDir string
	var dcDir string
	var confDir string
	if len(os.Args) >= 4 {
		supDir = os.Args[1]
		dcDir = os.Args[2]
		confDir = os.Args[3]
		fmt.Println("supDir: " + supDir)
		fmt.Println("dcDir: " + dcDir)
		fmt.Println("confDir: " + confDir)
		var del d.DcCartDelegate
		var fdel d.DcCartFileDelegate
		del = &fdel
		del.BuildDcCartFiles(supDir, dcDir, confDir)
		fmt.Println("Created dc files at " + dcDir)
	} else {
		fmt.Println("No params sent at command line")
	}

}
