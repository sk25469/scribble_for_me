package model

type Room struct {
	RoomID string   `json:"room_id"`
	Group1 []string `json:"grp1"`
	Group2 []string `json:"grp2"`
	index  int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Room

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return len(pq[i].Group1)+len(pq[i].Group2) < len(pq[j].Group1)+len(pq[j].Group2)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Room)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Room, ID string, Group1, Group2 []string) {
// 	item.RoomID = ID
// 	item.TotalClients = TotalClients
// 	heap.Fix(pq, item.index)
// }
