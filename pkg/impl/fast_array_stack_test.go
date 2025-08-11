package impl

import "testing"

func TestFastArrayStackAddAndGet(t *testing.T) {
	fas := NewFastArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	for i := range dummyDataLen {
		dummyData[i] = i
		fas.Add(i, dummyData[i])
	}

	for i := range dummyDataLen {
		if fas.Get(i) != dummyData[i] {
			t.Errorf("TestFastArrayStackAddAndGet has an error. expected: %d, value: %d\n", dummyData[i], fas.Get(i))
		}
	}
}

func TestFastArrayStackAddAndSet(t *testing.T) {
	fas := NewFastArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	dummyInterval := 15
	for i := range dummyDataLen {
		dummyData[i] = i + dummyInterval
		fas.Add(i, 100)
		fas.Set(i, dummyData[i])
	}

	for i := range dummyDataLen {
		if fas.Get(i) != dummyData[i] {
			t.Errorf("TestFastArrayStackAddAndSet has an error. expected: %d, value: %d\n", dummyData[i], fas.Get(i))
		}
	}
}

func TestFastArrayStackAddAndSize(t *testing.T) {
	fas := NewFastArrayStack[int]()

	dummyDataLen := 100
	for i := range dummyDataLen {
		fas.Add(i, i)
	}

	if fas.Size() != dummyDataLen {
		t.Errorf("TestFastArrayStackAddAndSize has an error. expected: %d, value: %d\n", dummyDataLen, fas.Size())
	}
}

func TestFastArrayStackAddAndRemove(t *testing.T) {
	fas := NewFastArrayStack[int]()

	dummyDataLen := 100
	dummyData := make([]int, dummyDataLen)
	for i := range dummyDataLen {
		dummyData[i] = i
		fas.Add(i, dummyData[i])
		fas.Add(i+1, dummyDataLen+100)
		fas.Remove(i + 1)
	}

	for i := range dummyDataLen {
		if fas.Get(i) != dummyData[i] {
			t.Errorf("TestFastArrayStackAddAndRemove has an error. expected: %d, value: %d\n", dummyData[i], fas.Get(i))
		}
	}
}
