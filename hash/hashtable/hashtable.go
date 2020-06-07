package hashtable

import "hash/fnv"

type TableItem struct {
	key string
	val interface{}
	next *TableItem
}

type HashTable struct {
	data [256]*TableItem
}

func NewHashTable() *HashTable{
	return &HashTable{}
}

func (t *HashTable) Add(k string, v interface{}){
	position := generateHash(k)
	if t.data[position] == nil{
		t.data[position] = &TableItem{
			key:  k,
			val:  v,
			next: nil,
		}
		return
	}
	current := t.data[position]
	for current.next != nil{
		current = current.next
	}
	current.next = &TableItem{
		key:  k,
		val:  v,
		next: nil,
	}
}

func (t *HashTable) Get(key string) (v interface{}, exist bool){
	position := generateHash(key)
	current := t.data[position]
	for current != nil{
		if current.key == key{
			return current.val, true
		}
		current = current.next
	}
	return nil, false
}

func (t *HashTable) Set(key string, v interface{}) bool{
	position := generateHash(key)
	current := t.data[position]
	for current != nil{
		if current.key == key{
			current.val = v
			return true
		}
		current = current.next
	}
	return false
}

func (t *HashTable) Remove(key string) bool{
	position := generateHash(key)
	if t.data[position].key == key{
		t.data[position] = t.data[position].next
		return true
	}
	current := t.data[position]
	for current.next != nil{
		if current.next.key == key{
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}
func generateHash(s string) uint8 {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(s))
	return uint8(hash.Sum32() % 256)
}
