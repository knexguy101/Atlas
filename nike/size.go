package nike

import (
	"errors"
	"fmt"
	"math"
)

/*
	Size

	Using a BaseValue int vs just the raw string itself to make converting between sizes easier.
	If we swap the region, it auto swaps without a lot of weird conversion.
	Downside to this being user input must follow this non-standard format, but eh, that's for future me to figure out.
 */

type Size struct {
	BaseValue float64
	Kids bool
}

//GetSizeStringByRegion gets the size used by the nike website
func (s *Size) GetSizeStringByRegion(region string) (str string, err error) {
	switch region {
	case "US":
		str = s.ToNAMens()
		if s.Kids {
			str += "Y"
		}
		break
	case "UK":
		str = s.ToUK()
		break
	case "EU":
		str = s.ToEU()
		break
	default:
		err = errors.New("region not found")
		break
	}
	return
}

//ToNAMens gets the size in US mens
func (s *Size) ToNAMens() string {
	return s.formatSize(3.5 + (s.BaseValue * .5))
}

//ToNAWomens gets the size in US womens, used for display purposes
func (s *Size) ToNAWomens() string {
	return s.formatSize(5 + (s.BaseValue * .5))
}

//ToKids gets the size in US kids
func (s *Size) ToKids() string {
	return s.formatSize(5 + (s.BaseValue * .5))
}

//ToEU gets the size in EU
func (s *Size) ToEU() string {
	return s.formatSize(35.5 + (s.BaseValue * .5))
}

//ToUK gets the size in UK
func (s *Size) ToUK() string {
	return s.formatSize(3 + (s.BaseValue * .5))
}

//formatSize formats the given size float64 into a string, if it has a remainder, it rounds to 1 decimal point, otherwise it floors it to no decimals
func (s *Size) formatSize(x float64) (str string) {
	if x == math.Trunc(x) {
		str = fmt.Sprintf("%.0f", x)
	} else {
		str = fmt.Sprintf("%.1f", x)
	}
	return
}