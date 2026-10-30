package model

import (
	"cmp"
	"slices"
	"strings"
)

type CMDBNode struct {
	Name       string    `json:"name"`
	CIID       int64     `json:"ci_id,string"`
	Department string    `json:"department"`
	Children   CMDBNodes `json:"children,omitzero"`
}

type CMDBNodes []*CMDBNode

func (cs CMDBNodes) Sort() {
	sortFn := func(a, b *CMDBNode) int {
		return cmp.Compare(a.CIID, b.CIID)
	}
	slices.SortFunc(cs, sortFn)

	for _, c := range cs {
		c.Children.Sort()
	}
}

type CMDBs []*CMDB

func (cs CMDBs) Treeing() CMDBNodes {
	idx := newCMDBNodeIndexMap()
	for _, c := range cs {
		idx.append(c)
	}

	roots := idx.getRoots()
	roots.Sort()

	return roots
}

type cmdbNodeIndexMap struct {
	index  map[string]*CMDBNode
	parent map[string]string // child -> parent mapping
}

func newCMDBNodeIndexMap() *cmdbNodeIndexMap {
	return &cmdbNodeIndexMap{
		index:  make(map[string]*CMDBNode, 64),
		parent: make(map[string]string, 64),
	}
}

func (im *cmdbNodeIndexMap) append(c *CMDB) {
	paths := strings.Split(c.Department, "/")
	node := im.lookup(paths)
	node.CIID = c.CIID
}

func (im *cmdbNodeIndexMap) lookup(paths []string) *CMDBNode {
	size := len(paths)
	if size == 0 {
		node := im.index[""]
		if node == nil {
			node = &CMDBNode{}
			im.index[""] = node
		}

		return node
	}

	full := strings.Join(paths, "/")
	if node := im.index[full]; node != nil {
		return node
	}

	var prev string
	for i := range size {
		current := strings.Join(paths[:i+1], "/")
		if im.index[current] == nil {
			node := &CMDBNode{
				Name:       paths[i],
				Department: current,
			}
			im.index[current] = node
			if prev != "" {
				im.parent[current] = prev
				im.index[prev].Children = append(im.index[prev].Children, node)
			}
		}
		prev = current
	}

	return im.index[full]
}

func (im *cmdbNodeIndexMap) getRoots() CMDBNodes {
	var roots CMDBNodes
	for path, node := range im.index {
		if im.parent[path] == "" {
			roots = append(roots, node)
		}
	}
	return roots
}
