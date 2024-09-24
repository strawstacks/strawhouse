package main

import (
	"backend/type/enum"
	"backend/util/signature"
	"bytes"
	"encoding/gob"
	uu "github.com/bsthun/goutils"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func main() {
	sign := signature.New("S+XTcv93IgYZFsVxU/WBUXDC66imggy6MEDUS9L3TWlSYEmZJDD0k41WH82ShZF4")
	attribute := &signature.ExampleAttribute{
		UploaderId:  uu.Ptr[uint64](20),
		SessionName: uu.Ptr("abcd"),
	}
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	gob.Register(signature.ExampleAttribute{})
	_ = enc.Encode(attribute)

	token := sign.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionGet, 2, time.Now().Add(20*time.Minute), "/photo/2024", buffer.Bytes())

	spew.Dump(token)
}
