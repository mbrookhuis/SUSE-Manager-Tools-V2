package base64conv

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "equal",
			args: args{data: "hello from golang"},
			want: "aGVsbG8gZnJvbSBnb2xhbmc=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.data); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "equal",
			args:       args{data: "aGVsbG8gZnJvbSBnb2xhbmc="},
			wantOutput: "hello from golang",
			wantErr:    false,
		},
		{
			name:       "equal",
			args:       args{data: "aGVsbG8gZnJvbSBnb2xhbmc="},
			wantOutput: "hello from golang",
			wantErr:    false,
		},
		{
			name:       "Negative",
			args:       args{data: "-1"},
			wantOutput: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("Decode() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
