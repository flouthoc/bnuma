package bnumamaps

import(
	"fmt"
	"github.com/cespare/xxhash"
	"hash/fnv"
)

type Node struct{

	key string
	value interface{}
	scattercount uint32
}

type Bnumamap struct{
	buckets []Node
	count uint32
}


func (m *Bnumamap) Size() int {
	return len(m.buckets)
}

func NewBnumamap(size int) *Bnumamap {
	e := Bnumamap{buckets: make([]Node, size, size)}
	return &e
}

func (m *Bnumamap) set(k string, v interface{}){

	idx := XxHash(key) % len(m.buckets)
	e := Node{key: k, value: v, scattercount: 0}
	for c := i; ; c = (c + 1) % m.Size() {
		if m.buckets[c].value == nil {

			m.buckets[c] = e
			m.count += 1
			return
		} else {
			if m.values[c].scattercount < e.scattercount {
				
				tmp := e
				e = Node{
					key:   m.buckets[c].key,
					value: m.buckets[c].value,
					scattercount:   m.buckets[c].scattercount,
				}
				
				m.buckets[c] = Node{
					key:   tmp.key,
					value: tmp.value,
					scattercount:   tmp.scattercount,
				}

			}
		}
		
		e.scattercount += 1
	}

}


func (m Bnumamap) Get(k string) interface{} {
	i := XxHash(k) % m.Size()
	e := m.buckets[i]
	for e.key != k {
		i = (i + 1) % m.Size()
		e = m.buckets[i]
		if i == XxHash(k)%m.Size() {
			// Not Found
			return nil
		}
	}
	return e.value
}

func FnvHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func XxHash(s string) uint64{
	return xxhash.Sum64([]byte("hello"))
}

func main(){

	fmt.Println(xxhash.Sum64([]byte("hello")))
	FnvHash("hello")
}