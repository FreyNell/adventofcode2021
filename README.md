```go test -cover ./...```

# HTML REPORT
```go test -cover -coverprofile=c.out ./...
go tool cover -html=c.out -o coverage.html```

# FOR BENCHMARKS
```go test -bench=. -run=x -benchmem -memprofile mem.prof -cpuprofile cpu.prof -benchtime=5s```
```go tool pprof cpu.prof ```
1. here we can run ```top 20``` and it is going to show us the top 20 lines using the most cpu.
2. ```web``` to show it on web mode.
3. ```list package/class/method``` shows the consuming of that method.
4. ```go tool pprof``` --svg cpu.prof to show svg output