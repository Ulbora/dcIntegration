package delegates

import (
	"fmt"

	sfb "github.com/Ulbora/intgFileBuilder"
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

//DcCartFileDelegate DcCartFileDelegate
type DcCartFileDelegate interface {
	BuildDcCartFiles()
}

//BuildDcCartFiles BuildDcCartFiles
func BuildDcCartFiles(sfdir string, dcdir string) {
	var b sfb.Builder
	var csvb sfb.CsvFileBuilder
	b = &csvb
	files := b.ReadAllSupplierDirs("../sfFileTest")
	for _, filed := range *files {
		fmt.Println("filed: ", filed)
		fmt.Println("spdir: ", filed.Name)
		for _, file := range filed.Files {
			fmt.Println("file full name: ", file.FullName)
			fcont := b.ReadSourceFile(file.FullName)
			//fmt.Println("sup file: ", fcont.)
			fmt.Println("len: ", len(fcont))
		}
	}
	//fmt.Println("sf files: ", files)
	//fmt.Println("sf file dir", sfdir)
	//fmt.Println("dc file dir", dcdir)
}
