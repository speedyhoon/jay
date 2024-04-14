// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() (b []byte) {
	b = make([]byte, 2)
	jay.WriteUint16(b, uint16(o.One))
	return
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 2 {
		return jay.ErrUnexpectedEOB
	}
	o.One = small(jay.ReadUint16(b))
	return nil
}

func (t *Two) MarshalJ() (b []byte) {
	b = make([]byte, 4)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:], uint16(t.Two))
	return
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 4 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:]))
	return nil
}

func (t *Three) MarshalJ() (b []byte) {
	b = make([]byte, 6)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:], uint16(t.Three))
	return
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 6 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:]))
	return nil
}

func (f *Four) MarshalJ() (b []byte) {
	b = make([]byte, 8)
	jay.WriteUint16(b[:2], uint16(f.One))
	jay.WriteUint16(b[2:4], uint16(f.Two))
	jay.WriteUint16(b[4:6], uint16(f.Three))
	jay.WriteUint16(b[6:], uint16(f.Four))
	return
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	f.One = small(jay.ReadUint16(b[:2]))
	f.Two = small(jay.ReadUint16(b[2:4]))
	f.Three = small(jay.ReadUint16(b[4:6]))
	f.Four = small(jay.ReadUint16(b[6:]))
	return nil
}

func (f *Five) MarshalJ() (b []byte) {
	b = make([]byte, 10)
	jay.WriteUint16(b[:2], uint16(f.One))
	jay.WriteUint16(b[2:4], uint16(f.Two))
	jay.WriteUint16(b[4:6], uint16(f.Three))
	jay.WriteUint16(b[6:8], uint16(f.Four))
	jay.WriteUint16(b[8:], uint16(f.Five))
	return
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 10 {
		return jay.ErrUnexpectedEOB
	}
	f.One = small(jay.ReadUint16(b[:2]))
	f.Two = small(jay.ReadUint16(b[2:4]))
	f.Three = small(jay.ReadUint16(b[4:6]))
	f.Four = small(jay.ReadUint16(b[6:8]))
	f.Five = small(jay.ReadUint16(b[8:]))
	return nil
}

func (s *Six) MarshalJ() (b []byte) {
	b = make([]byte, 12)
	jay.WriteUint16(b[:2], uint16(s.One))
	jay.WriteUint16(b[2:4], uint16(s.Two))
	jay.WriteUint16(b[4:6], uint16(s.Three))
	jay.WriteUint16(b[6:8], uint16(s.Four))
	jay.WriteUint16(b[8:10], uint16(s.Five))
	jay.WriteUint16(b[10:], uint16(s.Six))
	return
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 12 {
		return jay.ErrUnexpectedEOB
	}
	s.One = small(jay.ReadUint16(b[:2]))
	s.Two = small(jay.ReadUint16(b[2:4]))
	s.Three = small(jay.ReadUint16(b[4:6]))
	s.Four = small(jay.ReadUint16(b[6:8]))
	s.Five = small(jay.ReadUint16(b[8:10]))
	s.Six = small(jay.ReadUint16(b[10:]))
	return nil
}

func (s *Seven) MarshalJ() (b []byte) {
	b = make([]byte, 14)
	jay.WriteUint16(b[:2], uint16(s.One))
	jay.WriteUint16(b[2:4], uint16(s.Two))
	jay.WriteUint16(b[4:6], uint16(s.Three))
	jay.WriteUint16(b[6:8], uint16(s.Four))
	jay.WriteUint16(b[8:10], uint16(s.Five))
	jay.WriteUint16(b[10:12], uint16(s.Six))
	jay.WriteUint16(b[12:], uint16(s.Seven))
	return
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 14 {
		return jay.ErrUnexpectedEOB
	}
	s.One = small(jay.ReadUint16(b[:2]))
	s.Two = small(jay.ReadUint16(b[2:4]))
	s.Three = small(jay.ReadUint16(b[4:6]))
	s.Four = small(jay.ReadUint16(b[6:8]))
	s.Five = small(jay.ReadUint16(b[8:10]))
	s.Six = small(jay.ReadUint16(b[10:12]))
	s.Seven = small(jay.ReadUint16(b[12:]))
	return nil
}

func (e *Eight) MarshalJ() (b []byte) {
	b = make([]byte, 16)
	jay.WriteUint16(b[:2], uint16(e.One))
	jay.WriteUint16(b[2:4], uint16(e.Two))
	jay.WriteUint16(b[4:6], uint16(e.Three))
	jay.WriteUint16(b[6:8], uint16(e.Four))
	jay.WriteUint16(b[8:10], uint16(e.Five))
	jay.WriteUint16(b[10:12], uint16(e.Six))
	jay.WriteUint16(b[12:14], uint16(e.Seven))
	jay.WriteUint16(b[14:], uint16(e.Eight))
	return
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	e.One = small(jay.ReadUint16(b[:2]))
	e.Two = small(jay.ReadUint16(b[2:4]))
	e.Three = small(jay.ReadUint16(b[4:6]))
	e.Four = small(jay.ReadUint16(b[6:8]))
	e.Five = small(jay.ReadUint16(b[8:10]))
	e.Six = small(jay.ReadUint16(b[10:12]))
	e.Seven = small(jay.ReadUint16(b[12:14]))
	e.Eight = small(jay.ReadUint16(b[14:]))
	return nil
}

func (n *Nine) MarshalJ() (b []byte) {
	b = make([]byte, 18)
	jay.WriteUint16(b[:2], uint16(n.One))
	jay.WriteUint16(b[2:4], uint16(n.Two))
	jay.WriteUint16(b[4:6], uint16(n.Three))
	jay.WriteUint16(b[6:8], uint16(n.Four))
	jay.WriteUint16(b[8:10], uint16(n.Five))
	jay.WriteUint16(b[10:12], uint16(n.Six))
	jay.WriteUint16(b[12:14], uint16(n.Seven))
	jay.WriteUint16(b[14:16], uint16(n.Eight))
	jay.WriteUint16(b[16:], uint16(n.Nine))
	return
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 18 {
		return jay.ErrUnexpectedEOB
	}
	n.One = small(jay.ReadUint16(b[:2]))
	n.Two = small(jay.ReadUint16(b[2:4]))
	n.Three = small(jay.ReadUint16(b[4:6]))
	n.Four = small(jay.ReadUint16(b[6:8]))
	n.Five = small(jay.ReadUint16(b[8:10]))
	n.Six = small(jay.ReadUint16(b[10:12]))
	n.Seven = small(jay.ReadUint16(b[12:14]))
	n.Eight = small(jay.ReadUint16(b[14:16]))
	n.Nine = small(jay.ReadUint16(b[16:]))
	return nil
}

func (t *Ten) MarshalJ() (b []byte) {
	b = make([]byte, 20)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:], uint16(t.Ten))
	return
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 20 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:]))
	return nil
}

func (e *Eleven) MarshalJ() (b []byte) {
	b = make([]byte, 22)
	jay.WriteUint16(b[:2], uint16(e.One))
	jay.WriteUint16(b[2:4], uint16(e.Two))
	jay.WriteUint16(b[4:6], uint16(e.Three))
	jay.WriteUint16(b[6:8], uint16(e.Four))
	jay.WriteUint16(b[8:10], uint16(e.Five))
	jay.WriteUint16(b[10:12], uint16(e.Six))
	jay.WriteUint16(b[12:14], uint16(e.Seven))
	jay.WriteUint16(b[14:16], uint16(e.Eight))
	jay.WriteUint16(b[16:18], uint16(e.Nine))
	jay.WriteUint16(b[18:20], uint16(e.Ten))
	jay.WriteUint16(b[20:], uint16(e.Eleven))
	return
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 22 {
		return jay.ErrUnexpectedEOB
	}
	e.One = small(jay.ReadUint16(b[:2]))
	e.Two = small(jay.ReadUint16(b[2:4]))
	e.Three = small(jay.ReadUint16(b[4:6]))
	e.Four = small(jay.ReadUint16(b[6:8]))
	e.Five = small(jay.ReadUint16(b[8:10]))
	e.Six = small(jay.ReadUint16(b[10:12]))
	e.Seven = small(jay.ReadUint16(b[12:14]))
	e.Eight = small(jay.ReadUint16(b[14:16]))
	e.Nine = small(jay.ReadUint16(b[16:18]))
	e.Ten = small(jay.ReadUint16(b[18:20]))
	e.Eleven = small(jay.ReadUint16(b[20:]))
	return nil
}

func (t *Twelve) MarshalJ() (b []byte) {
	b = make([]byte, 24)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:], uint16(t.Twelve))
	return
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 24 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:]))
	return nil
}

func (t *Thirteen) MarshalJ() (b []byte) {
	b = make([]byte, 26)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:24], uint16(t.Twelve))
	jay.WriteUint16(b[24:], uint16(t.Thirteen))
	return
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 26 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:24]))
	t.Thirteen = small(jay.ReadUint16(b[24:]))
	return nil
}

func (f *Fourteen) MarshalJ() (b []byte) {
	b = make([]byte, 28)
	jay.WriteUint16(b[:2], uint16(f.One))
	jay.WriteUint16(b[2:4], uint16(f.Two))
	jay.WriteUint16(b[4:6], uint16(f.Three))
	jay.WriteUint16(b[6:8], uint16(f.Four))
	jay.WriteUint16(b[8:10], uint16(f.Five))
	jay.WriteUint16(b[10:12], uint16(f.Six))
	jay.WriteUint16(b[12:14], uint16(f.Seven))
	jay.WriteUint16(b[14:16], uint16(f.Eight))
	jay.WriteUint16(b[16:18], uint16(f.Nine))
	jay.WriteUint16(b[18:20], uint16(f.Ten))
	jay.WriteUint16(b[20:22], uint16(f.Eleven))
	jay.WriteUint16(b[22:24], uint16(f.Twelve))
	jay.WriteUint16(b[24:26], uint16(f.Thirteen))
	jay.WriteUint16(b[26:], uint16(f.Fourteen))
	return
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 28 {
		return jay.ErrUnexpectedEOB
	}
	f.One = small(jay.ReadUint16(b[:2]))
	f.Two = small(jay.ReadUint16(b[2:4]))
	f.Three = small(jay.ReadUint16(b[4:6]))
	f.Four = small(jay.ReadUint16(b[6:8]))
	f.Five = small(jay.ReadUint16(b[8:10]))
	f.Six = small(jay.ReadUint16(b[10:12]))
	f.Seven = small(jay.ReadUint16(b[12:14]))
	f.Eight = small(jay.ReadUint16(b[14:16]))
	f.Nine = small(jay.ReadUint16(b[16:18]))
	f.Ten = small(jay.ReadUint16(b[18:20]))
	f.Eleven = small(jay.ReadUint16(b[20:22]))
	f.Twelve = small(jay.ReadUint16(b[22:24]))
	f.Thirteen = small(jay.ReadUint16(b[24:26]))
	f.Fourteen = small(jay.ReadUint16(b[26:]))
	return nil
}

func (f *Fifteen) MarshalJ() (b []byte) {
	b = make([]byte, 30)
	jay.WriteUint16(b[:2], uint16(f.One))
	jay.WriteUint16(b[2:4], uint16(f.Two))
	jay.WriteUint16(b[4:6], uint16(f.Three))
	jay.WriteUint16(b[6:8], uint16(f.Four))
	jay.WriteUint16(b[8:10], uint16(f.Five))
	jay.WriteUint16(b[10:12], uint16(f.Six))
	jay.WriteUint16(b[12:14], uint16(f.Seven))
	jay.WriteUint16(b[14:16], uint16(f.Eight))
	jay.WriteUint16(b[16:18], uint16(f.Nine))
	jay.WriteUint16(b[18:20], uint16(f.Ten))
	jay.WriteUint16(b[20:22], uint16(f.Eleven))
	jay.WriteUint16(b[22:24], uint16(f.Twelve))
	jay.WriteUint16(b[24:26], uint16(f.Thirteen))
	jay.WriteUint16(b[26:28], uint16(f.Fourteen))
	jay.WriteUint16(b[28:], uint16(f.Fifteen))
	return
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 30 {
		return jay.ErrUnexpectedEOB
	}
	f.One = small(jay.ReadUint16(b[:2]))
	f.Two = small(jay.ReadUint16(b[2:4]))
	f.Three = small(jay.ReadUint16(b[4:6]))
	f.Four = small(jay.ReadUint16(b[6:8]))
	f.Five = small(jay.ReadUint16(b[8:10]))
	f.Six = small(jay.ReadUint16(b[10:12]))
	f.Seven = small(jay.ReadUint16(b[12:14]))
	f.Eight = small(jay.ReadUint16(b[14:16]))
	f.Nine = small(jay.ReadUint16(b[16:18]))
	f.Ten = small(jay.ReadUint16(b[18:20]))
	f.Eleven = small(jay.ReadUint16(b[20:22]))
	f.Twelve = small(jay.ReadUint16(b[22:24]))
	f.Thirteen = small(jay.ReadUint16(b[24:26]))
	f.Fourteen = small(jay.ReadUint16(b[26:28]))
	f.Fifteen = small(jay.ReadUint16(b[28:]))
	return nil
}

func (s *Sixteen) MarshalJ() (b []byte) {
	b = make([]byte, 32)
	jay.WriteUint16(b[:2], uint16(s.One))
	jay.WriteUint16(b[2:4], uint16(s.Two))
	jay.WriteUint16(b[4:6], uint16(s.Three))
	jay.WriteUint16(b[6:8], uint16(s.Four))
	jay.WriteUint16(b[8:10], uint16(s.Five))
	jay.WriteUint16(b[10:12], uint16(s.Six))
	jay.WriteUint16(b[12:14], uint16(s.Seven))
	jay.WriteUint16(b[14:16], uint16(s.Eight))
	jay.WriteUint16(b[16:18], uint16(s.Nine))
	jay.WriteUint16(b[18:20], uint16(s.Ten))
	jay.WriteUint16(b[20:22], uint16(s.Eleven))
	jay.WriteUint16(b[22:24], uint16(s.Twelve))
	jay.WriteUint16(b[24:26], uint16(s.Thirteen))
	jay.WriteUint16(b[26:28], uint16(s.Fourteen))
	jay.WriteUint16(b[28:30], uint16(s.Fifteen))
	jay.WriteUint16(b[30:], uint16(s.Sixteen))
	return
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 32 {
		return jay.ErrUnexpectedEOB
	}
	s.One = small(jay.ReadUint16(b[:2]))
	s.Two = small(jay.ReadUint16(b[2:4]))
	s.Three = small(jay.ReadUint16(b[4:6]))
	s.Four = small(jay.ReadUint16(b[6:8]))
	s.Five = small(jay.ReadUint16(b[8:10]))
	s.Six = small(jay.ReadUint16(b[10:12]))
	s.Seven = small(jay.ReadUint16(b[12:14]))
	s.Eight = small(jay.ReadUint16(b[14:16]))
	s.Nine = small(jay.ReadUint16(b[16:18]))
	s.Ten = small(jay.ReadUint16(b[18:20]))
	s.Eleven = small(jay.ReadUint16(b[20:22]))
	s.Twelve = small(jay.ReadUint16(b[22:24]))
	s.Thirteen = small(jay.ReadUint16(b[24:26]))
	s.Fourteen = small(jay.ReadUint16(b[26:28]))
	s.Fifteen = small(jay.ReadUint16(b[28:30]))
	s.Sixteen = small(jay.ReadUint16(b[30:]))
	return nil
}

func (s *Seventeen) MarshalJ() (b []byte) {
	b = make([]byte, 34)
	jay.WriteUint16(b[:2], uint16(s.One))
	jay.WriteUint16(b[2:4], uint16(s.Two))
	jay.WriteUint16(b[4:6], uint16(s.Three))
	jay.WriteUint16(b[6:8], uint16(s.Four))
	jay.WriteUint16(b[8:10], uint16(s.Five))
	jay.WriteUint16(b[10:12], uint16(s.Six))
	jay.WriteUint16(b[12:14], uint16(s.Seven))
	jay.WriteUint16(b[14:16], uint16(s.Eight))
	jay.WriteUint16(b[16:18], uint16(s.Nine))
	jay.WriteUint16(b[18:20], uint16(s.Ten))
	jay.WriteUint16(b[20:22], uint16(s.Eleven))
	jay.WriteUint16(b[22:24], uint16(s.Twelve))
	jay.WriteUint16(b[24:26], uint16(s.Thirteen))
	jay.WriteUint16(b[26:28], uint16(s.Fourteen))
	jay.WriteUint16(b[28:30], uint16(s.Fifteen))
	jay.WriteUint16(b[30:32], uint16(s.Sixteen))
	jay.WriteUint16(b[32:], uint16(s.Seventeen))
	return
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 34 {
		return jay.ErrUnexpectedEOB
	}
	s.One = small(jay.ReadUint16(b[:2]))
	s.Two = small(jay.ReadUint16(b[2:4]))
	s.Three = small(jay.ReadUint16(b[4:6]))
	s.Four = small(jay.ReadUint16(b[6:8]))
	s.Five = small(jay.ReadUint16(b[8:10]))
	s.Six = small(jay.ReadUint16(b[10:12]))
	s.Seven = small(jay.ReadUint16(b[12:14]))
	s.Eight = small(jay.ReadUint16(b[14:16]))
	s.Nine = small(jay.ReadUint16(b[16:18]))
	s.Ten = small(jay.ReadUint16(b[18:20]))
	s.Eleven = small(jay.ReadUint16(b[20:22]))
	s.Twelve = small(jay.ReadUint16(b[22:24]))
	s.Thirteen = small(jay.ReadUint16(b[24:26]))
	s.Fourteen = small(jay.ReadUint16(b[26:28]))
	s.Fifteen = small(jay.ReadUint16(b[28:30]))
	s.Sixteen = small(jay.ReadUint16(b[30:32]))
	s.Seventeen = small(jay.ReadUint16(b[32:]))
	return nil
}

func (e *Eighteen) MarshalJ() (b []byte) {
	b = make([]byte, 36)
	jay.WriteUint16(b[:2], uint16(e.One))
	jay.WriteUint16(b[2:4], uint16(e.Two))
	jay.WriteUint16(b[4:6], uint16(e.Three))
	jay.WriteUint16(b[6:8], uint16(e.Four))
	jay.WriteUint16(b[8:10], uint16(e.Five))
	jay.WriteUint16(b[10:12], uint16(e.Six))
	jay.WriteUint16(b[12:14], uint16(e.Seven))
	jay.WriteUint16(b[14:16], uint16(e.Eight))
	jay.WriteUint16(b[16:18], uint16(e.Nine))
	jay.WriteUint16(b[18:20], uint16(e.Ten))
	jay.WriteUint16(b[20:22], uint16(e.Eleven))
	jay.WriteUint16(b[22:24], uint16(e.Twelve))
	jay.WriteUint16(b[24:26], uint16(e.Thirteen))
	jay.WriteUint16(b[26:28], uint16(e.Fourteen))
	jay.WriteUint16(b[28:30], uint16(e.Fifteen))
	jay.WriteUint16(b[30:32], uint16(e.Sixteen))
	jay.WriteUint16(b[32:34], uint16(e.Seventeen))
	jay.WriteUint16(b[34:], uint16(e.Eighteen))
	return
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 36 {
		return jay.ErrUnexpectedEOB
	}
	e.One = small(jay.ReadUint16(b[:2]))
	e.Two = small(jay.ReadUint16(b[2:4]))
	e.Three = small(jay.ReadUint16(b[4:6]))
	e.Four = small(jay.ReadUint16(b[6:8]))
	e.Five = small(jay.ReadUint16(b[8:10]))
	e.Six = small(jay.ReadUint16(b[10:12]))
	e.Seven = small(jay.ReadUint16(b[12:14]))
	e.Eight = small(jay.ReadUint16(b[14:16]))
	e.Nine = small(jay.ReadUint16(b[16:18]))
	e.Ten = small(jay.ReadUint16(b[18:20]))
	e.Eleven = small(jay.ReadUint16(b[20:22]))
	e.Twelve = small(jay.ReadUint16(b[22:24]))
	e.Thirteen = small(jay.ReadUint16(b[24:26]))
	e.Fourteen = small(jay.ReadUint16(b[26:28]))
	e.Fifteen = small(jay.ReadUint16(b[28:30]))
	e.Sixteen = small(jay.ReadUint16(b[30:32]))
	e.Seventeen = small(jay.ReadUint16(b[32:34]))
	e.Eighteen = small(jay.ReadUint16(b[34:]))
	return nil
}

func (n *Nineteen) MarshalJ() (b []byte) {
	b = make([]byte, 38)
	jay.WriteUint16(b[:2], uint16(n.One))
	jay.WriteUint16(b[2:4], uint16(n.Two))
	jay.WriteUint16(b[4:6], uint16(n.Three))
	jay.WriteUint16(b[6:8], uint16(n.Four))
	jay.WriteUint16(b[8:10], uint16(n.Five))
	jay.WriteUint16(b[10:12], uint16(n.Six))
	jay.WriteUint16(b[12:14], uint16(n.Seven))
	jay.WriteUint16(b[14:16], uint16(n.Eight))
	jay.WriteUint16(b[16:18], uint16(n.Nine))
	jay.WriteUint16(b[18:20], uint16(n.Ten))
	jay.WriteUint16(b[20:22], uint16(n.Eleven))
	jay.WriteUint16(b[22:24], uint16(n.Twelve))
	jay.WriteUint16(b[24:26], uint16(n.Thirteen))
	jay.WriteUint16(b[26:28], uint16(n.Fourteen))
	jay.WriteUint16(b[28:30], uint16(n.Fifteen))
	jay.WriteUint16(b[30:32], uint16(n.Sixteen))
	jay.WriteUint16(b[32:34], uint16(n.Seventeen))
	jay.WriteUint16(b[34:36], uint16(n.Eighteen))
	jay.WriteUint16(b[36:], uint16(n.Nineteen))
	return
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 38 {
		return jay.ErrUnexpectedEOB
	}
	n.One = small(jay.ReadUint16(b[:2]))
	n.Two = small(jay.ReadUint16(b[2:4]))
	n.Three = small(jay.ReadUint16(b[4:6]))
	n.Four = small(jay.ReadUint16(b[6:8]))
	n.Five = small(jay.ReadUint16(b[8:10]))
	n.Six = small(jay.ReadUint16(b[10:12]))
	n.Seven = small(jay.ReadUint16(b[12:14]))
	n.Eight = small(jay.ReadUint16(b[14:16]))
	n.Nine = small(jay.ReadUint16(b[16:18]))
	n.Ten = small(jay.ReadUint16(b[18:20]))
	n.Eleven = small(jay.ReadUint16(b[20:22]))
	n.Twelve = small(jay.ReadUint16(b[22:24]))
	n.Thirteen = small(jay.ReadUint16(b[24:26]))
	n.Fourteen = small(jay.ReadUint16(b[26:28]))
	n.Fifteen = small(jay.ReadUint16(b[28:30]))
	n.Sixteen = small(jay.ReadUint16(b[30:32]))
	n.Seventeen = small(jay.ReadUint16(b[32:34]))
	n.Eighteen = small(jay.ReadUint16(b[34:36]))
	n.Nineteen = small(jay.ReadUint16(b[36:]))
	return nil
}

func (t *Twenty) MarshalJ() (b []byte) {
	b = make([]byte, 40)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:24], uint16(t.Twelve))
	jay.WriteUint16(b[24:26], uint16(t.Thirteen))
	jay.WriteUint16(b[26:28], uint16(t.Fourteen))
	jay.WriteUint16(b[28:30], uint16(t.Fifteen))
	jay.WriteUint16(b[30:32], uint16(t.Sixteen))
	jay.WriteUint16(b[32:34], uint16(t.Seventeen))
	jay.WriteUint16(b[34:36], uint16(t.Eighteen))
	jay.WriteUint16(b[36:38], uint16(t.Nineteen))
	jay.WriteUint16(b[38:], uint16(t.Twenty))
	return
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 40 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:24]))
	t.Thirteen = small(jay.ReadUint16(b[24:26]))
	t.Fourteen = small(jay.ReadUint16(b[26:28]))
	t.Fifteen = small(jay.ReadUint16(b[28:30]))
	t.Sixteen = small(jay.ReadUint16(b[30:32]))
	t.Seventeen = small(jay.ReadUint16(b[32:34]))
	t.Eighteen = small(jay.ReadUint16(b[34:36]))
	t.Nineteen = small(jay.ReadUint16(b[36:38]))
	t.Twenty = small(jay.ReadUint16(b[38:]))
	return nil
}

func (t *TwentyOne) MarshalJ() (b []byte) {
	b = make([]byte, 42)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:24], uint16(t.Twelve))
	jay.WriteUint16(b[24:26], uint16(t.Thirteen))
	jay.WriteUint16(b[26:28], uint16(t.Fourteen))
	jay.WriteUint16(b[28:30], uint16(t.Fifteen))
	jay.WriteUint16(b[30:32], uint16(t.Sixteen))
	jay.WriteUint16(b[32:34], uint16(t.Seventeen))
	jay.WriteUint16(b[34:36], uint16(t.Eighteen))
	jay.WriteUint16(b[36:38], uint16(t.Nineteen))
	jay.WriteUint16(b[38:40], uint16(t.Twenty))
	jay.WriteUint16(b[40:], uint16(t.TwentyOne))
	return
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 42 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:24]))
	t.Thirteen = small(jay.ReadUint16(b[24:26]))
	t.Fourteen = small(jay.ReadUint16(b[26:28]))
	t.Fifteen = small(jay.ReadUint16(b[28:30]))
	t.Sixteen = small(jay.ReadUint16(b[30:32]))
	t.Seventeen = small(jay.ReadUint16(b[32:34]))
	t.Eighteen = small(jay.ReadUint16(b[34:36]))
	t.Nineteen = small(jay.ReadUint16(b[36:38]))
	t.Twenty = small(jay.ReadUint16(b[38:40]))
	t.TwentyOne = small(jay.ReadUint16(b[40:]))
	return nil
}

func (t *TwentyTwo) MarshalJ() (b []byte) {
	b = make([]byte, 44)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:24], uint16(t.Twelve))
	jay.WriteUint16(b[24:26], uint16(t.Thirteen))
	jay.WriteUint16(b[26:28], uint16(t.Fourteen))
	jay.WriteUint16(b[28:30], uint16(t.Fifteen))
	jay.WriteUint16(b[30:32], uint16(t.Sixteen))
	jay.WriteUint16(b[32:34], uint16(t.Seventeen))
	jay.WriteUint16(b[34:36], uint16(t.Eighteen))
	jay.WriteUint16(b[36:38], uint16(t.Nineteen))
	jay.WriteUint16(b[38:40], uint16(t.Twenty))
	jay.WriteUint16(b[40:42], uint16(t.TwentyOne))
	jay.WriteUint16(b[42:], uint16(t.TwentyTwo))
	return
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 44 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:24]))
	t.Thirteen = small(jay.ReadUint16(b[24:26]))
	t.Fourteen = small(jay.ReadUint16(b[26:28]))
	t.Fifteen = small(jay.ReadUint16(b[28:30]))
	t.Sixteen = small(jay.ReadUint16(b[30:32]))
	t.Seventeen = small(jay.ReadUint16(b[32:34]))
	t.Eighteen = small(jay.ReadUint16(b[34:36]))
	t.Nineteen = small(jay.ReadUint16(b[36:38]))
	t.Twenty = small(jay.ReadUint16(b[38:40]))
	t.TwentyOne = small(jay.ReadUint16(b[40:42]))
	t.TwentyTwo = small(jay.ReadUint16(b[42:]))
	return nil
}

func (t *TwentyThree) MarshalJ() (b []byte) {
	b = make([]byte, 46)
	jay.WriteUint16(b[:2], uint16(t.One))
	jay.WriteUint16(b[2:4], uint16(t.Two))
	jay.WriteUint16(b[4:6], uint16(t.Three))
	jay.WriteUint16(b[6:8], uint16(t.Four))
	jay.WriteUint16(b[8:10], uint16(t.Five))
	jay.WriteUint16(b[10:12], uint16(t.Six))
	jay.WriteUint16(b[12:14], uint16(t.Seven))
	jay.WriteUint16(b[14:16], uint16(t.Eight))
	jay.WriteUint16(b[16:18], uint16(t.Nine))
	jay.WriteUint16(b[18:20], uint16(t.Ten))
	jay.WriteUint16(b[20:22], uint16(t.Eleven))
	jay.WriteUint16(b[22:24], uint16(t.Twelve))
	jay.WriteUint16(b[24:26], uint16(t.Thirteen))
	jay.WriteUint16(b[26:28], uint16(t.Fourteen))
	jay.WriteUint16(b[28:30], uint16(t.Fifteen))
	jay.WriteUint16(b[30:32], uint16(t.Sixteen))
	jay.WriteUint16(b[32:34], uint16(t.Seventeen))
	jay.WriteUint16(b[34:36], uint16(t.Eighteen))
	jay.WriteUint16(b[36:38], uint16(t.Nineteen))
	jay.WriteUint16(b[38:40], uint16(t.Twenty))
	jay.WriteUint16(b[40:42], uint16(t.TwentyOne))
	jay.WriteUint16(b[42:44], uint16(t.TwentyTwo))
	jay.WriteUint16(b[44:], uint16(t.TwentyThree))
	return
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 46 {
		return jay.ErrUnexpectedEOB
	}
	t.One = small(jay.ReadUint16(b[:2]))
	t.Two = small(jay.ReadUint16(b[2:4]))
	t.Three = small(jay.ReadUint16(b[4:6]))
	t.Four = small(jay.ReadUint16(b[6:8]))
	t.Five = small(jay.ReadUint16(b[8:10]))
	t.Six = small(jay.ReadUint16(b[10:12]))
	t.Seven = small(jay.ReadUint16(b[12:14]))
	t.Eight = small(jay.ReadUint16(b[14:16]))
	t.Nine = small(jay.ReadUint16(b[16:18]))
	t.Ten = small(jay.ReadUint16(b[18:20]))
	t.Eleven = small(jay.ReadUint16(b[20:22]))
	t.Twelve = small(jay.ReadUint16(b[22:24]))
	t.Thirteen = small(jay.ReadUint16(b[24:26]))
	t.Fourteen = small(jay.ReadUint16(b[26:28]))
	t.Fifteen = small(jay.ReadUint16(b[28:30]))
	t.Sixteen = small(jay.ReadUint16(b[30:32]))
	t.Seventeen = small(jay.ReadUint16(b[32:34]))
	t.Eighteen = small(jay.ReadUint16(b[34:36]))
	t.Nineteen = small(jay.ReadUint16(b[36:38]))
	t.Twenty = small(jay.ReadUint16(b[38:40]))
	t.TwentyOne = small(jay.ReadUint16(b[40:42]))
	t.TwentyTwo = small(jay.ReadUint16(b[42:44]))
	t.TwentyThree = small(jay.ReadUint16(b[44:]))
	return nil
}