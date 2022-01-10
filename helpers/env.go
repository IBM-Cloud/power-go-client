package helpers

import "os"

//EnvFallBack ...
func EnvFallBack(envs []string, defaultValue string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return defaultValue
}

// GetPowerEndPoint
func GetPowerEndPoint(region string) string {
	defaultURL := region + ".power-iaas.cloud.ibm.com"
	return EnvFallBack([]string{"IBMCLOUD_POWER_API_ENDPOINT"}, defaultURL)
}
