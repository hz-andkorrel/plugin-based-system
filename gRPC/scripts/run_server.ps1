# Run the gRPC server
Push-Location (Split-Path -Parent $MyInvocation.MyCommand.Path)\..\

Write-Host "Starting RouteGuide gRPC server on :50051..." -ForegroundColor Cyan
go run ./server

Pop-Location
