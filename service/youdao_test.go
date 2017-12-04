package service

import (
	"reflect"
	"testing"
)

func TestYoudao_Translate(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		y    *Youdao
		args args
		want *ResponseYoudao
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &Youdao{}
			if got := y.Translate(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Youdao.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
