package hashtable

// entry — один элемент бакета. При коллизии элементы с одинаковым
// индексом связываются в список через next (separate chaining).
type entry struct {
	key   string
	value int
	next  *entry
}

// HashTable — хэш-таблица string -> int на separate chaining.
type HashTable struct {
	buckets []*entry
	size    int // количество хранимых пар ключ-значение
}

// New создаёт пустую таблицу, готовую к работе.
func New() *HashTable {
	panic("TODO: implement New")
}

// hash отображает ключ в индекс бакета в диапазоне [0, len(buckets)).
func (h *HashTable) hash(key string) int {
	panic("TODO: implement hash")
}

// Put кладёт пару ключ-значение. Если ключ уже есть — перезаписывает значение.
func (h *HashTable) Put(key string, value int) {
	panic("TODO: implement Put")
}

// Get возвращает значение по ключу и true, либо 0 и false если ключа нет.
func (h *HashTable) Get(key string) (int, bool) {
	panic("TODO: implement Get")
}

// Delete удаляет ключ. Возвращает true если ключ был, false если его не было.
func (h *HashTable) Delete(key string) bool {
	panic("TODO: implement Delete")
}

// Len возвращает количество пар в таблице.
func (h *HashTable) Len() int {
	panic("TODO: implement Len")
}
