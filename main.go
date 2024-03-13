package main

import (
	"fmt"
	"log"
	"github.com/k-stz/cbuffer/circular"
)

func main() {
	cbuffer := circular.NewBuffer(2)
	cbuffer.WriteByte(1)
	cbuffer.Overwrite(2)
	b, err := cbuffer.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("read:", b)
	cbuffer.Reset()
	fmt.Println(cbuffer)
	// // b, err := cbuffer.ReadByte()
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// cbuffer.Reset()
	// fmt.Println("last cbuffer:", cbuffer)
	// var b byte
	// cbuffer := circular.NewBuffer(5)
	// for _, v := range []byte{1, 2, 3, 4, 5} {
	// 	err := cbuffer.WriteByte(v)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// cbuffer.WriteByte(22)
	// b, _ = cbuffer.ReadByte()
	// fmt.Println("read byte:", b)
	// b, _ = cbuffer.ReadByte()
	// fmt.Println("read byte:", b)
	// fmt.Println(cbuffer)
}
