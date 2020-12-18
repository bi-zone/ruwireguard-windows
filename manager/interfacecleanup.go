/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2020 WireGuard LLC. All Rights Reserved.
 * Copyright (C) 2020 BI.ZONE LLC. All Rights Reserved.
 */

package manager

import (
	"log"

	"github.com/bi-zone/ruwireguard-go/tun/wintun"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"

	"github.com/bi-zone/ruwireguard-go/tun"
	"github.com/bi-zone/ruwireguard-windows/services"
)

func cleanupStaleWintunInterfaces() {
	m, err := mgr.Connect()
	if err != nil {
		return
	}
	defer m.Disconnect()

	tun.WintunPool.DeleteMatchingAdapters(func(wintun *wintun.Adapter) bool {
		interfaceName, err := wintun.Name()
		if err != nil {
			log.Printf("Removing Wintun interface because determining interface name failed: %v", err)
			return true
		}
		serviceName, err := services.ServiceNameOfTunnel(interfaceName)
		if err != nil {
			log.Printf("Removing Wintun interface ‘%s’ because determining tunnel service name failed: %v", interfaceName, err)
			return true
		}
		service, err := m.OpenService(serviceName)
		if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
			log.Printf("Removing Wintun interface ‘%s’ because no service for it exists", interfaceName)
			return true
		} else if err != nil {
			return false
		}
		defer service.Close()
		status, err := service.Query()
		if err != nil {
			return false
		}
		if status.State == svc.Stopped {
			log.Printf("Removing Wintun interface ‘%s’ because its service is stopped", interfaceName)
			return true
		}
		return false
	}, false)
}
