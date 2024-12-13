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
		return fmt.Errorf("%s is not a valid cover type", str)
	}

	return nil
}

// MarshalGQL writes the task status to the supplied writer as a quoted string
func (t CoverType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}

type CoverPeriod string

const (
	DailyCover   CoverPeriod = "DAILY"
	MonthlyCover CoverPeriod = "MONTHLY"
	AnnualCover  CoverPeriod = "ANNUAL"
)

var AllCoverPeriods = []CoverPeriod{
	DailyCover,
	MonthlyCover,
	AnnualCover,
}

// IsValid ...
func (t CoverPeriod) IsValid() bool {
	switch t {
	case DailyCover, MonthlyCover, AnnualCover:
		return true
	}

	return false
}

// String ...
func (t CoverPeriod) String() string {
	return string(t)
}

// UnmarshalGQL ...
func (e *CoverPeriod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CoverPeriod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid cover period", str)
	}

	return nil
}

// MarshalGQL writes the task status to the supplied writer as a quoted string
func (t CoverPeriod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
