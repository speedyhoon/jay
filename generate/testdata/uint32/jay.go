// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 4)
	jay.WriteUint32(b, o.One)
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 4 {
		return jay.ErrUnexpectedEOB
	}
	o.One = jay.ReadUint32(b)
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:], t.Two)
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:])
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 12)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:], t.Three)
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 12 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:])
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteUint32(b[:4], f.One)
	jay.WriteUint32(b[4:8], f.Two)
	jay.WriteUint32(b[8:12], f.Three)
	jay.WriteUint32(b[12:], f.Four)
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadUint32(b[:4])
	f.Two = jay.ReadUint32(b[4:8])
	f.Three = jay.ReadUint32(b[8:12])
	f.Four = jay.ReadUint32(b[12:])
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 20)
	jay.WriteUint32(b[:4], f.One)
	jay.WriteUint32(b[4:8], f.Two)
	jay.WriteUint32(b[8:12], f.Three)
	jay.WriteUint32(b[12:16], f.Four)
	jay.WriteUint32(b[16:], f.Five)
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 20 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadUint32(b[:4])
	f.Two = jay.ReadUint32(b[4:8])
	f.Three = jay.ReadUint32(b[8:12])
	f.Four = jay.ReadUint32(b[12:16])
	f.Five = jay.ReadUint32(b[16:])
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteUint32(b[:4], s.One)
	jay.WriteUint32(b[4:8], s.Two)
	jay.WriteUint32(b[8:12], s.Three)
	jay.WriteUint32(b[12:16], s.Four)
	jay.WriteUint32(b[16:20], s.Five)
	jay.WriteUint32(b[20:], s.Six)
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadUint32(b[:4])
	s.Two = jay.ReadUint32(b[4:8])
	s.Three = jay.ReadUint32(b[8:12])
	s.Four = jay.ReadUint32(b[12:16])
	s.Five = jay.ReadUint32(b[16:20])
	s.Six = jay.ReadUint32(b[20:])
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 28)
	jay.WriteUint32(b[:4], s.One)
	jay.WriteUint32(b[4:8], s.Two)
	jay.WriteUint32(b[8:12], s.Three)
	jay.WriteUint32(b[12:16], s.Four)
	jay.WriteUint32(b[16:20], s.Five)
	jay.WriteUint32(b[20:24], s.Six)
	jay.WriteUint32(b[24:], s.Seven)
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 28 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadUint32(b[:4])
	s.Two = jay.ReadUint32(b[4:8])
	s.Three = jay.ReadUint32(b[8:12])
	s.Four = jay.ReadUint32(b[12:16])
	s.Five = jay.ReadUint32(b[16:20])
	s.Six = jay.ReadUint32(b[20:24])
	s.Seven = jay.ReadUint32(b[24:])
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteUint32(b[:4], e.One)
	jay.WriteUint32(b[4:8], e.Two)
	jay.WriteUint32(b[8:12], e.Three)
	jay.WriteUint32(b[12:16], e.Four)
	jay.WriteUint32(b[16:20], e.Five)
	jay.WriteUint32(b[20:24], e.Six)
	jay.WriteUint32(b[24:28], e.Seven)
	jay.WriteUint32(b[28:], e.Eight)
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadUint32(b[:4])
	e.Two = jay.ReadUint32(b[4:8])
	e.Three = jay.ReadUint32(b[8:12])
	e.Four = jay.ReadUint32(b[12:16])
	e.Five = jay.ReadUint32(b[16:20])
	e.Six = jay.ReadUint32(b[20:24])
	e.Seven = jay.ReadUint32(b[24:28])
	e.Eight = jay.ReadUint32(b[28:])
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 36)
	jay.WriteUint32(b[:4], n.One)
	jay.WriteUint32(b[4:8], n.Two)
	jay.WriteUint32(b[8:12], n.Three)
	jay.WriteUint32(b[12:16], n.Four)
	jay.WriteUint32(b[16:20], n.Five)
	jay.WriteUint32(b[20:24], n.Six)
	jay.WriteUint32(b[24:28], n.Seven)
	jay.WriteUint32(b[28:32], n.Eight)
	jay.WriteUint32(b[32:], n.Nine)
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 36 {
		return jay.ErrUnexpectedEOB
	}
	n.One = jay.ReadUint32(b[:4])
	n.Two = jay.ReadUint32(b[4:8])
	n.Three = jay.ReadUint32(b[8:12])
	n.Four = jay.ReadUint32(b[12:16])
	n.Five = jay.ReadUint32(b[16:20])
	n.Six = jay.ReadUint32(b[20:24])
	n.Seven = jay.ReadUint32(b[24:28])
	n.Eight = jay.ReadUint32(b[28:32])
	n.Nine = jay.ReadUint32(b[32:])
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:], t.Ten)
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:])
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 44)
	jay.WriteUint32(b[:4], e.One)
	jay.WriteUint32(b[4:8], e.Two)
	jay.WriteUint32(b[8:12], e.Three)
	jay.WriteUint32(b[12:16], e.Four)
	jay.WriteUint32(b[16:20], e.Five)
	jay.WriteUint32(b[20:24], e.Six)
	jay.WriteUint32(b[24:28], e.Seven)
	jay.WriteUint32(b[28:32], e.Eight)
	jay.WriteUint32(b[32:36], e.Nine)
	jay.WriteUint32(b[36:40], e.Ten)
	jay.WriteUint32(b[40:], e.Eleven)
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 44 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadUint32(b[:4])
	e.Two = jay.ReadUint32(b[4:8])
	e.Three = jay.ReadUint32(b[8:12])
	e.Four = jay.ReadUint32(b[12:16])
	e.Five = jay.ReadUint32(b[16:20])
	e.Six = jay.ReadUint32(b[20:24])
	e.Seven = jay.ReadUint32(b[24:28])
	e.Eight = jay.ReadUint32(b[28:32])
	e.Nine = jay.ReadUint32(b[32:36])
	e.Ten = jay.ReadUint32(b[36:40])
	e.Eleven = jay.ReadUint32(b[40:])
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 48)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:], t.Twelve)
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 48 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:])
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 52)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:48], t.Twelve)
	jay.WriteUint32(b[48:], t.Thirteen)
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 52 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:48])
	t.Thirteen = jay.ReadUint32(b[48:])
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 56)
	jay.WriteUint32(b[:4], f.One)
	jay.WriteUint32(b[4:8], f.Two)
	jay.WriteUint32(b[8:12], f.Three)
	jay.WriteUint32(b[12:16], f.Four)
	jay.WriteUint32(b[16:20], f.Five)
	jay.WriteUint32(b[20:24], f.Six)
	jay.WriteUint32(b[24:28], f.Seven)
	jay.WriteUint32(b[28:32], f.Eight)
	jay.WriteUint32(b[32:36], f.Nine)
	jay.WriteUint32(b[36:40], f.Ten)
	jay.WriteUint32(b[40:44], f.Eleven)
	jay.WriteUint32(b[44:48], f.Twelve)
	jay.WriteUint32(b[48:52], f.Thirteen)
	jay.WriteUint32(b[52:], f.Fourteen)
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 56 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadUint32(b[:4])
	f.Two = jay.ReadUint32(b[4:8])
	f.Three = jay.ReadUint32(b[8:12])
	f.Four = jay.ReadUint32(b[12:16])
	f.Five = jay.ReadUint32(b[16:20])
	f.Six = jay.ReadUint32(b[20:24])
	f.Seven = jay.ReadUint32(b[24:28])
	f.Eight = jay.ReadUint32(b[28:32])
	f.Nine = jay.ReadUint32(b[32:36])
	f.Ten = jay.ReadUint32(b[36:40])
	f.Eleven = jay.ReadUint32(b[40:44])
	f.Twelve = jay.ReadUint32(b[44:48])
	f.Thirteen = jay.ReadUint32(b[48:52])
	f.Fourteen = jay.ReadUint32(b[52:])
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 60)
	jay.WriteUint32(b[:4], f.One)
	jay.WriteUint32(b[4:8], f.Two)
	jay.WriteUint32(b[8:12], f.Three)
	jay.WriteUint32(b[12:16], f.Four)
	jay.WriteUint32(b[16:20], f.Five)
	jay.WriteUint32(b[20:24], f.Six)
	jay.WriteUint32(b[24:28], f.Seven)
	jay.WriteUint32(b[28:32], f.Eight)
	jay.WriteUint32(b[32:36], f.Nine)
	jay.WriteUint32(b[36:40], f.Ten)
	jay.WriteUint32(b[40:44], f.Eleven)
	jay.WriteUint32(b[44:48], f.Twelve)
	jay.WriteUint32(b[48:52], f.Thirteen)
	jay.WriteUint32(b[52:56], f.Fourteen)
	jay.WriteUint32(b[56:], f.Fifteen)
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 60 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadUint32(b[:4])
	f.Two = jay.ReadUint32(b[4:8])
	f.Three = jay.ReadUint32(b[8:12])
	f.Four = jay.ReadUint32(b[12:16])
	f.Five = jay.ReadUint32(b[16:20])
	f.Six = jay.ReadUint32(b[20:24])
	f.Seven = jay.ReadUint32(b[24:28])
	f.Eight = jay.ReadUint32(b[28:32])
	f.Nine = jay.ReadUint32(b[32:36])
	f.Ten = jay.ReadUint32(b[36:40])
	f.Eleven = jay.ReadUint32(b[40:44])
	f.Twelve = jay.ReadUint32(b[44:48])
	f.Thirteen = jay.ReadUint32(b[48:52])
	f.Fourteen = jay.ReadUint32(b[52:56])
	f.Fifteen = jay.ReadUint32(b[56:])
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 64)
	jay.WriteUint32(b[:4], s.One)
	jay.WriteUint32(b[4:8], s.Two)
	jay.WriteUint32(b[8:12], s.Three)
	jay.WriteUint32(b[12:16], s.Four)
	jay.WriteUint32(b[16:20], s.Five)
	jay.WriteUint32(b[20:24], s.Six)
	jay.WriteUint32(b[24:28], s.Seven)
	jay.WriteUint32(b[28:32], s.Eight)
	jay.WriteUint32(b[32:36], s.Nine)
	jay.WriteUint32(b[36:40], s.Ten)
	jay.WriteUint32(b[40:44], s.Eleven)
	jay.WriteUint32(b[44:48], s.Twelve)
	jay.WriteUint32(b[48:52], s.Thirteen)
	jay.WriteUint32(b[52:56], s.Fourteen)
	jay.WriteUint32(b[56:60], s.Fifteen)
	jay.WriteUint32(b[60:], s.Sixteen)
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 64 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadUint32(b[:4])
	s.Two = jay.ReadUint32(b[4:8])
	s.Three = jay.ReadUint32(b[8:12])
	s.Four = jay.ReadUint32(b[12:16])
	s.Five = jay.ReadUint32(b[16:20])
	s.Six = jay.ReadUint32(b[20:24])
	s.Seven = jay.ReadUint32(b[24:28])
	s.Eight = jay.ReadUint32(b[28:32])
	s.Nine = jay.ReadUint32(b[32:36])
	s.Ten = jay.ReadUint32(b[36:40])
	s.Eleven = jay.ReadUint32(b[40:44])
	s.Twelve = jay.ReadUint32(b[44:48])
	s.Thirteen = jay.ReadUint32(b[48:52])
	s.Fourteen = jay.ReadUint32(b[52:56])
	s.Fifteen = jay.ReadUint32(b[56:60])
	s.Sixteen = jay.ReadUint32(b[60:])
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 68)
	jay.WriteUint32(b[:4], s.One)
	jay.WriteUint32(b[4:8], s.Two)
	jay.WriteUint32(b[8:12], s.Three)
	jay.WriteUint32(b[12:16], s.Four)
	jay.WriteUint32(b[16:20], s.Five)
	jay.WriteUint32(b[20:24], s.Six)
	jay.WriteUint32(b[24:28], s.Seven)
	jay.WriteUint32(b[28:32], s.Eight)
	jay.WriteUint32(b[32:36], s.Nine)
	jay.WriteUint32(b[36:40], s.Ten)
	jay.WriteUint32(b[40:44], s.Eleven)
	jay.WriteUint32(b[44:48], s.Twelve)
	jay.WriteUint32(b[48:52], s.Thirteen)
	jay.WriteUint32(b[52:56], s.Fourteen)
	jay.WriteUint32(b[56:60], s.Fifteen)
	jay.WriteUint32(b[60:64], s.Sixteen)
	jay.WriteUint32(b[64:], s.Seventeen)
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 68 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadUint32(b[:4])
	s.Two = jay.ReadUint32(b[4:8])
	s.Three = jay.ReadUint32(b[8:12])
	s.Four = jay.ReadUint32(b[12:16])
	s.Five = jay.ReadUint32(b[16:20])
	s.Six = jay.ReadUint32(b[20:24])
	s.Seven = jay.ReadUint32(b[24:28])
	s.Eight = jay.ReadUint32(b[28:32])
	s.Nine = jay.ReadUint32(b[32:36])
	s.Ten = jay.ReadUint32(b[36:40])
	s.Eleven = jay.ReadUint32(b[40:44])
	s.Twelve = jay.ReadUint32(b[44:48])
	s.Thirteen = jay.ReadUint32(b[48:52])
	s.Fourteen = jay.ReadUint32(b[52:56])
	s.Fifteen = jay.ReadUint32(b[56:60])
	s.Sixteen = jay.ReadUint32(b[60:64])
	s.Seventeen = jay.ReadUint32(b[64:])
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 72)
	jay.WriteUint32(b[:4], e.One)
	jay.WriteUint32(b[4:8], e.Two)
	jay.WriteUint32(b[8:12], e.Three)
	jay.WriteUint32(b[12:16], e.Four)
	jay.WriteUint32(b[16:20], e.Five)
	jay.WriteUint32(b[20:24], e.Six)
	jay.WriteUint32(b[24:28], e.Seven)
	jay.WriteUint32(b[28:32], e.Eight)
	jay.WriteUint32(b[32:36], e.Nine)
	jay.WriteUint32(b[36:40], e.Ten)
	jay.WriteUint32(b[40:44], e.Eleven)
	jay.WriteUint32(b[44:48], e.Twelve)
	jay.WriteUint32(b[48:52], e.Thirteen)
	jay.WriteUint32(b[52:56], e.Fourteen)
	jay.WriteUint32(b[56:60], e.Fifteen)
	jay.WriteUint32(b[60:64], e.Sixteen)
	jay.WriteUint32(b[64:68], e.Seventeen)
	jay.WriteUint32(b[68:], e.Eighteen)
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 72 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadUint32(b[:4])
	e.Two = jay.ReadUint32(b[4:8])
	e.Three = jay.ReadUint32(b[8:12])
	e.Four = jay.ReadUint32(b[12:16])
	e.Five = jay.ReadUint32(b[16:20])
	e.Six = jay.ReadUint32(b[20:24])
	e.Seven = jay.ReadUint32(b[24:28])
	e.Eight = jay.ReadUint32(b[28:32])
	e.Nine = jay.ReadUint32(b[32:36])
	e.Ten = jay.ReadUint32(b[36:40])
	e.Eleven = jay.ReadUint32(b[40:44])
	e.Twelve = jay.ReadUint32(b[44:48])
	e.Thirteen = jay.ReadUint32(b[48:52])
	e.Fourteen = jay.ReadUint32(b[52:56])
	e.Fifteen = jay.ReadUint32(b[56:60])
	e.Sixteen = jay.ReadUint32(b[60:64])
	e.Seventeen = jay.ReadUint32(b[64:68])
	e.Eighteen = jay.ReadUint32(b[68:])
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 76)
	jay.WriteUint32(b[:4], n.One)
	jay.WriteUint32(b[4:8], n.Two)
	jay.WriteUint32(b[8:12], n.Three)
	jay.WriteUint32(b[12:16], n.Four)
	jay.WriteUint32(b[16:20], n.Five)
	jay.WriteUint32(b[20:24], n.Six)
	jay.WriteUint32(b[24:28], n.Seven)
	jay.WriteUint32(b[28:32], n.Eight)
	jay.WriteUint32(b[32:36], n.Nine)
	jay.WriteUint32(b[36:40], n.Ten)
	jay.WriteUint32(b[40:44], n.Eleven)
	jay.WriteUint32(b[44:48], n.Twelve)
	jay.WriteUint32(b[48:52], n.Thirteen)
	jay.WriteUint32(b[52:56], n.Fourteen)
	jay.WriteUint32(b[56:60], n.Fifteen)
	jay.WriteUint32(b[60:64], n.Sixteen)
	jay.WriteUint32(b[64:68], n.Seventeen)
	jay.WriteUint32(b[68:72], n.Eighteen)
	jay.WriteUint32(b[72:], n.Nineteen)
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 76 {
		return jay.ErrUnexpectedEOB
	}
	n.One = jay.ReadUint32(b[:4])
	n.Two = jay.ReadUint32(b[4:8])
	n.Three = jay.ReadUint32(b[8:12])
	n.Four = jay.ReadUint32(b[12:16])
	n.Five = jay.ReadUint32(b[16:20])
	n.Six = jay.ReadUint32(b[20:24])
	n.Seven = jay.ReadUint32(b[24:28])
	n.Eight = jay.ReadUint32(b[28:32])
	n.Nine = jay.ReadUint32(b[32:36])
	n.Ten = jay.ReadUint32(b[36:40])
	n.Eleven = jay.ReadUint32(b[40:44])
	n.Twelve = jay.ReadUint32(b[44:48])
	n.Thirteen = jay.ReadUint32(b[48:52])
	n.Fourteen = jay.ReadUint32(b[52:56])
	n.Fifteen = jay.ReadUint32(b[56:60])
	n.Sixteen = jay.ReadUint32(b[60:64])
	n.Seventeen = jay.ReadUint32(b[64:68])
	n.Eighteen = jay.ReadUint32(b[68:72])
	n.Nineteen = jay.ReadUint32(b[72:])
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 80)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:48], t.Twelve)
	jay.WriteUint32(b[48:52], t.Thirteen)
	jay.WriteUint32(b[52:56], t.Fourteen)
	jay.WriteUint32(b[56:60], t.Fifteen)
	jay.WriteUint32(b[60:64], t.Sixteen)
	jay.WriteUint32(b[64:68], t.Seventeen)
	jay.WriteUint32(b[68:72], t.Eighteen)
	jay.WriteUint32(b[72:76], t.Nineteen)
	jay.WriteUint32(b[76:], t.Twenty)
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 80 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:48])
	t.Thirteen = jay.ReadUint32(b[48:52])
	t.Fourteen = jay.ReadUint32(b[52:56])
	t.Fifteen = jay.ReadUint32(b[56:60])
	t.Sixteen = jay.ReadUint32(b[60:64])
	t.Seventeen = jay.ReadUint32(b[64:68])
	t.Eighteen = jay.ReadUint32(b[68:72])
	t.Nineteen = jay.ReadUint32(b[72:76])
	t.Twenty = jay.ReadUint32(b[76:])
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 84)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:48], t.Twelve)
	jay.WriteUint32(b[48:52], t.Thirteen)
	jay.WriteUint32(b[52:56], t.Fourteen)
	jay.WriteUint32(b[56:60], t.Fifteen)
	jay.WriteUint32(b[60:64], t.Sixteen)
	jay.WriteUint32(b[64:68], t.Seventeen)
	jay.WriteUint32(b[68:72], t.Eighteen)
	jay.WriteUint32(b[72:76], t.Nineteen)
	jay.WriteUint32(b[76:80], t.Twenty)
	jay.WriteUint32(b[80:], t.TwentyOne)
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 84 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:48])
	t.Thirteen = jay.ReadUint32(b[48:52])
	t.Fourteen = jay.ReadUint32(b[52:56])
	t.Fifteen = jay.ReadUint32(b[56:60])
	t.Sixteen = jay.ReadUint32(b[60:64])
	t.Seventeen = jay.ReadUint32(b[64:68])
	t.Eighteen = jay.ReadUint32(b[68:72])
	t.Nineteen = jay.ReadUint32(b[72:76])
	t.Twenty = jay.ReadUint32(b[76:80])
	t.TwentyOne = jay.ReadUint32(b[80:])
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 88)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:48], t.Twelve)
	jay.WriteUint32(b[48:52], t.Thirteen)
	jay.WriteUint32(b[52:56], t.Fourteen)
	jay.WriteUint32(b[56:60], t.Fifteen)
	jay.WriteUint32(b[60:64], t.Sixteen)
	jay.WriteUint32(b[64:68], t.Seventeen)
	jay.WriteUint32(b[68:72], t.Eighteen)
	jay.WriteUint32(b[72:76], t.Nineteen)
	jay.WriteUint32(b[76:80], t.Twenty)
	jay.WriteUint32(b[80:84], t.TwentyOne)
	jay.WriteUint32(b[84:], t.TwentyTwo)
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 88 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:48])
	t.Thirteen = jay.ReadUint32(b[48:52])
	t.Fourteen = jay.ReadUint32(b[52:56])
	t.Fifteen = jay.ReadUint32(b[56:60])
	t.Sixteen = jay.ReadUint32(b[60:64])
	t.Seventeen = jay.ReadUint32(b[64:68])
	t.Eighteen = jay.ReadUint32(b[68:72])
	t.Nineteen = jay.ReadUint32(b[72:76])
	t.Twenty = jay.ReadUint32(b[76:80])
	t.TwentyOne = jay.ReadUint32(b[80:84])
	t.TwentyTwo = jay.ReadUint32(b[84:])
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 92)
	jay.WriteUint32(b[:4], t.One)
	jay.WriteUint32(b[4:8], t.Two)
	jay.WriteUint32(b[8:12], t.Three)
	jay.WriteUint32(b[12:16], t.Four)
	jay.WriteUint32(b[16:20], t.Five)
	jay.WriteUint32(b[20:24], t.Six)
	jay.WriteUint32(b[24:28], t.Seven)
	jay.WriteUint32(b[28:32], t.Eight)
	jay.WriteUint32(b[32:36], t.Nine)
	jay.WriteUint32(b[36:40], t.Ten)
	jay.WriteUint32(b[40:44], t.Eleven)
	jay.WriteUint32(b[44:48], t.Twelve)
	jay.WriteUint32(b[48:52], t.Thirteen)
	jay.WriteUint32(b[52:56], t.Fourteen)
	jay.WriteUint32(b[56:60], t.Fifteen)
	jay.WriteUint32(b[60:64], t.Sixteen)
	jay.WriteUint32(b[64:68], t.Seventeen)
	jay.WriteUint32(b[68:72], t.Eighteen)
	jay.WriteUint32(b[72:76], t.Nineteen)
	jay.WriteUint32(b[76:80], t.Twenty)
	jay.WriteUint32(b[80:84], t.TwentyOne)
	jay.WriteUint32(b[84:88], t.TwentyTwo)
	jay.WriteUint32(b[88:], t.TwentyThree)
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 92 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadUint32(b[:4])
	t.Two = jay.ReadUint32(b[4:8])
	t.Three = jay.ReadUint32(b[8:12])
	t.Four = jay.ReadUint32(b[12:16])
	t.Five = jay.ReadUint32(b[16:20])
	t.Six = jay.ReadUint32(b[20:24])
	t.Seven = jay.ReadUint32(b[24:28])
	t.Eight = jay.ReadUint32(b[28:32])
	t.Nine = jay.ReadUint32(b[32:36])
	t.Ten = jay.ReadUint32(b[36:40])
	t.Eleven = jay.ReadUint32(b[40:44])
	t.Twelve = jay.ReadUint32(b[44:48])
	t.Thirteen = jay.ReadUint32(b[48:52])
	t.Fourteen = jay.ReadUint32(b[52:56])
	t.Fifteen = jay.ReadUint32(b[56:60])
	t.Sixteen = jay.ReadUint32(b[60:64])
	t.Seventeen = jay.ReadUint32(b[64:68])
	t.Eighteen = jay.ReadUint32(b[68:72])
	t.Nineteen = jay.ReadUint32(b[72:76])
	t.Twenty = jay.ReadUint32(b[76:80])
	t.TwentyOne = jay.ReadUint32(b[80:84])
	t.TwentyTwo = jay.ReadUint32(b[84:88])
	t.TwentyThree = jay.ReadUint32(b[88:])
	return nil
}
