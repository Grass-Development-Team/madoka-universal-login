Set-Location ..

$platform = @(
    "windows",
    "linux"
)

$arch = @(
    "amd64",
    "386",
    "arm"
)


foreach ($p in $platform) {
    foreach ($a in $arch) {
        try {
            go build -o dist/release/mul-$p-$a$($p -eq "windows" ? ".exe" : '')
        }
        catch {
            Write-Warning "Error: $_"
        }
    }
}