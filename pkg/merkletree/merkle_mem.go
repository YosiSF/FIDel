package merkletree

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/google/btree"
	"github.com/petar/GoLLRB/llrb"
)

var (
	size   = flag.Int("size", 1000000, "size of the tree to build")
	degree = flag.Int("degree", 8, "degree of btree")
	gollrb = flag.Bool("llrb", false, "use llrb instead of btree")
)

func main() {
	flag.Parse()
	vals := rand.Perm(*size)
	var t, v interface{}
	v = vals
	var stats runtime.MemStats
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- BEFORE ----------")
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	start := time.Now()
	if *gollrb {
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
