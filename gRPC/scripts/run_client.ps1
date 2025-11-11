# Run the gRPC client
Push-Location (Split-Path -Parent $MyInvocation.MyCommand.Path)\..\

Write-Host "Running RouteGuide gRPC client..." -ForegroundColor Cyan
go run ./client

Pop-Location
