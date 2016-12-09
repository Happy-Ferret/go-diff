# go-diff [WIP] [![Build Status](https://travis-ci.org/magicshui/go-diff.svg?branch=master)](https://travis-ci.org/magicshui/go-diff)

using [structs](github.com/fatih/structs)

```
package main

import (
    "github.com/magicshui/go-diff"
    "log"
)
func main(){
    type M map[string]interface{}
    var a = M{
        "int_n": 123, "int_c": 123,
        "str_n": "hello", "str_c": "hello",

        "array_n": []string{"1231"}, "array_c": []string{"12313"},
    }
    var b = M{
        "int_n": 123, "int_c": 1231,
        "str_n": "hello", "str_c": "hello world",

        "array_n": []string{"1231"}, "array_c": []interface{}{"12313", 123},
    }

    for k, v := range DiffMaps(a, b).ChangedItems() {
        log.Printf("Key: %s , Value: %t", k, v.Changed)
    }
}

```