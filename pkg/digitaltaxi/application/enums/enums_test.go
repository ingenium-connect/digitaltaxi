package enums

import (
	"bytes"
	"strconv"
	"testing"
)

func TestCoverTypes_String(t *testing.T) {
	tests := []struct {
		name string
		e    CoverType
		want string
	}{
		{
			name: "TPO",
			e:    ThirdParty,
			want: "TPO",
		},
		{
			name: "COMPREHENSIVE",
			e:    Comprehensive,
			want: "COMPREHENSIVE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("CoverType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoverTypes_UnmarshalGQL(t *testing.T) {
	value := ThirdParty
	invalidType := CoverType("invalid")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *CoverType
		args    args
		wantErr bool
	}{
		{
			name: "valid type",
			e:    &value,
			args: args{
				v: "TPO",
			},
			wantErr: false,
		},
		{
			name: "invalid type",
			e:    &invalidType,
			args: args{
				v: "this is not a valid type",
			},
			wantErr: true,
		},
		{
			name: "non string type",
			e:    &invalidType,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("CoverType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCoverTypes_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     CoverType
		wantW string
	}{
		{
			name:  "TPO",
			e:     ThirdParty,
			wantW: strconv.Quote("TPO"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if got := w.String(); got != tt.wantW {
				t.Errorf("CoverType.MarshalGQL() = %v, want %v", got, tt.wantW)
			}
		})
	}
}

func TestCoverTypes_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    CoverType
		want bool
	}{
		{
			name: "Valid cover type - COMPREHENSIVE",
			e:    Comprehensive,
			want: true,
		},
		{
			name: "Invalid cover type",
			e:    "invalid",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("CoverType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
