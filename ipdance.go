package ipdance

func GetIp() (string, error) {
	IPInfo, err := getIpAddressInfo()
	if err != nil {
		return "", err
	}
	ip := IPInfo.Ip
	return *ip, nil
}

func GetHttpVersion() (string, error) {
	IPInfo, err := getIpAddressInfo()
	if err != nil {
		return "", err
	}
	ip := IPInfo.HttpProtocol
	return *ip, nil
}

func GetAsn() (string, error) {
	IPInfo, err := getIpAddressInfo()
	if err != nil {
		return "", err
	}
	ip := IPInfo.Asn
	return *ip, nil
}

func GetLocation() (string, error) {
	IPInfo, err := getIpAddressInfo()
	if err != nil {
		return "", err
	}
	ip := IPInfo.Location
	return *ip, nil
}

func GetRegionCode() (string, error) {
	IPInfo, err := getIpAddressInfo()
	if err != nil {
		return "", err
	}
	ip := IPInfo.RegionCode
	return *ip, nil
}
