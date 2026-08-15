/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2026 MyAmneziaWG contributors. All Rights Reserved.
 */

package manager

import (
	"strings"

	"github.com/amnezia-vpn/amneziawg-windows/v3/services"
)

// standaloneServiceNameOfTunnel keeps tunnel Windows service names separate
// from the upstream AmneziaWG installation while preserving its validation.
func standaloneServiceNameOfTunnel(tunnelName string) (string, error) {
	name, err := services.ServiceNameOfTunnel(tunnelName)
	if err != nil {
		return "", err
	}
	return strings.Replace(name, "AmneziaWGTunnel$", "MyAmneziaWGTunnel$", 1), nil
}

// standalonePipePathOfTunnel keeps the tunnel UAPI named pipe separate from
// the upstream installation. The upstream helper remains responsible for
// validating the tunnel name and constructing the canonical pipe path.
func standalonePipePathOfTunnel(tunnelName string) (string, error) {
	path, err := services.PipePathOfTunnel(tunnelName)
	if err != nil {
		return "", err
	}
	return strings.Replace(path, `\\.\pipe\AmneziaWG\`, `\\.\pipe\MyAmneziaWG\`, 1), nil
}
