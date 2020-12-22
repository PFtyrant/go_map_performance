package main

import (
	"fmt"
	elli "github.com/elliotchance/orderedmap"
	iter "github.com/mantyr/iterator"
	"github.com/mitchellh/hashstructure"
	wk8 "github.com/wk8/go-ordered-map"
	"strconv"
	"math/rand"
	)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
type human struct {
	Name string
	Age int
	Job string
}
func makeData() human {
	ret := human{
		Name: RandStringRunes(10),
		Age: rand.Int(),
		Job: RandStringRunes(5),
	}
	return ret
}

const SIZE = 10

func main() {

	lookup := make(map[string]human, SIZE)
	stringmap := make([]string, 0)
	wk8om := wk8.New()
	elli := elli.NewOrderedMap()
	iterom := iter.New()
	for i := 0; i < SIZE; i += 1 {
		data := makeData()
		hash, _ := hashstructure.Hash(data, nil)
		key := strconv.FormatUint(hash, 10)
		stringmap = append(stringmap, key)
		lookup[key] = data
		wk8om.Set(key, data)
		elli.Set(key, data)
		iterom.Add(key, data)
	}
	fmt.Println(lookup)
	keymap := make([]int, 0)
	for i:=0;i<SIZE;i++ {
		keymap = append(keymap, i)
	}
	rand.Shuffle(len(keymap),func(i, j int) {
		keymap [i], keymap[j] = keymap [j], keymap[i]
	})
	fmt.Println(wk8om)
	for key := range keymap {
		wk8om.Delete(stringmap[key])
		elli.Delete(stringmap[key])
		iterom.Del(stringmap[key])
	}
	fmt.Println(wk8om)
	fmt.Println(elli)
	fmt.Println(iterom)
}