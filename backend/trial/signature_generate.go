package main

import (
	"backend/util/signature"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func main() {
	sign := signature.New("S+XTcv93IgYZFsVxU/WBUXDC66imggy6MEDUS9L3TWlSYEmZJDD0k41WH82ShZF4")
	token := sign.Generate(1, 2, 0, time.Now().Add(20*time.Minute), "/photo/test1")
	spew.Dump(token)
}
