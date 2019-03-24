package tests

import (
	"github.com/saichler/utils/golang"
	"testing"
)

func TestTrim(t *testing.T) {
	input:="\r\n  \n\rHello World \r\n \n\r"
	output:=utils.Trim(input)
	if output!="Hello World" {
		t.Fail()
	}
}
