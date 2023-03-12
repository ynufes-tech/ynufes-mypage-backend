go build -o ynufes-mypage-backend.exe ./svc/cmd/dev/main.go
if ($LASTEXITCODE -eq 0)
{
    .\ynufes-mypage-backend.exe
}
