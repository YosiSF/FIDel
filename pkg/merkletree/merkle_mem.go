package merkletree

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	_ "sync"
	_ "sync/atomic"
	"time"
	_ "unsafe"
)


const DISPLAY_NAME = "Headless Sink Tree"

func (t *mvsr) String() string {
	if t.root == nil {
		return ""

	}

	return t.root.String()
}




type mvsr struct {
	root *mvsrNode



}




type mvsrNode struct {
	val interface{}
	left, right *mvsrNode

}


func (n mvsrNode) String() string {
	return fmt.Sprintf("%v", n.val)
}

/*
Bug Report
Deviation from expected behavior:
Disk zapping script sometimes returns error when removing the device mapper:

Creating new GPT entries in memory.
GPT data structures destroyed! You may now partition the disk using fdisk or
other utilities.
100+0 records in
100+0 records out
104857600 bytes (105 MB, 100 MiB) copied, 0.397611 s, 264 MB/s
blkdiscard: /dev/vdb: BLKDISCARD ioctl failed: Operation not supported
device-mapper: remove ioctl on ceph--27387f93--51fa--411b--b01c--b4f91e2962bd-osd--block--67ecef68--8e98--4d9b--96d9--b1d9659e2188  failed: Device or resource busy
Command failed.
For now I added set -e to the script, so I still would be able to unlock the device if it fails, otherwise the rest of the script deletes the device mappers. But We can use lsof to find the open files of the device and kill them to make the device free and then try to remove it.*/



type MerkleTree struct {

	root *mvsr

}




var (


	tr   *mvsr
	v    int
	size int
)



func init() {
	//set fields
	var 	_ []string 	= []string{"CR3", "NtBuildNumber", "KernBase", "KDBG", "KPCR00", "KPCR01", "KPCR02", "KPCR03", "KPCR04", "KPCR05", "KPCR06", "KPCR07", "KPCR08", "KPCR09", "KPCR10", "KPCR11", "KPCR12", "KPCR13", "KPCR14", "KPCR15", "KPCR16", "KPCR17", "KPCR18", "KPCR19", "KPCR20", "KPCR21", "KPCR22", "KPCR23", "KPCR24", "KPCR25", "KPCR26", "KPCR27", "KPCR28", "KPCR29", "KPCR30", "KPCR31", "PfnDataBase", "PsLoadedModuleList", "PsActiveProcessHead", "Padding0", "Padding1", "Padding2", "Padding3", "Padding4", "Padding5", "Padding6", "Padding7", "Padding8", "Padding9", "Padding10", "Padding11", "Padding12", "Padding13", "Padding14", "Padding15", "Padding16", "Padding17", "Padding18", "Padding19", "Padding20", "Padding21", "Padding22", "Padding23", "Padding24", "Padding25", "Padding26", "Padding27", "Padding28", "Padding29", "Padding30", "Padding31", "Padding32", "Padding33", "Padding34", "Padding35", "Padding36", "Padding37", "Padding38", "Padding39", "Padding40", "Padding41", "Padding42", "Padding43", "Padding44", "Padding45", "Padding46", "Padding47", "Padding48", "Padding49", "Padding50", "Padding51", "Padding52", "Padding53", "Padding54", "Padding55", "Padding56", "Padding57", "Padding58", "Padding59", "Padding60", "Padding61", "Padding62", "Padding63", "Padding64", "Padding65", "Padding66", "Padding67", "Padding68", "Padding69", "Padding70", "Padding71", "Padding72", "Padding73", "Padding74", "Padding75", "Padding76", "Padding77", "Padding78", "Padding79", "Padding80", "Padding81", "Padding82", "Padding83", "Padding84", "Padding85", "Padding86", "Padding87", "Padding88", "Padding89", "Padding90", "Padding91", "Padding92", "Padding93", "Padding94", "Padding95", "Padding96", "Padding97", "Padding98", "Padding99", "Padding100", "Padding101", "Padding102", "Padding103", "Padding104", "Padding105", "Padding106", "Padding107", "Padding108", "Padding109", "Padding110", "Padding111", "Padding112", "Padding113", "Padding114", "Padding115", "Padding116", "Padding117", "Padding118", "Padding119", "Padding120", "Padding121", "Padding122", "Padding123", "Padding124", "Padding125", "Padding126", "Padding127", "Padding128", "Padding129", "Padding130", "Padding131", "Padding132", "Padding133", "Padding134", "Padding135", "Padding136", "Padding137", "Padding138", "Padding139", "Padding140", "Padding141", "Padding142", "Padding143", "Padding144", "Padding145", "Padding146", "Padding147", "Padding148", "Padding149", "Padding150", "Padding151", "Padding152", "Padding153", "Padding154", "Padding155", "Padding156", "Padding157", "Padding158", "Padding159", "Padding160", "Padding161", "Padding162", "Padding163", "Padding164", "Padding165", "Padding166", "Padding167", "Padding168", "Padding169", "Padding170", "Padding171", "Padding172", "Padding173", "Padding174", "Padding175", "Padding176", "Padding177", "Padding178", "Padding179", "Padding180", "Padding181", "Padding182", "Padding183", "Padding184", "Padding185", "Padding186", "Padding187", "Padding188", "Padding189", "Padding190", "Padding191", "Padding192", "Padding193", "Padding194", "Padding195", "Padding196", "Padding197", "Padding198", "Padding199", "Padding200", "Padding201", "Padding202", "Padding203", "Padding204", "Padding205", "Padding206", "Padding207", "Padding208", "Padding209", "Padding210", "Padding211", "Padding212", "Padding213", "Padding214", "Padding215", "Padding216", "Padding217", "Padding218", "Padding219", "Padding220", "Padding221", "Padding222", "Padding223", "Padding224", "Padding225", "Padding226", "Padding227", "Padding228", "Padding229", "Padding230", "Padding231", "Padding232", "Padding233", "Padding234", "Padding235", "Padding236", "Padding237", "Padding238", "Padding239", "Padding240", "Padding241", "Padding242", "Padding243", "Padding244", "Padding245", "Padding246", "Padding247", "Padding248", "Padding249", "Padding250", "Padding251", "Padding252", "Padding253", "Padding254", "NumberOfRuns"} //issue driver command
	_ = make([]byte, 102400)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       //buffer for output
	var (
		_ error
	)
	var _ int
	var _ *os.File
	var (
		_ *os.File
		_ uint32 = 0
		_ uint32 = 0
	)
	//err := syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,
	//	uint32(len(inBuffer)),
	//	outBuffer,
	//	uint32(len(outBuffer)),
	//	&bytesReturned,
	//	nil,
	//)
	//if err != nil {
	//	return err
	//}
	//return nil

	//err = syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,
	//	uint32(len(inBuffer)),
	//	outBuffer,
	//	uint32(len(outBuffer)),
	//	&bytesReturned,
	//	nil,
	//)
	//if err != nil {
	//	return err
	//}
	//return nil

	//err = syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,





}


	flag.IntVar(&v, "v", 0, "value to insert")
	flag.IntVar(&size, "size", 0, "size of the tree")
}

func compare(a, b interface{}) int {
	flag.Parse()
	if a.(int) < b.(int) {
		return -1
	}
	if a.(int) > b.(int) {
		return 1

	}

	return 0

}

func (t *mvsr) ReplaceOrInsert(key interface{}) (replaced bool) {
	if t.root == nil {
		t.root = mvsrNode{val: key}
		return false
	}
	return t.root.ReplaceOrInsert(key)
}

func main() {
	flag.Parse()
	values := rand.Perm(*size)
	var _, _ interface{}
	_ = values
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

	tr = New(compare)
	for i := 0; i < *size; i++ {
		tr.ReplaceOrInsert(llrb.Int(rand.Int()))

	}
}










	//var (


	func (t *mvsr) Delete(key interface{}) (deleted bool) {
	return t.root.Delete(key)
}

func (n mvsrNode) Delete(key interface{}) bool {
	//persist on ipfs but delete on EinsteinDB
	return false
}

func (n mvsrNode) ReplaceOrInsert(key interface{}) (replaced bool) {
	if n.val == key {
		return true
	}
	if n.val < key {
		if n.right == nil {
			n.right = mvsrNode{val: key}
			return false
		}
		return n.right.ReplaceOrInsert(key)
	}
	if n.left == nil {
		n.left = mvsrNode{val: key}
		return false
	}
	return n.left.ReplaceOrInsert(key)
}

func (n mvsrNode) String() string {
	var b bytes.Buffer
	if n.left != nil {
		b.WriteString(n.left.String())
	}
	b.WriteString(fmt.Sprintf("%v ", n.val))
	if n.right != nil {
		b.WriteString(n.right.String())
	}
	return b.String()
}

func (t *mvsr) String() string {
	if t.root == nil {
		return ""
	}
	return t.root.String()
}

func (t *mvsr) Len() int {
	if t.root == nil {
		return 0
	}
	return t.root.Len()
}

func (t *mvsr) Max() interface{} {
	if t.root == nil {
		return nil
	}
	return t.root.Max()
}

func (t *mvsr) Min() interface{} {
	if t.root == nil {
		return nil
	}
	return t.root.Min()
}


func (t *mvsr) Get(key interface{}) (value interface{}, ok bool) {
if t.root == nil {
	if *size > 0 {
		return nil, false
	}
	return nil, true
		fmt.Println("-------- BEFORE ----------")
		runtime.ReadMemStats(&stats)
		fmt.Println(tr.String())
	if t.root == nil {
		return nil, true
	}
	return t.root.Get(key)
}

	return t.root.Get(key)
}

	func (n mvsrNode) Get(key interface{}) (value interface{}, ok bool) {
	if n.val == key {
		fmt.Println("-------- AFTER ----------")

		runtime.ReadMemStats(&stats)
		fmt.Printf("%+v\n", stats)
		return n.val, true
	}
	if n.val < key {
		fmt.Println(tr.String())
		if n.right == nil {


		fmt.Println("-------- DONE ----------")
		return nil, false
	} else {
		return n.right.Get(key)
	}


		}
		if n.left == nil {
			return nil, false
		}
		return n.left.Get(key)
		fmt.Println("-------- BEFORE ----------")


	//if n.val < key {
	//	if n.right == nil {
	//		return nil, false
	//	}
	//	return n.right.Get(key)
	//}

	//if n.left == nil {
	//	return nil, false
	//}




	//if n.val < key {
	//	if n.right == nil {
	//		return nil, false
	//	}
	//	return n.right.Get(key)
	//}







		fmt.Println("-------- AFTER ----------")

		runtime.ReadMemStats(&stats)

fmt.Printf("%+v\n", stats)
		fmt.Printf("%+v\n", stats)
		fmt.Println(tr.String())
		return nil, false
		fmt.Println("-------- AFTER ----------")
		runtime.ReadMemStats(&stats)
		fmt.Printf("%+v\n", stats)
		fmt.Println(tr.String())
		return nil, false
		fmt.Println("-------- DONE ----------")
		return nil, false
		fmt.Println("-------- BEFORE ----------")
		runtime.ReadMemStats(&stats)
		fmt.Println(tr.String())
		if n.val < key {


			if n.right == nil {
				return nil, false
			}
			return n.right.Get(key)




		}
	return n.left.Get(key)
}

type Node struct {
	Val   interface{}
	Left  *Node
	//merkle
	//prefix

	right *Node
	val   int
}



func (n mvsrNode) Len() int {
	return n.left.Len() + n.right.Len() + 1
}

func (n mvsrNode) Max() interface{} {
	if n.right == nil {
		return n.val
	}
	return n.right.Max()
}

func (n mvsrNode) Min() interface{} {
	if n.left == nil {
		return n.val
	}
	return n.left.Min()
}