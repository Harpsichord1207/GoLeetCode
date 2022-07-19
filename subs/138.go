package subs

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {

	m := make(map[*Node138]*Node138)

	h := head
	for h != nil {
		n := &Node138{Val: h.Val}
		m[h] = n
		h = h.Next
	}

	h = head
	for h != nil {
		if h.Next != nil {
			// h的Next指向h.Next, 因此h对应的value的Next指向h.Next对应的value
			m[h].Next = m[h.Next]
		}
		if h.Random != nil {
			m[h].Random = m[h.Random]
		}
		h = h.Next
	}
	return m[head]

}
