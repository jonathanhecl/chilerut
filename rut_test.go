package chilerut

import (
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"12.667.869-k",
			args{
				rut: "12.667.869.k",
			},
			"12667869-K",
		},
		{
			"12.667.869-k",
			args{
				rut: "12.667.869.k",
			},
			"12667869-K",
		},
		{
			"12-667-869-k",
			args{
				rut: "12-667-869-k",
			},
			"12667869-K",
		},
		{
			"12*667*869*K",
			args{
				rut: "12*667*869*K",
			},
			"12667869-K",
		},
		{
			"12 667 869 k",
			args{
				rut: "12 667 869 k",
			},
			"12667869-K",
		},
		{
			"   000012667869k   ",
			args{
				rut: "   000012667869k   ",
			},
			"12667869-K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.rut); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerificationDigit(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name   string
		args   args
		wantDv string
	}{
		{
			"9868503(0)",
			args{
				rut: "9868503",
			},
			"0",
		},
		{
			"12667869(K)",
			args{
				rut: "12667869",
			},
			"K",
		},
		{
			"16.647.869(3)",
			args{
				rut: "16.647.869",
			},
			"3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDv := VerificationDigit(tt.args.rut); gotDv != tt.wantDv {
				t.Errorf("VerificationDigit() = %v, want %v", gotDv, tt.wantDv)
			}
		})
	}
}

func TestValid(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"12345678-9",
			args{
				rut: "12345678-9",
			},
			false,
		},
		{
			"6265837-1",
			args{
				rut: "6265837-1",
			},
			true,
		},
		{
			"98685030",
			args{
				rut: "98685030",
			},
			true,
		},
		{
			"9868503-0",
			args{
				rut: "9868503-0",
			},
			true,
		},
		{
			"9.868.503-0",
			args{
				rut: "9.868.503-0",
			},
			true,
		},
		{
			"12.667.869-K",
			args{
				rut: "12.667.869-K",
			},
			true,
		},
		{
			"12.667.869-k",
			args{
				rut: "12.667.869-k",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.rut); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Compare(t *testing.T) {
	type args struct {
		rut1 string
		rut2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"12.667.869-K == 12.667.869-k",
			args{
				rut1: "12.667.869-K",
				rut2: "12.667.869-k",
			},
			true,
		},
		{
			"12.667.869-K != 12.667.861-K",
			args{
				rut1: "12.667.869-K",
				rut2: "12.667.861-K",
			},
			false,
		},
		{
			"12667869K == 12.667.869-k",
			args{
				rut1: "12667869K",
				rut2: "12.667.869-k",
			},
			true,
		},
		{
			"12.667.869-K == 12667869k",
			args{
				rut1: "12.667.869-K",
				rut2: "12667869k",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.rut1, tt.args.rut2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
