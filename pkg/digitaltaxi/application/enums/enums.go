package enums

import (
	"fmt"
	"io"
	"strconv"
)

type CoverType string

const (
	ThirdParty    CoverType = "TPO"
	Comprehensive CoverType = "COMPREHENSIVE"
)

// AllCoverTypes is a list of all cover types
var AllCoverTypes = []CoverType{
	ThirdParty,
	Comprehensive,
}

// IsValid ...
func (t CoverType) IsValid() bool {
	switch t {
	case ThirdParty, Comprehensive:
		return true
	}

	return false
}

// String ...
func (t CoverType) String() string {
	return string(t)
}

// UnmarshalGQL ...
func (e *CoverType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CoverType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid task status", str)
	}

	return nil
}

// MarshalGQL writes the task status to the supplied writer as a quoted string
func (t CoverType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
