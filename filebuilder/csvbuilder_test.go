package filebuilder

import (
	//fb "github.com/Ulbora/dcIntegration/filebuilder"
	"testing"
)

func TestParseFiles(t *testing.T) {
	ParseFiles("testdir")
}
func TestReadFileDir(t *testing.T) {
	res := readFileDir("../testdir")
	if (*res)[0].Name != "test1" {
		t.Fail()
	}
	if (*res)[0].Files[0] != "test11.csv" {
		t.Fail()
	}
}
