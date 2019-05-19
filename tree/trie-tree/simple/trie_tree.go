package simple

import "fmt"

type TrieNode struct {
	Child *map[string]TrieNode
	Exist bool
}

func main() {
	root := TrieNode{nil, false}
	root.addWord("a")
	root.addWord("b")
	root.addWord("apple")
	root.addWord("apple watch")
	root.addWord("banana")
	root.addWord("water")
	root.addWord("c")
	root.addWord("你好世界")

	root.traversal(0)

	testCases := [...]string{
		"b", "water", "watermelon",
		"app", "apple tv",
		"mango", "你好世界", "你好朋友", "你好世界上的朋友",
	}

	for _, str := range testCases {
		isExist, existStr := root.isExist(str)
		fmt.Println(str + " isExist ", isExist, existStr)
		fmt.Println(str + " isMatch ", root.isMatch(str))
	}
}

func (n *TrieNode) isMatch(words string) bool {
	runes := []rune(words)
	key := string(runes[0])
	theMap := *n.Child

	if _, ok := theMap[key]; !ok {
		return false
	} else {
		theNode := theMap[key]
		runesLen := len(runes)

		if runesLen == 1 {
			return theNode.Exist
		} else {
			if theNode.Child != nil {
				return theNode.isMatch(string(runes[1:]))
			} else {
				return false
			}
		}
	}
}

func (n *TrieNode) isExist(words string) (bool, string) {
	runes := []rune(words)
	key := string(runes[0])
	theMap := *n.Child
	existStr := key

	if _, ok := theMap[key]; !ok {
		return false, ""
	} else {
		theNode := theMap[key]
		runesLen := len(runes)

		if theNode.Exist || runesLen == 1 {
			if theNode.Exist {
				return true, existStr
			} else {
				return false, ""
			}
		} else {
			if theNode.Child != nil {
				bo, str := theNode.isExist(string(runes[1:]))
				if bo {
					return bo, existStr + str
				} else {
					return false, ""
				}
			} else {
				return false, ""
			}
		}
	}
}

func (n *TrieNode) traversal(deep int) {
	for k, v := range *n.Child {
		fmt.Println(deep, k, v)
		if v.Child != nil {
			v.traversal(deep + 1)
		}
	}
}

func (n *TrieNode) addWord(words string) {

	runes := []rune(words)
	keyStr := string(runes[0])

	var exist bool
	var restStr string
	if len(runes) == 1 {
		exist = true
		restStr = ""
	} else {
		exist = false
		restStr = string(runes[1:])
	}

	if n.Child == nil {
		tm := make(map[string]TrieNode)
		n.Child = &tm
	}

	tmpMap := *n.Child

	if _, ok := tmpMap[keyStr]; !ok {
		tmpMap[keyStr] = TrieNode{nil, exist}
	} else {
		if exist {
			tm := tmpMap[keyStr]
			tm.Exist = exist
			tmpMap[keyStr] = tm
		}
	}

	n.Child = &tmpMap

	if len(restStr) > 0 {
		tm := tmpMap[keyStr]
		tm.addWord(restStr)
		tmpMap[keyStr] = tm
	}
}
