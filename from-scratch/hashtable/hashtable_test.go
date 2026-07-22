package hashtable

import (
	"fmt"
	"testing"
)

func TestGetOnEmpty(t *testing.T) {
	h := New()
	if v, ok := h.Get("nope"); ok || v != 0 {
		t.Fatalf("Get on empty table: got (%d, %v), want (0, false)", v, ok)
	}
	if h.Len() != 0 {
		t.Fatalf("Len on empty table = %d, want 0", h.Len())
	}
}

func TestPutGet(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value int
	}{
		{"simple", "foo", 1},
		{"empty key", "", 42},
		{"long key", "a-very-long-key-with-dashes", 7},
		{"unicode key", "ключ", 99},
		{"negative value", "neg", -5},
		{"zero value", "zero", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New()
			h.Put(tt.key, tt.value)
			v, ok := h.Get(tt.key)
			if !ok {
				t.Fatalf("Get(%q) after Put: ok=false, want true", tt.key)
			}
			if v != tt.value {
				t.Fatalf("Get(%q) = %d, want %d", tt.key, v, tt.value)
			}
			if h.Len() != 1 {
				t.Fatalf("Len() = %d after one Put, want 1", h.Len())
			}
		})
	}
}

func TestPutOverwrite(t *testing.T) {
	h := New()
	h.Put("key", 1)
	h.Put("key", 2)
	h.Put("key", 3)

	v, ok := h.Get("key")
	if !ok || v != 3 {
		t.Fatalf("Get after overwrite = (%d, %v), want (3, true)", v, ok)
	}
	if h.Len() != 1 {
		t.Fatalf("Len() = %d after overwriting same key, want 1", h.Len())
	}
}

func TestGetMissing(t *testing.T) {
	h := New()
	h.Put("present", 10)

	if v, ok := h.Get("absent"); ok || v != 0 {
		t.Fatalf("Get(absent) = (%d, %v), want (0, false)", v, ok)
	}
}

func TestDelete(t *testing.T) {
	h := New()
	h.Put("a", 1)
	h.Put("b", 2)

	if ok := h.Delete("a"); !ok {
		t.Fatalf("Delete(a) = false, want true")
	}
	if v, ok := h.Get("a"); ok || v != 0 {
		t.Fatalf("Get(a) after delete = (%d, %v), want (0, false)", v, ok)
	}
	if h.Len() != 1 {
		t.Fatalf("Len() = %d after deleting one of two, want 1", h.Len())
	}
	// оставшийся ключ не должен пострадать
	if v, ok := h.Get("b"); !ok || v != 2 {
		t.Fatalf("Get(b) after deleting a = (%d, %v), want (2, true)", v, ok)
	}
}

func TestDeleteMissing(t *testing.T) {
	h := New()
	h.Put("a", 1)

	if ok := h.Delete("missing"); ok {
		t.Fatalf("Delete(missing) = true, want false")
	}
	if h.Len() != 1 {
		t.Fatalf("Len() = %d after deleting missing key, want 1", h.Len())
	}
}

func TestDeleteThenPut(t *testing.T) {
	h := New()
	h.Put("k", 1)
	h.Delete("k")
	h.Put("k", 2)

	if v, ok := h.Get("k"); !ok || v != 2 {
		t.Fatalf("Get(k) after delete+put = (%d, %v), want (2, true)", v, ok)
	}
	if h.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", h.Len())
	}
}

// TestManyKeys гоняет много ключей: заведомо больше, чем стартовое число
// бакетов, чтобы поймать и коллизии в цепочках, и (если реализуешь) ресайз.
func TestManyKeys(t *testing.T) {
	h := New()
	const n = 1000

	for i := 0; i < n; i++ {
		h.Put(fmt.Sprintf("key-%d", i), i)
	}
	if h.Len() != n {
		t.Fatalf("Len() = %d after %d puts, want %d", h.Len(), n, n)
	}
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key-%d", i)
		v, ok := h.Get(key)
		if !ok || v != i {
			t.Fatalf("Get(%q) = (%d, %v), want (%d, true)", key, v, ok, i)
		}
	}
	// удаляем половину, вторая половина должна остаться нетронутой
	for i := 0; i < n; i += 2 {
		if ok := h.Delete(fmt.Sprintf("key-%d", i)); !ok {
			t.Fatalf("Delete(key-%d) = false, want true", i)
		}
	}
	if h.Len() != n/2 {
		t.Fatalf("Len() = %d after deleting half, want %d", h.Len(), n/2)
	}
	for i := 1; i < n; i += 2 {
		key := fmt.Sprintf("key-%d", i)
		v, ok := h.Get(key)
		if !ok || v != i {
			t.Fatalf("Get(%q) after partial delete = (%d, %v), want (%d, true)", key, v, ok, i)
		}
	}
}
