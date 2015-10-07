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
	array          []byte // TODO: need concurrency protection
	numBytesPerRow uint
}

// Returns the byte with the bit and that bits position in that byte
func (t *Xbm) getByte(x, y uint) (*byte, uint, error) {
	// TODO: better
	if x >= t.Width {
		return nil, 0x0, errors.New("X larger than the image width")
	}
	if y >= t.Height {
		return nil, 0x0, errors.New("Y larger than the image height")
	}

	const bitsPerByte uint = 8

	bite := (y - 1) * uint(t.numBytesPerRow)
	bite += (x / bitsPerByte)

	var bit uint
	bit = x % bitsPerByte

	// bite := bit / 8
	// 8,8
	// 8,8
	return &t.array[bite], bit, nil
}

func (t *Xbm) SetBit(x uint, y uint, val bool) error {
	bite, pos, err := t.getByte(x, y)
	if err != nil {
		return err
	}

	// TODO ?
	if val {
		*bite |= 1 << pos
	} else {
		*bite &= 1 << pos
	}

	return nil
}

func (t *Xbm) GetBit(x, y uint) (bool, error) {
	bite, pos, err := t.getByte(x, y)
	if err != nil {
		return false, err
	}

	return ((*bite & (1 << pos)) != 0), nil
	// return (*bite&pos != 0), nil
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

// func Decode(array *Xbm, location io.Reader) error {
//    return nil
// }

// func Encode(w io.Writer) (*Xbm, error) {
//    return nil,nil
// }

func New(name string, width, height uint) (*Xbm, error) {
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

	if ((img.Width * img.Height) % 8) != 0 {
		numBytes++
	}

	img.array = make([]byte, numBytes)

	return img, nil
}
