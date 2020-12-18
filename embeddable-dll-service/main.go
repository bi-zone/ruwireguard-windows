/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2020 WireGuard LLC. All Rights Reserved.
 * Copyright (C) 2020 BI.ZONE LLC. All Rights Reserved.
 */

package main

import (
	"C"

	"golang.org/x/crypto/curve25519"
	"golang.org/x/sys/windows"

	"github.com/bi-zone/ruwireguard-windows/conf"
	"github.com/bi-zone/ruwireguard-windows/tunnel"

	"crypto/rand"
	"log"
	"path/filepath"
	"unsafe"
)

//export WireGuardTunnelService
func WireGuardTunnelService(confFile16 *uint16) bool {
	confFile := windows.UTF16PtrToString(confFile16)
	conf.PresetRootDirectory(filepath.Dir(confFile))
	tunnel.UseFixedGUIDInsteadOfDeterministic = true
	err := tunnel.Run(confFile)
	if err != nil {
		log.Printf("Service run error: %v", err)
	}
	return err == nil
}

//export WireGuardGenerateKeypair
func WireGuardGenerateKeypair(publicKey *byte, privateKey *byte) {
	publicKeyArray := (*[32]byte)(unsafe.Pointer(publicKey))
	privateKeyArray := (*[32]byte)(unsafe.Pointer(privateKey))
	n, err := rand.Read(privateKeyArray[:])
	if err != nil || n != len(privateKeyArray) {
		panic("Unable to generate random bytes")
	}
	privateKeyArray[0] &= 248
	privateKeyArray[31] = (privateKeyArray[31] & 127) | 64

	curve25519.ScalarBaseMult(publicKeyArray, privateKeyArray)
}

func main() {}
