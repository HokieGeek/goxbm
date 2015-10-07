package goxbm

import "testing"

func TestCreationName(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 1

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	if xbm.Name != testName {
		t.Errorf("Retrieved name '%s' does not match expected name '%s'\n", xbm.Name, testName)
	}
}

func TestCreationNameError(t *testing.T) {
	t.Skip("TODO")
}

func TestCreationWidth(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 1

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	if xbm.Width != testWidth {
		t.Errorf("Retrieved width '%s' does not match expected width '%s'\n", xbm.Width, testWidth)
	}
}

func TestCreationWidthError(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 0
	testHeight = 1

	_, err := New(testName, testWidth, testHeight)
	if err == nil {
		t.Errorf("Constructor allowed me to create an image with width %d\n", testWidth)
	}
}

func TestCreationHeight(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 1

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	if xbm.Height != testHeight {
		t.Errorf("Retrieved height '%s' does not match expected height '%s'\n", xbm.Height, testHeight)
	}
}

func TestCreationHeightError(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 0

	_, err := New(testName, testWidth, testHeight)
	if err == nil {
		t.Errorf("Constructor allowed me to create an image with height %d\n", testHeight)
	}
}

func TestCreationArray(t *testing.T) {
	t.Skip("The logic for this one is not sound")
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 5
	testHeight = 5

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	expectedArrayLength := int(testWidth * testHeight)
	if len(xbm.array) != expectedArrayLength {
		// TODO: hmm....
		t.Fatal("Array size is not the expected size")
	}
}

func TestSettingBit(t *testing.T) {
	t.Skip("TODO")
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 1

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	// Flip to true
	testVal := true
	err = xbm.SetBit(0, 0, testVal)
	if err != nil {
		t.Fatalf("Received unexpected error when setting bit: %s\n", err)
	}

	var val bool
	val, err = xbm.GetBit(0, 0)
	if val != testVal {
		t.Fatalf("Bit was not set to %t as expected\n", testVal)
	}

	// Now reverse it
	testVal = !testVal
	err = xbm.SetBit(0, 0, testVal)
	if err != nil {
		t.Fatalf("Received unexpected error when setting bit: %s\n", err)
	}

	val, err = xbm.GetBit(0, 0)
	if val != testVal {
		t.Fatalf("Bit was not set to %t as expected\n", testVal)
	}
}

func TestSetBitError(t *testing.T) {
	t.Skip("TODO")
}

func TestGetBitError(t *testing.T) {
	t.Skip("TODO")
}

func TestString(t *testing.T) {
	testName := "foobar"
	var testWidth, testHeight uint
	testWidth = 1
	testHeight = 1

	xbm, err := New(testName, testWidth, testHeight)
	if err != nil {
		t.Fatal("Unable to create simple object")
	}

	if len(xbm.String()) <= 0 {
		t.Error("String function returned empty string")
	}
}

func TestEncode(t *testing.T) {
	t.Skip("TODO")
}

func TestEncodeEerror(t *testing.T) {
	t.Skip("TODO")
}

func TestDecode(t *testing.T) {
	t.Skip("TODO")
}

func TestDecodeError(t *testing.T) {
	t.Skip("TODO")
}
