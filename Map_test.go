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

var Lookup map[string]human
var notsort map[string]human
var wk8om *wk8.OrderedMap
var elliom *elli.OrderedMap
var iterom *iter.Items

var keymap [] string
var bitmap [] int

func init() {
	rand.Seed(time.Now().UnixNano())

	keymap = make([]string, TEST_SIZE)
	bitmap = make([]int, TEST_SIZE)
	for i:=0; i<TEST_SIZE; i++ {
		bitmap = append(bitmap, i)
	}
	rand.Shuffle(len(keymap),func(i, j int) {
		bitmap [i], bitmap[j] = bitmap [j], bitmap[i]
	})

	Lookup = make(map[string]human, TEST_SIZE)
	notsort = make(map[string]human, TEST_SIZE)

	wk8om = wk8.New()
	elliom = elli.NewOrderedMap()
	iterom = iter.New()

	for i:=0; i<TEST_SIZE; i++ {
		data := makeData()
		hash, _ := hashstructure.Hash(data, nil)
		key := strconv.FormatUint(hash, 10)
		Lookup[key] = data
		notsort[key] = data // for deletion test
		keymap = append(keymap, key)
		wk8om.Set(hash, data)
		elliom.Set(hash, data)
		iterom.Add(hash,data)
	}
}

func BenchmarkNotSortMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range Lookup {
			str := fmt.Sprintf("name is %s", v.Name)
			str += fmt.Sprintf("age is %d", v.Age)
			str += fmt.Sprintf("job is %s", v.Job)
		}
	}
}

func BenchmarkDeleteNotSortMap(b *testing.B) {
	//for n := 0; n < b.N; n++ {
		for k := range bitmap {
			delete(notsort, keymap[k])
		}
	//}
}

func BenchmarkSortMap(b *testing.B) {
	keys := make([]string, len(Lookup))
	i := 0
	for k := range Lookup {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for n := 0; n < b.N; n++ {
		for _, v := range keys {
			str := fmt.Sprintf("name is %s", Lookup[v].Name)
			str += fmt.Sprintf("age is %d", Lookup[v].Age)
			str += fmt.Sprintf("job is %s", Lookup[v].Job)
		}
	}
}

func BenchmarkDeleteSortMap(b *testing.B) {
	//for n := 0; n < b.N; n++ {
		for k := range bitmap {
			delete(Lookup, keymap[k])
		}
	//}
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

func BenchmarkDeleteWK8OrderedMap(b *testing.B) {
	for k := range bitmap {
		wk8om.Delete(keymap[k])
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

func BenchmarkDeleteElliOrderedMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for k := range bitmap {
			elliom.Delete(keymap[k])
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

func BenchmarkDeleteIterOrderedMap(b *testing.B) {
	for k := range bitmap {
		iterom.Del(keymap[k])
	}
}