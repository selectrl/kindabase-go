package kindabase

import (
	"bytes"
	"math"
	"testing"
)

func TestBase26Alphabet(t *testing.T) {
	alphabet := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}
	if !bytes.Equal(Base26, alphabet) {
		t.Error("Base26 alphabet not equal")
	}
}

func TestBase36Alphabet(t *testing.T) {
	alphabet := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}
	if !bytes.Equal(Base36, alphabet) {
		t.Error("Base36 alphabet not equal")
	}
}

func TestBase62Alphabet(t *testing.T) {
	alphabet := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z'}
	if !bytes.Equal(Base62, alphabet) {
		t.Error("Base62 alphabet not equal")
	}
}

func TestDecode(t *testing.T) {
	i := Decode("0000", Base36)
	expected := 47988
	if i != expected {
		t.Errorf(`expected "%d" got "%d"`, expected, i)
	}
}

func TestEncode(t *testing.T) {
	s := Encode(47988, Base36)
	expected := "0000"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestDecodeNegative(t *testing.T) {
	i := Decode("-0000", Base36)
	expected := -47988
	if i != expected {
		t.Errorf(`expected "%d" got "%d"`, expected, i)
	}
}

func TestEncodeNegative(t *testing.T) {
	s := Encode(-47988, Base36)
	expected := "-0000"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

/*
Test base26 decode and encode with max ints and units
*/

func TestDecode26MaxInt8(t *testing.T) {
	i := Decode26("dx")
	if i != math.MaxInt8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt8, i)
	}
}

func TestDecode26MaxInt16(t *testing.T) {
	i := Decode26("avlh")
	if i != math.MaxInt16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt16, i)
	}
}

func TestDecode26MaxInt32(t *testing.T) {
	i := Decode26("fxshrxx")
	if i != math.MaxInt32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt32, i)
	}
}

func TestDecode26MaxInt64(t *testing.T) {
	i := Decode26("crpxnlskvljfhh")
	if i != math.MaxInt64 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt64, i)
	}
}

func TestDecode26MaxUint8(t *testing.T) {
	i := Decode26("iv")
	if i != math.MaxUint8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint8, i)
	}
}

func TestDecode26MaxUint16(t *testing.T) {
	i := Decode26("crxp")
	if i != math.MaxUint16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint16, i)
	}
}

func TestDecode26MaxUint32(t *testing.T) {
	i := Decode26("mwlqkwv")
	if i != math.MaxUint32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint32, i)
	}
}

func TestEncode26MaxInt8(t *testing.T) {
	s := Encode26(math.MaxInt8)
	const expected = "dx"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxInt16(t *testing.T) {
	s := Encode26(math.MaxInt16)
	const expected = "avlh"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxInt32(t *testing.T) {
	s := Encode26(math.MaxInt32)
	const expected = "fxshrxx"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxInt64(t *testing.T) {
	s := Encode26(math.MaxInt64)
	const expected = "crpxnlskvljfhh"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxUint8(t *testing.T) {
	s := Encode26(uint8(math.MaxUint8))
	const expected = "iv"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxUint16(t *testing.T) {
	s := Encode26(uint16(math.MaxUint16))
	const expected = "crxp"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode26MaxUint32(t *testing.T) {
	s := Encode26(uint32(math.MaxUint32))
	const expected = "mwlqkwv"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

/*
Test base36 decode and encode with max ints and units
*/

func TestDecode36MaxInt8(t *testing.T) {
	i := Decode36("2j")
	if i != math.MaxInt8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt8, i)
	}
}

func TestDecode36MaxInt16(t *testing.T) {
	i := Decode36("o97")
	if i != math.MaxInt16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt16, i)
	}
}

func TestDecode36MaxInt32(t *testing.T) {
	i := Decode36("yhizyj")
	if i != math.MaxInt32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt32, i)
	}
}

func TestDecode36MaxInt64(t *testing.T) {
	i := Decode36("0x1nzhi21d7d7")
	if i != math.MaxInt64 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt64, i)
	}
}

func TestDecode36MaxUint8(t *testing.T) {
	i := Decode36("63")
	if i != math.MaxUint8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint8, i)
	}
}

func TestDecode36MaxUint16(t *testing.T) {
	i := Decode36("0djf")
	if i != math.MaxUint16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint16, i)
	}
}

func TestDecode36MaxUint32(t *testing.T) {
	i := Decode36("0y030y3")
	if i != math.MaxUint32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint32, i)
	}
}

func TestEncode36MaxInt8(t *testing.T) {
	s := Encode36(math.MaxInt8)
	const expected = "2j"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxInt16(t *testing.T) {
	s := Encode36(math.MaxInt16)
	const expected = "o97"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxInt32(t *testing.T) {
	s := Encode36(math.MaxInt32)
	const expected = "yhizyj"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxInt64(t *testing.T) {
	s := Encode36(math.MaxInt64)
	const expected = "0x1nzhi21d7d7"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxUint8(t *testing.T) {
	s := Encode36(uint8(math.MaxUint8))
	const expected = "63"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxUint16(t *testing.T) {
	s := Encode36(uint16(math.MaxUint16))
	const expected = "0djf"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode36MaxUint32(t *testing.T) {
	s := Encode36(uint32(math.MaxUint32))
	const expected = "0y030y3"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

/*
Test base62 decode and encode with max ints and units
*/

func TestDecode62MaxInt8(t *testing.T) {
	i := Decode62("13")
	if i != math.MaxInt8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt8, i)
	}
}

func TestDecode62MaxInt16(t *testing.T) {
	i := Decode62("7VV")
	if i != math.MaxInt16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt16, i)
	}
}

func TestDecode62MaxInt32(t *testing.T) {
	i := Decode62("1KJba1")
	if i != math.MaxInt32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt32, i)
	}
}

func TestDecode62MaxInt64(t *testing.T) {
	i := Decode62("9yK7lzX47l7")
	if i != math.MaxInt64 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxInt64, i)
	}
}

func TestDecode62MaxUint8(t *testing.T) {
	i := Decode62("37")
	if i != math.MaxUint8 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint8, i)
	}
}

func TestDecode62MaxUint16(t *testing.T) {
	i := Decode62("G21")
	if i != math.MaxUint16 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint16, i)
	}
}

func TestDecode62MaxUint32(t *testing.T) {
	i := Decode62("3feEB3")
	if i != math.MaxUint32 {
		t.Errorf(`expected "%d" got "%d"`, math.MaxUint32, i)
	}
}

func TestEncode62MaxInt8(t *testing.T) {
	s := Encode62(math.MaxInt8)
	const expected = "13"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxInt16(t *testing.T) {
	s := Encode62(math.MaxInt16)
	const expected = "7VV"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxInt32(t *testing.T) {
	s := Encode62(math.MaxInt32)
	const expected = "1KJba1"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxInt64(t *testing.T) {
	s := Encode62(math.MaxInt64)
	const expected = "9yK7lzX47l7"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxUint8(t *testing.T) {
	s := Encode62(uint8(math.MaxUint8))
	const expected = "37"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxUint16(t *testing.T) {
	s := Encode62(uint16(math.MaxUint16))
	const expected = "G21"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}

func TestEncode62MaxUint32(t *testing.T) {
	s := Encode62(uint32(math.MaxUint32))
	const expected = "3feEB3"
	if s != expected {
		t.Errorf(`expected "%s" got "%s"`, expected, s)
	}
}
