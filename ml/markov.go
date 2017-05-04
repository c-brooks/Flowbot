// Markov Chain
// This is a probabilistic model of language that treats groups of letters
// as nodes (3-letter groups for 3rd-order chain) and transitions
// between groups as edges

package ml

import (
	"fmt"
	"strings"
)


/*  TransitionTable
 * map[string]^Order float
 * EG: 2nd-order model
 * {
 * 	word: {
 * 	 word: occurrence,
 * 	 }
 * }
 */

type TransitionTable map[string]map[string]map[string]float64

//func normalize(tt TransitionTable, length float64) {
//	for key, val := range tt  {
//		for j := range tt[key] {
//			tt[i][j] = tt[i][j]/length
//		}
//	}
//}

func Train(song string, order int) {
	// This is super cheap so IDK about looping twice
	words := strings.Split(song, " ")
	tt := make(TransitionTable)

	for i := range words {

		if i < order + 1 {
			// Special case when just starting, I don't care for now.
			//tt["\n"] = make(map[string]map[string]float64)
			//tt["\n"]["\n"][words[i]] = 1
		} else {
			fmt.Println(words[i-2], words[i-1], words[i])

			if _, ok1 := tt[words[i-1]]; ok1 {
				if _, ok2 := tt[words[i-1]][words[i]]; ok2 {
					if _, ok3 := tt[words[i-2]][words[i-1]][words[i]]; ok3 {
						tt[words[i-2]][words[i-1]][words[i]]++
					} else {
						tt[words[i-2]] = make(map[string]map[string]float64)
						tt[words[i-2]][words[i-1]] = make(map[string]float64)
						tt[words[i-2]][words[i-1]][words[i]] = 1
					}
				}
			} else {
				tt[words[i-2]] = make(map[string]map[string]float64)
				tt[words[i-2]][words[i-1]] = make(map[string]float64)
				tt[words[i-2]][words[i-1]][words[i]] = 1
			}
		}
	}

	//normalize(&tt, (float64)(len(song)))

	Predict(tt)
}

/**
 * Predict
 * @param {TransitionTable} tt
 *
 * Approach (so I don't forget):
 * Take occurence of result
 * If second order prediction yields results, double them.
 * If third order prediction yields results, triple them.
 * ...etc
 */


func Predict(tt TransitionTable) {
		var max float64
		var ret []string
		for i := 0; i < 50; i++ {
			// find max score for word in tt
			for _, val1 := range tt {
				for _, val2 := range val1 {
					for thirdWord, occ := range val2 {
						if occ > max {
							max = occ
							ret = append(ret, thirdWord)
							max = 0
						}
					}
				}
			}
		}
	fmt.Println(tt)
	fmt.Println("======================")
	fmt.Println(ret)
}
