// Code generated by bebopc-go; DO NOT EDIT.
package alternatives

import (
	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
	"io"
)

var _ bebop.Record = &Foo{}

//
//type Foo struct {
//	Make string
//	Model string
//	Badge string
//	Price int16
//	Wheels uint8
//	EngineCC uint16
//}

func (bbp Foo) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.Make)))
	copy(buf[at+4:at+4+len(bbp.Make)], []byte(bbp.Make))
	at += 4 + len(bbp.Make)
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.Model)))
	copy(buf[at+4:at+4+len(bbp.Model)], []byte(bbp.Model))
	at += 4 + len(bbp.Model)
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.Badge)))
	copy(buf[at+4:at+4+len(bbp.Badge)], []byte(bbp.Badge))
	at += 4 + len(bbp.Badge)
	iohelp.WriteInt16Bytes(buf[at:], bbp.Price)
	at += 2
	iohelp.WriteUint8Bytes(buf[at:], bbp.Wheels)
	at += 1
	iohelp.WriteUint16Bytes(buf[at:], bbp.EngineCC)
	at += 2
	return at
}

func (bbp *Foo) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	bbp.Make, err = iohelp.ReadStringBytes(buf[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Make)
	bbp.Model, err = iohelp.ReadStringBytes(buf[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Model)
	bbp.Badge, err = iohelp.ReadStringBytes(buf[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Badge)
	if len(buf[at:]) < 2 {
		return io.ErrUnexpectedEOF
	}
	bbp.Price = iohelp.ReadInt16Bytes(buf[at:])
	at += 2
	if len(buf[at:]) < 1 {
		return io.ErrUnexpectedEOF
	}
	bbp.Wheels = iohelp.ReadUint8Bytes(buf[at:])
	at += 1
	if len(buf[at:]) < 2 {
		return io.ErrUnexpectedEOF
	}
	bbp.EngineCC = iohelp.ReadUint16Bytes(buf[at:])
	at += 2
	return nil
}

func (bbp Foo) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.Make)))
	w.Write([]byte(bbp.Make))
	iohelp.WriteUint32(w, uint32(len(bbp.Model)))
	w.Write([]byte(bbp.Model))
	iohelp.WriteUint32(w, uint32(len(bbp.Badge)))
	w.Write([]byte(bbp.Badge))
	iohelp.WriteInt16(w, bbp.Price)
	iohelp.WriteUint8(w, bbp.Wheels)
	iohelp.WriteUint16(w, bbp.EngineCC)
	return w.Err
}

func (bbp *Foo) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.Make = iohelp.ReadString(r)
	bbp.Model = iohelp.ReadString(r)
	bbp.Badge = iohelp.ReadString(r)
	bbp.Price = iohelp.ReadInt16(r)
	bbp.Wheels = iohelp.ReadUint8(r)
	bbp.EngineCC = iohelp.ReadUint16(r)
	return r.Err
}

func (bbp Foo) Size() int {
	bodyLen := 0
	bodyLen += 4 + len(bbp.Make)
	bodyLen += 4 + len(bbp.Model)
	bodyLen += 4 + len(bbp.Badge)
	bodyLen += 2
	bodyLen += 1
	bodyLen += 2
	return bodyLen
}

func (bbp Foo) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

//func MakeFoo(r iohelp.ErrorReader) (Foo, error) {
//	v := Foo{}
//	err := v.DecodeBebop(r)
//	return v, err
//}

func MakeFooFromBytes(buf []byte) (Foo, error) {
	v := Foo{}
	err := v.UnmarshalBebop(buf)
	return v, err
}
