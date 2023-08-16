package ibmpisession

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/runtime"
)

func TestNewIBMPISession(t *testing.T) {
	bearerTokenAuth := &core.BearerTokenAuthenticator{BearerToken: "sample"}
	o1 := &IBMPIOptions{
		Authenticator: bearerTokenAuth,
		UserAccount:   "1234",
		Region:        "dal",
		Zone:          "dal12",
		URL:           "dal.power-iaas.test.cloud.ibm.com",
	}
	type args struct {
		o *IBMPIOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *IBMPISession
		wantErr bool
	}{
		{
			name: "Simple BearerTokenAuthenticator",
			args: args{
				o: o1,
			},
			want: &IBMPISession{
				Options:   o1,
				CRNFormat: "crn:v1:staging:public:power-iaas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Without Options",
			args: args{
				o: nil,
			},
			wantErr: true,
		},
		{
			name: "Invalid Authenticator",
			args: args{
				o: &IBMPIOptions{
					Authenticator: &core.BearerTokenAuthenticator{},
				},
			},
			wantErr: true,
		},
		{
			name: "Without Authenticator",
			args: args{
				o: &IBMPIOptions{
					UserAccount: "1234",
					Zone:        "dal12",
					URL:         "dal.power-iaas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without UserAccount",
			args: args{
				o: &IBMPIOptions{
					Authenticator: bearerTokenAuth,
					Zone:          "dal12",
					URL:           "dal.power-iaas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without Zone",
			args: args{
				o: &IBMPIOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					URL:           "dal.power-iaas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without URL and Region",
			args: args{
				o: &IBMPIOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					Zone:          "dal12",
				},
			},
			want: &IBMPISession{
				Options:   o1,
				CRNFormat: "crn:v1:bluemix:public:power-iaas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Without URL but with region",
			args: args{
				o: &IBMPIOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					Zone:          "dal12",
					Region:        "dal",
				},
			},
			want: &IBMPISession{
				Options:   o1,
				CRNFormat: "crn:v1:bluemix:public:power-iaas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Simple URL with https",
			args: args{
				o: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Region:        "dal",
					Zone:          "dal12",
					URL:           "https://dal.power-iaas.test.cloud.ibm.com",
				},
			},
			want: &IBMPISession{
				Options: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:power-iaas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Simple URL with http",
			args: args{
				o: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Region:        "dal",
					Zone:          "dal12",
					URL:           "http://dal.power-iaas.test.cloud.ibm.com",
				},
			},
			want: &IBMPISession{
				Options: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:power-iaas:dal12:a/1234:%s::",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIBMPISession(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIBMPISession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if same, msg := isIBMPISessionEqual(got, tt.want); !same {
				t.Error(msg)
			}
		})
	}
}

func TestNewIBMPISessionViaEnv(t *testing.T) {
	os.Setenv("IBMCLOUD_PI_API_ENDPOINT", "power-iaas.test.cloud.ibm.com")
	type args struct {
		o *IBMPIOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *IBMPISession
		wantErr bool
	}{
		{
			name: "URL from Env Without Region",
			args: args{
				o: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Zone:          "dal12",
				},
			},
			want: &IBMPISession{
				Options: &IBMPIOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:power-iaas:dal12:a/1234:%s::",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIBMPISession(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIBMPISession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if same, msg := isIBMPISessionEqual(got, tt.want); !same {
				t.Error(msg)
			}
		})
	}
}

func isIBMPISessionEqual(x *IBMPISession, y *IBMPISession) (bool, string) {
	if x == nil && y == nil {
		return true, ""
	}
	if (x == nil && y != nil) || (x != nil && y == nil) {
		return true, fmt.Sprintf("NewIBMPISession() = %v, want %v", x, y)
	}
	if x.Options.Authenticator != y.Options.Authenticator {
		return false, fmt.Sprintf("NewIBMPISession() Authenticator = %v, want %v", x.Options.Authenticator, y.Options.Authenticator)
	}
	if x.CRNFormat != y.CRNFormat {
		return false, fmt.Sprintf("NewIBMPISession() CRNFormat = %v, want %v", x.CRNFormat, y.CRNFormat)
	}
	return true, ""
}

func TestIBMPISession_AuthInfo(t *testing.T) {
	type args struct {
		cloudInstanceID string
	}
	type wantHeaders struct {
		auth string
		crn  string
	}
	tests := []struct {
		name        string
		s           *IBMPISession
		args        args
		wantHeaders wantHeaders
		wantErr     bool
	}{
		{
			name: "Auth",
			s: &IBMPISession{
				Options: &IBMPIOptions{
					Authenticator: &core.BearerTokenAuthenticator{BearerToken: "sample"},
					Zone:          "dal12",
					URL:           "dal.power-iaas.test.cloud.ibm.com",
				},
				CRNFormat: "crn...:%s::",
			},
			args: args{
				cloudInstanceID: "1234",
			},
			wantHeaders: wantHeaders{
				auth: "Bearer sample",
				crn:  "crn...:1234::",
			},
		},
		{
			name: "Incorrect Auth",
			s: &IBMPISession{
				Options: &IBMPIOptions{
					Authenticator: &core.IamAuthenticator{ApiKey: "sample"},
					Zone:          "dal12",
					URL:           "dal.power-iaas.test.cloud.ibm.com",
				},
				CRNFormat: "crn...:%s::",
			},
			args: args{
				cloudInstanceID: "1234",
			},
			wantHeaders: wantHeaders{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.AuthInfo(tt.args.cloudInstanceID)
			if got == nil {
				t.Errorf("NewAuth() = %v", got)
			}
			// Authenticate the request
			r := &runtime.TestClientRequest{}
			err := got.AuthenticateRequest(r, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthInfo().AuthenticateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Test the header values
			auth := r.Headers.Get("Authorization")
			if auth != tt.wantHeaders.auth {
				t.Errorf("AuthInfo().AuthenticateRequest() Authorization = %v, want %v", auth, tt.wantHeaders.auth)
			}
			crn := r.Headers.Get("CRN")
			if crn != tt.wantHeaders.crn {
				t.Errorf("AuthInfo().AuthenticateRequest() CRN = %v, want %v", crn, tt.wantHeaders.crn)
			}
		})
	}
}
