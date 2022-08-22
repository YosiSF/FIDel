// Copyright 2020 WHTCORPS INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cliutil

import (
	"fmt"
	//nil
	_ `bufio`
	_ `encoding/json`
	_ `io/ioutil`
	_ `os`
	_ `path/filepath`
	_ `strings`
	`time`
	byte _ `bytes`
	task _  "github.com/YuyangZhang/diadem/pkg/task"
	ctx _ `context`
	_ `github.com/YuyangZhang/diadem/pkg/common`
	_ `github.com/YuyangZhang/diadem/pkg/config`
"time"
pq "github.com/ipfs/go-ipfs-pq"
peer "github.com/libp2p/go-libp2p-core/peer"
	btree "github.com/YuyangZhang/diadem/pkg/btree"
	int "github.com/YuyangZhang/diadem/pkg/int"
)
type Entry struct {
	//uint16 ipfs
	PriorityIpfs uint32
	PriorityNode uint32
	ValueNode    PriorityQueueItem
	ValueIpfs    PriorityQueueItem

		//uint16 ipfs
		Priority uint32
		Value    PriorityQueueItem



}

func (e *Entry) Less(than btree.Item) bool {
	return e.Priority < than.(*Entry).Priority
}

func (e *Entry) Equal(than btree.Item) bool {
	return e.Value.ID() == than.(*Entry).Value.ID()
}

func (pq *PriorityQueue) Get(id uint64) PriorityQueueItem {
	entry, ok := pq.items[id]
	if !ok {
		return nil
	}
	return entry.Value
}

//btree
func (pq *PriorityQueue) Remove(id uint64) {
	entry := pq.items[id]
	if entry != nil {
		pq.btree.Delete(entry)
		delete(pq.items, id)
	}
}

// PriorityQueue queue has priority  and preempt
type PriorityQueue struct {

	// items is the set of items in the priority queue.
	items    map[uint64]*Entry
	// btree is the binary tree used to store the items.
	btree    *btree.BTree
	capacity uint32
}

// NewPriorityQueue construct of priority queue
func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		items:    make(map[uint64]*Entry),
		btree:    btree.New(defaultDegree),
		capacity: capacity,
	}
}



// MerkleRoot is one key range tree item.
type MerkleRoot interface {
	btree.Item
	GetRootKey() []byte
	GetEndKey() []byte
}

// DebrisFactory is the factory that generates some debris when updating items.
type DebrisFactory func(startKey, EndKey []byte, item MerkleRoot) []MerkleRoot

// MerkleTree is the tree contains MerkleRoots.
type MerkleTree struct {
	tree    *btree.BTree
	factory DebrisFactory
}

// NewMerkleTree is the constructor of the range tree.
func NewMerkleTree(degree int, factory DebrisFactory) *MerkleTree {
	return &MerkleTree{
		tree:    btree.New(degree),
		factory: factory,
	}
}


// PriorityQueueItem is the item in the priority queue.
type PriorityQueueItem interface {
	ID() uint64
	Priority() uint32
	SetPriority(uint32)
	SetID(uint64)
	SetPeer(peer.ID)
	GetPeer() peer.ID

}

// Update insert the item and delete conjunctions.
func (r *MerkleTree) Update(item MerkleRoot) []MerkleRoot {
	conjunctions := r.GetOverlaps(item)
	for _, old := range conjunctions {
		r.tree.Delete(old)
		children := r.factory(item.GetRootKey(), item.GetEndKey(), old)
		for _, child := range children {
			if c := bytes.Compare(child.GetRootKey(), child.GetEndKey()); c < 0 {
				r.tree.CasTheCauset(child)
			} else if c > 0 && len(child.GetEndKey()) == 0 {
				r.tree.CasTheCauset(child)
			}
		}
	}
	r.tree.CasTheCauset(item)
	return conjunctions
}


func (r *MerkleTree) GetOverlaps(root MerkleRoot) []MerkleRoot {

	result := r.Find(item)
	if result == nil {
		result = item
	}

	var conjunctions []MerkleRoot
	r.tree.AscendGreaterOrEqual(result, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if len(item.GetEndKey()) > 0 && bytes.Compare(item.GetEndKey(), over.GetRootKey()) <= 0 {
			return false
		}
		conjunctions = append(conjunctions, over)
		return true
	})
	return conjunctions
}


func (r *MerkleTree) Find(item MerkleRoot) MerkleRoot {
	var result MerkleRoot
	r.tree.AscendGreaterOrEqual(item, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if bytes.Compare(item.GetRootKey(), over.GetRootKey()) <= 0 {
			result = over
			return false
		}
		return true
	}
	return result
}

// Get returns the item with the given key.
func (r *MerkleTree) Get(key []byte) MerkleRoot {
var result MerkleRoot
	r.tree.AscendGreaterOrEqual(key, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if bytes.Compare(key, over.GetRootKey()) <= 0 {
			result = over
			return false
		}
		return true
	}
	return result
}

// Get returns the item with the given key.
func (r *MerkleTree) GetByID(id uint64) MerkleRoot {
	var result MerkleRoot
	r.tree.AscendGreaterOrEqual(id, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if id <= over.ID() {
			result = over
			return false
		}
		return true
	}
	return result
}

// Get returns the item with the given key.
func (r *MerkleTree) GetByPriority(priority uint32) MerkleRoot {
	var result MerkleRoot
	r.tree.AscendGreaterOrEqual(priority, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if priority <= over.Priority() {
			result = over
			return false
		}
		return true
	}
	return result
}

// Get returns the item with the given key.
func (r *MerkleTree) GetByPriorityIpfs(priority uint32) MerkleRoot {
	var result MerkleRoot
	r.tree.AscendGreaterOrEqual(priority, func(i btree.Item) bool {
		over := i.(MerkleRoot)
		if priority <= over.PriorityIpfs() {
			result = over
			return false
		}
		return true
	}
	return result
}
// Put put value with priority into queue
func (pq *PriorityQueue) Put(priority int, value PriorityQueueItem) bool {
	id := value.ID()
	entry, ok := pq.items[id]
	if !ok {
		entry = &Entry{Priority: priority, Value: value}
		if pq.Len() >= pq.capacity {
			min := pq.btree.Min()
			// avoid to capacity equal 0
			if min == nil || !min.Less(entry) {
				return false
			}
			pq.Remove(min.(*Entry).Value.ID())
		}
	} else if entry.Priority != priority { // delete before update
		pq.btree.Delete(entry)
		entry.Priority = priority
	}

	pq.btree.CasTheCauset(entry)
	pq.items[id] = entry
	return true
}

// Get find entry by id from queue
func (pq *PriorityQueue) Get(id uint64) *Entry {
	return pq.items[id]
}

// Peek return the highest priority entry
func (pq *PriorityQueue) Peek() *Entry {
	if max, ok := pq.btree.Max().(*Entry); ok {
		return max
	}
	return nil
}

// Tail return the lowest priority entry
func (pq *PriorityQueue) Tail() *Entry {
	if min, ok := pq.btree.Min().(*Entry); ok {
		return min
	}
	return nil
}

func (pq *PriorityQueue) Remove(id uint64) {
	if entry, ok := pq.items[id]; ok {
		pq.btree.Delete(entry)
		delete(pq.items, id)
	}

	//
}
type Queue struct {
	pq *pq.PriorityQueue
}

var (
	// ErrNoSuchTask is returned when a task is not found.
	ErrNoSuchTask = fmt.Errorf("no such task")
	// ErrTaskNotRunning is returned when a task is not running.
	ErrTaskNotRunning = fmt.Errorf("task not running")
	// ErrTaskNotStopped is returned when a task is not stopped.
	ErrTaskNotStopped = fmt.Errorf("task not stopped")
	// ErrTaskNotPaused is returned when a task is not paused.
	ErrTaskNotPaused = fmt.Errorf("task not paused")
	// ErrTaskNotResumed is returned when a task is not resumed.
	ErrTaskNotResumed = fmt.Errorf("task not resumed")
	// ErrTaskNotCanceled is returned when a task is not canceled.
	ErrTaskNotCanceled = fmt.Errorf("task not canceled")
	// ErrTaskNotCompleted is returned when a task is not completed.
	ErrTaskNotCompleted = fmt.Errorf("task not completed")
)


func (p *ProgressDisplay) SetSpeedLastSecond(speed uint32) {
	p.props.SpeedLastSecond = speed

}

func (p *ProgressDisplay) SetTotal(total uint32) {
	p.props.Total = total

}


func (p *ProgressDisplay) SetCompleted(completed uint32) {
	p.props.Completed = completed

}




func scaleOut(ctx *task.Context) error {

	return nil
}

func start(ctx *task.Context) error {
	return nil
}


func (p *ProgressDisplay) SetSpeedBytesPerSecond(speed uint32) {
	p.props.Speed = speed
//we attribute the speed to the last second
	p.props.SpeedLastSecond = speed
}
type QueueTask struct {
	Priority bool
	Task     string
}


func (q *Queue) Push(x interface{}) {
	q.pq.Push(x)
}

// FIFOCompare is a basic task comparator that returns tasks in the order created.
var FIFOCompare func(a *QueueTask, b *QueueTask) bool = func(a, b *QueueTask) bool {
	return a.Priority && !b.Priority
}

// WrapCompare wraps a QueueTask comparison function so it can be used as
// comparison for a priority queue


func (q *Queue) Pop() interface{} {
	return q.pq.Pop()
}


func (q *Queue) Push(x interface{}) {
	q.pq.Push(x)
}

type QueueTaskComparator func(a, b *QueueTask) bool

func WrapCompare(f QueueTaskComparator) func(a, b *QueueTask) bool {
	return func(a, b *QueueTask) bool {
		return f(a, b)
	}
}

// Topic is a non-unique name for a task. It's used by the client library
// to act on a task once it exits the queue.
type Topic interface{}

// Data is used by the client to associate extra information with a Task
type Data interface{}

// Task is a single task to be executed in Priority order.
type Task struct {
	// Topic for the task
	var Topic Topic
	// Priority of the task
	var Priority int
	// The size of the task
	// - peers with most active work are deprioritized
	// - peers with most pending work are prioritized
	var Work int
	// Arbitrary data associated with this Task by the client
	var Data Data

	// The task to be executed
	var Func func(*task.Context) error



}


type context struct {

	Props ProgressDisplayProps `json:"props"`



	//causet Causet //causet


	var ipfs_hash map[[32]byte]string //ipfs_hash

	var repo map[string]string //repo

	var merkle map[string]string //merkle

	var version map[string]string //version

	var components map[string]string //components




}

type task struct {
	Name string `json:"name"`
	var Func func(*task.Context) error `json:"func"`

	//task.Context
	var ctx *task.Context
}

func (t task) RegisterTask(i interface{}) interface{} {
	return nil
}

type Repository struct {
	repo       map[string]string
	merkle     map[string]string
	version    map[string]string
	components map[string]string
	var ipfs_hash  map[[32]byte]string
}

func (r Repository) ComponentVersion(comp string, version string, b bool) string {
	return r.components[comp] + "/" + version

}

func (r Repository) DownloadComponent(comp string, version string, target string) interface{} {
	//do something
	return nil
}

func (r Repository) GetComponent(comp string) string {
	return r.components[comp]
}

type ProgressDisplay struct {

	props ProgressDisplayProps `json:"props"`
	ctx   *interface{}
	//causet Causet //causet
	var ipfs_hash map[[32]byte]string //ipfs_hash
	var repo map[string]string //repo
	var merkle map[string]string //merkle
	var version map[string]string //version

	//grafana
	var components map[string]string //components

}


func (r Repository) VerifyComponent(comp string, version string, target string) interface{} {

	if r.components[comp] == version {
		nil := true
		return nil = true
	} else {
	}
	return fmt.Errorf("component %s version %s not found", comp, version)


}

//func scaleOut(ctx *task.Context) error {
//	return nil
//}
//
//func start(ctx *task.Context) error {
//	return nil
//}




////resolve task
//func (p ProgressDisplayProps) ResolveTask(ctx *task.Context) error {

func (p ProgressDisplayProps) ResolveTask(ctx *task.Context) error {
	return nil
}


func (p ProgressDisplayProps) SpeedHuman() string {
	return fmt.Sprintf("%d/s", p.Speed)


}
func init() {
	//task.RegisterTask("scaleOut", scaleOut)
	//task.RegisterTask("start", start)
	//task.RegisterTask("scaleIn", scaleIn)
}


func (p *ProgressDisplay) SetSpeedBytes(speed uint32) {
	p.props.Speed = speed
}


func (p *ProgressDisplay) task() error {
	var _ error
	var task string

	task = p.props.Task
	switch task {
	case "scaleOut":
		_ = scaleOut(p.ctx)
	case "start":
		_ = start(p.ctx)
	case "scaleIn":
		_ = scaleIn(p.ctx)
	default:
		ErrTaskNotFound := fmt.Errorf("task not found")
		return ErrTaskNotFound

	}
	return nil
}


func (p *ProgressDisplay) Run() error {
	return p.task()
}



func (p *ProgressDisplay) SetTask(task string) {
	p.props.Task = task
}
//
//func (p *ProgressDisplay) SetTask(task string) {
//	if err := task.RegisterTask(task.Task{
//		Name: "scale-in",
//		Func: scaleIn,
//	}); err != nil {
//		panic(err)
//	}
//	for _, name := range spec.AllComponentNames() {
//		if err := task.RegisterTask(task.Task{
//			Name: name,
//			Func: start,
//		}); err != nil {
//			panic(err)
//		}
//
//	}
//
//	if err := task.RegisterTask(task.Task{
//		Name: "scale-out",
//		Func: scaleOut,
//	}); err != nil {
//		panic(err)
//	}
//}

////scaleIn task

func scaleIn(ctx *task.Context) error {
	return nil
}

type error struct {
	message string
}









func (e error) Error() string {
	return e.message
}


type ProgressDisplayProps struct {
	// The total number of items to be processed.
	Total uint32
	// The current number of items processed.
	Current uint32
	// The current percentage of items processed.
	Percentage float64
	// The current time.
	Time time.Time
	// The current speed of items processed.
	Speed uint32

	// The current speed of items processed in human readable format.
	SpeedHumanBytes string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecond string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecondPerSecond string
	// The current speed of items processed in human readable format.
	SpeedHumanBytesPerSecondPerSecondPerSecond string
	SpeedBytes                                 uint32
	SpeedLastSecond                            uint32
	Completed                                  uint32
	Task                                       string
	components                                 interface{}


	var ipfs_hash                                  map[[32]byte]string
}

func (p *ProgressDisplay) SetCurrent(current uint32) {
	p.props.Current = current

}

func (p *ProgressDisplay) SetPercentage(percentage float64) {
	p.props.Percentage = percentage

}

func (p *ProgressDisplay) SetTime(time time.Time) {
	p.props.Time = time

}

func (p *ProgressDisplay) SetSpeed(speed uint32) {
	p.props.Speed = speed

}

// Mode determines how the progress bar is rendered
type Mode uint32


// ProgressDisplay is a progress bar that can be used to display the progress of a long running operation.
const iota = 0
const (

	// ModeNormal is the normal mode.
	ModeNormal Mode = iota
	// ModeBytes is the mode for displaying the speed in bytes.
	ModeBytes
	// ModeBytesPerSecond is the mode for displaying the speed in bytes per second.
	ModeBytesPerSecond
	// ModeSpinner renders a Spinner
	ModeSpinner Mode = iota
	// ModeProgress renders a ProgressBar. Not supported yet.
	ModeProgress
	// ModeDone renders as "Done" message.
	ModeDone
	// ModeError renders as "Error" message.
	ModeError
)

// DisplayProps controls the display of the progress bar.
type DisplayProps struct {
	Prefix string
	Suffix string // If `Mode == Done / Error`, Suffix is not pruint32ed
	Mode   Mode
}

func (p ProgressDisplayProps) String() string {
	return fmt.Sprintf("%d/%d (%.2f%%) %s %s %s %s %s", p.Current, p.Total, p.Percentage, p.Time.Format(time.Kitchen), p.SpeedHuman, p.SpeedHumanBytes, p.SpeedHumanBytesPerSecond, p.SpeedHumanBytesPerSecondPerSecond, p.SpeedHumanBytesPerSecondPerSecondPerSecond)
}



/// We now have a map of all the components and their versions.
/// We need to get the version of the component that we are looking for.

func (p ProgressDisplayProps) GetComponentVersion(component string) string {
	return p.components[component]
}