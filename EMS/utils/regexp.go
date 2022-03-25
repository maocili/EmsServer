package utils

import (
	"regexp"
)

//str = "i\n网: 6, Ip=10.0.2.15  子网=255.255.255.0  网关=10.0.2.2\n净值: 6, Ip=fe80::31b3:f0d:e77d:b479\n"
func RegNetDivceID(res string) string {

	reg := regexp.MustCompile(`值: (\d+), Ip`)
	if reg == nil {
		return ""
	}
	divceId := reg.FindAllStringSubmatch(res, -1)
	return divceId[0][1]
}
