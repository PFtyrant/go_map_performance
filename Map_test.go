package main

import (
	"fmt"
	elli "github.com/elliotchance/orderedmap"
	iter "github.com/mantyr/iterator"
	wk8 "github.com/wk8/go-ordered-map"
	"github.com/mitchellh/hashstructure"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"sort"
)
const TEST_SIZE = 10000

var sortedom map[string]human
var notsort map[string]human
var wk8om *wk8.OrderedMap
var elliom *elli.OrderedMap
var iterom *iter.Items

var keymap [] string
var bitmap [] int

func init() {
	rand.Seed(time.Now().UnixNano())

	// bitmap : shuffled number array
	// keymap[bitmap[i]] = key
	keymap = make([]string, 0)
	bitmap = make([]int, 0)
	for i:=0; i<TEST_SIZE; i++ {
		bitmap  = append(bitmap, i)
	}
	rand.Shuffle(len(keymap),func(i, j int) {
		bitmap [i], bitmap[j] = bitmap [j], bitmap[i]
	})

	// define map
	notsort = make(map[string]human, TEST_SIZE)
	sortedom = make(map[string]human, TEST_SIZE)
	wk8om = wk8.New()
	elliom = elli.NewOrderedMap()
	iterom = iter.New()

	for i:=0; i<TEST_SIZE; i++ {
		// make the data
		data := makeData()
		hash, _ := hashstructure.Hash(data, nil)
		key := strconv.FormatUint(hash, 10)

		keymap = append(keymap, key)
		// insert the data into maps
		sortedom[key] = data
		notsort[key] = data
		wk8om.Set(key, data)
		elliom.Set(key, data)
		iterom.Add(key,data)
	}
}

func BenchmarkNotSortMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range notsort {
			str := fmt.Sprintf("name is %s", v.Name)
			str += fmt.Sprintf("age is %d", v.Age)
			str += fmt.Sprintf("job is %s", v.Job)
		}
	}
}

func BenchmarkSortMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		keys := make([]string, len(sortedom))
		i := 0
		for k := range sortedom {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, v := range keys {
			str := fmt.Sprintf("name is %s", sortedom[v].Name)
			str += fmt.Sprintf("age is %d", sortedom[v].Age)
			str += fmt.Sprintf("job is %s", sortedom[v].Job)
		}
	}
}

// wk8/go-ordered-map
func BenchmarkWK8OrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for pair := wk8om.Oldest(); pair != nil; pair = pair.Next() {
			str := fmt.Sprintf("name is %s", pair.Value.(human).Name)
			str += fmt.Sprintf("age is %d", pair.Value.(human).Age)
			str += fmt.Sprintf("job is %s", pair.Value.(human).Job)
		}
	}
}

// elliotchance/orderedmap
func BenchmarkElliOrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for el := elliom.Front(); el != nil; el = el.Next() {
			str := fmt.Sprintf("name is %s", el.Value.(human).Name)
			str += fmt.Sprintf("age is %d", el.Value.(human).Age)
			str += fmt.Sprintf("job is %s", el.Value.(human).Job)
		}
	}
}

// mantyr/iterator
func BenchmarkIterOrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, value := range iterom.Items {
			str := fmt.Sprintf("name is %s", value.(human).Name)
			str += fmt.Sprintf("age is %d", value.(human).Age)
			str += fmt.Sprintf("job is %s", value.(human).Job)
		}
	}
}

func BenchmarkDeleteNotSortMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copymap := make(map[string]human, TEST_SIZE)
		for k, v := range notsort {
			copymap[k] = v
		}
		b.StartTimer()
		for k := range bitmap {
			delete(copymap, keymap[k])
		}
	}
}

func BenchmarkDeleteSortMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copymap := make(map[string]human, TEST_SIZE)
		for k, v := range sortedom {
			copymap[k] = v
		}
		b.StartTimer()
		for k := range bitmap {
			delete(copymap, keymap[k])
		}
	}
}

func BenchmarkDeleteWK8OrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copymap := wk8.New()
		for pair := wk8om.Oldest(); pair != nil; pair = pair.Next() {
			copymap.Set(pair.Key, pair.Value)
		}
		b.StartTimer()
		for k := range bitmap {
			copymap.Delete(keymap[k])
		}
	}
}

func BenchmarkDeleteElliOrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copymap := elli.NewOrderedMap()
		for el := elliom.Front(); el != nil; el = el.Next() {
			copymap.Set(el.Key, el.Value)
		}
		b.StartTimer()
		for k := range bitmap {
			copymap.Delete(keymap[k])
		}
	}
}
func BenchmarkDeleteIterOrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copymap := iter.New()
		for key, value := range iterom.Items {
			copymap.Add(key, value)
		}
		b.StartTimer()
		for k := range bitmap {
			copymap.Del(keymap[k])
		}
	}
}