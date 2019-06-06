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

//Elem Elem
type Elem struct {
	ColumnName string
	Value      string
}

//BuildDcCartFiles BuildDcCartFiles
func BuildDcCartFiles(supdir string, dcartdir string, confdir string) {
	var cf DcConfigFileDelegate
	var conf DcConfigFiles
	cf = &conf
	cfiles := cf.GetDcConfigs(confdir)
	fmt.Println("cfiles: ", cfiles)

	var b sfb.Builder
	var csvb sfb.CsvFileBuilder
	csvb.OutputDir = dcartdir
	b = &csvb
	files := b.ReadAllSupplierDirs(supdir)
	for _, filed := range *files {
		fmt.Println("filed: ", filed)
		fmt.Println("spdir: ", filed.Name)
		for _, file := range filed.Files {
			fmt.Println("file name: ", file.Name)
			fmt.Println("file full name: ", file.FullName)
			fcont := b.ReadSourceFile(file.FullName)
			//fmt.Println("sup file: ", fcont.)
			fmt.Println("fcont len: ", len(fcont))
			cfil := (*cfiles)[filed.Name]
			dccont := buildCartFile(&fcont, &cfil)
			fmt.Println("dccont len: ", len(*dccont))
			var df sfb.CartCsvFile
			df.FileName = file.Name
			df.Content = *dccont
			b.SaveCartFile(df)

		}
	}
	//fmt.Println("sf files: ", files)
	//fmt.Println("sf file dir", sfdir)
	//fmt.Println("dc file dir", dcdir)
}

func buildCartFile(sourceFile *[][]string, conf *ConfFile) *[][]string {
	var rtn [][]string
	var dccol = []string{"distributor", "id", "mfgid", "name", "manufacturer", "categories", "cost", "price", "price2", "stock", "weight", "free_shipping", "date_created", "description", "extended_description", "keywords", "hide", "sorting", "thumbnail", "image1", "image2", "image3", "image4", "related", "distributor", "shipcost", "homespecial", "categoryspecial", "title", "metatags"}
	rtn = append(rtn, dccol)
	// var dcrow []map[string]string
	var scol []string
	for c, row := range *sourceFile {
		if c == 0 {
			fmt.Println("row c : ", c)
			scol = row
			fmt.Println("header row: ", scol)
		} else {
			var dcrow []map[string]string
			var elemMap = make(map[string]string)
			for cc, elem := range row {
				//fmt.Println("elem c : ", cc)
				//var e Elem
				//e.ColumnName = scol[cc]
				//e.Value = elem
				//fmt.Println("Elem : ", e)
				elemMap[scol[cc]] = elem
				//fmt.Println("elemMap : ", elemMap)
				dcrow = append(dcrow, elemMap)
			}
			var dcvr []string
			dcvr = append(dcvr, conf.Distributor)
			fmt.Println("elemMap : ", elemMap)
			for colc, dck := range dccol {
				if colc == 0 {
					continue
				} else {
					dce := (*conf.Fields)[dck]
					if dce.CartKey != "" {
						fmt.Println("dce : ", dce)
						fcnt := elemMap[dce.SpfKey]
						fmt.Println("fcnt : ", fcnt)
						if dce.Required && fcnt == "" {
							continue
						} else {
							var cont = ""
							if dce.Prefix != "" {
								cont = dce.Prefix
							}
							if len(dce.SpfSubKeys) > 0 {
								fcnt2 := elemMap[dce.SpfSubKeys[0]]
								cont += (fcnt + "/" + fcnt2)
							} else {
								cont += fcnt
							}
							dcvr = append(dcvr, cont)
						}
					} else {
						dcvr = append(dcvr, "")
					}
				}
				//dcvr = append(dcvr, dce.)
			}
			rtn = append(rtn, dcvr)

		}
		//fmt.Println("dcrow : ", dcrow)
		//fmt.Println("col : ", scol)
		//fmt.Println("dccol : ", dccol)
		//fmt.Println("row : ", row)

	}
	fmt.Println("cart file : ", rtn)
	return &rtn
}
