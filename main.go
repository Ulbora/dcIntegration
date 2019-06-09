package main

import (
	"flag"
	"fmt"
	d "github.com/Ulbora/dcIntegration/delegates"
	"os"
)

func main() {
	var supDir string
	var dcDir string
	var confDir string
	if len(os.Args) >= 4 {
		cFlag := flag.Bool("c", false, "clean csv file")
		flag.Parse()
		fmt.Println("flag :", *cFlag)
		if !*cFlag {
			supDir = os.Args[1]
			dcDir = os.Args[2]
			confDir = os.Args[3]
		} else {
			supDir = os.Args[2]
			dcDir = os.Args[3]
			confDir = os.Args[4]
		}

		fmt.Println("supDir: " + supDir)
		fmt.Println("dcDir: " + dcDir)
		fmt.Println("confDir: " + confDir)
		var del d.DcCartDelegate
		var fdel d.DcCartFileDelegate
		del = &fdel
		del.BuildDcCartFiles(supDir, dcDir, confDir, *cFlag)
		fmt.Println("Created dc files at " + dcDir)
	} else {
		fmt.Println("No params sent at command line")
	}

}
