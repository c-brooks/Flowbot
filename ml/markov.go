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
	rt := make(RootTable)
	lt := make(LeafTable)

	for i := range words {
		var root *TransitionNode

		if i < order + 1 {
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
				rt[words[i]].occ ++
				root = rt[words[i]]
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
			// When traverser is at leaf node, add it to LeafTable
			lt[traverser.word] = append(lt[traverser.word], traverser)
		}
	}
	fmt.Println(lt)
	Predict(lt)
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
func Predict(lt LeafTable) {
	var max int
	var ret []string

	node := findFirst(lt)
	for i := 0; i < 50; i++ {
		// find max score for word in lt
		for _, val1 := range node.children {
			for _, val2 := range val1.children {
				for _, thirdWord := range val2.children {
					if thirdWord.occ > max {
						max = thirdWord.occ
						ret = append(ret, thirdWord.word)
					}
				}
				max = 0
			}
		}
	}
	fmt.Println(lt)
	fmt.Println("======================")
	fmt.Println(ret)
}

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

/**
 *	findFirst
 *	@param {LeafTable} lt
 *
 *	For now, this just returns a random entry
 */
func findFirst(lt LeafTable) *TransitionNode {
	for _, v := range lt {
		if len(v) > 0 {
			return v[0]
		}
	}
	return nil
}
