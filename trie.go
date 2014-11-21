package cart

import "regexp"

type branch struct {
	children []*branch
	key      *route
	regexp   *regexp.Regexp
}

func (b *branch) add(node regexp.Regexp, rt *route, nodeIndex, nodeLength int) *branch {
	present := false
	var child *branch
	for i, c := range b.children {
		if (*c.regexp).MatchString(node.String()) {
			present = true
			child = b.children[i]
			if nodeIndex == (nodeLength - 1) {
				child.key = rt
			}
			break
		}
	}
	if !present {
		child = &branch{
			key:      rt,
			children: make([]*branch, 0),
			regexp:   &node,
		}
		b.children = append(b.children, child)
	}
	return child
}

func (b *branch) Insert(rt *route) {

	parent := b

	nodeLength := len(rt.nodes) + 1
	r := regexp.MustCompile(rt.method)
	parent = parent.add(*r, rt, 0, nodeLength)

	for i, n := range rt.nodes {
		r = regexp.MustCompile(n)
		parent = parent.add(*r, rt, i+1, nodeLength)
	}
}

func (b *branch) Find(path []string, method string) *branch {

	parent := b

	parent = parent.findChild(method)

	for _, node := range path {
		parent = parent.findChild(node)
	}

	return parent
}

func (b *branch) findChild(node string) *branch {
	for _, child := range b.children {
		if child.regexp.MatchString(node) {
			return child
		}
	}
	return nil
}

func newTrie() *branch {
	t := &branch{
		children: []*branch{
			&branch{regexp: regexp.MustCompile("POST")},
			&branch{regexp: regexp.MustCompile("GET")},
		},
		key:    new(route),
		regexp: new(regexp.Regexp),
	}
	return t
}
