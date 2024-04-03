// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 4)
	jay.WriteInt32(b, rune(o.One))
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 4 {
		return jay.ErrUnexpectedEOB
	}
	o.One = char(jay.ReadInt32(b[:4]))
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:], rune(t.Two))
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 12)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:], rune(t.Three))
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 12 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteInt32(b[:4], rune(f.One))
	jay.WriteInt32(b[4:8], rune(f.Two))
	jay.WriteInt32(b[8:12], rune(f.Three))
	jay.WriteInt32(b[12:], rune(f.Four))
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	f.One = char(jay.ReadInt32(b[:4]))
	f.Two = char(jay.ReadInt32(b[4:8]))
	f.Three = char(jay.ReadInt32(b[8:12]))
	f.Four = char(jay.ReadInt32(b[12:16]))
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 20)
	jay.WriteInt32(b[:4], rune(f.One))
	jay.WriteInt32(b[4:8], rune(f.Two))
	jay.WriteInt32(b[8:12], rune(f.Three))
	jay.WriteInt32(b[12:16], rune(f.Four))
	jay.WriteInt32(b[16:], rune(f.Five))
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 20 {
		return jay.ErrUnexpectedEOB
	}
	f.One = char(jay.ReadInt32(b[:4]))
	f.Two = char(jay.ReadInt32(b[4:8]))
	f.Three = char(jay.ReadInt32(b[8:12]))
	f.Four = char(jay.ReadInt32(b[12:16]))
	f.Five = char(jay.ReadInt32(b[16:20]))
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteInt32(b[:4], rune(s.One))
	jay.WriteInt32(b[4:8], rune(s.Two))
	jay.WriteInt32(b[8:12], rune(s.Three))
	jay.WriteInt32(b[12:16], rune(s.Four))
	jay.WriteInt32(b[16:20], rune(s.Five))
	jay.WriteInt32(b[20:], rune(s.Six))
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	s.One = char(jay.ReadInt32(b[:4]))
	s.Two = char(jay.ReadInt32(b[4:8]))
	s.Three = char(jay.ReadInt32(b[8:12]))
	s.Four = char(jay.ReadInt32(b[12:16]))
	s.Five = char(jay.ReadInt32(b[16:20]))
	s.Six = char(jay.ReadInt32(b[20:24]))
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 28)
	jay.WriteInt32(b[:4], rune(s.One))
	jay.WriteInt32(b[4:8], rune(s.Two))
	jay.WriteInt32(b[8:12], rune(s.Three))
	jay.WriteInt32(b[12:16], rune(s.Four))
	jay.WriteInt32(b[16:20], rune(s.Five))
	jay.WriteInt32(b[20:24], rune(s.Six))
	jay.WriteInt32(b[24:], rune(s.Seven))
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 28 {
		return jay.ErrUnexpectedEOB
	}
	s.One = char(jay.ReadInt32(b[:4]))
	s.Two = char(jay.ReadInt32(b[4:8]))
	s.Three = char(jay.ReadInt32(b[8:12]))
	s.Four = char(jay.ReadInt32(b[12:16]))
	s.Five = char(jay.ReadInt32(b[16:20]))
	s.Six = char(jay.ReadInt32(b[20:24]))
	s.Seven = char(jay.ReadInt32(b[24:28]))
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteInt32(b[:4], rune(e.One))
	jay.WriteInt32(b[4:8], rune(e.Two))
	jay.WriteInt32(b[8:12], rune(e.Three))
	jay.WriteInt32(b[12:16], rune(e.Four))
	jay.WriteInt32(b[16:20], rune(e.Five))
	jay.WriteInt32(b[20:24], rune(e.Six))
	jay.WriteInt32(b[24:28], rune(e.Seven))
	jay.WriteInt32(b[28:], rune(e.Eight))
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	e.One = char(jay.ReadInt32(b[:4]))
	e.Two = char(jay.ReadInt32(b[4:8]))
	e.Three = char(jay.ReadInt32(b[8:12]))
	e.Four = char(jay.ReadInt32(b[12:16]))
	e.Five = char(jay.ReadInt32(b[16:20]))
	e.Six = char(jay.ReadInt32(b[20:24]))
	e.Seven = char(jay.ReadInt32(b[24:28]))
	e.Eight = char(jay.ReadInt32(b[28:32]))
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 36)
	jay.WriteInt32(b[:4], rune(n.One))
	jay.WriteInt32(b[4:8], rune(n.Two))
	jay.WriteInt32(b[8:12], rune(n.Three))
	jay.WriteInt32(b[12:16], rune(n.Four))
	jay.WriteInt32(b[16:20], rune(n.Five))
	jay.WriteInt32(b[20:24], rune(n.Six))
	jay.WriteInt32(b[24:28], rune(n.Seven))
	jay.WriteInt32(b[28:32], rune(n.Eight))
	jay.WriteInt32(b[32:], rune(n.Nine))
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 36 {
		return jay.ErrUnexpectedEOB
	}
	n.One = char(jay.ReadInt32(b[:4]))
	n.Two = char(jay.ReadInt32(b[4:8]))
	n.Three = char(jay.ReadInt32(b[8:12]))
	n.Four = char(jay.ReadInt32(b[12:16]))
	n.Five = char(jay.ReadInt32(b[16:20]))
	n.Six = char(jay.ReadInt32(b[20:24]))
	n.Seven = char(jay.ReadInt32(b[24:28]))
	n.Eight = char(jay.ReadInt32(b[28:32]))
	n.Nine = char(jay.ReadInt32(b[32:36]))
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:], rune(t.Ten))
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 44)
	jay.WriteInt32(b[:4], rune(e.One))
	jay.WriteInt32(b[4:8], rune(e.Two))
	jay.WriteInt32(b[8:12], rune(e.Three))
	jay.WriteInt32(b[12:16], rune(e.Four))
	jay.WriteInt32(b[16:20], rune(e.Five))
	jay.WriteInt32(b[20:24], rune(e.Six))
	jay.WriteInt32(b[24:28], rune(e.Seven))
	jay.WriteInt32(b[28:32], rune(e.Eight))
	jay.WriteInt32(b[32:36], rune(e.Nine))
	jay.WriteInt32(b[36:40], rune(e.Ten))
	jay.WriteInt32(b[40:], rune(e.Eleven))
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 44 {
		return jay.ErrUnexpectedEOB
	}
	e.One = char(jay.ReadInt32(b[:4]))
	e.Two = char(jay.ReadInt32(b[4:8]))
	e.Three = char(jay.ReadInt32(b[8:12]))
	e.Four = char(jay.ReadInt32(b[12:16]))
	e.Five = char(jay.ReadInt32(b[16:20]))
	e.Six = char(jay.ReadInt32(b[20:24]))
	e.Seven = char(jay.ReadInt32(b[24:28]))
	e.Eight = char(jay.ReadInt32(b[28:32]))
	e.Nine = char(jay.ReadInt32(b[32:36]))
	e.Ten = char(jay.ReadInt32(b[36:40]))
	e.Eleven = char(jay.ReadInt32(b[40:44]))
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 48)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:], rune(t.Twelve))
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 48 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 52)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:48], rune(t.Twelve))
	jay.WriteInt32(b[48:], rune(t.Thirteen))
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 52 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	t.Thirteen = char(jay.ReadInt32(b[48:52]))
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 56)
	jay.WriteInt32(b[:4], rune(f.One))
	jay.WriteInt32(b[4:8], rune(f.Two))
	jay.WriteInt32(b[8:12], rune(f.Three))
	jay.WriteInt32(b[12:16], rune(f.Four))
	jay.WriteInt32(b[16:20], rune(f.Five))
	jay.WriteInt32(b[20:24], rune(f.Six))
	jay.WriteInt32(b[24:28], rune(f.Seven))
	jay.WriteInt32(b[28:32], rune(f.Eight))
	jay.WriteInt32(b[32:36], rune(f.Nine))
	jay.WriteInt32(b[36:40], rune(f.Ten))
	jay.WriteInt32(b[40:44], rune(f.Eleven))
	jay.WriteInt32(b[44:48], rune(f.Twelve))
	jay.WriteInt32(b[48:52], rune(f.Thirteen))
	jay.WriteInt32(b[52:], rune(f.Fourteen))
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 56 {
		return jay.ErrUnexpectedEOB
	}
	f.One = char(jay.ReadInt32(b[:4]))
	f.Two = char(jay.ReadInt32(b[4:8]))
	f.Three = char(jay.ReadInt32(b[8:12]))
	f.Four = char(jay.ReadInt32(b[12:16]))
	f.Five = char(jay.ReadInt32(b[16:20]))
	f.Six = char(jay.ReadInt32(b[20:24]))
	f.Seven = char(jay.ReadInt32(b[24:28]))
	f.Eight = char(jay.ReadInt32(b[28:32]))
	f.Nine = char(jay.ReadInt32(b[32:36]))
	f.Ten = char(jay.ReadInt32(b[36:40]))
	f.Eleven = char(jay.ReadInt32(b[40:44]))
	f.Twelve = char(jay.ReadInt32(b[44:48]))
	f.Thirteen = char(jay.ReadInt32(b[48:52]))
	f.Fourteen = char(jay.ReadInt32(b[52:56]))
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 60)
	jay.WriteInt32(b[:4], rune(f.One))
	jay.WriteInt32(b[4:8], rune(f.Two))
	jay.WriteInt32(b[8:12], rune(f.Three))
	jay.WriteInt32(b[12:16], rune(f.Four))
	jay.WriteInt32(b[16:20], rune(f.Five))
	jay.WriteInt32(b[20:24], rune(f.Six))
	jay.WriteInt32(b[24:28], rune(f.Seven))
	jay.WriteInt32(b[28:32], rune(f.Eight))
	jay.WriteInt32(b[32:36], rune(f.Nine))
	jay.WriteInt32(b[36:40], rune(f.Ten))
	jay.WriteInt32(b[40:44], rune(f.Eleven))
	jay.WriteInt32(b[44:48], rune(f.Twelve))
	jay.WriteInt32(b[48:52], rune(f.Thirteen))
	jay.WriteInt32(b[52:56], rune(f.Fourteen))
	jay.WriteInt32(b[56:], rune(f.Fifteen))
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 60 {
		return jay.ErrUnexpectedEOB
	}
	f.One = char(jay.ReadInt32(b[:4]))
	f.Two = char(jay.ReadInt32(b[4:8]))
	f.Three = char(jay.ReadInt32(b[8:12]))
	f.Four = char(jay.ReadInt32(b[12:16]))
	f.Five = char(jay.ReadInt32(b[16:20]))
	f.Six = char(jay.ReadInt32(b[20:24]))
	f.Seven = char(jay.ReadInt32(b[24:28]))
	f.Eight = char(jay.ReadInt32(b[28:32]))
	f.Nine = char(jay.ReadInt32(b[32:36]))
	f.Ten = char(jay.ReadInt32(b[36:40]))
	f.Eleven = char(jay.ReadInt32(b[40:44]))
	f.Twelve = char(jay.ReadInt32(b[44:48]))
	f.Thirteen = char(jay.ReadInt32(b[48:52]))
	f.Fourteen = char(jay.ReadInt32(b[52:56]))
	f.Fifteen = char(jay.ReadInt32(b[56:60]))
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 64)
	jay.WriteInt32(b[:4], rune(s.One))
	jay.WriteInt32(b[4:8], rune(s.Two))
	jay.WriteInt32(b[8:12], rune(s.Three))
	jay.WriteInt32(b[12:16], rune(s.Four))
	jay.WriteInt32(b[16:20], rune(s.Five))
	jay.WriteInt32(b[20:24], rune(s.Six))
	jay.WriteInt32(b[24:28], rune(s.Seven))
	jay.WriteInt32(b[28:32], rune(s.Eight))
	jay.WriteInt32(b[32:36], rune(s.Nine))
	jay.WriteInt32(b[36:40], rune(s.Ten))
	jay.WriteInt32(b[40:44], rune(s.Eleven))
	jay.WriteInt32(b[44:48], rune(s.Twelve))
	jay.WriteInt32(b[48:52], rune(s.Thirteen))
	jay.WriteInt32(b[52:56], rune(s.Fourteen))
	jay.WriteInt32(b[56:60], rune(s.Fifteen))
	jay.WriteInt32(b[60:], rune(s.Sixteen))
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 64 {
		return jay.ErrUnexpectedEOB
	}
	s.One = char(jay.ReadInt32(b[:4]))
	s.Two = char(jay.ReadInt32(b[4:8]))
	s.Three = char(jay.ReadInt32(b[8:12]))
	s.Four = char(jay.ReadInt32(b[12:16]))
	s.Five = char(jay.ReadInt32(b[16:20]))
	s.Six = char(jay.ReadInt32(b[20:24]))
	s.Seven = char(jay.ReadInt32(b[24:28]))
	s.Eight = char(jay.ReadInt32(b[28:32]))
	s.Nine = char(jay.ReadInt32(b[32:36]))
	s.Ten = char(jay.ReadInt32(b[36:40]))
	s.Eleven = char(jay.ReadInt32(b[40:44]))
	s.Twelve = char(jay.ReadInt32(b[44:48]))
	s.Thirteen = char(jay.ReadInt32(b[48:52]))
	s.Fourteen = char(jay.ReadInt32(b[52:56]))
	s.Fifteen = char(jay.ReadInt32(b[56:60]))
	s.Sixteen = char(jay.ReadInt32(b[60:64]))
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 68)
	jay.WriteInt32(b[:4], rune(s.One))
	jay.WriteInt32(b[4:8], rune(s.Two))
	jay.WriteInt32(b[8:12], rune(s.Three))
	jay.WriteInt32(b[12:16], rune(s.Four))
	jay.WriteInt32(b[16:20], rune(s.Five))
	jay.WriteInt32(b[20:24], rune(s.Six))
	jay.WriteInt32(b[24:28], rune(s.Seven))
	jay.WriteInt32(b[28:32], rune(s.Eight))
	jay.WriteInt32(b[32:36], rune(s.Nine))
	jay.WriteInt32(b[36:40], rune(s.Ten))
	jay.WriteInt32(b[40:44], rune(s.Eleven))
	jay.WriteInt32(b[44:48], rune(s.Twelve))
	jay.WriteInt32(b[48:52], rune(s.Thirteen))
	jay.WriteInt32(b[52:56], rune(s.Fourteen))
	jay.WriteInt32(b[56:60], rune(s.Fifteen))
	jay.WriteInt32(b[60:64], rune(s.Sixteen))
	jay.WriteInt32(b[64:], rune(s.Seventeen))
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 68 {
		return jay.ErrUnexpectedEOB
	}
	s.One = char(jay.ReadInt32(b[:4]))
	s.Two = char(jay.ReadInt32(b[4:8]))
	s.Three = char(jay.ReadInt32(b[8:12]))
	s.Four = char(jay.ReadInt32(b[12:16]))
	s.Five = char(jay.ReadInt32(b[16:20]))
	s.Six = char(jay.ReadInt32(b[20:24]))
	s.Seven = char(jay.ReadInt32(b[24:28]))
	s.Eight = char(jay.ReadInt32(b[28:32]))
	s.Nine = char(jay.ReadInt32(b[32:36]))
	s.Ten = char(jay.ReadInt32(b[36:40]))
	s.Eleven = char(jay.ReadInt32(b[40:44]))
	s.Twelve = char(jay.ReadInt32(b[44:48]))
	s.Thirteen = char(jay.ReadInt32(b[48:52]))
	s.Fourteen = char(jay.ReadInt32(b[52:56]))
	s.Fifteen = char(jay.ReadInt32(b[56:60]))
	s.Sixteen = char(jay.ReadInt32(b[60:64]))
	s.Seventeen = char(jay.ReadInt32(b[64:68]))
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 72)
	jay.WriteInt32(b[:4], rune(e.One))
	jay.WriteInt32(b[4:8], rune(e.Two))
	jay.WriteInt32(b[8:12], rune(e.Three))
	jay.WriteInt32(b[12:16], rune(e.Four))
	jay.WriteInt32(b[16:20], rune(e.Five))
	jay.WriteInt32(b[20:24], rune(e.Six))
	jay.WriteInt32(b[24:28], rune(e.Seven))
	jay.WriteInt32(b[28:32], rune(e.Eight))
	jay.WriteInt32(b[32:36], rune(e.Nine))
	jay.WriteInt32(b[36:40], rune(e.Ten))
	jay.WriteInt32(b[40:44], rune(e.Eleven))
	jay.WriteInt32(b[44:48], rune(e.Twelve))
	jay.WriteInt32(b[48:52], rune(e.Thirteen))
	jay.WriteInt32(b[52:56], rune(e.Fourteen))
	jay.WriteInt32(b[56:60], rune(e.Fifteen))
	jay.WriteInt32(b[60:64], rune(e.Sixteen))
	jay.WriteInt32(b[64:68], rune(e.Seventeen))
	jay.WriteInt32(b[68:], rune(e.Eighteen))
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 72 {
		return jay.ErrUnexpectedEOB
	}
	e.One = char(jay.ReadInt32(b[:4]))
	e.Two = char(jay.ReadInt32(b[4:8]))
	e.Three = char(jay.ReadInt32(b[8:12]))
	e.Four = char(jay.ReadInt32(b[12:16]))
	e.Five = char(jay.ReadInt32(b[16:20]))
	e.Six = char(jay.ReadInt32(b[20:24]))
	e.Seven = char(jay.ReadInt32(b[24:28]))
	e.Eight = char(jay.ReadInt32(b[28:32]))
	e.Nine = char(jay.ReadInt32(b[32:36]))
	e.Ten = char(jay.ReadInt32(b[36:40]))
	e.Eleven = char(jay.ReadInt32(b[40:44]))
	e.Twelve = char(jay.ReadInt32(b[44:48]))
	e.Thirteen = char(jay.ReadInt32(b[48:52]))
	e.Fourteen = char(jay.ReadInt32(b[52:56]))
	e.Fifteen = char(jay.ReadInt32(b[56:60]))
	e.Sixteen = char(jay.ReadInt32(b[60:64]))
	e.Seventeen = char(jay.ReadInt32(b[64:68]))
	e.Eighteen = char(jay.ReadInt32(b[68:72]))
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 76)
	jay.WriteInt32(b[:4], rune(n.One))
	jay.WriteInt32(b[4:8], rune(n.Two))
	jay.WriteInt32(b[8:12], rune(n.Three))
	jay.WriteInt32(b[12:16], rune(n.Four))
	jay.WriteInt32(b[16:20], rune(n.Five))
	jay.WriteInt32(b[20:24], rune(n.Six))
	jay.WriteInt32(b[24:28], rune(n.Seven))
	jay.WriteInt32(b[28:32], rune(n.Eight))
	jay.WriteInt32(b[32:36], rune(n.Nine))
	jay.WriteInt32(b[36:40], rune(n.Ten))
	jay.WriteInt32(b[40:44], rune(n.Eleven))
	jay.WriteInt32(b[44:48], rune(n.Twelve))
	jay.WriteInt32(b[48:52], rune(n.Thirteen))
	jay.WriteInt32(b[52:56], rune(n.Fourteen))
	jay.WriteInt32(b[56:60], rune(n.Fifteen))
	jay.WriteInt32(b[60:64], rune(n.Sixteen))
	jay.WriteInt32(b[64:68], rune(n.Seventeen))
	jay.WriteInt32(b[68:72], rune(n.Eighteen))
	jay.WriteInt32(b[72:], rune(n.Nineteen))
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 76 {
		return jay.ErrUnexpectedEOB
	}
	n.One = char(jay.ReadInt32(b[:4]))
	n.Two = char(jay.ReadInt32(b[4:8]))
	n.Three = char(jay.ReadInt32(b[8:12]))
	n.Four = char(jay.ReadInt32(b[12:16]))
	n.Five = char(jay.ReadInt32(b[16:20]))
	n.Six = char(jay.ReadInt32(b[20:24]))
	n.Seven = char(jay.ReadInt32(b[24:28]))
	n.Eight = char(jay.ReadInt32(b[28:32]))
	n.Nine = char(jay.ReadInt32(b[32:36]))
	n.Ten = char(jay.ReadInt32(b[36:40]))
	n.Eleven = char(jay.ReadInt32(b[40:44]))
	n.Twelve = char(jay.ReadInt32(b[44:48]))
	n.Thirteen = char(jay.ReadInt32(b[48:52]))
	n.Fourteen = char(jay.ReadInt32(b[52:56]))
	n.Fifteen = char(jay.ReadInt32(b[56:60]))
	n.Sixteen = char(jay.ReadInt32(b[60:64]))
	n.Seventeen = char(jay.ReadInt32(b[64:68]))
	n.Eighteen = char(jay.ReadInt32(b[68:72]))
	n.Nineteen = char(jay.ReadInt32(b[72:76]))
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 80)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:48], rune(t.Twelve))
	jay.WriteInt32(b[48:52], rune(t.Thirteen))
	jay.WriteInt32(b[52:56], rune(t.Fourteen))
	jay.WriteInt32(b[56:60], rune(t.Fifteen))
	jay.WriteInt32(b[60:64], rune(t.Sixteen))
	jay.WriteInt32(b[64:68], rune(t.Seventeen))
	jay.WriteInt32(b[68:72], rune(t.Eighteen))
	jay.WriteInt32(b[72:76], rune(t.Nineteen))
	jay.WriteInt32(b[76:], rune(t.Twenty))
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 80 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	t.Thirteen = char(jay.ReadInt32(b[48:52]))
	t.Fourteen = char(jay.ReadInt32(b[52:56]))
	t.Fifteen = char(jay.ReadInt32(b[56:60]))
	t.Sixteen = char(jay.ReadInt32(b[60:64]))
	t.Seventeen = char(jay.ReadInt32(b[64:68]))
	t.Eighteen = char(jay.ReadInt32(b[68:72]))
	t.Nineteen = char(jay.ReadInt32(b[72:76]))
	t.Twenty = char(jay.ReadInt32(b[76:80]))
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 84)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:48], rune(t.Twelve))
	jay.WriteInt32(b[48:52], rune(t.Thirteen))
	jay.WriteInt32(b[52:56], rune(t.Fourteen))
	jay.WriteInt32(b[56:60], rune(t.Fifteen))
	jay.WriteInt32(b[60:64], rune(t.Sixteen))
	jay.WriteInt32(b[64:68], rune(t.Seventeen))
	jay.WriteInt32(b[68:72], rune(t.Eighteen))
	jay.WriteInt32(b[72:76], rune(t.Nineteen))
	jay.WriteInt32(b[76:80], rune(t.Twenty))
	jay.WriteInt32(b[80:], rune(t.TwentyOne))
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 84 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	t.Thirteen = char(jay.ReadInt32(b[48:52]))
	t.Fourteen = char(jay.ReadInt32(b[52:56]))
	t.Fifteen = char(jay.ReadInt32(b[56:60]))
	t.Sixteen = char(jay.ReadInt32(b[60:64]))
	t.Seventeen = char(jay.ReadInt32(b[64:68]))
	t.Eighteen = char(jay.ReadInt32(b[68:72]))
	t.Nineteen = char(jay.ReadInt32(b[72:76]))
	t.Twenty = char(jay.ReadInt32(b[76:80]))
	t.TwentyOne = char(jay.ReadInt32(b[80:84]))
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 88)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:48], rune(t.Twelve))
	jay.WriteInt32(b[48:52], rune(t.Thirteen))
	jay.WriteInt32(b[52:56], rune(t.Fourteen))
	jay.WriteInt32(b[56:60], rune(t.Fifteen))
	jay.WriteInt32(b[60:64], rune(t.Sixteen))
	jay.WriteInt32(b[64:68], rune(t.Seventeen))
	jay.WriteInt32(b[68:72], rune(t.Eighteen))
	jay.WriteInt32(b[72:76], rune(t.Nineteen))
	jay.WriteInt32(b[76:80], rune(t.Twenty))
	jay.WriteInt32(b[80:84], rune(t.TwentyOne))
	jay.WriteInt32(b[84:], rune(t.TwentyTwo))
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 88 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	t.Thirteen = char(jay.ReadInt32(b[48:52]))
	t.Fourteen = char(jay.ReadInt32(b[52:56]))
	t.Fifteen = char(jay.ReadInt32(b[56:60]))
	t.Sixteen = char(jay.ReadInt32(b[60:64]))
	t.Seventeen = char(jay.ReadInt32(b[64:68]))
	t.Eighteen = char(jay.ReadInt32(b[68:72]))
	t.Nineteen = char(jay.ReadInt32(b[72:76]))
	t.Twenty = char(jay.ReadInt32(b[76:80]))
	t.TwentyOne = char(jay.ReadInt32(b[80:84]))
	t.TwentyTwo = char(jay.ReadInt32(b[84:88]))
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 92)
	jay.WriteInt32(b[:4], rune(t.One))
	jay.WriteInt32(b[4:8], rune(t.Two))
	jay.WriteInt32(b[8:12], rune(t.Three))
	jay.WriteInt32(b[12:16], rune(t.Four))
	jay.WriteInt32(b[16:20], rune(t.Five))
	jay.WriteInt32(b[20:24], rune(t.Six))
	jay.WriteInt32(b[24:28], rune(t.Seven))
	jay.WriteInt32(b[28:32], rune(t.Eight))
	jay.WriteInt32(b[32:36], rune(t.Nine))
	jay.WriteInt32(b[36:40], rune(t.Ten))
	jay.WriteInt32(b[40:44], rune(t.Eleven))
	jay.WriteInt32(b[44:48], rune(t.Twelve))
	jay.WriteInt32(b[48:52], rune(t.Thirteen))
	jay.WriteInt32(b[52:56], rune(t.Fourteen))
	jay.WriteInt32(b[56:60], rune(t.Fifteen))
	jay.WriteInt32(b[60:64], rune(t.Sixteen))
	jay.WriteInt32(b[64:68], rune(t.Seventeen))
	jay.WriteInt32(b[68:72], rune(t.Eighteen))
	jay.WriteInt32(b[72:76], rune(t.Nineteen))
	jay.WriteInt32(b[76:80], rune(t.Twenty))
	jay.WriteInt32(b[80:84], rune(t.TwentyOne))
	jay.WriteInt32(b[84:88], rune(t.TwentyTwo))
	jay.WriteInt32(b[88:], rune(t.TwentyThree))
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 92 {
		return jay.ErrUnexpectedEOB
	}
	t.One = char(jay.ReadInt32(b[:4]))
	t.Two = char(jay.ReadInt32(b[4:8]))
	t.Three = char(jay.ReadInt32(b[8:12]))
	t.Four = char(jay.ReadInt32(b[12:16]))
	t.Five = char(jay.ReadInt32(b[16:20]))
	t.Six = char(jay.ReadInt32(b[20:24]))
	t.Seven = char(jay.ReadInt32(b[24:28]))
	t.Eight = char(jay.ReadInt32(b[28:32]))
	t.Nine = char(jay.ReadInt32(b[32:36]))
	t.Ten = char(jay.ReadInt32(b[36:40]))
	t.Eleven = char(jay.ReadInt32(b[40:44]))
	t.Twelve = char(jay.ReadInt32(b[44:48]))
	t.Thirteen = char(jay.ReadInt32(b[48:52]))
	t.Fourteen = char(jay.ReadInt32(b[52:56]))
	t.Fifteen = char(jay.ReadInt32(b[56:60]))
	t.Sixteen = char(jay.ReadInt32(b[60:64]))
	t.Seventeen = char(jay.ReadInt32(b[64:68]))
	t.Eighteen = char(jay.ReadInt32(b[68:72]))
	t.Nineteen = char(jay.ReadInt32(b[72:76]))
	t.Twenty = char(jay.ReadInt32(b[76:80]))
	t.TwentyOne = char(jay.ReadInt32(b[80:84]))
	t.TwentyTwo = char(jay.ReadInt32(b[84:88]))
	t.TwentyThree = char(jay.ReadInt32(b[88:92]))
	return nil
}
