$ErrorActionPreference = 'Stop'

$files = @(
    'manager/install.go',
    'manager/ipc_server.go',
    'manager/ipc_uapi.go'
)

foreach ($relativePath in $files) {
    $path = Join-Path $PSScriptRoot $relativePath
    $text = [System.IO.File]::ReadAllText($path)
    $text = $text.Replace('services.ServiceNameOfTunnel(', 'standaloneServiceNameOfTunnel(')
    $text = $text.Replace('services.PipePathOfTunnel(', 'standalonePipePathOfTunnel(')
    $text = [regex]::Replace($text, '(?m)^[\t ]*"github\.com/amnezia-vpn/amneziawg-windows/v3/services"\r?\n', '')
    [System.IO.File]::WriteAllText($path, $text, [System.Text.UTF8Encoding]::new($false))
}
