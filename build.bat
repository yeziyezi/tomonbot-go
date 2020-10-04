SET TOMON_GO_SRC_MAIN=.\src\one\yezii\tomon\main.go
SET TOMON_GO_BUILD_PATH=.\build
SET GOOS = windows
go build -o %TOMON_GO_BUILD_PATH%\tomon-go-win.exe %TOMON_GO_SRC_MAIN%
SET GOOS = linux
go build -o %TOMON_GO_BUILD_PATH%\tomon-go-linux %TOMON_GO_SRC_MAIN%
