/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2020 WireGuard LLC. All Rights Reserved.
 * Copyright (C) 2020 BI.ZONE LLC. All Rights Reserved.
 */

package services

import (
	"errors"

	"github.com/bi-zone/ruwireguard-windows/conf"
)

func ServiceNameOfTunnel(tunnelName string) (string, error) {
	if !conf.TunnelNameIsValid(tunnelName) {
		return "", errors.New("Tunnel name is not valid")
	}
	return "WireGuardTunnel$" + tunnelName, nil
}

func PipePathOfTunnel(tunnelName string) (string, error) {
	if !conf.TunnelNameIsValid(tunnelName) {
		return "", errors.New("Tunnel name is not valid")
	}
	return `\\.\pipe\ProtectedPrefix\Administrators\WireGuard\` + tunnelName, nil
}
