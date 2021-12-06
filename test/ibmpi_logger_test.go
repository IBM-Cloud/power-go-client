package test

import (
	"reflect"
	"testing"

	session "github.com/IBM-Cloud/power-go-client/ibmpisession"
)

func Test_SanatizeArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		wantOut []interface{}
	}{
		{
			name:    "Mask the Authorization token",
			args:    []interface{}{"Authorization: bearer supersecrettoken", 123},
			wantOut: []interface{}{"Authorization: [PRIVATE DATA HIDDEN]", 123},
		},
		{
			name:    "Does not mask anything",
			args:    []interface{}{"some random text", 123},
			wantOut: []interface{}{"some random text", 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := session.SanatizeArgs(tt.args); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("SanatizeArgs() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
