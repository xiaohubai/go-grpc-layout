package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	type args struct {
		str []byte
		b   []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Md5",
			args: args{
				str: []byte("123456+t!dV$"),
			},
			want: "ce8c73b00381ac97ceba4feaf6f67b4d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.str, tt.args.b...); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "RandString",
			args: args{
				n: 6,
			},
			want: "ss",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandString(tt.args.n)
			fmt.Println(got)
		})
	}
}
