# provider-go
###### For CDC exercise

How to run API: 
```
go run provider.go
```

How to build and run in Docker:
```
docker build . -t provider-go
docker run -p 8080:8080 provider-go
```

To check:
`open in browser http://localhost:8080/`

How to run contract test: 
```
go test -v -run TestProvider
```