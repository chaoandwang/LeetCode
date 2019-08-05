package container

type queue struct {
	que []interface{}
	index int
}

func (q *queue) Push(i interface{}) {
	q.que = append(q.que, i)
}

func (q *queue) Pop() interface{} {
	q.index += 1
	ret := q.que[q.index]
	return ret
}

func (q *queue) IsEmpty() bool {
	if len(q.que) == q.index + 1 {
		return true
	}
	return false
}

func NewQueue() *queue {
	return &queue{que:make([]interface{}, 0), index:-1}
}
