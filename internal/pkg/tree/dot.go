// Package tree .
package tree

import (
	"fmt"
	"github.com/cozmo-zh/zearches/internal/pkg/tree/treenode"
	"github.com/cozmo-zh/zearches/pkg/siface"
	"html/template"
	"io"
	"strconv"
)

type PNode struct {
	Nodes    []*Elem
	Entities []*Elem
	Edges    []*Pair
}

type Elem struct {
	Name  string
	Label string
}

type Pair struct {
	Parent, Child *Elem
}

func parseNode(e *Elem, n *treenode.TreeNode) {
	if n.Parent() == nil {
		e.Name = "root"
		e.Label = "root"
	} else {
		e.Name = fmt.Sprintf("node_%d_%d_%d", n.Parent().Index(), n.Depth(), n.Index())
		e.Label = fmt.Sprintf("node_%d_%d", n.Depth(), n.Index())
	}
}

func ToDot(path string, root *treenode.TreeNode, output io.Writer) error {
	if tpl, err := template.ParseFiles(path); err != nil {
		return err
	} else if err = tpl.Execute(output, ToPNod(root)); err != nil {
		return err
	} else {
		return nil
	}
}

func ToPNod(root *treenode.TreeNode) *PNode {
	nodes := make([]*Elem, 0)
	entities := make([]*Elem, 0)
	edges := make([]*Pair, 0)
	root.Range(func(n *treenode.TreeNode) bool {
		node := &Elem{}
		parseNode(node, n)
		if n.Parent() != nil {
			pair := &Pair{
				Child: node,
			}
			parent := &Elem{}
			parseNode(parent, n.Parent())
			pair.Parent = parent
			edges = append(edges, pair)
		}
		nodes = append(nodes, node)
		if n.IsLeaf() {
			n.RangeEntities(func(entity siface.ISpatial) bool {
				el := &Elem{
					Name:  fmt.Sprintf("entity_%d", entity.GetID()),
					Label: strconv.Itoa(int(entity.GetID())),
				}
				entities = append(entities, el)
				pair := &Pair{
					Child: el,
				}
				parent := &Elem{}
				parseNode(parent, n)
				pair.Parent = parent
				edges = append(edges, pair)
				return true
			})
		}
		return true
	})
	return &PNode{
		Nodes:    nodes,
		Entities: entities,
		Edges:    edges,
	}
}
