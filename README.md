# go_map_performance

Since the native map of the GO is not an order stable structure, we must sort a map or use an ordered map structure to resolve the problem that we faced.
So, I decide to test the performance of the varying map.

Test the performance of the map structure with Go native map, ordered map(made by users on Github) 

Test five map structure:
```
1. map
2. map(sorted by key)
3. Ordered map("github.com/wk8/go-ordered-map")
4. Ordered map("github.com/elliotchance/orderedmap")
5. Ordered map("github.com/mantyr/iterator")
```

TEST data:
```
type human struct {
	Name string
	Age int
	Job string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func makeData() human {
	ret := human{
		Name: RandStringRunes(10),
		Age: rand.Int(),
		Job: RandStringRunes(5),
	}
	return ret
}
```

The type of the key is string type, and the type of the value is the human struct.

Made a key by the hashed human struct what we make with makedata() function.


## Test

The size value of the data structure is 10000.
```bigquery
TESTSIZE=10000
```

Ran the test with flag -benchtime=5s
```
go test -bench=. -benchmem -run=Bench -benchtime=5s
```
The result as below:
```
goos: linux
goarch: amd64
pkg: test
BenchmarkNotSortMap-4                        590           9969112 ns/op         2320169 B/op      80001 allocs/op
BenchmarkSortMap-4                           363          16478182 ns/op         2484053 B/op      80003 allocs/op
BenchmarkWK8OrderedMap-4                     572          10223365 ns/op         2320161 B/op      80000 allocs/op
BenchmarkElliOrderedMap-4                    535          11215278 ns/op         2800152 B/op      90000 allocs/op
BenchmarkIterOrderedMap-4                    555          10834914 ns/op         2320175 B/op      80001 allocs/op
BenchmarkDeleteNotSortMap-4                 5874           1037779 ns/op               0 B/op          0 allocs/op
BenchmarkDeleteSortMap-4                    5796           1023836 ns/op               0 B/op          0 allocs/op
BenchmarkDeleteWK8OrderedMap-4              2310           2501172 ns/op               0 B/op          0 allocs/op
BenchmarkDeleteElliOrderedMap-4             2792           2208096 ns/op               0 B/op          0 allocs/op
BenchmarkDeleteIterOrderedMap-4               10         507257492 ns/op               0 B/op          0 allocs/op
PASS
ok      test    131.101s
```
