package state

type PQ []State

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].heuristic == nil {
		pq[i].heuristic = pq[i].heurisitcFunc(&pq[i])
	}
	if pq[j].heuristic == nil {
		pq[j].heuristic = pq[j].heurisitcFunc(&pq[j])
	}
	return *(pq[i].heuristic) < *(pq[j].heuristic)
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
