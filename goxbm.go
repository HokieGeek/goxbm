package goxbm

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
)

// #define volume_100_width 23
// #define volume_100_height 16
// static unsigned char volume_100_bits[] = {
//    0x00, 0x00, 0x60, 0x00, 0x00, 0x60, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x6c,
//    0x00, 0x80, 0x6d, 0x00, 0x80, 0x6d, 0x00, 0xb0, 0x6d, 0x00, 0xb0, 0x6d,
//    0x00, 0xb6, 0x6d, 0x00, 0xb6, 0x6d, 0xc0, 0xb6, 0x6d, 0xc0, 0xb6, 0x6d,
//    0xd8, 0xb6, 0x6d, 0xd8, 0xb6, 0x6d, 0xdb, 0xb6, 0x6d, 0xdb, 0xb6, 0x6d };

const bitsPerByte uint = 8

type Xbm struct {
	Name           string
	Width          uint
	Height         uint
	numBytesPerRow uint
	array          []byte // TODO: need concurrency protection
}

// Returns the byte with the bit and that bits position in that byte
func (t *Xbm) getByte(x, y uint) (*byte, uint, error) {
	if x >= t.Width {
		return nil, 0x0, errors.New("X not without allowable range")
	}
	if y >= t.Height {
		return nil, 0x0, errors.New("Y not without allowable range")
	}

	const bitsPerByte uint = 8

	bite := y * uint(t.numBytesPerRow)
	bite += (x / bitsPerByte)

	var bit uint
	bit = x % bitsPerByte

	return &t.array[bite], bit, nil
}

func (t *Xbm) SetBit(x uint, y uint, val bool) error {
	fmt.Printf("SetBit(%d, %d, %t)\n", x, y, val)
	bite, pos, err := t.getByte(x, y)
	if err != nil {
		return err
	}

	if val {
		*bite |= 1 << pos
	} else {
		*bite &= 1 << pos
	}

	return nil
}

func (t *Xbm) GetBit(x, y uint) (bool, error) {
	// fmt.Printf("GetBit(%d, %d)\n", x, y)
	bite, pos, err := t.getByte(x, y)
	if err != nil {
		return false, err
	}

	return ((*bite & (1 << pos)) != 0), nil
}

func (t *Xbm) String() string {
	var buf bytes.Buffer

	// Write out the width
	buf.WriteString("#define ")
	buf.WriteString(t.Name)
	buf.WriteString("_width ")
	buf.WriteString(strconv.Itoa(int(t.Width)))
	buf.WriteString("\n")

	// Write out the height
	buf.WriteString("#define ")
	buf.WriteString(t.Name)
	buf.WriteString("_height ")
	buf.WriteString(strconv.Itoa(int(t.Height)))
	buf.WriteString("\n")

	// Write out the byte array as hexes
	buf.WriteString("static unsigned char ")
	buf.WriteString(t.Name)
	buf.WriteString("_bits[] = {\n")
	for i, b := range t.array {
		// TODO: columnize?
		buf.WriteString(fmt.Sprintf("0x%02x", b))
		if i != len(t.array)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(" };")
	//    0x00, 0x00, 0x60, 0x00, 0x00, 0x60, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x6c,
	//    0x00, 0x80, 0x6d, 0x00, 0x80, 0x6d, 0x00, 0xb0, 0x6d, 0x00, 0xb0, 0x6d,
	//    0x00, 0xb6, 0x6d, 0x00, 0xb6, 0x6d, 0xc0, 0xb6, 0x6d, 0xc0, 0xb6, 0x6d,
	//    0xd8, 0xb6, 0x6d, 0xd8, 0xb6, 0x6d, 0xdb, 0xb6, 0x6d, 0xdb, 0xb6, 0x6d };

	return buf.String()
}

func (t *Xbm) Draw() {
	var bitcount int
	for bite := 0; bite < len(t.array); bite++ {
		var pos uint
		for pos = 0; pos < 8; pos++ {
			if (t.array[bite] & (1 << pos)) != 0 {
				fmt.Printf("0")
			} else {
				fmt.Printf("-")
			}

			bitcount++
			if bitcount >= int(t.Width) {
				fmt.Println("")
				bitcount = 0
				break
			}
		}
	}
	fmt.Println("")
}

// func Decode(array *Xbm, location io.Reader) error {
//    return nil
// }

// func Encode(w io.Writer) (*Xbm, error) {
//    return nil,nil
// }

func New(name string, width, height uint) (*Xbm, error) {
	// fmt.Printf("goxbm.New(%s, %d, %d)\n", name, width, height)
	// TODO: Make smarter
	if width <= 0 {
		return nil, errors.New("Width must be greater than 0")
	}
	if height <= 0 {
		return nil, errors.New("Height must be greater than 0")
	}

	img := new(Xbm)
	img.Name = name
	img.Width = width
	img.Height = height

	// Determine the number of bytes needed
	img.numBytesPerRow = uint(math.Ceil(float64(17) / float64(bitsPerByte)))
	numBytes := img.numBytesPerRow * height
	img.array = make([]byte, numBytes)

	fmt.Printf("numBytes: %d\n", numBytes)
	fmt.Printf("array (%d): %v\n", len(img.array), img.array)

	return img, nil
}
