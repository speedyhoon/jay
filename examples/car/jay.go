// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (c *Car) MarshalJ() (b []byte) {
	l0, l1, l2, l3 := len(c.Name), len(c.CC), len(c.Gearbox.Model), len(c.Gearbox.Manufacturer)
	b = make([]byte, 41+l0+l1+l2+l3)
	b[0] = jay.Bool3(c.Auto, c.Gearbox.Sequential, c.Gearbox.Automatic)
	jay.WriteUint64(b[1:9], c.ID)
	jay.WriteUintArch64(b[9:17], c.Row)
	jay.WriteUint16(b[17:19], c.RedLine)
	jay.WriteTime(b[19:27], c.Expiry)
	jay.WriteIntArch64(b[27:35], c.Gearbox.Gears)
	b[35] = c.Gearbox.Reverse
	b[36] = byte(c.Gearbox.LinkageDelta)
	at := jay.WriteStringAt(b[37:], c.Name, l0, 37)
	at = jay.WriteStringAt(b[at:], c.CC, l1, at)
	at = jay.WriteStringAt(b[at:], c.Gearbox.Model, l2, at)
	jay.WriteString(b[at:], c.Gearbox.Manufacturer, l3)
	return
}

func (c *Car) UnmarshalJ(b []byte) error {
	c.Auto, c.Gearbox.Sequential, c.Gearbox.Automatic = jay.ReadBool3(b[0])
	c.ID = jay.ReadUint64(b[1:9])
	c.Row = jay.ReadUintArch64(b[9:17])
	c.RedLine = jay.ReadUint16(b[17:19])
	c.Expiry = jay.ReadTime(b[19:27])
	c.Gearbox.Gears = jay.ReadIntArch64(b[27:35])
	c.Gearbox.Reverse = b[35]
	c.Gearbox.LinkageDelta = int8(b[36])
	var ok bool
	at := 37
	c.Name, at, ok = jay.ReadStringAt(b[at:], at)
	if !ok {
		return jay.ErrUnexpectedEOB
	}
	c.CC, at, ok = jay.ReadStringAt(b[at:], at)
	if !ok {
		return jay.ErrUnexpectedEOB
	}
	c.Gearbox.Model, at, ok = jay.ReadStringAt(b[at:], at)
	if !ok {
		return jay.ErrUnexpectedEOB
	}
	return jay.ReadStringPtrErr(b[at:], &c.Gearbox.Manufacturer)
}
