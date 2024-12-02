package ipaddress

import (
	"atom-go/internal/common/curl"
	"atom-go/internal/common/utils"
	"encoding/json"
	"net"
)

// ip地址
type IpAddress struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro"`
	ProCode    string `json:"proCode"`
	City       string `json:"city"`
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
}

// 根据ip获取地址
func GetAddress(ip string) *IpAddress {

	var ipAddress IpAddress

	var internalIp = "(((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))\\.){3}((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))$"

	if netIp := net.ParseIP(ip); netIp == nil || netIp.IsLoopback() {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	if utils.CheckRegex(internalIp, ip) {
		ipAddress.Ip = ip
		ipAddress.Addr = "内网地址"
		return &ipAddress
	}

	body, err := curl.DefaultClient().Send(&curl.RequestParam{
		Url: "http://whois.pconline.com.cn/ipJson.jsp",
		Query: map[string]interface{}{
			"ip":   ip,
			"json": true,
		},
	})

	if err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	if err := json.Unmarshal([]byte(body), &ipAddress); err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	return &ipAddress
}
