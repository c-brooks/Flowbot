// Markov Chain
// This is a probabilistic model of language that treats groups of letters
// as nodes (3-letter groups for 3rd-order chain) and transitions
// between groups as edges

package ml

import (
	"fmt"
)

/*  TransitionTable
 * A 28x28 matrix M (26 letters + space + apostrophe) where M[i,j] represents
 * the likelihood of transitioning from the ith letter of the alphabet
 * to the jth letter of the alphabet.
 * Initialized a zeroes matrix.
 */
type TransitionTable [28][28]float64

func normalize(tt *TransitionTable, length float64) {
	for i := range tt  {
		for j := range tt[i] {
			tt[i][j] = tt[i][j]/length
		}
	}
}

func Train(song string, order int) {
	var tt TransitionTable
	AlphaLookup := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
		'i': 8,
		'j': 9,
		'k': 10,
		'l': 11,
		'm': 12,
		'n': 13,
		'o': 14,
		'p': 15,
		'q': 16,
		'r': 17,
		's': 18,
		't': 19,
		'u': 20,
		'v': 21,
		'w': 22,
		'x': 23,
		'y': 24,
		'z': 25,
		' ': 26,
		byte('\''): 27,
  	}

  	for i := range song {
		if i < len(song)-1 {
			tt[AlphaLookup[song[i]]][AlphaLookup[song[i+1]]]++
		}
  	}

	//normalize(&tt, (float64)(len(song)))

  // Print out tt
  //	for i := range tt {
	//	fmt.Println( i, tt[i] )
	//}

	Predict(&tt)
}

/** @param TransitionTable
  *
  */
func Predict(tt *TransitionTable) {

	AlphaLookup := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
		'i': 8,
		'j': 9,
		'k': 10,
		'l': 11,
		'm': 12,
		'n': 13,
		'o': 14,
		'p': 15,
		'q': 16,
		'r': 17,
		's': 18,
		't': 19,
		'u': 20,
		'v': 21,
		'w': 22,
		'x': 23,
		'y': 24,
		'z': 25,
		' ': 26,
		byte('\''): 27,
	}

	NumLookup := map[int]byte{
		 0:  'a',
		 1:  'b',
		 2:  'c',
		 3:  'd',
		 4:  'e',
		 5:  'f',
		 6:  'g',
		 7:  'h',
		 8:  'i',
		 9:  'j',
		 10: 'k',
		 11: 'l',
		 12: 'm',
		 13: 'n',
		 14: 'o',
		 15: 'p',
		 16: 'q',
		 17: 'r',
		 18: 's',
		 19: 't',
		 20: 'u',
		 21: 'v',
		 22: 'w',
		 23: 'x',
		 24: 'y',
		 25: 'z',
		 26: ' ',
		 27: byte('\''),
	}

	letter := byte('n')
	var max float64
	var maxAddr int
	var ret []byte
	for i := 0; i < 100; i++ {
		// find max score in letter's row in tt

		for j := range tt[AlphaLookup[letter]] {
			if tt[AlphaLookup[letter]][j] > max {
				max = tt[AlphaLookup[letter]][j]
				maxAddr = j
			}
		}
		letter = NumLookup[maxAddr]

		fmt.Println(string(letter))
		ret = append(ret, letter)
		max = 0
		maxAddr = 0
	}
	fmt.Println(string(ret))
}
