package set

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueInt64(t *testing.T) {
	set, err := New[int64](func(item int64) string {
		return fmt.Sprintf("%d", item)
	}, func(a, b int64) bool {
		return a > b
	})
	if err != nil {
		t.Fatalf("initialize set: %v", err)
	}

	for _, v := range []int64{1, 2, 1, -5, -9, 4} {
		set.Add(v)
	}

	if err := set.Sort(); err != nil {
		t.Fatalf("sort: %v", err)
	}

	got := set.Unique()
	want := []int64{-9, -5, 1, 2, 4}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %+v got: %+v", want, got)
	}
}
