package strawhouse

import "backend/util/signature"

type Driver struct {
	Server    string
	Signature *signature.Signature
}

func New(key string, server string) *Driver {
	sgn := signature.New(key)
	return &Driver{
		Server:    server,
		Signature: sgn,
	}
}
