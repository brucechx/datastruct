package hashtable

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	h := NewHashTable()
	h.Add("1", 1)
	h.Add("1", 2)
	h.Add("1", 3)
	h.Add("2", 3)
	fmt.Println(h.Get("1"))
	fmt.Println(h.Get("1"))
}


func TestDemo(t *testing.T){
	fmt.Println(5/2)
}