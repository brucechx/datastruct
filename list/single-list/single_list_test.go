package single_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkList(t *testing.T) {
	l := NewLinkList()
	l.Append(1)
	l.Append(2)
	assert.Equal(t, "12", l.String())
	l.InsertHead(0)
	assert.Equal(t, "012", l.String())
	assert.Equal(t, 3, l.size)
	assert.Equal(t, true, l.Exist(1))
	err := l.InsertAfterData(1, 3)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0132", l.String())
	err = l.InsertNodeByPosition(2, 4)
	assert.Equal(t, nil, err)
	assert.Equal(t, "01342", l.String())
	_ = l.Delete(0)
	assert.Equal(t, "1342", l.String())
	assert.Equal(t, 4, l.size)
	_ = l.Delete(2)
	assert.Equal(t, "134", l.String())
	_ = l.Delete(3)
	assert.Equal(t, "14", l.String())
	assert.Equal(t, 2, l.size)
	l.Append(3)
	l.Append(5)
	l.Append(4)
	assert.Equal(t, "14354", l.String())
	l.Reverse()
	assert.Equal(t, "45341", l.String())
	l.reverse2()
	assert.Equal(t, "14354", l.String())
	assert.Equal(t, 5, l.size)
}

