// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://github.com/speedyhoon/jay

package main

import "github.com/speedyhoon/jay"

func (o *One) MarshalJ() []byte {
	return []byte{byte(o.One)}
}

func (o *One) UnmarshalJ(b []byte) error {
	if len(b) < 1 {
		return jay.ErrUnexpectedEOB
	}
	o.One = int8(b[0])
	return nil
}

func (t *Two) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two)}
}

func (t *Two) UnmarshalJ(b []byte) error {
	if len(b) < 2 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	return nil
}

func (t *Three) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three)}
}

func (t *Three) UnmarshalJ(b []byte) error {
	if len(b) < 3 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	return nil
}

func (f *Four) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four)}
}

func (f *Four) UnmarshalJ(b []byte) error {
	if len(b) < 4 {
		return jay.ErrUnexpectedEOB
	}
	f.One = int8(b[0])
	f.Two = int8(b[1])
	f.Three = int8(b[2])
	f.Four = int8(b[3])
	return nil
}

func (f *Five) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five)}
}

func (f *Five) UnmarshalJ(b []byte) error {
	if len(b) < 5 {
		return jay.ErrUnexpectedEOB
	}
	f.One = int8(b[0])
	f.Two = int8(b[1])
	f.Three = int8(b[2])
	f.Four = int8(b[3])
	f.Five = int8(b[4])
	return nil
}

func (s *Six) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six)}
}

func (s *Six) UnmarshalJ(b []byte) error {
	if len(b) < 6 {
		return jay.ErrUnexpectedEOB
	}
	s.One = int8(b[0])
	s.Two = int8(b[1])
	s.Three = int8(b[2])
	s.Four = int8(b[3])
	s.Five = int8(b[4])
	s.Six = int8(b[5])
	return nil
}

func (s *Seven) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven)}
}

func (s *Seven) UnmarshalJ(b []byte) error {
	if len(b) < 7 {
		return jay.ErrUnexpectedEOB
	}
	s.One = int8(b[0])
	s.Two = int8(b[1])
	s.Three = int8(b[2])
	s.Four = int8(b[3])
	s.Five = int8(b[4])
	s.Six = int8(b[5])
	s.Seven = int8(b[6])
	return nil
}

func (e *Eight) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight)}
}

func (e *Eight) UnmarshalJ(b []byte) error {
	if len(b) < 8 {
		return jay.ErrUnexpectedEOB
	}
	e.One = int8(b[0])
	e.Two = int8(b[1])
	e.Three = int8(b[2])
	e.Four = int8(b[3])
	e.Five = int8(b[4])
	e.Six = int8(b[5])
	e.Seven = int8(b[6])
	e.Eight = int8(b[7])
	return nil
}

func (n *Nine) MarshalJ() []byte {
	return []byte{byte(n.One), byte(n.Two), byte(n.Three), byte(n.Four), byte(n.Five), byte(n.Six), byte(n.Seven), byte(n.Eight), byte(n.Nine)}
}

func (n *Nine) UnmarshalJ(b []byte) error {
	if len(b) < 9 {
		return jay.ErrUnexpectedEOB
	}
	n.One = int8(b[0])
	n.Two = int8(b[1])
	n.Three = int8(b[2])
	n.Four = int8(b[3])
	n.Five = int8(b[4])
	n.Six = int8(b[5])
	n.Seven = int8(b[6])
	n.Eight = int8(b[7])
	n.Nine = int8(b[8])
	return nil
}

func (t *Ten) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten)}
}

func (t *Ten) UnmarshalJ(b []byte) error {
	if len(b) < 10 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	return nil
}

func (e *Eleven) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight), byte(e.Nine), byte(e.Ten), byte(e.Eleven)}
}

func (e *Eleven) UnmarshalJ(b []byte) error {
	if len(b) < 11 {
		return jay.ErrUnexpectedEOB
	}
	e.One = int8(b[0])
	e.Two = int8(b[1])
	e.Three = int8(b[2])
	e.Four = int8(b[3])
	e.Five = int8(b[4])
	e.Six = int8(b[5])
	e.Seven = int8(b[6])
	e.Eight = int8(b[7])
	e.Nine = int8(b[8])
	e.Ten = int8(b[9])
	e.Eleven = int8(b[10])
	return nil
}

func (t *Twelve) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve)}
}

func (t *Twelve) UnmarshalJ(b []byte) error {
	if len(b) < 12 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	return nil
}

func (t *Thirteen) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen)}
}

func (t *Thirteen) UnmarshalJ(b []byte) error {
	if len(b) < 13 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	t.Thirteen = int8(b[12])
	return nil
}

func (f *Fourteen) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five), byte(f.Six), byte(f.Seven), byte(f.Eight), byte(f.Nine), byte(f.Ten), byte(f.Eleven), byte(f.Twelve), byte(f.Thirteen), byte(f.Fourteen)}
}

func (f *Fourteen) UnmarshalJ(b []byte) error {
	if len(b) < 14 {
		return jay.ErrUnexpectedEOB
	}
	f.One = int8(b[0])
	f.Two = int8(b[1])
	f.Three = int8(b[2])
	f.Four = int8(b[3])
	f.Five = int8(b[4])
	f.Six = int8(b[5])
	f.Seven = int8(b[6])
	f.Eight = int8(b[7])
	f.Nine = int8(b[8])
	f.Ten = int8(b[9])
	f.Eleven = int8(b[10])
	f.Twelve = int8(b[11])
	f.Thirteen = int8(b[12])
	f.Fourteen = int8(b[13])
	return nil
}

func (f *Fifteen) MarshalJ() []byte {
	return []byte{byte(f.One), byte(f.Two), byte(f.Three), byte(f.Four), byte(f.Five), byte(f.Six), byte(f.Seven), byte(f.Eight), byte(f.Nine), byte(f.Ten), byte(f.Eleven), byte(f.Twelve), byte(f.Thirteen), byte(f.Fourteen), byte(f.Fifteen)}
}

func (f *Fifteen) UnmarshalJ(b []byte) error {
	if len(b) < 15 {
		return jay.ErrUnexpectedEOB
	}
	f.One = int8(b[0])
	f.Two = int8(b[1])
	f.Three = int8(b[2])
	f.Four = int8(b[3])
	f.Five = int8(b[4])
	f.Six = int8(b[5])
	f.Seven = int8(b[6])
	f.Eight = int8(b[7])
	f.Nine = int8(b[8])
	f.Ten = int8(b[9])
	f.Eleven = int8(b[10])
	f.Twelve = int8(b[11])
	f.Thirteen = int8(b[12])
	f.Fourteen = int8(b[13])
	f.Fifteen = int8(b[14])
	return nil
}

func (s *Sixteen) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven), byte(s.Eight), byte(s.Nine), byte(s.Ten), byte(s.Eleven), byte(s.Twelve), byte(s.Thirteen), byte(s.Fourteen), byte(s.Fifteen), byte(s.Sixteen)}
}

func (s *Sixteen) UnmarshalJ(b []byte) error {
	if len(b) < 16 {
		return jay.ErrUnexpectedEOB
	}
	s.One = int8(b[0])
	s.Two = int8(b[1])
	s.Three = int8(b[2])
	s.Four = int8(b[3])
	s.Five = int8(b[4])
	s.Six = int8(b[5])
	s.Seven = int8(b[6])
	s.Eight = int8(b[7])
	s.Nine = int8(b[8])
	s.Ten = int8(b[9])
	s.Eleven = int8(b[10])
	s.Twelve = int8(b[11])
	s.Thirteen = int8(b[12])
	s.Fourteen = int8(b[13])
	s.Fifteen = int8(b[14])
	s.Sixteen = int8(b[15])
	return nil
}

func (s *Seventeen) MarshalJ() []byte {
	return []byte{byte(s.One), byte(s.Two), byte(s.Three), byte(s.Four), byte(s.Five), byte(s.Six), byte(s.Seven), byte(s.Eight), byte(s.Nine), byte(s.Ten), byte(s.Eleven), byte(s.Twelve), byte(s.Thirteen), byte(s.Fourteen), byte(s.Fifteen), byte(s.Sixteen), byte(s.Seventeen)}
}

func (s *Seventeen) UnmarshalJ(b []byte) error {
	if len(b) < 17 {
		return jay.ErrUnexpectedEOB
	}
	s.One = int8(b[0])
	s.Two = int8(b[1])
	s.Three = int8(b[2])
	s.Four = int8(b[3])
	s.Five = int8(b[4])
	s.Six = int8(b[5])
	s.Seven = int8(b[6])
	s.Eight = int8(b[7])
	s.Nine = int8(b[8])
	s.Ten = int8(b[9])
	s.Eleven = int8(b[10])
	s.Twelve = int8(b[11])
	s.Thirteen = int8(b[12])
	s.Fourteen = int8(b[13])
	s.Fifteen = int8(b[14])
	s.Sixteen = int8(b[15])
	s.Seventeen = int8(b[16])
	return nil
}

func (e *Eighteen) MarshalJ() []byte {
	return []byte{byte(e.One), byte(e.Two), byte(e.Three), byte(e.Four), byte(e.Five), byte(e.Six), byte(e.Seven), byte(e.Eight), byte(e.Nine), byte(e.Ten), byte(e.Eleven), byte(e.Twelve), byte(e.Thirteen), byte(e.Fourteen), byte(e.Fifteen), byte(e.Sixteen), byte(e.Seventeen), byte(e.Eighteen)}
}

func (e *Eighteen) UnmarshalJ(b []byte) error {
	if len(b) < 18 {
		return jay.ErrUnexpectedEOB
	}
	e.One = int8(b[0])
	e.Two = int8(b[1])
	e.Three = int8(b[2])
	e.Four = int8(b[3])
	e.Five = int8(b[4])
	e.Six = int8(b[5])
	e.Seven = int8(b[6])
	e.Eight = int8(b[7])
	e.Nine = int8(b[8])
	e.Ten = int8(b[9])
	e.Eleven = int8(b[10])
	e.Twelve = int8(b[11])
	e.Thirteen = int8(b[12])
	e.Fourteen = int8(b[13])
	e.Fifteen = int8(b[14])
	e.Sixteen = int8(b[15])
	e.Seventeen = int8(b[16])
	e.Eighteen = int8(b[17])
	return nil
}

func (n *Nineteen) MarshalJ() []byte {
	return []byte{byte(n.One), byte(n.Two), byte(n.Three), byte(n.Four), byte(n.Five), byte(n.Six), byte(n.Seven), byte(n.Eight), byte(n.Nine), byte(n.Ten), byte(n.Eleven), byte(n.Twelve), byte(n.Thirteen), byte(n.Fourteen), byte(n.Fifteen), byte(n.Sixteen), byte(n.Seventeen), byte(n.Eighteen), byte(n.Nineteen)}
}

func (n *Nineteen) UnmarshalJ(b []byte) error {
	if len(b) < 19 {
		return jay.ErrUnexpectedEOB
	}
	n.One = int8(b[0])
	n.Two = int8(b[1])
	n.Three = int8(b[2])
	n.Four = int8(b[3])
	n.Five = int8(b[4])
	n.Six = int8(b[5])
	n.Seven = int8(b[6])
	n.Eight = int8(b[7])
	n.Nine = int8(b[8])
	n.Ten = int8(b[9])
	n.Eleven = int8(b[10])
	n.Twelve = int8(b[11])
	n.Thirteen = int8(b[12])
	n.Fourteen = int8(b[13])
	n.Fifteen = int8(b[14])
	n.Sixteen = int8(b[15])
	n.Seventeen = int8(b[16])
	n.Eighteen = int8(b[17])
	n.Nineteen = int8(b[18])
	return nil
}

func (t *Twenty) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty)}
}

func (t *Twenty) UnmarshalJ(b []byte) error {
	if len(b) < 20 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	t.Thirteen = int8(b[12])
	t.Fourteen = int8(b[13])
	t.Fifteen = int8(b[14])
	t.Sixteen = int8(b[15])
	t.Seventeen = int8(b[16])
	t.Eighteen = int8(b[17])
	t.Nineteen = int8(b[18])
	t.Twenty = int8(b[19])
	return nil
}

func (t *TwentyOne) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne)}
}

func (t *TwentyOne) UnmarshalJ(b []byte) error {
	if len(b) < 21 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	t.Thirteen = int8(b[12])
	t.Fourteen = int8(b[13])
	t.Fifteen = int8(b[14])
	t.Sixteen = int8(b[15])
	t.Seventeen = int8(b[16])
	t.Eighteen = int8(b[17])
	t.Nineteen = int8(b[18])
	t.Twenty = int8(b[19])
	t.TwentyOne = int8(b[20])
	return nil
}

func (t *TwentyTwo) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne), byte(t.TwentyTwo)}
}

func (t *TwentyTwo) UnmarshalJ(b []byte) error {
	if len(b) < 22 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	t.Thirteen = int8(b[12])
	t.Fourteen = int8(b[13])
	t.Fifteen = int8(b[14])
	t.Sixteen = int8(b[15])
	t.Seventeen = int8(b[16])
	t.Eighteen = int8(b[17])
	t.Nineteen = int8(b[18])
	t.Twenty = int8(b[19])
	t.TwentyOne = int8(b[20])
	t.TwentyTwo = int8(b[21])
	return nil
}

func (t *TwentyThree) MarshalJ() []byte {
	return []byte{byte(t.One), byte(t.Two), byte(t.Three), byte(t.Four), byte(t.Five), byte(t.Six), byte(t.Seven), byte(t.Eight), byte(t.Nine), byte(t.Ten), byte(t.Eleven), byte(t.Twelve), byte(t.Thirteen), byte(t.Fourteen), byte(t.Fifteen), byte(t.Sixteen), byte(t.Seventeen), byte(t.Eighteen), byte(t.Nineteen), byte(t.Twenty), byte(t.TwentyOne), byte(t.TwentyTwo), byte(t.TwentyThree)}
}

func (t *TwentyThree) UnmarshalJ(b []byte) error {
	if len(b) < 23 {
		return jay.ErrUnexpectedEOB
	}
	t.One = int8(b[0])
	t.Two = int8(b[1])
	t.Three = int8(b[2])
	t.Four = int8(b[3])
	t.Five = int8(b[4])
	t.Six = int8(b[5])
	t.Seven = int8(b[6])
	t.Eight = int8(b[7])
	t.Nine = int8(b[8])
	t.Ten = int8(b[9])
	t.Eleven = int8(b[10])
	t.Twelve = int8(b[11])
	t.Thirteen = int8(b[12])
	t.Fourteen = int8(b[13])
	t.Fifteen = int8(b[14])
	t.Sixteen = int8(b[15])
	t.Seventeen = int8(b[16])
	t.Eighteen = int8(b[17])
	t.Nineteen = int8(b[18])
	t.Twenty = int8(b[19])
	t.TwentyOne = int8(b[20])
	t.TwentyTwo = int8(b[21])
	t.TwentyThree = int8(b[22])
	return nil
}