package main

import (
	"reflect"
	"testing"
)

func Test_weakHash(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Should return '[113 119 101 212 29 140 217 143 0 178 4 233 128 9 152 236 248 66 126]' when value is 'qwe'",
			args: args{value: "qwe"},
			want: []byte{113, 119, 101, 212, 29, 140, 217, 143, 0, 178, 4, 233, 128, 9, 152, 236, 248, 66, 126},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := weakHash(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("weakHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
