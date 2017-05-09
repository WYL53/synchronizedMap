package synchronizedMap

import (
	"fmt"
	"testing"
)

func TestSynchronizedMap(t *testing.T) {
	//	t.Log("start")
	sm := New()
	for i := 0; i < 100; i++ {
		sm.Set(i, fmt.Sprintf("i = %d", i))
	}
	if sm.Len() != 100 {
		t.Fatal(" sm.Len() != 100")
	}
	for i := 0; i < 100; i++ {
		if !sm.IsContain(i) {
			t.Fatal(fmt.Sprint("!sm.IsContain(%d)", i))
		}
	}
	if sm.Get(0).(string) != "i = 0" {
		t.Fatal("sm.Get(0) != 'i = 0'")
	}
	sm.Each(func(k, v interface{}) {
		t.Log(fmt.Sprintf("key is %v,value is %v", k, v))
	})
	sm.Clear()
	if sm.Len() != 0 {
		t.Fatal("sm.Clear() not working")
	}
}

func BenchmarkSynchronizedMap(b *testing.B) {
	sm := New()
	//	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sm.Set(i, i)
	}
	//	b.StopTimer()
}

func BenchmarkMap(b *testing.B) {
	m := make(map[interface{}]interface{})
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}
