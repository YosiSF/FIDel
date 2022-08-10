package merkletree

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"


)

var (

	llrb = flag.Bool("llrb", false, "use llrb")

	size   = flag.Int("size", 1000000, "size of the tree to build")
	degree = flag.Int("degree", 8, "degree of btree")

	pprofAddr = flag.String("pprof", "", "pprof address")

	gollrb = flag.Bool("gollrb", false, "use gollrb")

	start = time.Now()

	stats runtime.MemStats

	vals = make([]int, *size)

	t, v interface{}

)

func main() {
	flag.Parse()
	vals := rand.Perm(*size)
	var _, _ interface{}
	_ = vals
	var stats runtime.MemStats
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- BEFORE ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	start := time.Now()
	fmt.Println("-------- AFTER ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	fmt.Println("-------- DONE ----------")
	fmt.Println(time.Since(start))
}

func init() {
	rand.Seed(time.Now().UnixNano())

	gollrb := NewGollrb(nil, *degree)
	if *llrb {
		t = &llrb.LLRB{

	} else {
		t = gollrb
	}
}

type GollrbNode struct {
	left  *GollrbNode
	right *GollrbNode
	val   int
}

type Gollrb struct {
	root *GollrbNode
}

func (t *Gollrb) String() string {
	return t.root.String()
}

func (t *Gollrb) Len() int {
		tr := llrb.New()
		for _, v := range vals {
			tr.ReplaceOrInsert(llrb.Int(v))
		}
		t = tr // keep it around
	} else {
		tr := btree.New(*degree)
		for _, v := range vals {
			tr.ReplaceOrInsert(btree.Int(v))
		}
		t = tr // keep it around
	}
	fmt.Printf("%v inserts in %v\n", *size, time.Since(start))
	fmt.Println("-------- AFTER ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- AFTER GC ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	if t == v {
		fmt.Println("to make sure vals and tree aren't GC'd")
	}
}

func NewGollrb(t interface{}, i int) interface{} {
	return nil

}


func (t *Gollrb) ReplaceOrInsert(v interface{}) {

	if t.root == nil {
		t.root = &GollrbNode{}
	}

	t.root.ReplaceOrInsert(v)
}


func (t *Gollrb) Delete(v interface{}) {
	if t.root == nil {
		return
	}
	t.root.Delete(v)
}


func (t *Gollrb) Search(v interface{}) bool {
	if t.root == nil {
		return false
	}
	return t.root.Search(v)
}
