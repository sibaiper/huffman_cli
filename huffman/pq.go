package huffman

// PriorityQueue implements a priority queue for HuffmanNodes.
type PriorityQueue []*HuffmanNode

// Len is part of sort.Interface and returns the length of the queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less is part of sort.Interface and determines the priority order.
// Lower frequencies have higher priority.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

// Swap is part of sort.Interface and swaps two elements.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(node *HuffmanNode) {
	*pq = append(*pq, node)
	for i := len(*pq) - 1; i > 0; {
		parent := (i - 1) / 2
		if (*pq)[i].Freq >= (*pq)[parent].Freq {
			break
		}
		(*pq)[i], (*pq)[parent] = (*pq)[parent], (*pq)[i]
		i = parent
	}
}

func (pq *PriorityQueue) Pop() *HuffmanNode {
	if len(*pq) == 0 {
		return nil
	}
	root := (*pq)[0]
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]

	if len(*pq) == 0 {
		return root
	}

	(*pq)[0] = last
	for i := 0; 2*i+1 < len(*pq); {
		child := 2*i + 1
		if child+1 < len(*pq) && (*pq)[child+1].Freq < (*pq)[child].Freq {
			child++
		}
		if (*pq)[i].Freq <= (*pq)[child].Freq {
			break
		}
		(*pq)[i], (*pq)[child] = (*pq)[child], (*pq)[i]
		i = child
	}
	return root
}
