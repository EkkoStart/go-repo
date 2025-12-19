package main

import (
	"github.com/EkkoStart/go-repo/werrors"
	"github.com/EkkoStart/go-repo/wlog"
)

func main() {

	err := werrors.WithCode(101101, "unreconized Authorization header")
	wlog.Errorf("%v", err)
}
