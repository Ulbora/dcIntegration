package delegates

import (
	"fmt"
	"strconv"

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

//DcCartDelegate DcCartDelegate
type DcCartDelegate interface {
	BuildDcCartFiles(supdir string, dcartdir string, confdir string, clean bool)
}

//Elem Elem
type Elem struct {
	ColumnName string
	Value      string
}

//DcCartFileDelegate DcCartFileDelegate
type DcCartFileDelegate struct {
}

//BuildDcCartFiles BuildDcCartFiles
func (d *DcCartFileDelegate) BuildDcCartFiles(supdir string, dcartdir string, confdir string, clean bool) {
	var cf DcConfigFileDelegate
	var conf DcConfigFiles
	cf = &conf
	cfiles := cf.GetDcConfigs(confdir)
	//fmt.Println("cfiles: ", cfiles)

	var b sfb.Builder
	var csvb sfb.CsvFileBuilder
	csvb.OutputDir = dcartdir
	b = &csvb
	files := b.ReadAllSupplierDirs(supdir)
	for _, filed := range *files {
		//fmt.Println("filed: ", filed)
		//fmt.Println("spdir: ", filed.Name)
		for _, file := range filed.Files {
			//fmt.Println("file name: ", file.Name)
			//fmt.Println("file full name: ", file.FullName)
			fcont := b.ReadSourceFile(file.FullName, clean)
			//fmt.Println("sup file: ", fcont)
			//fmt.Println("fcont len: ", len(fcont))
			cfil := (*cfiles)[filed.Name]
			dccont := buildCartFile(&fcont, &cfil)
			//fmt.Println("dccont len: ", len(*dccont))
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
	var dccol = conf.CartHeader // []string{"distributor", "id", "mfgid", "name", "manufacturer", "categories", "cost", "price", "price2", "stock", "weight", "description", "extended_description", "thumbnail", "image1"}
	//fmt.Println("dccol : ", dccol)
	rtn = append(rtn, dccol)
	// var dcrow []map[string]string
	var scol []string
	for c, row := range *sourceFile {
		if c == 0 {
			//fmt.Println("row c : ", c)
			scol = row
			//fmt.Println("header row: ", scol)
		} else {
			//var dcrow []map[string]string
			var elemMap = make(map[string]string)
			for cc, elem := range row {
				//fmt.Println("elem c : ", cc)
				//var e Elem
				//e.ColumnName = scol[cc]
				//e.Value = elem
				//fmt.Println("Elem : ", e)
				elemMap[scol[cc]] = elem
				//fmt.Println("elemMap : ", elemMap)
				//dcrow = append(dcrow, elemMap)
			}
			var dcvr []string
			dcvr = append(dcvr, conf.Distributor)
			//fmt.Println("elemMap : ", elemMap)
			var foundErr = false
			//fmt.Println("dccol : ", dccol)
			for colc, dck := range dccol {
				//fmt.Println("dccol : ", dccol)
				//fmt.Println("dck : ", dck)
				//fmt.Println("colc : ", colc)
				if colc == 0 {
					continue
				} else {
					dce := (*conf.Fields)[dck]
					//fmt.Println("dce : ", dce)
					//fmt.Println("CartKey : ", dce.CartKey)
					if dce.CartKey != "" {
						//fmt.Println("dce : ", dce)
						//fmt.Println("supply Key : ", dce.CartKey)
						fcnt := elemMap[dce.SpfKey]
						//fmt.Println("dce.SpfKey: ", dce.SpfKey)
						//fmt.Println("fcnt: ", fcnt)
						if dce.CartKey == "price" {
							if fcnt == "" || fcnt == "0" {
								fcnt = elemMap[dce.SpfAltKey]
								cst := elemMap[dce.SpfCost]
								factor := dce.CartPriceFactor
								// fmt.Println("fcnt : ", fcnt)
								//fmt.Println("factor : ", factor)
								if msrp, err := strconv.ParseFloat(fcnt, 64); err == nil {
									//fmt.Println("msrp before 0: ", msrp)
									if mulFact, err2 := strconv.ParseFloat(factor, 64); err2 == nil {
										//fmt.Println("msrp before 0 after factor: ", msrp)
										if cost, err3 := strconv.ParseFloat(cst, 64); err3 == nil {
											//fmt.Println("msrp before test for 0: ", msrp)
											if msrp != 0 {
												msrp = msrp * mulFact
												//fmt.Println("msrp : ", msrp)
												//fmt.Println("cost : ", cost+2)
												if msrp > cost+2 {
													fcnt = fmt.Sprintf("%.2f", msrp)
													//fmt.Println("msrp %: ", fcnt)
												}
											} else {
												// cost = elemMap[dce.SpfAltKey2]
												markuFactor := dce.CostMarkupFactor
												if mupFact, err2 := strconv.ParseFloat(markuFactor, 64); err2 == nil {
													prc := cost * mupFact
													fcnt = fmt.Sprintf("%.2f", prc)
													//fmt.Println("price from cost %: ", fcnt)
												}
											}

										}

									}
								}
							}
						}
						if dce.CartKey == "stock" {
							if _, err := strconv.Atoi(fcnt); err != nil {
								foundErr = true
								break
							}
						} else if dce.CartKey == "cost" || dce.CartKey == "price" || dce.CartKey == "price2" || dce.CartKey == "weight" {
							if _, err := strconv.ParseFloat(fcnt, 64); err != nil {
								foundErr = true
								break
							}
						}
						if dce.CartKey == "weight" {
							if wtg, err := strconv.Atoi(fcnt); err == nil {
								if wtg == 0 {
									fcnt = "1"
								}
							}
						}
						if dce.CartKey == "price2" {
							if price2, err := strconv.ParseFloat(fcnt, 64); err == nil {
								if price2 == 0 {
									foundErr = true
									break
								}
							}
						}
						//fmt.Println("elem found : ", dce.SpfKey, " ", fcnt)
						//fmt.Println("dce.CartKey", dce.CartKey, fcnt)
						if dce.Required && fcnt == "" {
							//fmt.Println("required cont missing : ", fcnt)
							foundErr = true
							break
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
							//fmt.Println("row ready to be appended : ", dcvr)
						}
					}
					// else {
					// 	dcvr = append(dcvr, "")
					// }
				}
				//dcvr = append(dcvr, dce.)
			}
			//fmt.Println("foundErr : ", foundErr)
			//fmt.Println("row ready to be added : ", dcvr)
			if foundErr {
				continue
			}
			rtn = append(rtn, dcvr)

		}
		//fmt.Println("dcrow : ", dcrow)
		//fmt.Println("col : ", scol)
		//fmt.Println("dccol : ", dccol)
		//fmt.Println("row : ", row)

	}
	//fmt.Println("cart file : ", rtn)
	return &rtn
}
