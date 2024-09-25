package main

import (
	"backend/type/enum"
	"backend/util/signature"
	"bytes"
	"encoding/base64"
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
	gob.Register(new(signature.ExampleAttribute))
	_ = enc.Encode(attribute)

	token := sign.Generate(1, enum.SignatureModeDirectory, enum.SignatureActionUpload, 1, time.Now().Add(20*time.Minute), "/photo/2024", buffer.Bytes())

	// Encode attribute to base64
	base64Attr := base64.StdEncoding.EncodeToString(buffer.Bytes())
	signature.ReplaceChar(&base64Attr, '+', '*')
	spew.Dump(token, base64Attr)
}
