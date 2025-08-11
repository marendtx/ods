package impl

import "testing"

func TestArrayStackAddAndGet(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	for i := range dummyDataLen {
		dummyData[i] = i
		as.Add(i, dummyData[i])
	}

	for i := range dummyDataLen {
		if as.Get(i) != dummyData[i] {
			t.Errorf("TestArrayStackAddAndGet has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}

func TestArrayStackAddAndSet(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	dummyInterval := 15
	for i := range dummyDataLen {
		dummyData[i] = i + dummyInterval
		as.Add(i, 100)
		as.Set(i, dummyData[i])
	}

	for i := range dummyDataLen {
		if as.Get(i) != dummyData[i] {
			t.Errorf("TestArrayStackAddAndSet has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}

func TestArrayStackAddAndSize(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	for i := range dummyDataLen {
		as.Add(i, i)
	}

	if as.Size() != dummyDataLen {
		t.Errorf("TestArrayStackAddAndSize has an error. expected: %d, value: %d\n", dummyDataLen, as.Size())
	}
}

func TestArrayStackAddAndRemove(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	for i := range dummyDataLen {
		dummyData[i] = i
		as.Add(i, dummyData[i])
		as.Add(i+1, dummyDataLen+100)
		as.Remove(i + 1)
	}

	for i := range dummyDataLen {
		if as.Get(i) != dummyData[i] {
			t.Errorf("TestArrayStackAddAndRemove has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}
