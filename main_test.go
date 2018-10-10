package main

import (
	"reflect"
	"testing"
)

func TestDecodeLogo(t *testing.T) {
	type args struct {
		logo string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success decode logo",
			args:    args{logo: "Zm9v"},
			want:    []byte("foo"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeLogo(tt.args.logo)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeLogo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeLogo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveLogo(t *testing.T) {
	type args struct {
		f string
		l []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success save logo",
			args:    args{f: "/tmp/test.jpeg", l: []byte("test")},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SaveLogo(tt.args.f, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveLogo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveLogo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTranslate(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name:    "success translate",
			args:    args{text: []string{"hello"}},
			want:    "привет",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTranslate(tt.args.text); got != tt.want {
				t.Errorf("GetTranslate() = %v, want %v", got, tt.want)
			}
		})
	}
}
