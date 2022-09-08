package utils

import (
	"testing"
)

func TestIs_Email(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "name1",
			args: args{email: "1234@qq.com"},
			want: true,
		},
		{
			name: "name2",
			args: args{email: "1234qq.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is_Email(tt.args.email); got != tt.want {
				t.Errorf("Is_Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs_Number(t *testing.T) {
	t.Log(Is_Number("123456"))
	t.Log(Is_Number("a123456"))
}

func TestIs_UserName(t *testing.T) {
	t.Log(Is_UserName("a123"))
	t.Log(Is_UserName("123456"))
	t.Log(Is_UserName("a12345"))
	t.Log(Is_UserName("a12345*"))
}
