package delegates

import "testing"

func TestBuildDcCartFiles(t *testing.T) {
	var del DcCartDelegate
	var fdel DcCartFileDelegate
	del = &fdel
	del.BuildDcCartFiles("../sfFileTest", "../cartFileTest", "../confFileTest")
}
