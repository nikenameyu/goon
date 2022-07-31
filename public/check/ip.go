package check

import (
	"goon3/public"
	"net"
	"strconv"
	"strings"
)

/*
整理端口:
eg:
10.0.0.0/24
10.0.0.0-10.0.0.255
10.0.0.0
10.0.0-123.120
10.0.0.0-123
*/
// Check IP format
func GetIp(ips string) []string{
	//fmt.Println(ips)
	ipsResult := []string{}
	// if ips Contains "-"
	if find := strings.Count(ips, "-"); find==1 {
		// if ips likes 10.0.0.0-10.0.0.255
		if find2 := strings.Count(ips, "."); find2==6 {
			ips_Arr := strings.SplitN(ips,"-",-1)
			ips1_Arr, ips2_Arr := strings.SplitN(ips_Arr[0],".",-1),strings.SplitN(ips_Arr[1],".",-1)
			ipsResult := IpGet(ips1_Arr,ips2_Arr)
			return ipsResult
		} else if find2==3 {
			// if ips likes 10.0.0.0-255 or 10.10.10-11.23
			ips_Arr := strings.SplitN(ips,"-",-1)
			ips1_Arr, ips2_Arr :=  strings.SplitN(ips_Arr[0],".",-1),strings.SplitN(ips_Arr[1],".",-1)
			if len(ips1_Arr[0]) != 0 && len(ips2_Arr[0]) != 0{
				ips_Arr_Temp := make([]string,len(ips1_Arr)-1,4)
				copy(ips_Arr_Temp,ips1_Arr[:len(ips1_Arr)-1])
				for i:=0;i<len(ips2_Arr)-1;i++{
					ips1_Arr = append(ips1_Arr,"0")
				}
				ips2_Arr = append(ips_Arr_Temp,ips2_Arr...)
				ipsResult := IpGet(ips1_Arr,ips2_Arr)
				return ipsResult
			} else {
				public.Error.Printf("ip:%s is error!\n",ips)
			}
		} else {
			public.Error.Printf("ip:%s is error!\n",ips)
		}
	} else if find := strings.Count(ips, "/"); find==1 {
		// if ips likes 10.0.0.0/24
		addr := getSubNet(ips)
		if addr != "err"{
			ips_Arr := strings.SplitN(addr,"-",-1)
			ips1_Arr, ips2_Arr := strings.SplitN(ips_Arr[0],".",-1),strings.SplitN(ips_Arr[1],".",-1)
			ipsResult := IpGet(ips1_Arr,ips2_Arr)
			return ipsResult
		}
	} else if is_Ipv4 := net.ParseIP(ips);is_Ipv4!=nil{
		ipsResult = append(ipsResult,ips)
		return ipsResult
	} else {
		public.Error.Printf("ip:%s is error!\n",ips)
	}
	return ipsResult
}

// get ip
func IpGet(ips1,ips2 []string) []string{
	ipsResult := []string{}
	ips11, _ := strconv.Atoi(ips1[0])
	ips12, _ := strconv.Atoi(ips1[1])
	ips13, _ := strconv.Atoi(ips1[2])
	ips14, _ := strconv.Atoi(ips1[3])
	ips21, _ := strconv.Atoi(ips2[0])
	ips22, _ := strconv.Atoi(ips2[1])
	ips23, _ := strconv.Atoi(ips2[2])
	ips24, _ := strconv.Atoi(ips2[3])
	for a := ips11; a <= ips21; a++{
		for b := ips12;b <= ips22; b++{
			for c:= ips13; c <= ips23; c++{
				for d := ips14; d <= ips24; d++{
					if d>0 && d<256{
						// int to string
						ip := strconv.Itoa(a)+"."+strconv.Itoa(b)+"."+strconv.Itoa(c)+"."+strconv.Itoa(d)
						ipsResult = append(ipsResult,ip)
					}
				}
			}
		}
	}
	return ipsResult
}

// check likes 10.10.10.0/24
func getSubNet(ips string) string{
	_, ipNet, err := net.ParseCIDR(ips)
	if err != nil {
		public.Error.Printf("ip:%s is error!\n",ips)
		return "err"
	} else {
		val := make([]byte, len(ipNet.Mask))
		copy(val, ipNet.Mask)
		var ip_Mask  []string
		for _, i := range val[:] {
			ip_Mask = append(ip_Mask, strconv.Itoa(int(i)))
		}
		ip_Net := strings.SplitN(ipNet.IP.String(),".",-1)
		ip_Net1, _ := strconv.Atoi(ip_Net[0])
		ip_Net2, _ := strconv.Atoi(ip_Net[1])
		ip_Net3, _ := strconv.Atoi(ip_Net[2])
		ip_Net4, _ := strconv.Atoi(ip_Net[3])
		ip_Mask1, _ := strconv.Atoi(ip_Mask[0])
		ip_Mask2, _ := strconv.Atoi(ip_Mask[1])
		ip_Mask3, _ := strconv.Atoi(ip_Mask[2])
		ip_Mask4, _ := strconv.Atoi(ip_Mask[3])
		a, b, c, d := ip_Net1 & ip_Mask1, ip_Net2 & ip_Mask2, ip_Net3 & ip_Mask3, ip_Net4 & ip_Mask4
		/* net_addr */
		net_addr := strconv.Itoa(a)+"."+strconv.Itoa(b)+"."+strconv.Itoa(c)+"."+strconv.Itoa(d+1)
		/* bd_addr */
		a1, b1, c1, d1 := a^(ip_Mask1^255), b^(ip_Mask2^255), c^(ip_Mask3^255), d^(ip_Mask4^255)
		gb_addr := strconv.Itoa(a1)+"."+strconv.Itoa(b1)+"."+strconv.Itoa(c1)+"."+strconv.Itoa(d1-1)
		return net_addr+"-"+gb_addr
	}
}