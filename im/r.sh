swag init
cd front
npm run build
cd ..
clear
go build .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build .
./im
