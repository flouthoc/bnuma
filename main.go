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


func (m *Bnumamap) Size() uint64 {
	return uint64(len(m.buckets))
}

func NewBnumamap(size int) *Bnumamap {
	e := Bnumamap{buckets: make([]Node, size, size)}
	return &e
}

func (m *Bnumamap) set(k string, v interface{}){

	idx := XxHash(k) % m.Size()
	e := Node{key: k, value: v, scattercount: 0}
	for c := idx; ; c = (c + 1) % m.Size() {
		if m.buckets[c].value == nil {

			m.buckets[c] = e
			m.count += 1
			return
		} else {

			if m.buckets[c].scattercount < e.scattercount {
				
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
	idx := XxHash(k) % m.Size()
	e := m.buckets[idx]
	for e.key != k {
		idx = (idx + 1) % m.Size()
		e = m.buckets[idx]
		if idx == XxHash(k)%m.Size() {
			// Not Found
			return nil
		}
	}
	return e.value
}

func (m *Bnumamap) LoadFactor() float32 {
	return float32(m.count) / float32(m.Size())
}

func (m *Bnumamap) DibAverage() float32 {
	sum := uint32(0)
	for _, v := range m.buckets {
		sum += v.scattercount
	}
	return float32(sum) / float32(m.count)
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