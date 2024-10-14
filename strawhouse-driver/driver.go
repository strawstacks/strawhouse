package strawhouse

type Driver struct {
	Signature Signaturer
	Client    Clienter
}

type Option struct {
	Secure bool
}

func New(key string, server string, option *Option) *Driver {
	sgn := NewSignature(key)
	cnt := NewClient(key, server, option)

	return &Driver{
		Signature: sgn,
		Client:    cnt,
	}
}

func (r *Driver) Close() {
	_ = r.Client.Close()
}
