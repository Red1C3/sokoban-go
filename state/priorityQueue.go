package state

// PQ is the underlying states priority queue.
type PQ []State

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].cost == nil {
		pq[i].cost = pq[i].costFunc(&pq[i])
	}
	if pq[j].cost == nil {
		pq[j].cost = pq[j].costFunc(&pq[j])
	}
	return *(pq[i].cost) < *(pq[j].cost)
}

func (pq PQ) Swap(i, j int) {
	temp := pq[i]
	pq[i] = pq[j]
	pq[j] = temp
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(State))
}

func (pq *PQ) Pop() any {
	v := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return v
}
