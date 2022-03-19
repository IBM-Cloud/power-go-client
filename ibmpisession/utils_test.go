package ibmpisession

import (
	"bytes"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
)

func Test_fetchAuthorizationData(t *testing.T) {
	type args struct {
		a core.Authenticator
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Fetch Sample Bearer Token",
			args: args{
				a: &core.BearerTokenAuthenticator{
					BearerToken: "sample",
				},
			},
			want: "Bearer sample",
		},
		{
			name: "Fetch Incorrect IAM Token",
			args: args{
				a: &core.IamAuthenticator{
					ApiKey: "sample",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchAuthorizationData(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchAuthorizationData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fetchAuthorizationData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crnBuilder(t *testing.T) {
	type args struct {
		useraccount string
		regionZone  string
		host        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Generate for Prod",
			args: args{"12345", "dal12", "dal.power-iaas.cloud.ibm.com"},
			want: "crn:v1:bluemix:public:power-iaas:dal12:a/12345:%s::",
		},
		{
			name: "Generate for Staging",
			args: args{"12345", "dal12", "dal.power-iaas.test.cloud.ibm.com"},
			want: "crn:v1:staging:public:power-iaas:dal12:a/12345:%s::",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crnBuilder(tt.args.useraccount, tt.args.regionZone, tt.args.host); got != tt.want {
				t.Errorf("crnBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_powerJSONConsumer(t *testing.T) {
	var data struct {
		Name string
		ID   int
	}
	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name:    "Simple powerJSONConsumer",
			json:    `{"name":"Somebody","id":1}`,
			wantErr: false,
		},
		{
			name:    "Null value",
			json:    "null",
			wantErr: false,
		},
		{
			name:    "Nil value",
			json:    "nil",
			wantErr: true,
		},
		{
			name:    "Empty value",
			json:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerJSONConsumer()
			err := got.Consume(bytes.NewBuffer([]byte(tt.json)), &data)
			if (err != nil) != tt.wantErr {
				t.Errorf("powerJSONConsumer().Consume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getPIClient(t *testing.T) {
	type args struct {
		debug  bool
		host   string
		scheme string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Scheme HTTTP",
			args: args{debug: true, host: "", scheme: "http"},
		},
		{
			name: "Scheme HTTTPS",
			args: args{debug: true, host: "", scheme: "https"},
		},
		{
			name: "Scheme Empty",
			args: args{debug: true, host: "", scheme: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPIClient(tt.args.debug, tt.args.host, tt.args.scheme); got == nil {
				t.Errorf("getPIClient() = %v", got)
			}
		})
	}
}

func Test_costructRegionFromZone(t *testing.T) {
	type args struct {
		zone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "DC Zone",
			args: args{
				zone: "dal12",
			},
			want: "dal",
		},
		{
			name: "AZ Zone",
			args: args{
				zone: "eu-de-1",
			},
			want: "eu-de",
		},
		{
			name: "Region Zone",
			args: args{
				zone: "us-south",
			},
			want: "us-south",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := costructRegionFromZone(tt.args.zone); got != tt.want {
				t.Errorf("costructRegionFromZone() = %v, want %v", got, tt.want)
			}
		})
	}
}
