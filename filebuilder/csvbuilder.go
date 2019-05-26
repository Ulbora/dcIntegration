package filebuilder

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//ParseFiles ParseFiles
func ParseFiles(dir string) {
	fmt.Println(dir)
}
func readFileDir(dir string) *[]SupplierDir {
	var rtn []SupplierDir
	fmt.Println("dir: ", dir)
	files, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, file := range files {
			fmt.Println("file name: ", file.Name())
			if file.IsDir() {
				var sd SupplierDir
				sd.Name = file.Name()
				rtn = append(rtn, sd)
			}
		}
		for c, spd := range rtn {
			fmt.Println("spd: ", spd)
			var sdirname = dir + string(filepath.Separator) + spd.Name
			fmt.Println("dir name: ", sdirname)
			sfiles, err := ioutil.ReadDir(sdirname)
			if err == nil {
				for _, sfile := range sfiles {
					if !sfile.IsDir() {
						//fmt.Println("sfile: ", sfile)
						spd.Files = append(spd.Files, sfile.Name())
					}
				}
			}
			rtn[c].Files = spd.Files
			fmt.Println("spd: ", spd)
		}
		fmt.Println("rtn: ", rtn)
	}
	return &rtn
}
