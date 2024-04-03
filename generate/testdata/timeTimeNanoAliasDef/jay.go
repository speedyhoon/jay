// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import (
	"time"

	"github.com/speedyhoon/jay"
)

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteTimeNano(b, time.Time(o.One))
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	o.One = tim(jay.ReadTimeNano(b))
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:], time.Time(t.Two))
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:]))
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:], time.Time(t.Three))
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:]))
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteTimeNano(b[:8], time.Time(f.One))
	jay.WriteTimeNano(b[8:16], time.Time(f.Two))
	jay.WriteTimeNano(b[16:24], time.Time(f.Three))
	jay.WriteTimeNano(b[24:], time.Time(f.Four))
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tim(jay.ReadTimeNano(b[:8]))
	f.Two = tim(jay.ReadTimeNano(b[8:16]))
	f.Three = tim(jay.ReadTimeNano(b[16:24]))
	f.Four = tim(jay.ReadTimeNano(b[24:]))
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteTimeNano(b[:8], time.Time(f.One))
	jay.WriteTimeNano(b[8:16], time.Time(f.Two))
	jay.WriteTimeNano(b[16:24], time.Time(f.Three))
	jay.WriteTimeNano(b[24:32], time.Time(f.Four))
	jay.WriteTimeNano(b[32:], time.Time(f.Five))
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tim(jay.ReadTimeNano(b[:8]))
	f.Two = tim(jay.ReadTimeNano(b[8:16]))
	f.Three = tim(jay.ReadTimeNano(b[16:24]))
	f.Four = tim(jay.ReadTimeNano(b[24:32]))
	f.Five = tim(jay.ReadTimeNano(b[32:]))
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 48)
	jay.WriteTimeNano(b[:8], time.Time(s.One))
	jay.WriteTimeNano(b[8:16], time.Time(s.Two))
	jay.WriteTimeNano(b[16:24], time.Time(s.Three))
	jay.WriteTimeNano(b[24:32], time.Time(s.Four))
	jay.WriteTimeNano(b[32:40], time.Time(s.Five))
	jay.WriteTimeNano(b[40:], time.Time(s.Six))
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 48 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tim(jay.ReadTimeNano(b[:8]))
	s.Two = tim(jay.ReadTimeNano(b[8:16]))
	s.Three = tim(jay.ReadTimeNano(b[16:24]))
	s.Four = tim(jay.ReadTimeNano(b[24:32]))
	s.Five = tim(jay.ReadTimeNano(b[32:40]))
	s.Six = tim(jay.ReadTimeNano(b[40:]))
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 56)
	jay.WriteTimeNano(b[:8], time.Time(s.One))
	jay.WriteTimeNano(b[8:16], time.Time(s.Two))
	jay.WriteTimeNano(b[16:24], time.Time(s.Three))
	jay.WriteTimeNano(b[24:32], time.Time(s.Four))
	jay.WriteTimeNano(b[32:40], time.Time(s.Five))
	jay.WriteTimeNano(b[40:48], time.Time(s.Six))
	jay.WriteTimeNano(b[48:], time.Time(s.Seven))
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 56 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tim(jay.ReadTimeNano(b[:8]))
	s.Two = tim(jay.ReadTimeNano(b[8:16]))
	s.Three = tim(jay.ReadTimeNano(b[16:24]))
	s.Four = tim(jay.ReadTimeNano(b[24:32]))
	s.Five = tim(jay.ReadTimeNano(b[32:40]))
	s.Six = tim(jay.ReadTimeNano(b[40:48]))
	s.Seven = tim(jay.ReadTimeNano(b[48:]))
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 64)
	jay.WriteTimeNano(b[:8], time.Time(e.One))
	jay.WriteTimeNano(b[8:16], time.Time(e.Two))
	jay.WriteTimeNano(b[16:24], time.Time(e.Three))
	jay.WriteTimeNano(b[24:32], time.Time(e.Four))
	jay.WriteTimeNano(b[32:40], time.Time(e.Five))
	jay.WriteTimeNano(b[40:48], time.Time(e.Six))
	jay.WriteTimeNano(b[48:56], time.Time(e.Seven))
	jay.WriteTimeNano(b[56:], time.Time(e.Eight))
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 64 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tim(jay.ReadTimeNano(b[:8]))
	e.Two = tim(jay.ReadTimeNano(b[8:16]))
	e.Three = tim(jay.ReadTimeNano(b[16:24]))
	e.Four = tim(jay.ReadTimeNano(b[24:32]))
	e.Five = tim(jay.ReadTimeNano(b[32:40]))
	e.Six = tim(jay.ReadTimeNano(b[40:48]))
	e.Seven = tim(jay.ReadTimeNano(b[48:56]))
	e.Eight = tim(jay.ReadTimeNano(b[56:]))
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 72)
	jay.WriteTimeNano(b[:8], time.Time(n.One))
	jay.WriteTimeNano(b[8:16], time.Time(n.Two))
	jay.WriteTimeNano(b[16:24], time.Time(n.Three))
	jay.WriteTimeNano(b[24:32], time.Time(n.Four))
	jay.WriteTimeNano(b[32:40], time.Time(n.Five))
	jay.WriteTimeNano(b[40:48], time.Time(n.Six))
	jay.WriteTimeNano(b[48:56], time.Time(n.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(n.Eight))
	jay.WriteTimeNano(b[64:], time.Time(n.Nine))
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 72 {
		return jay.ErrUnexpectedEOB
	}
	n.One = tim(jay.ReadTimeNano(b[:8]))
	n.Two = tim(jay.ReadTimeNano(b[8:16]))
	n.Three = tim(jay.ReadTimeNano(b[16:24]))
	n.Four = tim(jay.ReadTimeNano(b[24:32]))
	n.Five = tim(jay.ReadTimeNano(b[32:40]))
	n.Six = tim(jay.ReadTimeNano(b[40:48]))
	n.Seven = tim(jay.ReadTimeNano(b[48:56]))
	n.Eight = tim(jay.ReadTimeNano(b[56:64]))
	n.Nine = tim(jay.ReadTimeNano(b[64:]))
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 80)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:], time.Time(t.Ten))
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 80 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:]))
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 88)
	jay.WriteTimeNano(b[:8], time.Time(e.One))
	jay.WriteTimeNano(b[8:16], time.Time(e.Two))
	jay.WriteTimeNano(b[16:24], time.Time(e.Three))
	jay.WriteTimeNano(b[24:32], time.Time(e.Four))
	jay.WriteTimeNano(b[32:40], time.Time(e.Five))
	jay.WriteTimeNano(b[40:48], time.Time(e.Six))
	jay.WriteTimeNano(b[48:56], time.Time(e.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(e.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(e.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(e.Ten))
	jay.WriteTimeNano(b[80:], time.Time(e.Eleven))
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 88 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tim(jay.ReadTimeNano(b[:8]))
	e.Two = tim(jay.ReadTimeNano(b[8:16]))
	e.Three = tim(jay.ReadTimeNano(b[16:24]))
	e.Four = tim(jay.ReadTimeNano(b[24:32]))
	e.Five = tim(jay.ReadTimeNano(b[32:40]))
	e.Six = tim(jay.ReadTimeNano(b[40:48]))
	e.Seven = tim(jay.ReadTimeNano(b[48:56]))
	e.Eight = tim(jay.ReadTimeNano(b[56:64]))
	e.Nine = tim(jay.ReadTimeNano(b[64:72]))
	e.Ten = tim(jay.ReadTimeNano(b[72:80]))
	e.Eleven = tim(jay.ReadTimeNano(b[80:]))
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 96)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:], time.Time(t.Twelve))
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 96 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:]))
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 104)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(t.Twelve))
	jay.WriteTimeNano(b[96:], time.Time(t.Thirteen))
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 104 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	t.Thirteen = tim(jay.ReadTimeNano(b[96:]))
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 112)
	jay.WriteTimeNano(b[:8], time.Time(f.One))
	jay.WriteTimeNano(b[8:16], time.Time(f.Two))
	jay.WriteTimeNano(b[16:24], time.Time(f.Three))
	jay.WriteTimeNano(b[24:32], time.Time(f.Four))
	jay.WriteTimeNano(b[32:40], time.Time(f.Five))
	jay.WriteTimeNano(b[40:48], time.Time(f.Six))
	jay.WriteTimeNano(b[48:56], time.Time(f.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(f.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(f.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(f.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(f.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(f.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(f.Thirteen))
	jay.WriteTimeNano(b[104:], time.Time(f.Fourteen))
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 112 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tim(jay.ReadTimeNano(b[:8]))
	f.Two = tim(jay.ReadTimeNano(b[8:16]))
	f.Three = tim(jay.ReadTimeNano(b[16:24]))
	f.Four = tim(jay.ReadTimeNano(b[24:32]))
	f.Five = tim(jay.ReadTimeNano(b[32:40]))
	f.Six = tim(jay.ReadTimeNano(b[40:48]))
	f.Seven = tim(jay.ReadTimeNano(b[48:56]))
	f.Eight = tim(jay.ReadTimeNano(b[56:64]))
	f.Nine = tim(jay.ReadTimeNano(b[64:72]))
	f.Ten = tim(jay.ReadTimeNano(b[72:80]))
	f.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	f.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	f.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	f.Fourteen = tim(jay.ReadTimeNano(b[104:]))
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 120)
	jay.WriteTimeNano(b[:8], time.Time(f.One))
	jay.WriteTimeNano(b[8:16], time.Time(f.Two))
	jay.WriteTimeNano(b[16:24], time.Time(f.Three))
	jay.WriteTimeNano(b[24:32], time.Time(f.Four))
	jay.WriteTimeNano(b[32:40], time.Time(f.Five))
	jay.WriteTimeNano(b[40:48], time.Time(f.Six))
	jay.WriteTimeNano(b[48:56], time.Time(f.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(f.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(f.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(f.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(f.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(f.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(f.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(f.Fourteen))
	jay.WriteTimeNano(b[112:], time.Time(f.Fifteen))
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 120 {
		return jay.ErrUnexpectedEOB
	}
	f.One = tim(jay.ReadTimeNano(b[:8]))
	f.Two = tim(jay.ReadTimeNano(b[8:16]))
	f.Three = tim(jay.ReadTimeNano(b[16:24]))
	f.Four = tim(jay.ReadTimeNano(b[24:32]))
	f.Five = tim(jay.ReadTimeNano(b[32:40]))
	f.Six = tim(jay.ReadTimeNano(b[40:48]))
	f.Seven = tim(jay.ReadTimeNano(b[48:56]))
	f.Eight = tim(jay.ReadTimeNano(b[56:64]))
	f.Nine = tim(jay.ReadTimeNano(b[64:72]))
	f.Ten = tim(jay.ReadTimeNano(b[72:80]))
	f.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	f.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	f.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	f.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	f.Fifteen = tim(jay.ReadTimeNano(b[112:]))
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 128)
	jay.WriteTimeNano(b[:8], time.Time(s.One))
	jay.WriteTimeNano(b[8:16], time.Time(s.Two))
	jay.WriteTimeNano(b[16:24], time.Time(s.Three))
	jay.WriteTimeNano(b[24:32], time.Time(s.Four))
	jay.WriteTimeNano(b[32:40], time.Time(s.Five))
	jay.WriteTimeNano(b[40:48], time.Time(s.Six))
	jay.WriteTimeNano(b[48:56], time.Time(s.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(s.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(s.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(s.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(s.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(s.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(s.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(s.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(s.Fifteen))
	jay.WriteTimeNano(b[120:], time.Time(s.Sixteen))
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 128 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tim(jay.ReadTimeNano(b[:8]))
	s.Two = tim(jay.ReadTimeNano(b[8:16]))
	s.Three = tim(jay.ReadTimeNano(b[16:24]))
	s.Four = tim(jay.ReadTimeNano(b[24:32]))
	s.Five = tim(jay.ReadTimeNano(b[32:40]))
	s.Six = tim(jay.ReadTimeNano(b[40:48]))
	s.Seven = tim(jay.ReadTimeNano(b[48:56]))
	s.Eight = tim(jay.ReadTimeNano(b[56:64]))
	s.Nine = tim(jay.ReadTimeNano(b[64:72]))
	s.Ten = tim(jay.ReadTimeNano(b[72:80]))
	s.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	s.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	s.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	s.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	s.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	s.Sixteen = tim(jay.ReadTimeNano(b[120:]))
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 136)
	jay.WriteTimeNano(b[:8], time.Time(s.One))
	jay.WriteTimeNano(b[8:16], time.Time(s.Two))
	jay.WriteTimeNano(b[16:24], time.Time(s.Three))
	jay.WriteTimeNano(b[24:32], time.Time(s.Four))
	jay.WriteTimeNano(b[32:40], time.Time(s.Five))
	jay.WriteTimeNano(b[40:48], time.Time(s.Six))
	jay.WriteTimeNano(b[48:56], time.Time(s.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(s.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(s.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(s.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(s.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(s.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(s.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(s.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(s.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(s.Sixteen))
	jay.WriteTimeNano(b[128:], time.Time(s.Seventeen))
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 136 {
		return jay.ErrUnexpectedEOB
	}
	s.One = tim(jay.ReadTimeNano(b[:8]))
	s.Two = tim(jay.ReadTimeNano(b[8:16]))
	s.Three = tim(jay.ReadTimeNano(b[16:24]))
	s.Four = tim(jay.ReadTimeNano(b[24:32]))
	s.Five = tim(jay.ReadTimeNano(b[32:40]))
	s.Six = tim(jay.ReadTimeNano(b[40:48]))
	s.Seven = tim(jay.ReadTimeNano(b[48:56]))
	s.Eight = tim(jay.ReadTimeNano(b[56:64]))
	s.Nine = tim(jay.ReadTimeNano(b[64:72]))
	s.Ten = tim(jay.ReadTimeNano(b[72:80]))
	s.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	s.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	s.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	s.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	s.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	s.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	s.Seventeen = tim(jay.ReadTimeNano(b[128:]))
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 144)
	jay.WriteTimeNano(b[:8], time.Time(e.One))
	jay.WriteTimeNano(b[8:16], time.Time(e.Two))
	jay.WriteTimeNano(b[16:24], time.Time(e.Three))
	jay.WriteTimeNano(b[24:32], time.Time(e.Four))
	jay.WriteTimeNano(b[32:40], time.Time(e.Five))
	jay.WriteTimeNano(b[40:48], time.Time(e.Six))
	jay.WriteTimeNano(b[48:56], time.Time(e.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(e.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(e.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(e.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(e.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(e.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(e.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(e.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(e.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(e.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(e.Seventeen))
	jay.WriteTimeNano(b[136:], time.Time(e.Eighteen))
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 144 {
		return jay.ErrUnexpectedEOB
	}
	e.One = tim(jay.ReadTimeNano(b[:8]))
	e.Two = tim(jay.ReadTimeNano(b[8:16]))
	e.Three = tim(jay.ReadTimeNano(b[16:24]))
	e.Four = tim(jay.ReadTimeNano(b[24:32]))
	e.Five = tim(jay.ReadTimeNano(b[32:40]))
	e.Six = tim(jay.ReadTimeNano(b[40:48]))
	e.Seven = tim(jay.ReadTimeNano(b[48:56]))
	e.Eight = tim(jay.ReadTimeNano(b[56:64]))
	e.Nine = tim(jay.ReadTimeNano(b[64:72]))
	e.Ten = tim(jay.ReadTimeNano(b[72:80]))
	e.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	e.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	e.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	e.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	e.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	e.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	e.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	e.Eighteen = tim(jay.ReadTimeNano(b[136:]))
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 152)
	jay.WriteTimeNano(b[:8], time.Time(n.One))
	jay.WriteTimeNano(b[8:16], time.Time(n.Two))
	jay.WriteTimeNano(b[16:24], time.Time(n.Three))
	jay.WriteTimeNano(b[24:32], time.Time(n.Four))
	jay.WriteTimeNano(b[32:40], time.Time(n.Five))
	jay.WriteTimeNano(b[40:48], time.Time(n.Six))
	jay.WriteTimeNano(b[48:56], time.Time(n.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(n.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(n.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(n.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(n.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(n.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(n.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(n.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(n.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(n.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(n.Seventeen))
	jay.WriteTimeNano(b[136:144], time.Time(n.Eighteen))
	jay.WriteTimeNano(b[144:], time.Time(n.Nineteen))
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 152 {
		return jay.ErrUnexpectedEOB
	}
	n.One = tim(jay.ReadTimeNano(b[:8]))
	n.Two = tim(jay.ReadTimeNano(b[8:16]))
	n.Three = tim(jay.ReadTimeNano(b[16:24]))
	n.Four = tim(jay.ReadTimeNano(b[24:32]))
	n.Five = tim(jay.ReadTimeNano(b[32:40]))
	n.Six = tim(jay.ReadTimeNano(b[40:48]))
	n.Seven = tim(jay.ReadTimeNano(b[48:56]))
	n.Eight = tim(jay.ReadTimeNano(b[56:64]))
	n.Nine = tim(jay.ReadTimeNano(b[64:72]))
	n.Ten = tim(jay.ReadTimeNano(b[72:80]))
	n.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	n.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	n.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	n.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	n.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	n.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	n.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	n.Eighteen = tim(jay.ReadTimeNano(b[136:144]))
	n.Nineteen = tim(jay.ReadTimeNano(b[144:]))
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 160)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(t.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(t.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(t.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(t.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(t.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(t.Seventeen))
	jay.WriteTimeNano(b[136:144], time.Time(t.Eighteen))
	jay.WriteTimeNano(b[144:152], time.Time(t.Nineteen))
	jay.WriteTimeNano(b[152:], time.Time(t.Twenty))
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 160 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	t.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	t.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	t.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	t.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	t.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	t.Eighteen = tim(jay.ReadTimeNano(b[136:144]))
	t.Nineteen = tim(jay.ReadTimeNano(b[144:152]))
	t.Twenty = tim(jay.ReadTimeNano(b[152:]))
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 168)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(t.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(t.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(t.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(t.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(t.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(t.Seventeen))
	jay.WriteTimeNano(b[136:144], time.Time(t.Eighteen))
	jay.WriteTimeNano(b[144:152], time.Time(t.Nineteen))
	jay.WriteTimeNano(b[152:160], time.Time(t.Twenty))
	jay.WriteTimeNano(b[160:], time.Time(t.TwentyOne))
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 168 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	t.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	t.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	t.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	t.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	t.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	t.Eighteen = tim(jay.ReadTimeNano(b[136:144]))
	t.Nineteen = tim(jay.ReadTimeNano(b[144:152]))
	t.Twenty = tim(jay.ReadTimeNano(b[152:160]))
	t.TwentyOne = tim(jay.ReadTimeNano(b[160:]))
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 176)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(t.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(t.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(t.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(t.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(t.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(t.Seventeen))
	jay.WriteTimeNano(b[136:144], time.Time(t.Eighteen))
	jay.WriteTimeNano(b[144:152], time.Time(t.Nineteen))
	jay.WriteTimeNano(b[152:160], time.Time(t.Twenty))
	jay.WriteTimeNano(b[160:168], time.Time(t.TwentyOne))
	jay.WriteTimeNano(b[168:], time.Time(t.TwentyTwo))
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 176 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	t.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	t.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	t.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	t.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	t.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	t.Eighteen = tim(jay.ReadTimeNano(b[136:144]))
	t.Nineteen = tim(jay.ReadTimeNano(b[144:152]))
	t.Twenty = tim(jay.ReadTimeNano(b[152:160]))
	t.TwentyOne = tim(jay.ReadTimeNano(b[160:168]))
	t.TwentyTwo = tim(jay.ReadTimeNano(b[168:]))
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 184)
	jay.WriteTimeNano(b[:8], time.Time(t.One))
	jay.WriteTimeNano(b[8:16], time.Time(t.Two))
	jay.WriteTimeNano(b[16:24], time.Time(t.Three))
	jay.WriteTimeNano(b[24:32], time.Time(t.Four))
	jay.WriteTimeNano(b[32:40], time.Time(t.Five))
	jay.WriteTimeNano(b[40:48], time.Time(t.Six))
	jay.WriteTimeNano(b[48:56], time.Time(t.Seven))
	jay.WriteTimeNano(b[56:64], time.Time(t.Eight))
	jay.WriteTimeNano(b[64:72], time.Time(t.Nine))
	jay.WriteTimeNano(b[72:80], time.Time(t.Ten))
	jay.WriteTimeNano(b[80:88], time.Time(t.Eleven))
	jay.WriteTimeNano(b[88:96], time.Time(t.Twelve))
	jay.WriteTimeNano(b[96:104], time.Time(t.Thirteen))
	jay.WriteTimeNano(b[104:112], time.Time(t.Fourteen))
	jay.WriteTimeNano(b[112:120], time.Time(t.Fifteen))
	jay.WriteTimeNano(b[120:128], time.Time(t.Sixteen))
	jay.WriteTimeNano(b[128:136], time.Time(t.Seventeen))
	jay.WriteTimeNano(b[136:144], time.Time(t.Eighteen))
	jay.WriteTimeNano(b[144:152], time.Time(t.Nineteen))
	jay.WriteTimeNano(b[152:160], time.Time(t.Twenty))
	jay.WriteTimeNano(b[160:168], time.Time(t.TwentyOne))
	jay.WriteTimeNano(b[168:176], time.Time(t.TwentyTwo))
	jay.WriteTimeNano(b[176:], time.Time(t.TwentyThree))
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 184 {
		return jay.ErrUnexpectedEOB
	}
	t.One = tim(jay.ReadTimeNano(b[:8]))
	t.Two = tim(jay.ReadTimeNano(b[8:16]))
	t.Three = tim(jay.ReadTimeNano(b[16:24]))
	t.Four = tim(jay.ReadTimeNano(b[24:32]))
	t.Five = tim(jay.ReadTimeNano(b[32:40]))
	t.Six = tim(jay.ReadTimeNano(b[40:48]))
	t.Seven = tim(jay.ReadTimeNano(b[48:56]))
	t.Eight = tim(jay.ReadTimeNano(b[56:64]))
	t.Nine = tim(jay.ReadTimeNano(b[64:72]))
	t.Ten = tim(jay.ReadTimeNano(b[72:80]))
	t.Eleven = tim(jay.ReadTimeNano(b[80:88]))
	t.Twelve = tim(jay.ReadTimeNano(b[88:96]))
	t.Thirteen = tim(jay.ReadTimeNano(b[96:104]))
	t.Fourteen = tim(jay.ReadTimeNano(b[104:112]))
	t.Fifteen = tim(jay.ReadTimeNano(b[112:120]))
	t.Sixteen = tim(jay.ReadTimeNano(b[120:128]))
	t.Seventeen = tim(jay.ReadTimeNano(b[128:136]))
	t.Eighteen = tim(jay.ReadTimeNano(b[136:144]))
	t.Nineteen = tim(jay.ReadTimeNano(b[144:152]))
	t.Twenty = tim(jay.ReadTimeNano(b[152:160]))
	t.TwentyOne = tim(jay.ReadTimeNano(b[160:168]))
	t.TwentyTwo = tim(jay.ReadTimeNano(b[168:176]))
	t.TwentyThree = tim(jay.ReadTimeNano(b[176:]))
	return nil
}
