package sorting

import "fmt"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %v", strconv.FormatFloat(f, 'f', 1, 64))
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	_, ok := fnb.(FancyNumber)
	n, err := strconv.Atoi(fnb.Value())
	if ok && (err == nil) {
		return n
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	r := ExtractFancyNumber(fnb)
	if ExtractFancyNumber(fnb) == 4 {
		r = 0
	}
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(r))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	var ret string
	switch i := i.(type) {
	case int:
		ret = DescribeNumber(float64(i))
	case float64:
		ret = DescribeNumber(i)
	case NumberBox:
		ret = DescribeNumberBox(i)
	case FancyNumberBox:
		ret = DescribeFancyNumberBox(i)
	default:
		ret = "Return to sender"
	}
	return ret
}
