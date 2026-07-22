package singleLinkedList_test

import (
	"testing"

	singleLinkedList "cs-fundamentals-go/from-scratch/linked-list"
)

func buildList(values []int) *singleLinkedList.List {
	l := &singleLinkedList.List{}
	for _, v := range values {
		l.Insert(v)
	}
	return l
}

func toSlice(l *singleLinkedList.List) []int {
	var out []int
	for n := l.Head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func makeRange(start, end int) []int {
	out := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		out = append(out, i)
	}
	return out
}

func reverseInts(in []int) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[len(in)-1-i] = v
	}
	return out
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, nil},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
		{"duplicates", []int{1, 1, 2}, []int{1, 1, 2}},
		{"large input", makeRange(1, 1000), makeRange(1, 1000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildList(tt.input)
			got := toSlice(l)
			if !equalSlice(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		wantOK   bool
		wantList []int
	}{
		{"empty list", []int{}, 1, false, nil},
		{"only element", []int{1}, 1, true, nil},
		{"head", []int{1, 2, 3}, 1, true, []int{2, 3}},
		{"middle", []int{1, 2, 3}, 2, true, []int{1, 3}},
		{"tail", []int{1, 2, 3}, 3, true, []int{1, 2}},
		{"value missing", []int{1, 2, 3}, 5, false, []int{1, 2, 3}},
		{"removes first duplicate only", []int{1, 2, 1}, 1, true, []int{2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildList(tt.input)
			ok := l.Delete(tt.target)
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
			got := toSlice(l)
			if !equalSlice(got, tt.wantList) {
				t.Errorf("list = %v, want %v", got, tt.wantList)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		target    int
		wantFound bool
		wantValue int
	}{
		{"empty list", []int{}, 1, false, 0},
		{"found head", []int{1, 2, 3}, 1, true, 1},
		{"found middle", []int{1, 2, 3}, 2, true, 2},
		{"found tail", []int{1, 2, 3}, 3, true, 3},
		{"not found", []int{1, 2, 3}, 5, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildList(tt.input)
			got := l.Find(tt.target)
			if !tt.wantFound {
				if got != nil {
					t.Errorf("got %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("got nil, want node with value %d", tt.wantValue)
			}
			if got.Value != tt.wantValue {
				t.Errorf("got value %d, want %d", got.Value, tt.wantValue)
			}
		})
	}
}

func TestMiddle(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantNil   bool
		wantValue int
	}{
		{"empty", []int{}, true, 0},
		{"single", []int{1}, false, 1},
		{"two elements", []int{1, 2}, false, 2},
		{"three elements", []int{1, 2, 3}, false, 2},
		{"four elements", []int{1, 2, 3, 4}, false, 3},
		{"five elements", []int{1, 2, 3, 4, 5}, false, 3},
		{"six elements", []int{1, 2, 3, 4, 5, 6}, false, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildList(tt.input)
			got := l.Middle()
			if tt.wantNil {
				if got != nil {
					t.Errorf("got %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("got nil, want node with value %d", tt.wantValue)
			}
			if got.Value != tt.wantValue {
				t.Errorf("got value %d, want %d", got.Value, tt.wantValue)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, nil},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"duplicates", []int{1, 2, 2, 3}, []int{3, 2, 2, 1}},
		{"large input", makeRange(1, 1000), reverseInts(makeRange(1, 1000))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := buildList(tt.input)
			l.Reverse()
			got := toSlice(l)
			if !equalSlice(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("double reverse returns original", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		l := buildList(input)
		l.Reverse()
		l.Reverse()
		got := toSlice(l)
		if !equalSlice(got, input) {
			t.Errorf("got %v, want %v", got, input)
		}
	})
}

func TestHasCycle(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := &singleLinkedList.List{}
		if l.HasCycle() {
			t.Error("empty list reported as having a cycle")
		}
	})

	t.Run("no cycle", func(t *testing.T) {
		l := buildList([]int{1, 2, 3, 4})
		if l.HasCycle() {
			t.Error("acyclic list reported as having a cycle")
		}
	})

	t.Run("single node without self loop", func(t *testing.T) {
		l := buildList([]int{1})
		if l.HasCycle() {
			t.Error("single node reported as having a cycle")
		}
	})

	t.Run("single node self loop", func(t *testing.T) {
		n := &singleLinkedList.Node{Value: 1}
		n.Next = n
		l := &singleLinkedList.List{Head: n}
		if !l.HasCycle() {
			t.Error("self-referencing node not detected as a cycle")
		}
	})

	t.Run("cycle back to head", func(t *testing.T) {
		n1 := &singleLinkedList.Node{Value: 1}
		n2 := &singleLinkedList.Node{Value: 2}
		n1.Next = n2
		n2.Next = n1
		l := &singleLinkedList.List{Head: n1}
		if !l.HasCycle() {
			t.Error("cycle back to head not detected")
		}
	})

	t.Run("cycle into tail (not at head)", func(t *testing.T) {
		n1 := &singleLinkedList.Node{Value: 1}
		n2 := &singleLinkedList.Node{Value: 2}
		n3 := &singleLinkedList.Node{Value: 3}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n2
		l := &singleLinkedList.List{Head: n1}
		if !l.HasCycle() {
			t.Error("cycle into the tail not detected")
		}
	})
}
