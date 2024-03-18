package github.com/ygy1997/talkPlat/AC

// ACNode 表示 AC 自动机的节点结构
type ACNode struct {
	children map[rune]*ACNode
	fail     *ACNode
	isEnd    bool
	pattern  string
}

// ACAutomaton 结构表示 AC 自动机
type ACAutomaton struct {
	root *ACNode
}

// NewACAutomaton 初始化 AC 自动机
func NewACAutomaton() *ACAutomaton {
	return &ACAutomaton{
		root: &ACNode{
			children: make(map[rune]*ACNode),
			fail:     nil,
			isEnd:    false,
			pattern:  "",
		},
	}
}

// Insert 向 AC 自动机中插入模式串
func (ac *ACAutomaton) Insert(pattern string) {
	node := ac.root
	for _, char := range pattern {
		if node.children[char] == nil {
			node.children[char] = &ACNode{
				children: make(map[rune]*ACNode),
				fail:     nil,
				isEnd:    false,
				pattern:  "",
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.pattern = pattern
}

// BuildFailPointers 构建失败指针
func (ac *ACAutomaton) BuildFailPointers() {
	queue := []*ACNode{}
	for _, child := range ac.root.children {
		queue = append(queue, child)
		child.fail = ac.root
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for char, child := range currentNode.children {
			failNode := currentNode.fail
			for failNode != nil && failNode.children[char] == nil {
				failNode = failNode.fail
			}
			if failNode == nil {
				child.fail = ac.root
			} else {
				child.fail = failNode.children[char]
			}
			queue = append(queue, child)
		}
	}
}

// Search 在 AC 自动机中搜索模式串
func (ac *ACAutomaton) Search(text string) map[string]string {
	result := make(map[string]string)
	currentNode := ac.root

	for _, char := range text {
		for currentNode != nil && currentNode.children[char] == nil {
			currentNode = currentNode.fail
		}
		if currentNode == nil {
			currentNode = ac.root
			continue
		}
		currentNode = currentNode.children[char]
		if currentNode.isEnd {
			result[currentNode.pattern] = "true"
		}
	}

	return result
}
