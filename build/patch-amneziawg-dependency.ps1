$ErrorActionPreference = 'Stop'

$module = 'github.com/amnezia-vpn/amneziawg-windows/v3'
$moduleDir = (& go list -m -f '{{.Dir}}' $module).Trim()
if (-not $moduleDir -or -not (Test-Path $moduleDir)) {
    throw "Unable to locate dependency module: $module"
}

$namesFile = Join-Path $moduleDir 'services\names.go'
if (-not (Test-Path $namesFile)) {
    throw "Dependency service names file not found: $namesFile"
}

$content = Get-Content $namesFile -Raw
$content = $content.Replace('AmneziaWGTunnel$" + tunnelName', 'MyAmneziaWGTunnel$" + tunnelName')
$content = $content.Replace('AmneziaWG\` + tunnelName', 'MyAmneziaWG\` + tunnelName')

if ($content -notmatch 'MyAmneziaWGTunnel\$') {
    throw 'Service name patch was not applied.'
}
if ($content -notmatch 'MyAmneziaWG\\') {
    throw 'Named pipe path patch was not applied.'
}

Set-Content -Path $namesFile -Value $content -NoNewline
Write-Host "Patched AmneziaWG Windows dependency: $namesFile"
