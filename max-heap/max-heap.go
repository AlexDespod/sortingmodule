package maxheap

import (
	"fmt"

	"github.com/AlexDespod/sortingmodule/structs"
)

type Maxheap struct {
	heapArray []MaxHeapNode
	size      int
	maxsize   int
}

type MaxHeapNode struct {
	structs.SortItem
	IndexOfFile int
}

func NewMaxHeap(maxsize int) *Maxheap {
	maxheap := &Maxheap{
		heapArray: []MaxHeapNode{},
		size:      0,
		maxsize:   maxsize,
	}
	return maxheap
}

func (m *Maxheap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *Maxheap) parent(index int) int {
	return (index - 1) / 2
}

func (m *Maxheap) leftchild(index int) int {
	return 2*index + 1
}

func (m *Maxheap) rightchild(index int) int {
	return 2*index + 2
}

func (m *Maxheap) Insert(item structs.SortItem, i int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("Heal is ful")
	}
	m.heapArray = append(m.heapArray, MaxHeapNode{item, i})
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}
func (m *Maxheap) GetMax() (MaxHeapNode, error) {
	if m.size != 0 {
		return m.heapArray[0], nil
	}
	return MaxHeapNode{}, fmt.Errorf("empty heap")
}

func (m *Maxheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *Maxheap) upHeapify(index int) {

	for m.heapArray[index].Num > m.heapArray[m.parent(index)].Num {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *Maxheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.size && m.heapArray[leftChildIndex].Num > m.heapArray[largest].Num {
		largest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex].Num > m.heapArray[largest].Num {
		largest = rightRightIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}

func (m *Maxheap) BuildMaxHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *Maxheap) Remove() (MaxHeapNode, error) {
	if m.size != 0 {
		top := m.heapArray[0]
		if m.size == 1 {
			m.size--
			m.heapArray = []MaxHeapNode{}
			return top, nil
		}
		m.heapArray[0] = m.heapArray[m.size-1]
		m.heapArray = m.heapArray[:(m.size)-1]
		m.size--
		m.downHeapify(0)
		return top, nil
	}
	return MaxHeapNode{}, fmt.Errorf("nothing to remove")
}
