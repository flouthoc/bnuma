package bnumamaps

import(
	"fmt"
	"github.com/cespare/xxhash"
)


func main(){

	fmt.Println(xxhash.Sum64([]byte("hello")))
}