package impl

import "testing"

func TestAddAndGet(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	for i := range dummyDataLen {
		dummyData[i] = i
		as.Add(i, dummyData[i])
	}

	for i := range dummyDataLen {
		if as.Get(i) != dummyData[i] {
			t.Errorf("TestAddAndGet has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}

func TestAddAndSet(t *testing.T) {
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
			t.Errorf("TestAddAndSet has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}

func TestAddAndSize(t *testing.T) {
	as := NewArrayStack[int]()

	dummyDataLen := 100
	for i := range dummyDataLen {
		as.Add(i, i)
	}

	if as.Size() != dummyDataLen {
		t.Errorf("TestAddAndSize has an error. expected: %d, value: %d\n", dummyDataLen, as.Size())
	}
}

func TestAddAndRemove(t *testing.T) {
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
			t.Errorf("TestAddAndRemove has an error. expected: %d, value: %d\n", dummyData[i], as.Get(i))
		}
	}
}
