package delegates

import "testing"

func TestBuildDcCartFiles(t *testing.T) {
	BuildDcCartFiles("../sfFileTest", "../cartFileTest", "../confFileTest")
}
