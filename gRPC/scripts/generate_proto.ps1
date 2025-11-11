# PowerShell script to generate Go gRPC code from proto files
$here = Split-Path -Parent $MyInvocation.MyCommand.Path
Push-Location $here\..\

Write-Host "Generating Go code from proto..." -ForegroundColor Cyan
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/route_guide.proto

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Code generation successful!" -ForegroundColor Green
} else {
    Write-Host "✗ Code generation failed!" -ForegroundColor Red
}

Pop-Location
