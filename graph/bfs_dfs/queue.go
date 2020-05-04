package bfs_dfs

type Item interface {

}

type Queue struct {
	data []Item
}

func NewQueue() *Queue{
	return &Queue{
		data: make([]Item, 0),
	}
}

func (q *Queue) EnQueue(i Item){
	q.data = append(q.data, i)
}

func (q *Queue) DeQueue() Item{
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) IsEmpty() bool{
	return len(q.data) == 0
}