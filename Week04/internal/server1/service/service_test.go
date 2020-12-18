package service

import (
	"testing"
)

func TestService(t *testing.T) {
	d := NewGetTestDataService(nil)
	d.GetData(nil,nil)
}
