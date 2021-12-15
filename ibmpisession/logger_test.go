package ibmpisession

import (
	"reflect"
	"testing"
)

func Test_sanatizeArgs(t *testing.T) {
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
			if gotOut := sanatizeArgs(tt.args); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("sanatizeArgs() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
