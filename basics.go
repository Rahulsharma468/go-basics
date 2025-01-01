package main

import (
    "fmt"
    "maps"
)

func main() {
    fmt.Println("Hello, World!")

    for i:= range 3 {
        fmt.Println("HEY " , i);
    }

    for {
        fmt.Println("loop");
        break
    }

    // Arrays

    var a [5]int
    a[4] = 10
    b := [5]int{1,2,3,4,5}
    fmt.Println("EMOT :: " , a);
    fmt.Println("EMOT :: " , b);

    // Maps

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 17

    fmt.Println("Map " , m , m["k1"]);

    delete(m , "k2")
    fmt.Println("Map " , m );

    clear(m)

    fmt.Println("Map " , m );


    mn := map[string]int{"foo" : 1 , "bar" : 2}

    fmt.Println(mn)

    if maps.Equal(m , mn) {fmt.Println("EQUAL MAPS")} else {fmt.Println("UNQUELA MAPS")
}}
