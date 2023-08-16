package helpers

import (
	"os"
	"testing"
)

func TestEnvFallBack(t *testing.T) {
	os.Setenv("MY_ENV_VAR", "hello")
	type args struct {
		envs         []string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Value for Known Environment Variable",
			args: args{[]string{"MY_ENV_VAR"}, "default"},
			want: "hello",
		},
		{
			name: "Get Value for Unknown Environment Variable",
			args: args{[]string{"NO_ENV_VAR"}, "default"},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnvFallBack(tt.args.envs, tt.args.defaultValue); got != tt.want {
				t.Errorf("EnvFallBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPowerEndPoint(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Get No Endpoint from Environment Variable",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowerEndPoint(); got != tt.want {
				t.Errorf("GetPowerEndPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPowerEndPointFromEnv(t *testing.T) {
	os.Setenv("IBMCLOUD_PI_API_ENDPOINT", "dal.power-iaas.cloud.ibm.com")
	type args struct {
		region string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Get Endpoint from Environment Variable",
			args: args{"dal"},
			want: "dal.power-iaas.cloud.ibm.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowerEndPoint(); got != tt.want {
				t.Errorf("GetPowerEndPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
