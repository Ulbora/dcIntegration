package delegates

import (
	"fmt"
	"testing"
)

func TestGetDcConfigs(t *testing.T) {
	var cf DcConfigFileDelegate
	var conf DcConfigFiles
	cf = &conf
	res := cf.GetDcConfigs("../confFileTest")
	fmt.Println("GetDcConfigs rtn: ", res)
	fmt.Println("GetDcConfigs rtn fields: ", (*res)["test1"].Fields)
}
