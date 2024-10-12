package strawhouse

type Driver struct {
	Signature Signaturer
	Client    Clienter
}

func New(key string, server string) *Driver {
	sgn := NewSignature(key)
	cnt := NewClient(key, server)

	return &Driver{
		Signature: sgn,
		Client:    cnt,
	}
}

func (r *Driver) Close() {
	_ = r.Client.Close()
}
