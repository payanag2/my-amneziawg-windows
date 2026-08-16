/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2026 MyAmneziaWG. All Rights Reserved.
 */

package manager

import (
	"errors"

	"github.com/amnezia-vpn/amneziawg-windows/v3/conf"
)

func standaloneServiceNameOfTunnel(tunnelName string) (string, error) {
	if !conf.TunnelNameIsValid(tunnelName) {
		return "", errors.New("Tunnel name is not valid")
	}
	return "MyAmneziaWGTunnel$" + tunnelName, nil
}

func standalonePipePathOfTunnel(tunnelName string) (string, error) {
	if !conf.TunnelNameIsValid(tunnelName) {
		return "", errors.New("Tunnel name is not valid")
	}
	return `\\.\pipe\ProtectedPrefix\Administrators\MyAmneziaWG\` + tunnelName, nil
}
