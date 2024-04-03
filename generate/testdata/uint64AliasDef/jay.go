// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteUint64(b, uint64(o.One))
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	o.One = long(jay.ReadUint64(b[:8]))
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:], uint64(t.Two))
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:], uint64(t.Three))
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteUint64(b[:8], uint64(f.One))
	jay.WriteUint64(b[8:16], uint64(f.Two))
	jay.WriteUint64(b[16:24], uint64(f.Three))
	jay.WriteUint64(b[24:], uint64(f.Four))
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	f.One = long(jay.ReadUint64(b[:8]))
	f.Two = long(jay.ReadUint64(b[8:16]))
	f.Three = long(jay.ReadUint64(b[16:24]))
	f.Four = long(jay.ReadUint64(b[24:32]))
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteUint64(b[:8], uint64(f.One))
	jay.WriteUint64(b[8:16], uint64(f.Two))
	jay.WriteUint64(b[16:24], uint64(f.Three))
	jay.WriteUint64(b[24:32], uint64(f.Four))
	jay.WriteUint64(b[32:], uint64(f.Five))
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	f.One = long(jay.ReadUint64(b[:8]))
	f.Two = long(jay.ReadUint64(b[8:16]))
	f.Three = long(jay.ReadUint64(b[16:24]))
	f.Four = long(jay.ReadUint64(b[24:32]))
	f.Five = long(jay.ReadUint64(b[32:40]))
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 48)
	jay.WriteUint64(b[:8], uint64(s.One))
	jay.WriteUint64(b[8:16], uint64(s.Two))
	jay.WriteUint64(b[16:24], uint64(s.Three))
	jay.WriteUint64(b[24:32], uint64(s.Four))
	jay.WriteUint64(b[32:40], uint64(s.Five))
	jay.WriteUint64(b[40:], uint64(s.Six))
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 48 {
		return jay.ErrUnexpectedEOB
	}
	s.One = long(jay.ReadUint64(b[:8]))
	s.Two = long(jay.ReadUint64(b[8:16]))
	s.Three = long(jay.ReadUint64(b[16:24]))
	s.Four = long(jay.ReadUint64(b[24:32]))
	s.Five = long(jay.ReadUint64(b[32:40]))
	s.Six = long(jay.ReadUint64(b[40:48]))
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 56)
	jay.WriteUint64(b[:8], uint64(s.One))
	jay.WriteUint64(b[8:16], uint64(s.Two))
	jay.WriteUint64(b[16:24], uint64(s.Three))
	jay.WriteUint64(b[24:32], uint64(s.Four))
	jay.WriteUint64(b[32:40], uint64(s.Five))
	jay.WriteUint64(b[40:48], uint64(s.Six))
	jay.WriteUint64(b[48:], uint64(s.Seven))
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 56 {
		return jay.ErrUnexpectedEOB
	}
	s.One = long(jay.ReadUint64(b[:8]))
	s.Two = long(jay.ReadUint64(b[8:16]))
	s.Three = long(jay.ReadUint64(b[16:24]))
	s.Four = long(jay.ReadUint64(b[24:32]))
	s.Five = long(jay.ReadUint64(b[32:40]))
	s.Six = long(jay.ReadUint64(b[40:48]))
	s.Seven = long(jay.ReadUint64(b[48:56]))
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 64)
	jay.WriteUint64(b[:8], uint64(e.One))
	jay.WriteUint64(b[8:16], uint64(e.Two))
	jay.WriteUint64(b[16:24], uint64(e.Three))
	jay.WriteUint64(b[24:32], uint64(e.Four))
	jay.WriteUint64(b[32:40], uint64(e.Five))
	jay.WriteUint64(b[40:48], uint64(e.Six))
	jay.WriteUint64(b[48:56], uint64(e.Seven))
	jay.WriteUint64(b[56:], uint64(e.Eight))
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 64 {
		return jay.ErrUnexpectedEOB
	}
	e.One = long(jay.ReadUint64(b[:8]))
	e.Two = long(jay.ReadUint64(b[8:16]))
	e.Three = long(jay.ReadUint64(b[16:24]))
	e.Four = long(jay.ReadUint64(b[24:32]))
	e.Five = long(jay.ReadUint64(b[32:40]))
	e.Six = long(jay.ReadUint64(b[40:48]))
	e.Seven = long(jay.ReadUint64(b[48:56]))
	e.Eight = long(jay.ReadUint64(b[56:64]))
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 72)
	jay.WriteUint64(b[:8], uint64(n.One))
	jay.WriteUint64(b[8:16], uint64(n.Two))
	jay.WriteUint64(b[16:24], uint64(n.Three))
	jay.WriteUint64(b[24:32], uint64(n.Four))
	jay.WriteUint64(b[32:40], uint64(n.Five))
	jay.WriteUint64(b[40:48], uint64(n.Six))
	jay.WriteUint64(b[48:56], uint64(n.Seven))
	jay.WriteUint64(b[56:64], uint64(n.Eight))
	jay.WriteUint64(b[64:], uint64(n.Nine))
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 72 {
		return jay.ErrUnexpectedEOB
	}
	n.One = long(jay.ReadUint64(b[:8]))
	n.Two = long(jay.ReadUint64(b[8:16]))
	n.Three = long(jay.ReadUint64(b[16:24]))
	n.Four = long(jay.ReadUint64(b[24:32]))
	n.Five = long(jay.ReadUint64(b[32:40]))
	n.Six = long(jay.ReadUint64(b[40:48]))
	n.Seven = long(jay.ReadUint64(b[48:56]))
	n.Eight = long(jay.ReadUint64(b[56:64]))
	n.Nine = long(jay.ReadUint64(b[64:72]))
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 80)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:], uint64(t.Ten))
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 80 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 88)
	jay.WriteUint64(b[:8], uint64(e.One))
	jay.WriteUint64(b[8:16], uint64(e.Two))
	jay.WriteUint64(b[16:24], uint64(e.Three))
	jay.WriteUint64(b[24:32], uint64(e.Four))
	jay.WriteUint64(b[32:40], uint64(e.Five))
	jay.WriteUint64(b[40:48], uint64(e.Six))
	jay.WriteUint64(b[48:56], uint64(e.Seven))
	jay.WriteUint64(b[56:64], uint64(e.Eight))
	jay.WriteUint64(b[64:72], uint64(e.Nine))
	jay.WriteUint64(b[72:80], uint64(e.Ten))
	jay.WriteUint64(b[80:], uint64(e.Eleven))
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 88 {
		return jay.ErrUnexpectedEOB
	}
	e.One = long(jay.ReadUint64(b[:8]))
	e.Two = long(jay.ReadUint64(b[8:16]))
	e.Three = long(jay.ReadUint64(b[16:24]))
	e.Four = long(jay.ReadUint64(b[24:32]))
	e.Five = long(jay.ReadUint64(b[32:40]))
	e.Six = long(jay.ReadUint64(b[40:48]))
	e.Seven = long(jay.ReadUint64(b[48:56]))
	e.Eight = long(jay.ReadUint64(b[56:64]))
	e.Nine = long(jay.ReadUint64(b[64:72]))
	e.Ten = long(jay.ReadUint64(b[72:80]))
	e.Eleven = long(jay.ReadUint64(b[80:88]))
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 96)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:], uint64(t.Twelve))
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 96 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 104)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:96], uint64(t.Twelve))
	jay.WriteUint64(b[96:], uint64(t.Thirteen))
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 104 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	t.Thirteen = long(jay.ReadUint64(b[96:104]))
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 112)
	jay.WriteUint64(b[:8], uint64(f.One))
	jay.WriteUint64(b[8:16], uint64(f.Two))
	jay.WriteUint64(b[16:24], uint64(f.Three))
	jay.WriteUint64(b[24:32], uint64(f.Four))
	jay.WriteUint64(b[32:40], uint64(f.Five))
	jay.WriteUint64(b[40:48], uint64(f.Six))
	jay.WriteUint64(b[48:56], uint64(f.Seven))
	jay.WriteUint64(b[56:64], uint64(f.Eight))
	jay.WriteUint64(b[64:72], uint64(f.Nine))
	jay.WriteUint64(b[72:80], uint64(f.Ten))
	jay.WriteUint64(b[80:88], uint64(f.Eleven))
	jay.WriteUint64(b[88:96], uint64(f.Twelve))
	jay.WriteUint64(b[96:104], uint64(f.Thirteen))
	jay.WriteUint64(b[104:], uint64(f.Fourteen))
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 112 {
		return jay.ErrUnexpectedEOB
	}
	f.One = long(jay.ReadUint64(b[:8]))
	f.Two = long(jay.ReadUint64(b[8:16]))
	f.Three = long(jay.ReadUint64(b[16:24]))
	f.Four = long(jay.ReadUint64(b[24:32]))
	f.Five = long(jay.ReadUint64(b[32:40]))
	f.Six = long(jay.ReadUint64(b[40:48]))
	f.Seven = long(jay.ReadUint64(b[48:56]))
	f.Eight = long(jay.ReadUint64(b[56:64]))
	f.Nine = long(jay.ReadUint64(b[64:72]))
	f.Ten = long(jay.ReadUint64(b[72:80]))
	f.Eleven = long(jay.ReadUint64(b[80:88]))
	f.Twelve = long(jay.ReadUint64(b[88:96]))
	f.Thirteen = long(jay.ReadUint64(b[96:104]))
	f.Fourteen = long(jay.ReadUint64(b[104:112]))
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 120)
	jay.WriteUint64(b[:8], uint64(f.One))
	jay.WriteUint64(b[8:16], uint64(f.Two))
	jay.WriteUint64(b[16:24], uint64(f.Three))
	jay.WriteUint64(b[24:32], uint64(f.Four))
	jay.WriteUint64(b[32:40], uint64(f.Five))
	jay.WriteUint64(b[40:48], uint64(f.Six))
	jay.WriteUint64(b[48:56], uint64(f.Seven))
	jay.WriteUint64(b[56:64], uint64(f.Eight))
	jay.WriteUint64(b[64:72], uint64(f.Nine))
	jay.WriteUint64(b[72:80], uint64(f.Ten))
	jay.WriteUint64(b[80:88], uint64(f.Eleven))
	jay.WriteUint64(b[88:96], uint64(f.Twelve))
	jay.WriteUint64(b[96:104], uint64(f.Thirteen))
	jay.WriteUint64(b[104:112], uint64(f.Fourteen))
	jay.WriteUint64(b[112:], uint64(f.Fifteen))
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 120 {
		return jay.ErrUnexpectedEOB
	}
	f.One = long(jay.ReadUint64(b[:8]))
	f.Two = long(jay.ReadUint64(b[8:16]))
	f.Three = long(jay.ReadUint64(b[16:24]))
	f.Four = long(jay.ReadUint64(b[24:32]))
	f.Five = long(jay.ReadUint64(b[32:40]))
	f.Six = long(jay.ReadUint64(b[40:48]))
	f.Seven = long(jay.ReadUint64(b[48:56]))
	f.Eight = long(jay.ReadUint64(b[56:64]))
	f.Nine = long(jay.ReadUint64(b[64:72]))
	f.Ten = long(jay.ReadUint64(b[72:80]))
	f.Eleven = long(jay.ReadUint64(b[80:88]))
	f.Twelve = long(jay.ReadUint64(b[88:96]))
	f.Thirteen = long(jay.ReadUint64(b[96:104]))
	f.Fourteen = long(jay.ReadUint64(b[104:112]))
	f.Fifteen = long(jay.ReadUint64(b[112:120]))
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 128)
	jay.WriteUint64(b[:8], uint64(s.One))
	jay.WriteUint64(b[8:16], uint64(s.Two))
	jay.WriteUint64(b[16:24], uint64(s.Three))
	jay.WriteUint64(b[24:32], uint64(s.Four))
	jay.WriteUint64(b[32:40], uint64(s.Five))
	jay.WriteUint64(b[40:48], uint64(s.Six))
	jay.WriteUint64(b[48:56], uint64(s.Seven))
	jay.WriteUint64(b[56:64], uint64(s.Eight))
	jay.WriteUint64(b[64:72], uint64(s.Nine))
	jay.WriteUint64(b[72:80], uint64(s.Ten))
	jay.WriteUint64(b[80:88], uint64(s.Eleven))
	jay.WriteUint64(b[88:96], uint64(s.Twelve))
	jay.WriteUint64(b[96:104], uint64(s.Thirteen))
	jay.WriteUint64(b[104:112], uint64(s.Fourteen))
	jay.WriteUint64(b[112:120], uint64(s.Fifteen))
	jay.WriteUint64(b[120:], uint64(s.Sixteen))
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 128 {
		return jay.ErrUnexpectedEOB
	}
	s.One = long(jay.ReadUint64(b[:8]))
	s.Two = long(jay.ReadUint64(b[8:16]))
	s.Three = long(jay.ReadUint64(b[16:24]))
	s.Four = long(jay.ReadUint64(b[24:32]))
	s.Five = long(jay.ReadUint64(b[32:40]))
	s.Six = long(jay.ReadUint64(b[40:48]))
	s.Seven = long(jay.ReadUint64(b[48:56]))
	s.Eight = long(jay.ReadUint64(b[56:64]))
	s.Nine = long(jay.ReadUint64(b[64:72]))
	s.Ten = long(jay.ReadUint64(b[72:80]))
	s.Eleven = long(jay.ReadUint64(b[80:88]))
	s.Twelve = long(jay.ReadUint64(b[88:96]))
	s.Thirteen = long(jay.ReadUint64(b[96:104]))
	s.Fourteen = long(jay.ReadUint64(b[104:112]))
	s.Fifteen = long(jay.ReadUint64(b[112:120]))
	s.Sixteen = long(jay.ReadUint64(b[120:128]))
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 136)
	jay.WriteUint64(b[:8], uint64(s.One))
	jay.WriteUint64(b[8:16], uint64(s.Two))
	jay.WriteUint64(b[16:24], uint64(s.Three))
	jay.WriteUint64(b[24:32], uint64(s.Four))
	jay.WriteUint64(b[32:40], uint64(s.Five))
	jay.WriteUint64(b[40:48], uint64(s.Six))
	jay.WriteUint64(b[48:56], uint64(s.Seven))
	jay.WriteUint64(b[56:64], uint64(s.Eight))
	jay.WriteUint64(b[64:72], uint64(s.Nine))
	jay.WriteUint64(b[72:80], uint64(s.Ten))
	jay.WriteUint64(b[80:88], uint64(s.Eleven))
	jay.WriteUint64(b[88:96], uint64(s.Twelve))
	jay.WriteUint64(b[96:104], uint64(s.Thirteen))
	jay.WriteUint64(b[104:112], uint64(s.Fourteen))
	jay.WriteUint64(b[112:120], uint64(s.Fifteen))
	jay.WriteUint64(b[120:128], uint64(s.Sixteen))
	jay.WriteUint64(b[128:], uint64(s.Seventeen))
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 136 {
		return jay.ErrUnexpectedEOB
	}
	s.One = long(jay.ReadUint64(b[:8]))
	s.Two = long(jay.ReadUint64(b[8:16]))
	s.Three = long(jay.ReadUint64(b[16:24]))
	s.Four = long(jay.ReadUint64(b[24:32]))
	s.Five = long(jay.ReadUint64(b[32:40]))
	s.Six = long(jay.ReadUint64(b[40:48]))
	s.Seven = long(jay.ReadUint64(b[48:56]))
	s.Eight = long(jay.ReadUint64(b[56:64]))
	s.Nine = long(jay.ReadUint64(b[64:72]))
	s.Ten = long(jay.ReadUint64(b[72:80]))
	s.Eleven = long(jay.ReadUint64(b[80:88]))
	s.Twelve = long(jay.ReadUint64(b[88:96]))
	s.Thirteen = long(jay.ReadUint64(b[96:104]))
	s.Fourteen = long(jay.ReadUint64(b[104:112]))
	s.Fifteen = long(jay.ReadUint64(b[112:120]))
	s.Sixteen = long(jay.ReadUint64(b[120:128]))
	s.Seventeen = long(jay.ReadUint64(b[128:136]))
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 144)
	jay.WriteUint64(b[:8], uint64(e.One))
	jay.WriteUint64(b[8:16], uint64(e.Two))
	jay.WriteUint64(b[16:24], uint64(e.Three))
	jay.WriteUint64(b[24:32], uint64(e.Four))
	jay.WriteUint64(b[32:40], uint64(e.Five))
	jay.WriteUint64(b[40:48], uint64(e.Six))
	jay.WriteUint64(b[48:56], uint64(e.Seven))
	jay.WriteUint64(b[56:64], uint64(e.Eight))
	jay.WriteUint64(b[64:72], uint64(e.Nine))
	jay.WriteUint64(b[72:80], uint64(e.Ten))
	jay.WriteUint64(b[80:88], uint64(e.Eleven))
	jay.WriteUint64(b[88:96], uint64(e.Twelve))
	jay.WriteUint64(b[96:104], uint64(e.Thirteen))
	jay.WriteUint64(b[104:112], uint64(e.Fourteen))
	jay.WriteUint64(b[112:120], uint64(e.Fifteen))
	jay.WriteUint64(b[120:128], uint64(e.Sixteen))
	jay.WriteUint64(b[128:136], uint64(e.Seventeen))
	jay.WriteUint64(b[136:], uint64(e.Eighteen))
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 144 {
		return jay.ErrUnexpectedEOB
	}
	e.One = long(jay.ReadUint64(b[:8]))
	e.Two = long(jay.ReadUint64(b[8:16]))
	e.Three = long(jay.ReadUint64(b[16:24]))
	e.Four = long(jay.ReadUint64(b[24:32]))
	e.Five = long(jay.ReadUint64(b[32:40]))
	e.Six = long(jay.ReadUint64(b[40:48]))
	e.Seven = long(jay.ReadUint64(b[48:56]))
	e.Eight = long(jay.ReadUint64(b[56:64]))
	e.Nine = long(jay.ReadUint64(b[64:72]))
	e.Ten = long(jay.ReadUint64(b[72:80]))
	e.Eleven = long(jay.ReadUint64(b[80:88]))
	e.Twelve = long(jay.ReadUint64(b[88:96]))
	e.Thirteen = long(jay.ReadUint64(b[96:104]))
	e.Fourteen = long(jay.ReadUint64(b[104:112]))
	e.Fifteen = long(jay.ReadUint64(b[112:120]))
	e.Sixteen = long(jay.ReadUint64(b[120:128]))
	e.Seventeen = long(jay.ReadUint64(b[128:136]))
	e.Eighteen = long(jay.ReadUint64(b[136:144]))
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 152)
	jay.WriteUint64(b[:8], uint64(n.One))
	jay.WriteUint64(b[8:16], uint64(n.Two))
	jay.WriteUint64(b[16:24], uint64(n.Three))
	jay.WriteUint64(b[24:32], uint64(n.Four))
	jay.WriteUint64(b[32:40], uint64(n.Five))
	jay.WriteUint64(b[40:48], uint64(n.Six))
	jay.WriteUint64(b[48:56], uint64(n.Seven))
	jay.WriteUint64(b[56:64], uint64(n.Eight))
	jay.WriteUint64(b[64:72], uint64(n.Nine))
	jay.WriteUint64(b[72:80], uint64(n.Ten))
	jay.WriteUint64(b[80:88], uint64(n.Eleven))
	jay.WriteUint64(b[88:96], uint64(n.Twelve))
	jay.WriteUint64(b[96:104], uint64(n.Thirteen))
	jay.WriteUint64(b[104:112], uint64(n.Fourteen))
	jay.WriteUint64(b[112:120], uint64(n.Fifteen))
	jay.WriteUint64(b[120:128], uint64(n.Sixteen))
	jay.WriteUint64(b[128:136], uint64(n.Seventeen))
	jay.WriteUint64(b[136:144], uint64(n.Eighteen))
	jay.WriteUint64(b[144:], uint64(n.Nineteen))
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 152 {
		return jay.ErrUnexpectedEOB
	}
	n.One = long(jay.ReadUint64(b[:8]))
	n.Two = long(jay.ReadUint64(b[8:16]))
	n.Three = long(jay.ReadUint64(b[16:24]))
	n.Four = long(jay.ReadUint64(b[24:32]))
	n.Five = long(jay.ReadUint64(b[32:40]))
	n.Six = long(jay.ReadUint64(b[40:48]))
	n.Seven = long(jay.ReadUint64(b[48:56]))
	n.Eight = long(jay.ReadUint64(b[56:64]))
	n.Nine = long(jay.ReadUint64(b[64:72]))
	n.Ten = long(jay.ReadUint64(b[72:80]))
	n.Eleven = long(jay.ReadUint64(b[80:88]))
	n.Twelve = long(jay.ReadUint64(b[88:96]))
	n.Thirteen = long(jay.ReadUint64(b[96:104]))
	n.Fourteen = long(jay.ReadUint64(b[104:112]))
	n.Fifteen = long(jay.ReadUint64(b[112:120]))
	n.Sixteen = long(jay.ReadUint64(b[120:128]))
	n.Seventeen = long(jay.ReadUint64(b[128:136]))
	n.Eighteen = long(jay.ReadUint64(b[136:144]))
	n.Nineteen = long(jay.ReadUint64(b[144:152]))
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 160)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:96], uint64(t.Twelve))
	jay.WriteUint64(b[96:104], uint64(t.Thirteen))
	jay.WriteUint64(b[104:112], uint64(t.Fourteen))
	jay.WriteUint64(b[112:120], uint64(t.Fifteen))
	jay.WriteUint64(b[120:128], uint64(t.Sixteen))
	jay.WriteUint64(b[128:136], uint64(t.Seventeen))
	jay.WriteUint64(b[136:144], uint64(t.Eighteen))
	jay.WriteUint64(b[144:152], uint64(t.Nineteen))
	jay.WriteUint64(b[152:], uint64(t.Twenty))
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 160 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	t.Thirteen = long(jay.ReadUint64(b[96:104]))
	t.Fourteen = long(jay.ReadUint64(b[104:112]))
	t.Fifteen = long(jay.ReadUint64(b[112:120]))
	t.Sixteen = long(jay.ReadUint64(b[120:128]))
	t.Seventeen = long(jay.ReadUint64(b[128:136]))
	t.Eighteen = long(jay.ReadUint64(b[136:144]))
	t.Nineteen = long(jay.ReadUint64(b[144:152]))
	t.Twenty = long(jay.ReadUint64(b[152:160]))
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 168)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:96], uint64(t.Twelve))
	jay.WriteUint64(b[96:104], uint64(t.Thirteen))
	jay.WriteUint64(b[104:112], uint64(t.Fourteen))
	jay.WriteUint64(b[112:120], uint64(t.Fifteen))
	jay.WriteUint64(b[120:128], uint64(t.Sixteen))
	jay.WriteUint64(b[128:136], uint64(t.Seventeen))
	jay.WriteUint64(b[136:144], uint64(t.Eighteen))
	jay.WriteUint64(b[144:152], uint64(t.Nineteen))
	jay.WriteUint64(b[152:160], uint64(t.Twenty))
	jay.WriteUint64(b[160:], uint64(t.TwentyOne))
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 168 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	t.Thirteen = long(jay.ReadUint64(b[96:104]))
	t.Fourteen = long(jay.ReadUint64(b[104:112]))
	t.Fifteen = long(jay.ReadUint64(b[112:120]))
	t.Sixteen = long(jay.ReadUint64(b[120:128]))
	t.Seventeen = long(jay.ReadUint64(b[128:136]))
	t.Eighteen = long(jay.ReadUint64(b[136:144]))
	t.Nineteen = long(jay.ReadUint64(b[144:152]))
	t.Twenty = long(jay.ReadUint64(b[152:160]))
	t.TwentyOne = long(jay.ReadUint64(b[160:168]))
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 176)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:96], uint64(t.Twelve))
	jay.WriteUint64(b[96:104], uint64(t.Thirteen))
	jay.WriteUint64(b[104:112], uint64(t.Fourteen))
	jay.WriteUint64(b[112:120], uint64(t.Fifteen))
	jay.WriteUint64(b[120:128], uint64(t.Sixteen))
	jay.WriteUint64(b[128:136], uint64(t.Seventeen))
	jay.WriteUint64(b[136:144], uint64(t.Eighteen))
	jay.WriteUint64(b[144:152], uint64(t.Nineteen))
	jay.WriteUint64(b[152:160], uint64(t.Twenty))
	jay.WriteUint64(b[160:168], uint64(t.TwentyOne))
	jay.WriteUint64(b[168:], uint64(t.TwentyTwo))
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 176 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	t.Thirteen = long(jay.ReadUint64(b[96:104]))
	t.Fourteen = long(jay.ReadUint64(b[104:112]))
	t.Fifteen = long(jay.ReadUint64(b[112:120]))
	t.Sixteen = long(jay.ReadUint64(b[120:128]))
	t.Seventeen = long(jay.ReadUint64(b[128:136]))
	t.Eighteen = long(jay.ReadUint64(b[136:144]))
	t.Nineteen = long(jay.ReadUint64(b[144:152]))
	t.Twenty = long(jay.ReadUint64(b[152:160]))
	t.TwentyOne = long(jay.ReadUint64(b[160:168]))
	t.TwentyTwo = long(jay.ReadUint64(b[168:176]))
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 184)
	jay.WriteUint64(b[:8], uint64(t.One))
	jay.WriteUint64(b[8:16], uint64(t.Two))
	jay.WriteUint64(b[16:24], uint64(t.Three))
	jay.WriteUint64(b[24:32], uint64(t.Four))
	jay.WriteUint64(b[32:40], uint64(t.Five))
	jay.WriteUint64(b[40:48], uint64(t.Six))
	jay.WriteUint64(b[48:56], uint64(t.Seven))
	jay.WriteUint64(b[56:64], uint64(t.Eight))
	jay.WriteUint64(b[64:72], uint64(t.Nine))
	jay.WriteUint64(b[72:80], uint64(t.Ten))
	jay.WriteUint64(b[80:88], uint64(t.Eleven))
	jay.WriteUint64(b[88:96], uint64(t.Twelve))
	jay.WriteUint64(b[96:104], uint64(t.Thirteen))
	jay.WriteUint64(b[104:112], uint64(t.Fourteen))
	jay.WriteUint64(b[112:120], uint64(t.Fifteen))
	jay.WriteUint64(b[120:128], uint64(t.Sixteen))
	jay.WriteUint64(b[128:136], uint64(t.Seventeen))
	jay.WriteUint64(b[136:144], uint64(t.Eighteen))
	jay.WriteUint64(b[144:152], uint64(t.Nineteen))
	jay.WriteUint64(b[152:160], uint64(t.Twenty))
	jay.WriteUint64(b[160:168], uint64(t.TwentyOne))
	jay.WriteUint64(b[168:176], uint64(t.TwentyTwo))
	jay.WriteUint64(b[176:], uint64(t.TwentyThree))
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 184 {
		return jay.ErrUnexpectedEOB
	}
	t.One = long(jay.ReadUint64(b[:8]))
	t.Two = long(jay.ReadUint64(b[8:16]))
	t.Three = long(jay.ReadUint64(b[16:24]))
	t.Four = long(jay.ReadUint64(b[24:32]))
	t.Five = long(jay.ReadUint64(b[32:40]))
	t.Six = long(jay.ReadUint64(b[40:48]))
	t.Seven = long(jay.ReadUint64(b[48:56]))
	t.Eight = long(jay.ReadUint64(b[56:64]))
	t.Nine = long(jay.ReadUint64(b[64:72]))
	t.Ten = long(jay.ReadUint64(b[72:80]))
	t.Eleven = long(jay.ReadUint64(b[80:88]))
	t.Twelve = long(jay.ReadUint64(b[88:96]))
	t.Thirteen = long(jay.ReadUint64(b[96:104]))
	t.Fourteen = long(jay.ReadUint64(b[104:112]))
	t.Fifteen = long(jay.ReadUint64(b[112:120]))
	t.Sixteen = long(jay.ReadUint64(b[120:128]))
	t.Seventeen = long(jay.ReadUint64(b[128:136]))
	t.Eighteen = long(jay.ReadUint64(b[136:144]))
	t.Nineteen = long(jay.ReadUint64(b[144:152]))
	t.Twenty = long(jay.ReadUint64(b[152:160]))
	t.TwentyOne = long(jay.ReadUint64(b[160:168]))
	t.TwentyTwo = long(jay.ReadUint64(b[168:176]))
	t.TwentyThree = long(jay.ReadUint64(b[176:184]))
	return nil
}