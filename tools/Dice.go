package tools

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// DiceThrowResult keep rolled dices and numeric value
type DiceThrowResult struct {
	resultString string
	result       int
	throws       []int
}

const modifiers = "+-*/"

func diceThrow(d int) int {
	return rand.Intn(d) + 1
}

func splitDiceString(diceString string) []string {
	var diceArray []string
	indexStart := 0
	indexEnd := 0
	i := 0
	for pos, char := range diceString {
		if strings.Index(modifiers, string(char)) >= 0 {
			indexEnd = pos
			diceArray = append(diceArray, diceString[indexStart:indexEnd])
			diceArray = append(diceArray, string(char))
			indexEnd++
			indexStart = indexEnd
		}
		i++
	}
	indexEnd = len(diceString)
	diceArray = append(diceArray, diceString[indexStart:indexEnd])
	return diceArray
}

func processDiceArray(diceArray []string) DiceThrowResult {

	result := DiceThrowResult{"", 0, []int{}}

	for _, item := range diceArray {
		exploding := 0
		if strings.Index(item, "!") >= 0 {
			exploding = 1
		}
		for pos, char := range item {
			if strings.Index("dDkK", string(char)) >= 0 {
				throw, _ := strconv.Atoi(item[0:pos])
				sides, _ := strconv.Atoi(item[pos+1 : len(item)-exploding])
				if throw == 0 {
					throw = 1
				}
				var sResult string
				sResult = "["
				for x := 1; x <= throw; x++ {
					// ToDo: Add 'exploding' implementation
					r := diceThrow(sides)
					result.result += r
					sResult += fmt.Sprintf("%d", r)
					if x < throw {
						sResult += ", "
					}
				}
				sResult += "]"
				result.resultString += sResult
				break
			}
		}

	}
	return result
}

// Throw dice
func Throw(diceString string) {
	rand.Seed(time.Now().UnixNano())
	da := splitDiceString(diceString)
	for _, item := range da {
		fmt.Println(item)
	}
	rr := processDiceArray(da)
	fmt.Println(rr.resultString, rr.result)
}
