package treesort

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *tree) String() string {
	return t.stringRecurse(0)
}

const indentTail = "  "
const indentHead = "+-"

// stringRecurse returns a string-representation of t, where the root node is indented depth depthIndents
func (t *tree) stringRecurse(depth int) string {
	tree := ""
	if depth >= 2 {
		tree += strings.Repeat(indentTail, depth-1)
	}
	if depth >= 1 {
		tree += indentHead
	}
	tree += strconv.Itoa(t.value)

	if t.left != nil {
		tree += "\n" + t.left.stringRecurse(depth+1)
	}
	if t.right != nil {
		tree += "\n" + t.right.stringRecurse(depth+1)
	}
	return tree
}

func Ugh() {
	root := &tree{
		value: 12,
		left: &tree{
			value: 3489201,
			left: &tree{
				value: 4981,
				right: &tree{
					value: 3,
				},
			},
		},
		right: &tree{
			value: 8,
		},
	}

	fmt.Printf("%v\n", root)
}
