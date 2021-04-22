package clientutils

func Zonemaps(region string) string {

	region_to_zone := make(map[string]string)
	region_to_zone["us-south"] = "us-south"
	region_to_zone["us-east"] = "us-east"
	region_to_zone["eu-de-1"] = "eu-de"
	region_to_zone["eu-de-2"] = "eu-de"
	region_to_zone["tok04"] = "tok"
	region_to_zone["tok02"] = "tok"
	region_to_zone["lon06"] = "lon"
	region_to_zone["lon04"] = "lon"
	region_to_zone["tor01"] = "tor"
	region_to_zone["mon01"] = "mon"
	region_to_zone["dal12"] = "us-south"
	region_to_zone["syd04"] = "au-syd"
	return region_to_zone[region]
}
