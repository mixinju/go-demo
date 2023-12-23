package model

type Education struct {
	Id             uint64
	Major          string // 专业
	Department     string // 单位
	City           string // 行政区划 省/市
	FirstCityCode  int    // 一级城市编码 eg: 河南省
	SecondCityCode int    // 二级城市编码 eg: 许昌市
	ThirdCityCode  int    // 三级城市编码 eg: 襄城县
	// https://www.mca.gov.cn/mzsj/xzqh/2022/202201xzqh.html
}
