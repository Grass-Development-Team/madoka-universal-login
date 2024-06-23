$env:GIN_MODE = "debug"

Set-Location ..

if ([System.Environment]::Is64BitOperatingSystem) {
    $env:GOARCH = "amd64"
}
else {
    $env:GOARCH = "386"
}

try {
    $location = "./dist/debug/"
    $output = "mul-windows-$($env:GOARCH).exe"
    go build -o $location$output
    Set-Location $location
    & ./$output
}
catch {
    Write-Warning "Error $_"
}