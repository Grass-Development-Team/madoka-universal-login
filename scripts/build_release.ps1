Set-Location ..

$platform = @(
    "windows",
    "linux"
)

$arch = @(
    "amd64",
    "386"
)


foreach ($p in $platform) {
    foreach ($a in $arch) {
        try {
            $env:GOOS = $p
            $env:GOARCH = $a
            $env:CGO_ENABLE = true
            go build -o dist/release/mul-$p-$a$($p -eq "windows" ? ".exe" : '')
        }
        catch {
            Write-Warning "Error: $_"
        }
    }
}