// Markov Chain
// This is a probabilistic model of language that treats words
// as nodes, chains of words as a root-to-leaf path (3-word groups for
// 3rd-order chain), and transitions between groups as edges

package ml

import (
	"fmt"
	"strings"
)

// RootTable contains entry points for forward traversal of the map of TransitionNodes
// Maps words -> pointers to the node (unique) that contains them
type RootTable map[string]*TransitionNode

// LeafTable contains entry points for backward traversal of the map of TransitionNodes
// Maps words -> pointers to a slice of nodes that contain them
type LeafTable map[string][]*TransitionNode

// TransitionNode is node in the transition tree with multiple children,
// one parent, a value "word", and an occurrence
type TransitionNode struct {
	children map[string]*TransitionNode
	parent   *TransitionNode

	occ  int
	word string
}

// Train Trains a Markov tree
func Train(song string, order int) {
	words := strings.Split(song, " ")
	rt := make(RootTable)
	lt := make(LeafTable)

	for i := range words {
		var root *TransitionNode

		if i < order+1 {
			// Special case when just starting, I don't care for now.
		} else {

			if _, ok := rt[words[i]]; !ok {
				// Make a new root TransitionNode
				root = &TransitionNode{
					children: nil,
					parent:   nil,
					occ:      1,
					word:     words[i],
				}
				rt[words[i]] = root
			} else {
				// TransitionNode with that word already exists
				rt[words[i]].occ++
				root = rt[words[i]]
			}

			traverser := root
			for j := 0; j < order; j++ {
				// TODO: This can be AddOrIncrementOcc
				if traverser.childExists(words[i-j]) {
					// If the child exists, update its occurrence and traverse deeper
					traverser.children[words[i-j]].occ++
				} else {
					// Child doesn't exist, so we make one
					traverser.add(words[i-j])
				}
				traverser = traverser.children[words[i-j]]
			}
			// When traverser is at leaf node, add it to LeafTable
			lt[traverser.word] = append(lt[traverser.word], traverser)
		}
	}
	predict(lt)
}

// Traverse the tree until its leaf for each letter to sum up its occurrences
// Weight more "specific" occurrences higher
func predict(lt LeafTable) {
	max := 0
	var ret []string
	var currentBestNode *TransitionNode
	word := findFirst(lt)

	for i := 0; i < 50; i++ {
		nodeArr := lt.findLeaf(word.word)

		// Explore each path
		for key, node := range nodeArr {
			fmt.Println(key, node.parent.word)
		}
		for _, node := range nodeArr {
			occAggr := 0 // Occurrence aggregate
			firstNode := node.parent
			for node.parent != nil {
				occAggr += node.occ
				node = node.parent
			}

			// Compare occurrence aggregates
			if occAggr > max {
				max = occAggr
				currentBestNode = firstNode
			}
		}

		ret = append(ret, currentBestNode.word)
		word = currentBestNode
		max = 0
	}

	fmt.Println("======================")
	fmt.Println(ret)
}

// Adds a child with a given word to the children hash of a given node.
func (tn *TransitionNode) add(word string) {
	if tn.children == nil {
		tn.children = map[string]*TransitionNode{}
	}
	tn.children[word] = &TransitionNode{
		children: nil,
		parent:   tn,
		occ:      1,
		word:     word,
	}
}

// Given a node, checks if a child e//ists in the children map
// with the specified word
func (tn TransitionNode) childExists(word string) bool {
	if _, ok := tn.children[word]; ok {
		return true
	}
	return false
}

// Returns a random leaf node
func findFirst(lt LeafTable) *TransitionNode {
	for _, v := range lt {
		if len(v) > 0 {
			return v[0]
		}
	}
	return nil
}

// Returns an array of leaf nodes for a given word
func (lt LeafTable) findLeaf(word string) []*TransitionNode {
	return lt[word]
}
