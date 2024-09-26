package strawhouse

type Driver struct {
	Server    string
	Signature *Signature
}

func New(key string, server string) *Driver {
	sgn := NewSignature(key)
	return &Driver{
		Server:    server,
		Signature: sgn,
	}
}
