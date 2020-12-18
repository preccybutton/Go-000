package init

import (
	"testing"
)

func TestInit(t *testing.T)  {
	pool, _, _ := InitStart("aaa")
	pool.Server().Serve(pool.Lis())
}
