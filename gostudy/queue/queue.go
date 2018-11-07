package queue

// An FIFO queue
type Queue []interface{}

// Pushes the element into the queue.
//		e.g. q.push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
	// q 是一个slice 的指针， *q 是 q 指针指向的slice 实际的存储内容， 这句是追加 q 指针指向的slice的内容
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
	// 判断 len(*q) == 0 为真 则，返回 ture。为假返回 false
}