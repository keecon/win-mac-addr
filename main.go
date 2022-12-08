// Copyright 2022 iwaltgen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"

	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
)

func main() {
	// https://learn.microsoft.com/en-us/windows/win32/api/iphlpapi/nf-iphlpapi-getadaptersaddresses
	addrs, err := winipcfg.GetAdaptersAddresses(windows.AF_INET, winipcfg.GAAFlagSkipAll|winipcfg.GAAFlagIncludeGateways)
	if err != nil {
		panic("GetAdaptersAddresses: " + err.Error())
	}

	for _, v := range addrs {
		if v.FirstGatewayAddress == nil || v.PhysicalAddress() == nil {
			continue
		}

		fmt.Printf("Name: %s\n", v.FriendlyName())
		fmt.Printf("AdapterName: %s\n", v.AdapterName())
		fmt.Printf("Description: %s\n", v.Description())
		fmt.Printf("PhysicalAddress: %v\n", net.HardwareAddr(v.PhysicalAddress()))
	}
}
