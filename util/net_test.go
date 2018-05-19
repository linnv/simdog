package util

import (
	"reflect"
	"testing"
)

func Test_loadNetState(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want NetState
	}{
		{"normal", args{"/Users/Jialin/golang/src/github.com/linnv/simdog/util/netstate.data"}, NetState{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadNetState(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadNetState() = %v, want %v", got, tt.want)
			}
		})
	}
}
