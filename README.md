# go_map_performance
Test golang map and ordered map performance
```
go test -bench=. -benchmem -run=Bench -benchtime=5s
```

```
goos: linux
goarch: amd64
pkg: test
BenchmarkNotSortMap-4                        590          10225141 ns/op         2320163 B/op      80000 allocs/op
BenchmarkDeleteNotSortMap-4             1000000000        0.000110 ns/op               0 B/op          0 allocs/op
BenchmarkSortMap-4                           450          13405330 ns/op         2320537 B/op      80001 allocs/op
BenchmarkDeleteSortMap-4                1000000000        0.000095 ns/op               0 B/op          0 allocs/op
BenchmarkWK8OrderedMap-4                     583          10299697 ns/op         2320170 B/op      80001 allocs/op
BenchmarkDeleteWK8OrderedMap-4          1000000000         0.00142 ns/op               0 B/op          0 allocs/op
BenchmarkElliOrderedMap-4                    511          11624683 ns/op         2800160 B/op      90000 allocs/op
BenchmarkDeleteElliOrderedMap-4             6656            886172 ns/op               0 B/op          0 allocs/op
BenchmarkIterOrderedMap-4                    550          11035224 ns/op         2320171 B/op      80001 allocs/op
BenchmarkDeleteIterOrderedMap-4         1000000000           0.146 ns/op               0 B/op          0 allocs/op
PASS
ok      test    43.383s
```
