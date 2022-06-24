package lcof_median_finder

type MedianFinder struct {
	minHeap *Heap
	maxHeap *Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: InitHeap([]int{}, Min),
		maxHeap: InitHeap([]int{}, Max),
	}
}

func (this *MedianFinder) AddNum(num int) {
	//先写入大顶堆
	this.maxHeap.Push(num)
	//将大顶堆中最大的数放入小顶堆
	this.minHeap.Push(this.maxHeap.Pop())
	//如果小顶堆的数据比大顶堆的数据多了不止一个，则将小顶堆的顶取出放入大顶堆中
	//这种方式可以保证小顶堆中的数据是所有数据中较大的一部分，且堆顶元素是这些数据中最小的
	//而大顶堆的数据是所有数据中较小的一部分，且堆顶元素是这些数据中最大的
	if this.maxHeap.Size()+1 < this.minHeap.Size() {
		this.maxHeap.Push(this.minHeap.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	maxHead := this.maxHeap.Peak()
	minHead := this.minHeap.Peak()

	if this.maxHeap.Size() == this.minHeap.Size() {
		if maxHead == nil && minHead == nil {
			return 0
		}

		return (float64(*maxHead) + float64(*minHead)) / 2
	}

	if this.maxHeap.Size() > this.minHeap.Size() {
		return float64(*maxHead)
	} else {
		return float64(*minHead)
	}
}

type HeapifyKind int

const (
	Undefined HeapifyKind = iota
	Min                   = iota
	Max
)

type Heap struct {
	data []int
	kind HeapifyKind
}

func InitHeap(data []int, kind HeapifyKind) *Heap {
	if kind == Undefined {
		kind = Max
	}
	h := &Heap{data: data, kind: kind}

	h.Heapify(kind)

	return h
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Push(data int) {
	h.data = append(h.data, data)

	h.up(len(h.data) - 1)
}

func (h *Heap) Peak() *int {
	if len(h.data) > 0 {
		return &h.data[0]
	}

	return nil
}

func (h *Heap) Pop() int {
	if len(h.data) == 0 {
		return -1
	}

	data := h.data[0]

	if len(h.data) > 1 {
		h.data[0] = h.data[len(h.data)-1]
		h.data = h.data[:len(h.data)-1]
		h.down(0)
	} else {
		h.data = []int{}
	}

	return data
}

func (h *Heap) up(i int) {
	//父节点
	i0 := (i - 1) >> 1

	for i0 >= 0 {
		//当大顶堆时，父节点大于当前节点，或当小顶堆时父节点小于当前节点，退出
		if (h.kind == Max && h.data[i0] > h.data[i]) || h.kind == Min && h.data[i0] < h.data[i] {
			break
		}

		h.data[i0], h.data[i] = h.data[i], h.data[i0]
		i = i0
		i0 = (i - 1) >> 1
	}
}

func (h *Heap) down(i int) {
	for i < len(h.data) {
		j := i*2 + 1 //左子节点

		if j > len(h.data)-1 {
			break
		}

		//当为大顶堆时，取出左右子节点的最大者和当前节点对比；当为小顶堆时，取出左右子节点的最小者和当前节点对比
		if j0 := j + 1; j0 <= len(h.data)-1 && ((h.kind == Max && h.data[j0] > h.data[j]) || (h.kind == Min && h.data[j0] < h.data[j])) {
			j = j0
		}

		if (h.kind == Max && h.data[i] > h.data[j]) || h.kind == Min && h.data[i] < h.data[j] {
			break
		}

		h.data[i], h.data[j] = h.data[j], h.data[i]

		i = j
	}
}

// Heapify 堆化一个数组，找到最后一个不是叶子的节点自下往上开始下沉
func (h *Heap) Heapify(kind HeapifyKind) {
	if len(h.data) == 0 {
		return
	}

	if kind != Undefined {
		h.kind = kind
	}

	//最后一个叶子的父节点为最后一个非叶子节点
	lastNonLeaf := (len(h.data) - 2) >> 1

	for i := lastNonLeaf; i >= 0; i-- {
		h.down(i)
	}
}
