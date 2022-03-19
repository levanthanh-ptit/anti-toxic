build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build block/anti-mun-block.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build open/anti-mun-open.go