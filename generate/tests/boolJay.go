// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package tests

import "github.com/speedyhoon/jay"

func (o *OnlyBools) MarshalJ() (b []byte) {
	b = make([]byte, 1)
	b[0] = jay.Bool5(o.Hidden, o.Deactivated, o.Selected, o.Modified, o.Updated)
	return
}

func (o *OnlyBools) UnmarshalJ(b []byte) error {
	if len(b) < 1 {
		return jay.ErrUnexpectedEOB
	}
	o.Hidden, o.Deactivated, o.Selected, o.Modified, o.Updated = jay.ReadBool5(b[0])
	return nil
}
