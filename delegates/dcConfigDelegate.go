package delegates

import (
	sfc "github.com/Ulbora/intgConvConfig"
)

//DcConfigFileDelegate DcConfigFileDelegate
type DcConfigFileDelegate interface {
	GetDcConfigs(dir string) *map[string]ConfFile
}

//DcConfigFiles DcConfigFiles
type DcConfigFiles struct {
}

//ConfFile ConfFile
type ConfFile struct {
	Directory   string
	Distributor string
	Fields      *map[string]sfc.SpfField
}

//GetDcConfigs GetDcConfigs
func (dcc *DcConfigFiles) GetDcConfigs(dir string) *map[string]ConfFile {
	var rtn = make(map[string]ConfFile)
	var c sfc.SpfConvConf
	var jc sfc.JSONConvConf
	c = &jc
	conf := c.GetSpfConversion(dir)
	for _, v := range *conf {
		//fmt.Printf("key[%s] value[%s]\n", k, v)
		var f ConfFile
		f.Directory = v.CfDirectory
		f.Distributor = v.Distributor
		f.Fields = buildFields(v.SpfFields)
		rtn[v.CfDirectory] = f
	}
	//fmt.Println("conf: ", conf)
	// fmt.Println("GetDcConfigs rtn: ", rtn)
	// fmt.Println("GetDcConfigs rtn fields: ", rtn["test1"].Fields)
	return &rtn
}

func buildFields(ff []sfc.SpfField) *map[string]sfc.SpfField {
	var rtn = make(map[string]sfc.SpfField)
	//var msrp1Found = false
	for _, f := range ff {
		// if f.SpfKey == "MSRP" && !msrp1Found {
		// 	var key = "MSRP1"
		// 	rtn[key] = f
		// 	msrp1Found = true
		// } else if f.SpfKey == "MSRP" && msrp1Found {
		// 	var key = "MSRP2"
		// 	rtn[key] = f
		// } else {
		rtn[f.CartKey] = f
		//}
	}
	return &rtn
}
