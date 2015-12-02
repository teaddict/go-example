package main

import "fmt"
import "math/rand"

/*
#####################################
#########SET#########################
*/

type Element interface {
	Equal(Element) bool
}

type Set []Element
type intEle int

// HasElement returns true if set s contains element e.
//
// Mathematically, this is the fundamental binary relation defining sets.
func (s Set) HasElement(e Element) bool {
	for _, ex := range s {
		if e.Equal(ex) {
			return true
		}
	}
	return false
}

// Equal satisfies the Element interface.
//
// An object of type Set can be an element of a set.
func (s Set) Equal(e Element) bool {
	t, ok := e.(Set)
	if !ok {
		return false
	}
	if len(s) != len(t) {
		return false
	}
	for _, se := range s {
		if !t.HasElement(se) {
			return false
		}
	}
	return true
}

// AddElement adds a single element to a set.
func (p *Set) AddElement(e Element) {
	if !p.HasElement(e) {
		// always allocate new backing array.
		s1 := *p
		s2 := make(Set, len(s1)+1)
		copy(s2, s1)
		s2[len(s1)] = e
		*p = s2
	}
}

// RemoveElement removes a single element from a set.
//
// If the element is not in the set, the method has no effect.
func (p *Set) RemoveElement(e Element) {
	for i, ex := range *p {
		if e.Equal(ex) {
			s := *p
			last := len(s) - 1
			s[i], s[last] = s[last], s[i]
			*p = s[:last]
			return
		}
	}
}

// String satisfies fmt.Stringer, providing a printable representation of a set.
func (s Set) String() string {
	r := "{"
	for i, j := range rand.Perm(len(s)) {
		if i > 0 {
			r += " "
		}
		r += fmt.Sprint(s[j])
	}
	return r + "}"
}

// PowerSet returns the power set of a set.
//
// The power set of s is the set of all possible subsets of s.
func (s Set) PowerSet() Set {
	r := Set{Set{}}
	for _, es := range s {
		var u Set
		for _, er := range r {
			u = append(u, append(er.(Set), es))
		}
		r = append(r, u...)
	}
	return r
}

// Refexive validates the reflexive property of the Equal method for an element.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Reflexive(e Element) bool {
	return e.Equal(e)
}

// Symetric validates the symetric property of the Equal method for an element.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Symetric(a, b Element) bool {
	return a.Equal(b) == b.Equal(a)
}

// Transitive validates the transitive property of the Equal method for an element.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Transitive(a, b, c Element) bool {
	if a.Equal(b) && b.Equal(c) {
		return a.Equal(c)
	}
	return true
}

/*################################################################
##################################################################
############################END OF SET############################*/

/*################################################################
##################################################################
############################DEQUE############################*/
type Deque struct {
	arr   []interface{}
	front int
	back  int
	n     int
}

func NewDeque() *Deque {
	return &Deque{arr: nil, front: 0, back: 0, n: 0}
}

func (d *Deque) PushFront(v interface{}) {
	d.growIfNeeded()

	if d.front == 0 {
		d.front = cap(d.arr) - 1
	} else {
		d.front--
	}

	d.arr[d.front] = v
	d.n++
}

func (d *Deque) PushBack(v interface{}) {
	d.growIfNeeded()

	d.arr[d.back] = v
	d.back++
	if d.back == cap(d.arr) {
		d.back = 0 // wrap-around
	}
	d.n++
}

func (d *Deque) PopFront() interface{} {
	if d.n <= 0 {
		return nil
	}

	v := d.arr[d.front]
	d.arr[d.front] = nil
	d.front++
	if d.front == cap(d.arr) {
		d.front = 0 // wrap-around
	}
	d.n--

	d.shrinkIfNeeded()
	return v
}

func (d *Deque) PopBack() interface{} {
	if d.n <= 0 {
		return nil
	}

	if d.back == 0 {
		d.back = cap(d.arr) - 1
	} else {
		d.back--
	}

	v := d.arr[d.back]
	d.arr[d.back] = nil
	d.n--

	d.shrinkIfNeeded()
	return v
}

func (d *Deque) PeekFront() interface{} {
	if d.n <= 0 {
		return nil
	}
	return d.arr[d.front]
}

func (d *Deque) PeekBack() interface{} {
	if d.n <= 0 {
		return nil
	}

	if d.back == 0 {
		return d.arr[cap(d.arr)-1]
	}
	return d.arr[d.back-1]
}

func (d *Deque) ClearDeque() {
	d.arr = nil
	d.front = 0
	d.back = 0
	d.n = 0
}

func (d *Deque) EmptyDeque() bool {
	return d.n == 0
}

func (d *Deque) LenDeque() int {
	return d.n
}

func (d *Deque) resize(n int) {
	newArr := make([]interface{}, n)
	for i := 0; i < d.n; i++ {
		newArr[i] = d.arr[(d.front+i)%cap(d.arr)]
	}
	d.arr = newArr
	d.front = 0
	d.back = d.n
}

func (d *Deque) growIfNeeded() {
	if d.n == cap(d.arr) {
		d.resize((d.n + 1) * 2)
	}
}

func (d *Deque) shrinkIfNeeded() {
	if d.n > 0 && d.n <= cap(d.arr)/4 {
		d.resize((cap(d.arr) - 1) / 2)
	}
}

/*################################################################
##################################################################
############################END OF DEQUE############################*/

func min(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func solution(A []int, B []int, K int) int {
	var N int = len(A)

	return 0
}

func main() {
	var s Set
	if s.Equal(nil) {
		fmt.Println("nil") // nil set != nil interface
	}
	s.AddElement(intEle(3))
	var s2 Set
	s2.AddElement(4)
	fmt.Println(s.HasElement(3))
	fmt.Println(s2.String())

	deque := NewDeque()
	//sol başa ekle
	deque.PushFront(5)
	deque.PushFront(3)
	deque.PushFront(2)
	//sağ başa ekle
	deque.PushBack(9)

	fmt.Println(deque.LenDeque())
	fmt.Println(deque.PopFront())
	fmt.Println(deque.PopBack())

	var A = []int{5, 1, 0, 2, 7, 0, 6, 6, 1}
	var B = []int{1, 0, 7, 4, 2, 6, 8, 3, 9}
	var K = 2
	fmt.Println(solution(A, B, K))
}
