// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay.

package alternatives

import (
	"github.com/200sc/bebop/iohelp"
	"github.com/speedyhoon/jay"
)

func (f *Foo) MarshalJ() (b []byte) {
	l0, l1, l2 := len(f.Make), len(f.Model), len(f.Badge)
	b = make([]byte, 8+l0+l1+l2)
	jay.WriteInt16(b[0:2], f.Price)
	b[2] = f.Wheels
	jay.WriteUint16(b[3:5], f.EngineCC)
	at := jay.WriteStringAt(b[5:], f.Make, l0, 5)
	at = jay.WriteStringAt(b[at:], f.Model, l1, at)
	jay.WriteString(b[at:], f.Badge, l2)

	return
}

const strSize = 1

func (f *Foo) MarshalJ2() (b []byte) {
	l0, l1, l2 := len(f.Make), len(f.Model), len(f.Badge)
	b = make([]byte, 8+l0+l1+l2)
	jay.WriteInt16(b[0:2], f.Price)
	b[2] = f.Wheels
	jay.WriteUint16(b[3:5], f.EngineCC)
	at := jay.WriteStringAt(b[5:5+strSize+l0], f.Make, l0, 5)
	at = jay.WriteStringAt(b[at:at+strSize+l1], f.Model, l1, at)
	jay.WriteString2(b[at:], f.Badge, l2)
	return
}

func (f *Foo) UnmarshalJ(b []byte) error {
	length := len(b)
	if length < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.Price = jay.ReadInt16(b[0:2])
	f.Wheels = b[2]
	f.EngineCC = jay.ReadUint16(b[3:5])

	var ok bool
	at := 5
	f.Model, at, ok = jay.ReadStringAt(b[at:], at)
	if !ok {
		return jay.ErrUnexpectedEOB
	}

	//var at int
	//size := int(b[5])
	//end := at + size + 1
	//if size != 0 {
	//	if end > length {
	//		return jay.ErrUnexpectedEOB
	//	}
	//	if length >= end {
	//		f.Make = string(b[at:end])
	//		//, at, ok = jay.ReadStringAt(b[at:], at)
	//	} else {
	//		return jay.ErrUnexpectedEOB
	//	}
	//}
	//at = end
	//size = int(b[end])
	//end = at + size + 1
	//if size != 0 {
	//	if length >= end {
	//		f.Model = string(b[at:end])
	//	} else {
	//		return jay.ErrUnexpectedEOB
	//	}
	//}
	//
	//at = end
	//size = int(b[end])
	//end = at + size + 1
	//if size != 0 {
	//	if length >= end {
	//		f.Badge = string(b[at:])
	//	} else {
	//		return jay.ErrUnexpectedEOB
	//	}
	//}

	//size = int(b[0]) + strSizeOf
	//if size == strSizeOf || len(b) < size {
	//	return
	//}

	//return string(b[strSizeOf:size]), size, true

	//size = b[at]
	//if size != 0 {
	f.Model, at, ok = jay.ReadStringAt(b[at:], at)
	if !ok {
		return jay.ErrUnexpectedEOB
	}
	//}
	f.Badge, at, ok = jay.ReadString(b[at:])
	if !ok {
		return jay.ErrUnexpectedEOB
	}

	return nil
}
func (f *Foo) UnmarshalJ2(b []byte) error {
	length := len(b)
	if length < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.Price = jay.ReadInt16(b[0:2])
	f.Wheels = b[2]
	f.EngineCC = jay.ReadUint16(b[3:5])
	//var ok bool
	//var at int
	at := 5
	size := int(b[5])
	end := at + size + 1
	if size != 0 {
		if end > length {
			return jay.ErrUnexpectedEOB
		}
		if length >= end {
			f.Make = string(b[at:end])
			//, at, ok = jay.ReadStringAt(b[at:], at)
		} else {
			return jay.ErrUnexpectedEOB
		}
	}
	at = end
	size = int(b[end])
	end = at + size + 1
	if size != 0 {
		if length >= end {
			f.Model = string(b[at:end])
		} else {
			return jay.ErrUnexpectedEOB
		}
	}

	at = end
	size = int(b[end])
	end = at + size + 1
	if size != 0 {
		if length >= end {
			f.Badge = string(b[at:])
		} else {
			return jay.ErrUnexpectedEOB
		}
	}

	//size = int(b[0]) + strSizeOf
	//if size == strSizeOf || len(b) < size {
	//	return
	//}

	//return string(b[strSizeOf:size]), size, true

	//size = b[at]
	//if size != 0 {
	//	f.Model, at, ok = jay.ReadStringAt(b[at:], at)
	//	if !ok {
	//		return jay.ErrUnexpectedEOB
	//	}
	//}
	//f.Badge, at, ok = jay.ReadString(b[at:])
	//if !ok {
	//	return jay.ErrUnexpectedEOB
	//}

	return nil
}

func (f *Foo) UnmarshalJ3(b []byte) error {
	length := len(b)
	if length < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.Price = jay.ReadInt16(b[0:2])
	f.Wheels = b[2]
	f.EngineCC = jay.ReadUint16(b[3:5])
	//var ok bool
	//var at int
	at := 5
	size := int(b[5])
	end := at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	if size != 0 {
		f.Make = string(b[at:end])
	}
	at = end
	size = int(b[end])
	end = at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	if size != 0 {
		f.Model = string(b[at:end])
	}

	at = end
	size = int(b[end])
	end = at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	if size != 0 {
		f.Badge = string(b[at:])
	}

	return nil
}

func (f *Foo) UnmarshalJ4(b []byte) error {
	length := len(b)
	if length < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.Price = jay.ReadInt16(b[0:2])
	f.Wheels = b[2]
	f.EngineCC = jay.ReadUint16(b[3:5])

	at := 5
	size := int(b[5])
	end := at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	f.Make = string(b[at:end])

	at = end
	size = int(b[end])
	end = at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	f.Model = string(b[at:end])

	at = end
	size = int(b[end])
	end = at + size + 1
	if end > length {
		return jay.ErrUnexpectedEOB
	}
	f.Badge = string(b[at:])

	return nil
}

func (f *Foo) UnmarshalJ5(b []byte) error {
	length := len(b)
	if length < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.Price = jay.ReadInt16(b[0:2])
	f.Wheels = b[2]
	f.EngineCC = jay.ReadUint16(b[3:5])

	at := 5
	sz := int(b[5])
	end := at + sz
	if length < end+2 {
		return jay.ErrUnexpectedEOB
	}
	f.Make = string(b[at:end])
	/*	if len(buf) < 4 {
			return "", io.ErrUnexpectedEOF
		}
		sz := ReadUint32Bytes(buf)
		if len(buf) < int(sz)+4 {
			return "", io.ErrUnexpectedEOF
		}
		return string(buf[4 : 4+sz]), nil*/
	bbp.Make, err = iohelp.ReadStringBytes(b[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Make)
	bbp.Model, err = iohelp.ReadStringBytes(b[at:])
	if err != nil {
		return err
	}
	at += 4 + len(bbp.Model)
	bbp.Badge, err = iohelp.ReadStringBytes(b[at:])
	if err != nil {
		return err
	}
	//at += 4 + len(bbp.Badge)

	//if len(b[at:]) < 2 {
	//	return io.ErrUnexpectedEOF
	//}
	//bbp.Price = iohelp.ReadInt16Bytes(b[at:])
	//at += 2
	//
	//
	//if len(b[at:]) < 1 {
	//	return io.ErrUnexpectedEOF
	//}
	//bbp.Wheels = iohelp.ReadUint8Bytes(b[at:])
	//at += 1
	//
	//
	//if len(b[at:]) < 2 {
	//	return io.ErrUnexpectedEOF
	//}
	//bbp.EngineCC = iohelp.ReadUint16Bytes(b[at:])
}