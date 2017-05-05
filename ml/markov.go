// Markov Chain
// This is a probabilistic model of language that treats words
// as nodes, chains of words as a root-to-leaf path (3-word groups for
// 3rd-order chain), and transitions between groups as edges

package ml

import (
	"fmt"
	"strings"
)


/**
 *	@type {map[string]*TransitionNode} RootTable
 *
 *	Entry points for forward traversal of the map of TransitionNodes
 *	Maps words -> pointers to the node (unique) that contains them
 */
type RootTable map[string]*TransitionNode

/**
 *	@type {map[string][]*TransitionNode} LeafTable
 *
 *	Entry points for backward traversal of the map of TransitionNodes
 *	Maps words -> pointers to a slice of nodes that contain them
 */
type LeafTable map[string][]*TransitionNode


/**
 *	@type {struct} TransitionNode
 *
 *	@member {map[string]*TransitionNode} children - Maps previous word -> pointer to its node
 *	@member {*TransitionNode}            parent   - Pointer to the node containing the next word
 *	@member {int}                        occ      - The occurrence of the node coming from its parent
 *	@member {string}                     word     - The value of the word
 */
type TransitionNode struct {
	children map[string]*TransitionNode
	parent *TransitionNode

	occ int
	word string
}




func Train(song string, order int) {
	words := strings.Split(song, " ")
	tt := make(RootTable)

	for i := range words {
		var root *TransitionNode

		if i < order + 1 {
			// Special case when just starting, I don't care for now.
			//tt["\n"] = make(map[string]map[string]float64)
			//tt["\n"]["\n"][words[i]] = 1
		} else {

			if _, ok := tt[words[i]]; !ok {
				// Make a new root TransitionNode
				root = &TransitionNode{
					children: nil,
					parent:   nil,
					occ:      1,
					word:     words[i],
				}
				tt[words[i]] = root
			} else {
				// TransitionNode with that word already exists
				tt[words[i]].occ ++
				root = tt[words[i]]
			}

			traverser := root
			for j := 0; j < order; j ++ {
				// TODO: This can be AddOrIncrementOcc
				if traverser.ChildExists(words[i-j]) {
					// If the child exists, update its occurrence and traverse deeper
					traverser.children[words[i-j]].occ ++
				} else {
					// Child doesn't exist, so we make one
					traverser.Add(words[i-j])
				}
				traverser = traverser.children[words[i-j]]
				fmt.Println( "\t", traverser.occ, traverser.word, traverser.children)
			}
		}
	}
//	Predict(tt)
}

/**
 * Predict
 * @param {RootTable} tt
 *
 * Approach (so I don't forget):
 * Take occurrence of result
 * If second order prediction yields results, double them.
 * If third order prediction yields results, triple them.
 * ...etc
 */


// Traverse the tree until its leaf for each letter to sum up it's occurrences
// Weight more "specific" occurrences higher
// func Predict(tt RootTable) {
//		var max int
//		var ret []string
//		for i := 0; i < 50; i++ {
//			// find max score for word in tt
//			for _, val1 := range tt {
//				for _, val2 := range val1.children {
//					for _, thirdWord := range val2.children {
//						if thirdWord.occ > max {
//							max = thirdWord.occ
//							ret = append(ret, thirdWord.word)
//						}
//					}
//					max = 0
//				}
//			}
//		}
//	fmt.Println(tt)
//	fmt.Println("======================")
//	fmt.Println(ret)
//}

/**
 * Add
 *
 * Adds a child with a given word to the children hash of a given node.
 */
func (tn *TransitionNode) Add(word string) {
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

/**
 * ChildExists
 *
 * Given a node, checks if a child exists in the children map
 * with the specified word
 */
func (tn TransitionNode) ChildExists(word string) bool {
	if _, ok := tn.children[word]; ok {
		return true
	}
	return false
}
