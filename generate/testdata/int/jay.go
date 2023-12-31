// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteIntArch64(b[:8], o.One)
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	o.One = jay.ReadIntArch64(b[:8])
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteIntArch64(b[:8], f.One)
	jay.WriteIntArch64(b[8:16], f.Two)
	jay.WriteIntArch64(b[16:24], f.Three)
	jay.WriteIntArch64(b[24:32], f.Four)
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadIntArch64(b[:8])
	f.Two = jay.ReadIntArch64(b[8:16])
	f.Three = jay.ReadIntArch64(b[16:24])
	f.Four = jay.ReadIntArch64(b[24:32])
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteIntArch64(b[:8], f.One)
	jay.WriteIntArch64(b[8:16], f.Two)
	jay.WriteIntArch64(b[16:24], f.Three)
	jay.WriteIntArch64(b[24:32], f.Four)
	jay.WriteIntArch64(b[32:40], f.Five)
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadIntArch64(b[:8])
	f.Two = jay.ReadIntArch64(b[8:16])
	f.Three = jay.ReadIntArch64(b[16:24])
	f.Four = jay.ReadIntArch64(b[24:32])
	f.Five = jay.ReadIntArch64(b[32:40])
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 48)
	jay.WriteIntArch64(b[:8], s.One)
	jay.WriteIntArch64(b[8:16], s.Two)
	jay.WriteIntArch64(b[16:24], s.Three)
	jay.WriteIntArch64(b[24:32], s.Four)
	jay.WriteIntArch64(b[32:40], s.Five)
	jay.WriteIntArch64(b[40:48], s.Six)
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 48 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadIntArch64(b[:8])
	s.Two = jay.ReadIntArch64(b[8:16])
	s.Three = jay.ReadIntArch64(b[16:24])
	s.Four = jay.ReadIntArch64(b[24:32])
	s.Five = jay.ReadIntArch64(b[32:40])
	s.Six = jay.ReadIntArch64(b[40:48])
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 56)
	jay.WriteIntArch64(b[:8], s.One)
	jay.WriteIntArch64(b[8:16], s.Two)
	jay.WriteIntArch64(b[16:24], s.Three)
	jay.WriteIntArch64(b[24:32], s.Four)
	jay.WriteIntArch64(b[32:40], s.Five)
	jay.WriteIntArch64(b[40:48], s.Six)
	jay.WriteIntArch64(b[48:56], s.Seven)
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 56 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadIntArch64(b[:8])
	s.Two = jay.ReadIntArch64(b[8:16])
	s.Three = jay.ReadIntArch64(b[16:24])
	s.Four = jay.ReadIntArch64(b[24:32])
	s.Five = jay.ReadIntArch64(b[32:40])
	s.Six = jay.ReadIntArch64(b[40:48])
	s.Seven = jay.ReadIntArch64(b[48:56])
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 64)
	jay.WriteIntArch64(b[:8], e.One)
	jay.WriteIntArch64(b[8:16], e.Two)
	jay.WriteIntArch64(b[16:24], e.Three)
	jay.WriteIntArch64(b[24:32], e.Four)
	jay.WriteIntArch64(b[32:40], e.Five)
	jay.WriteIntArch64(b[40:48], e.Six)
	jay.WriteIntArch64(b[48:56], e.Seven)
	jay.WriteIntArch64(b[56:64], e.Eight)
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 64 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadIntArch64(b[:8])
	e.Two = jay.ReadIntArch64(b[8:16])
	e.Three = jay.ReadIntArch64(b[16:24])
	e.Four = jay.ReadIntArch64(b[24:32])
	e.Five = jay.ReadIntArch64(b[32:40])
	e.Six = jay.ReadIntArch64(b[40:48])
	e.Seven = jay.ReadIntArch64(b[48:56])
	e.Eight = jay.ReadIntArch64(b[56:64])
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 72)
	jay.WriteIntArch64(b[:8], n.One)
	jay.WriteIntArch64(b[8:16], n.Two)
	jay.WriteIntArch64(b[16:24], n.Three)
	jay.WriteIntArch64(b[24:32], n.Four)
	jay.WriteIntArch64(b[32:40], n.Five)
	jay.WriteIntArch64(b[40:48], n.Six)
	jay.WriteIntArch64(b[48:56], n.Seven)
	jay.WriteIntArch64(b[56:64], n.Eight)
	jay.WriteIntArch64(b[64:72], n.Nine)
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 72 {
		return jay.ErrUnexpectedEOB
	}
	n.One = jay.ReadIntArch64(b[:8])
	n.Two = jay.ReadIntArch64(b[8:16])
	n.Three = jay.ReadIntArch64(b[16:24])
	n.Four = jay.ReadIntArch64(b[24:32])
	n.Five = jay.ReadIntArch64(b[32:40])
	n.Six = jay.ReadIntArch64(b[40:48])
	n.Seven = jay.ReadIntArch64(b[48:56])
	n.Eight = jay.ReadIntArch64(b[56:64])
	n.Nine = jay.ReadIntArch64(b[64:72])
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 80)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 80 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 88)
	jay.WriteIntArch64(b[:8], e.One)
	jay.WriteIntArch64(b[8:16], e.Two)
	jay.WriteIntArch64(b[16:24], e.Three)
	jay.WriteIntArch64(b[24:32], e.Four)
	jay.WriteIntArch64(b[32:40], e.Five)
	jay.WriteIntArch64(b[40:48], e.Six)
	jay.WriteIntArch64(b[48:56], e.Seven)
	jay.WriteIntArch64(b[56:64], e.Eight)
	jay.WriteIntArch64(b[64:72], e.Nine)
	jay.WriteIntArch64(b[72:80], e.Ten)
	jay.WriteIntArch64(b[80:88], e.Eleven)
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 88 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadIntArch64(b[:8])
	e.Two = jay.ReadIntArch64(b[8:16])
	e.Three = jay.ReadIntArch64(b[16:24])
	e.Four = jay.ReadIntArch64(b[24:32])
	e.Five = jay.ReadIntArch64(b[32:40])
	e.Six = jay.ReadIntArch64(b[40:48])
	e.Seven = jay.ReadIntArch64(b[48:56])
	e.Eight = jay.ReadIntArch64(b[56:64])
	e.Nine = jay.ReadIntArch64(b[64:72])
	e.Ten = jay.ReadIntArch64(b[72:80])
	e.Eleven = jay.ReadIntArch64(b[80:88])
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 96)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 96 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 104)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	jay.WriteIntArch64(b[96:104], t.Thirteen)
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 104 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	t.Thirteen = jay.ReadIntArch64(b[96:104])
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 112)
	jay.WriteIntArch64(b[:8], f.One)
	jay.WriteIntArch64(b[8:16], f.Two)
	jay.WriteIntArch64(b[16:24], f.Three)
	jay.WriteIntArch64(b[24:32], f.Four)
	jay.WriteIntArch64(b[32:40], f.Five)
	jay.WriteIntArch64(b[40:48], f.Six)
	jay.WriteIntArch64(b[48:56], f.Seven)
	jay.WriteIntArch64(b[56:64], f.Eight)
	jay.WriteIntArch64(b[64:72], f.Nine)
	jay.WriteIntArch64(b[72:80], f.Ten)
	jay.WriteIntArch64(b[80:88], f.Eleven)
	jay.WriteIntArch64(b[88:96], f.Twelve)
	jay.WriteIntArch64(b[96:104], f.Thirteen)
	jay.WriteIntArch64(b[104:112], f.Fourteen)
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 112 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadIntArch64(b[:8])
	f.Two = jay.ReadIntArch64(b[8:16])
	f.Three = jay.ReadIntArch64(b[16:24])
	f.Four = jay.ReadIntArch64(b[24:32])
	f.Five = jay.ReadIntArch64(b[32:40])
	f.Six = jay.ReadIntArch64(b[40:48])
	f.Seven = jay.ReadIntArch64(b[48:56])
	f.Eight = jay.ReadIntArch64(b[56:64])
	f.Nine = jay.ReadIntArch64(b[64:72])
	f.Ten = jay.ReadIntArch64(b[72:80])
	f.Eleven = jay.ReadIntArch64(b[80:88])
	f.Twelve = jay.ReadIntArch64(b[88:96])
	f.Thirteen = jay.ReadIntArch64(b[96:104])
	f.Fourteen = jay.ReadIntArch64(b[104:112])
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 120)
	jay.WriteIntArch64(b[:8], f.One)
	jay.WriteIntArch64(b[8:16], f.Two)
	jay.WriteIntArch64(b[16:24], f.Three)
	jay.WriteIntArch64(b[24:32], f.Four)
	jay.WriteIntArch64(b[32:40], f.Five)
	jay.WriteIntArch64(b[40:48], f.Six)
	jay.WriteIntArch64(b[48:56], f.Seven)
	jay.WriteIntArch64(b[56:64], f.Eight)
	jay.WriteIntArch64(b[64:72], f.Nine)
	jay.WriteIntArch64(b[72:80], f.Ten)
	jay.WriteIntArch64(b[80:88], f.Eleven)
	jay.WriteIntArch64(b[88:96], f.Twelve)
	jay.WriteIntArch64(b[96:104], f.Thirteen)
	jay.WriteIntArch64(b[104:112], f.Fourteen)
	jay.WriteIntArch64(b[112:120], f.Fifteen)
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 120 {
		return jay.ErrUnexpectedEOB
	}
	f.One = jay.ReadIntArch64(b[:8])
	f.Two = jay.ReadIntArch64(b[8:16])
	f.Three = jay.ReadIntArch64(b[16:24])
	f.Four = jay.ReadIntArch64(b[24:32])
	f.Five = jay.ReadIntArch64(b[32:40])
	f.Six = jay.ReadIntArch64(b[40:48])
	f.Seven = jay.ReadIntArch64(b[48:56])
	f.Eight = jay.ReadIntArch64(b[56:64])
	f.Nine = jay.ReadIntArch64(b[64:72])
	f.Ten = jay.ReadIntArch64(b[72:80])
	f.Eleven = jay.ReadIntArch64(b[80:88])
	f.Twelve = jay.ReadIntArch64(b[88:96])
	f.Thirteen = jay.ReadIntArch64(b[96:104])
	f.Fourteen = jay.ReadIntArch64(b[104:112])
	f.Fifteen = jay.ReadIntArch64(b[112:120])
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 128)
	jay.WriteIntArch64(b[:8], s.One)
	jay.WriteIntArch64(b[8:16], s.Two)
	jay.WriteIntArch64(b[16:24], s.Three)
	jay.WriteIntArch64(b[24:32], s.Four)
	jay.WriteIntArch64(b[32:40], s.Five)
	jay.WriteIntArch64(b[40:48], s.Six)
	jay.WriteIntArch64(b[48:56], s.Seven)
	jay.WriteIntArch64(b[56:64], s.Eight)
	jay.WriteIntArch64(b[64:72], s.Nine)
	jay.WriteIntArch64(b[72:80], s.Ten)
	jay.WriteIntArch64(b[80:88], s.Eleven)
	jay.WriteIntArch64(b[88:96], s.Twelve)
	jay.WriteIntArch64(b[96:104], s.Thirteen)
	jay.WriteIntArch64(b[104:112], s.Fourteen)
	jay.WriteIntArch64(b[112:120], s.Fifteen)
	jay.WriteIntArch64(b[120:128], s.Sixteen)
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 128 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadIntArch64(b[:8])
	s.Two = jay.ReadIntArch64(b[8:16])
	s.Three = jay.ReadIntArch64(b[16:24])
	s.Four = jay.ReadIntArch64(b[24:32])
	s.Five = jay.ReadIntArch64(b[32:40])
	s.Six = jay.ReadIntArch64(b[40:48])
	s.Seven = jay.ReadIntArch64(b[48:56])
	s.Eight = jay.ReadIntArch64(b[56:64])
	s.Nine = jay.ReadIntArch64(b[64:72])
	s.Ten = jay.ReadIntArch64(b[72:80])
	s.Eleven = jay.ReadIntArch64(b[80:88])
	s.Twelve = jay.ReadIntArch64(b[88:96])
	s.Thirteen = jay.ReadIntArch64(b[96:104])
	s.Fourteen = jay.ReadIntArch64(b[104:112])
	s.Fifteen = jay.ReadIntArch64(b[112:120])
	s.Sixteen = jay.ReadIntArch64(b[120:128])
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 136)
	jay.WriteIntArch64(b[:8], s.One)
	jay.WriteIntArch64(b[8:16], s.Two)
	jay.WriteIntArch64(b[16:24], s.Three)
	jay.WriteIntArch64(b[24:32], s.Four)
	jay.WriteIntArch64(b[32:40], s.Five)
	jay.WriteIntArch64(b[40:48], s.Six)
	jay.WriteIntArch64(b[48:56], s.Seven)
	jay.WriteIntArch64(b[56:64], s.Eight)
	jay.WriteIntArch64(b[64:72], s.Nine)
	jay.WriteIntArch64(b[72:80], s.Ten)
	jay.WriteIntArch64(b[80:88], s.Eleven)
	jay.WriteIntArch64(b[88:96], s.Twelve)
	jay.WriteIntArch64(b[96:104], s.Thirteen)
	jay.WriteIntArch64(b[104:112], s.Fourteen)
	jay.WriteIntArch64(b[112:120], s.Fifteen)
	jay.WriteIntArch64(b[120:128], s.Sixteen)
	jay.WriteIntArch64(b[128:136], s.Seventeen)
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 136 {
		return jay.ErrUnexpectedEOB
	}
	s.One = jay.ReadIntArch64(b[:8])
	s.Two = jay.ReadIntArch64(b[8:16])
	s.Three = jay.ReadIntArch64(b[16:24])
	s.Four = jay.ReadIntArch64(b[24:32])
	s.Five = jay.ReadIntArch64(b[32:40])
	s.Six = jay.ReadIntArch64(b[40:48])
	s.Seven = jay.ReadIntArch64(b[48:56])
	s.Eight = jay.ReadIntArch64(b[56:64])
	s.Nine = jay.ReadIntArch64(b[64:72])
	s.Ten = jay.ReadIntArch64(b[72:80])
	s.Eleven = jay.ReadIntArch64(b[80:88])
	s.Twelve = jay.ReadIntArch64(b[88:96])
	s.Thirteen = jay.ReadIntArch64(b[96:104])
	s.Fourteen = jay.ReadIntArch64(b[104:112])
	s.Fifteen = jay.ReadIntArch64(b[112:120])
	s.Sixteen = jay.ReadIntArch64(b[120:128])
	s.Seventeen = jay.ReadIntArch64(b[128:136])
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 144)
	jay.WriteIntArch64(b[:8], e.One)
	jay.WriteIntArch64(b[8:16], e.Two)
	jay.WriteIntArch64(b[16:24], e.Three)
	jay.WriteIntArch64(b[24:32], e.Four)
	jay.WriteIntArch64(b[32:40], e.Five)
	jay.WriteIntArch64(b[40:48], e.Six)
	jay.WriteIntArch64(b[48:56], e.Seven)
	jay.WriteIntArch64(b[56:64], e.Eight)
	jay.WriteIntArch64(b[64:72], e.Nine)
	jay.WriteIntArch64(b[72:80], e.Ten)
	jay.WriteIntArch64(b[80:88], e.Eleven)
	jay.WriteIntArch64(b[88:96], e.Twelve)
	jay.WriteIntArch64(b[96:104], e.Thirteen)
	jay.WriteIntArch64(b[104:112], e.Fourteen)
	jay.WriteIntArch64(b[112:120], e.Fifteen)
	jay.WriteIntArch64(b[120:128], e.Sixteen)
	jay.WriteIntArch64(b[128:136], e.Seventeen)
	jay.WriteIntArch64(b[136:144], e.Eighteen)
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 144 {
		return jay.ErrUnexpectedEOB
	}
	e.One = jay.ReadIntArch64(b[:8])
	e.Two = jay.ReadIntArch64(b[8:16])
	e.Three = jay.ReadIntArch64(b[16:24])
	e.Four = jay.ReadIntArch64(b[24:32])
	e.Five = jay.ReadIntArch64(b[32:40])
	e.Six = jay.ReadIntArch64(b[40:48])
	e.Seven = jay.ReadIntArch64(b[48:56])
	e.Eight = jay.ReadIntArch64(b[56:64])
	e.Nine = jay.ReadIntArch64(b[64:72])
	e.Ten = jay.ReadIntArch64(b[72:80])
	e.Eleven = jay.ReadIntArch64(b[80:88])
	e.Twelve = jay.ReadIntArch64(b[88:96])
	e.Thirteen = jay.ReadIntArch64(b[96:104])
	e.Fourteen = jay.ReadIntArch64(b[104:112])
	e.Fifteen = jay.ReadIntArch64(b[112:120])
	e.Sixteen = jay.ReadIntArch64(b[120:128])
	e.Seventeen = jay.ReadIntArch64(b[128:136])
	e.Eighteen = jay.ReadIntArch64(b[136:144])
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 152)
	jay.WriteIntArch64(b[:8], n.One)
	jay.WriteIntArch64(b[8:16], n.Two)
	jay.WriteIntArch64(b[16:24], n.Three)
	jay.WriteIntArch64(b[24:32], n.Four)
	jay.WriteIntArch64(b[32:40], n.Five)
	jay.WriteIntArch64(b[40:48], n.Six)
	jay.WriteIntArch64(b[48:56], n.Seven)
	jay.WriteIntArch64(b[56:64], n.Eight)
	jay.WriteIntArch64(b[64:72], n.Nine)
	jay.WriteIntArch64(b[72:80], n.Ten)
	jay.WriteIntArch64(b[80:88], n.Eleven)
	jay.WriteIntArch64(b[88:96], n.Twelve)
	jay.WriteIntArch64(b[96:104], n.Thirteen)
	jay.WriteIntArch64(b[104:112], n.Fourteen)
	jay.WriteIntArch64(b[112:120], n.Fifteen)
	jay.WriteIntArch64(b[120:128], n.Sixteen)
	jay.WriteIntArch64(b[128:136], n.Seventeen)
	jay.WriteIntArch64(b[136:144], n.Eighteen)
	jay.WriteIntArch64(b[144:152], n.Nineteen)
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 152 {
		return jay.ErrUnexpectedEOB
	}
	n.One = jay.ReadIntArch64(b[:8])
	n.Two = jay.ReadIntArch64(b[8:16])
	n.Three = jay.ReadIntArch64(b[16:24])
	n.Four = jay.ReadIntArch64(b[24:32])
	n.Five = jay.ReadIntArch64(b[32:40])
	n.Six = jay.ReadIntArch64(b[40:48])
	n.Seven = jay.ReadIntArch64(b[48:56])
	n.Eight = jay.ReadIntArch64(b[56:64])
	n.Nine = jay.ReadIntArch64(b[64:72])
	n.Ten = jay.ReadIntArch64(b[72:80])
	n.Eleven = jay.ReadIntArch64(b[80:88])
	n.Twelve = jay.ReadIntArch64(b[88:96])
	n.Thirteen = jay.ReadIntArch64(b[96:104])
	n.Fourteen = jay.ReadIntArch64(b[104:112])
	n.Fifteen = jay.ReadIntArch64(b[112:120])
	n.Sixteen = jay.ReadIntArch64(b[120:128])
	n.Seventeen = jay.ReadIntArch64(b[128:136])
	n.Eighteen = jay.ReadIntArch64(b[136:144])
	n.Nineteen = jay.ReadIntArch64(b[144:152])
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 160)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	jay.WriteIntArch64(b[96:104], t.Thirteen)
	jay.WriteIntArch64(b[104:112], t.Fourteen)
	jay.WriteIntArch64(b[112:120], t.Fifteen)
	jay.WriteIntArch64(b[120:128], t.Sixteen)
	jay.WriteIntArch64(b[128:136], t.Seventeen)
	jay.WriteIntArch64(b[136:144], t.Eighteen)
	jay.WriteIntArch64(b[144:152], t.Nineteen)
	jay.WriteIntArch64(b[152:160], t.Twenty)
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 160 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	t.Thirteen = jay.ReadIntArch64(b[96:104])
	t.Fourteen = jay.ReadIntArch64(b[104:112])
	t.Fifteen = jay.ReadIntArch64(b[112:120])
	t.Sixteen = jay.ReadIntArch64(b[120:128])
	t.Seventeen = jay.ReadIntArch64(b[128:136])
	t.Eighteen = jay.ReadIntArch64(b[136:144])
	t.Nineteen = jay.ReadIntArch64(b[144:152])
	t.Twenty = jay.ReadIntArch64(b[152:160])
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 168)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	jay.WriteIntArch64(b[96:104], t.Thirteen)
	jay.WriteIntArch64(b[104:112], t.Fourteen)
	jay.WriteIntArch64(b[112:120], t.Fifteen)
	jay.WriteIntArch64(b[120:128], t.Sixteen)
	jay.WriteIntArch64(b[128:136], t.Seventeen)
	jay.WriteIntArch64(b[136:144], t.Eighteen)
	jay.WriteIntArch64(b[144:152], t.Nineteen)
	jay.WriteIntArch64(b[152:160], t.Twenty)
	jay.WriteIntArch64(b[160:168], t.TwentyOne)
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 168 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	t.Thirteen = jay.ReadIntArch64(b[96:104])
	t.Fourteen = jay.ReadIntArch64(b[104:112])
	t.Fifteen = jay.ReadIntArch64(b[112:120])
	t.Sixteen = jay.ReadIntArch64(b[120:128])
	t.Seventeen = jay.ReadIntArch64(b[128:136])
	t.Eighteen = jay.ReadIntArch64(b[136:144])
	t.Nineteen = jay.ReadIntArch64(b[144:152])
	t.Twenty = jay.ReadIntArch64(b[152:160])
	t.TwentyOne = jay.ReadIntArch64(b[160:168])
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 176)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	jay.WriteIntArch64(b[96:104], t.Thirteen)
	jay.WriteIntArch64(b[104:112], t.Fourteen)
	jay.WriteIntArch64(b[112:120], t.Fifteen)
	jay.WriteIntArch64(b[120:128], t.Sixteen)
	jay.WriteIntArch64(b[128:136], t.Seventeen)
	jay.WriteIntArch64(b[136:144], t.Eighteen)
	jay.WriteIntArch64(b[144:152], t.Nineteen)
	jay.WriteIntArch64(b[152:160], t.Twenty)
	jay.WriteIntArch64(b[160:168], t.TwentyOne)
	jay.WriteIntArch64(b[168:176], t.TwentyTwo)
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 176 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	t.Thirteen = jay.ReadIntArch64(b[96:104])
	t.Fourteen = jay.ReadIntArch64(b[104:112])
	t.Fifteen = jay.ReadIntArch64(b[112:120])
	t.Sixteen = jay.ReadIntArch64(b[120:128])
	t.Seventeen = jay.ReadIntArch64(b[128:136])
	t.Eighteen = jay.ReadIntArch64(b[136:144])
	t.Nineteen = jay.ReadIntArch64(b[144:152])
	t.Twenty = jay.ReadIntArch64(b[152:160])
	t.TwentyOne = jay.ReadIntArch64(b[160:168])
	t.TwentyTwo = jay.ReadIntArch64(b[168:176])
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 184)
	jay.WriteIntArch64(b[:8], t.One)
	jay.WriteIntArch64(b[8:16], t.Two)
	jay.WriteIntArch64(b[16:24], t.Three)
	jay.WriteIntArch64(b[24:32], t.Four)
	jay.WriteIntArch64(b[32:40], t.Five)
	jay.WriteIntArch64(b[40:48], t.Six)
	jay.WriteIntArch64(b[48:56], t.Seven)
	jay.WriteIntArch64(b[56:64], t.Eight)
	jay.WriteIntArch64(b[64:72], t.Nine)
	jay.WriteIntArch64(b[72:80], t.Ten)
	jay.WriteIntArch64(b[80:88], t.Eleven)
	jay.WriteIntArch64(b[88:96], t.Twelve)
	jay.WriteIntArch64(b[96:104], t.Thirteen)
	jay.WriteIntArch64(b[104:112], t.Fourteen)
	jay.WriteIntArch64(b[112:120], t.Fifteen)
	jay.WriteIntArch64(b[120:128], t.Sixteen)
	jay.WriteIntArch64(b[128:136], t.Seventeen)
	jay.WriteIntArch64(b[136:144], t.Eighteen)
	jay.WriteIntArch64(b[144:152], t.Nineteen)
	jay.WriteIntArch64(b[152:160], t.Twenty)
	jay.WriteIntArch64(b[160:168], t.TwentyOne)
	jay.WriteIntArch64(b[168:176], t.TwentyTwo)
	jay.WriteIntArch64(b[176:184], t.TwentyThree)
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 184 {
		return jay.ErrUnexpectedEOB
	}
	t.One = jay.ReadIntArch64(b[:8])
	t.Two = jay.ReadIntArch64(b[8:16])
	t.Three = jay.ReadIntArch64(b[16:24])
	t.Four = jay.ReadIntArch64(b[24:32])
	t.Five = jay.ReadIntArch64(b[32:40])
	t.Six = jay.ReadIntArch64(b[40:48])
	t.Seven = jay.ReadIntArch64(b[48:56])
	t.Eight = jay.ReadIntArch64(b[56:64])
	t.Nine = jay.ReadIntArch64(b[64:72])
	t.Ten = jay.ReadIntArch64(b[72:80])
	t.Eleven = jay.ReadIntArch64(b[80:88])
	t.Twelve = jay.ReadIntArch64(b[88:96])
	t.Thirteen = jay.ReadIntArch64(b[96:104])
	t.Fourteen = jay.ReadIntArch64(b[104:112])
	t.Fifteen = jay.ReadIntArch64(b[112:120])
	t.Sixteen = jay.ReadIntArch64(b[120:128])
	t.Seventeen = jay.ReadIntArch64(b[128:136])
	t.Eighteen = jay.ReadIntArch64(b[136:144])
	t.Nineteen = jay.ReadIntArch64(b[144:152])
	t.Twenty = jay.ReadIntArch64(b[152:160])
	t.TwentyOne = jay.ReadIntArch64(b[160:168])
	t.TwentyTwo = jay.ReadIntArch64(b[168:176])
	t.TwentyThree = jay.ReadIntArch64(b[176:184])
	return nil
}
