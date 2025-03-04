package api

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_handlePost(t *testing.T) {
	type args struct {
		db map[string]string
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handlePost(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handlePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
