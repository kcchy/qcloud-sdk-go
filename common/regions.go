package common

type Region string

const (
	Beijing   = Region("bj")
	Guangzhou = Region("gz")
	Shanghai  = Region("sh")
	HongKong  = Region("hk")

	NorthAmerica = Region("ca")

	Singapore = Region("sg")

	ShanghaiFinance = Region("shjr")
	ShenzhenFinance = Region("szjr")

	GuangzhouOpen = Region("gzopen")
)

var ValidRegions = []Region{
	Beijing, Guangzhou, Shanghai, HongKong,
	NorthAmerica, Singapore,
	ShanghaiFinance, ShenzhenFinance,
	GuangzhouOpen,
}
