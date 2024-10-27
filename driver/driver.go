package strawhouse

import "github.com/bsthun/gut"

type Driver struct {
	Signature Signaturer
	Client    Clienter
}

type Option struct {
	Server string `validate:"required"`
	Key    string `validate:"required"`
	Secure bool   `validate:"omitempty"`
}

func New(option *Option) (*Driver, error) {
	if err := gut.Validator.Struct(option); err != nil {
		return nil, gut.Err(false, "invalid option", err)
	}

	sgn := NewSignature(option.Key)
	cnt := NewClient(option)

	return &Driver{
		Signature: sgn,
		Client:    cnt,
	}, nil
}

func (r *Driver) Close() {
	_ = r.Client.Close()
}
