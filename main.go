// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"strconv"
)

type CountRunes struct {
	Character rune
	TimesAppeared int
}

func main() {
	result:= makeMap("5577")
	stringResult := MakeString(result)
	fmt.Println(stringResult)
}

func makeMap(inputString string) []CountRunes {
	input:= []rune(inputString)
	runeArray := make([]CountRunes,0)


	for  i := 0; i<len(input); i++ {
		var character CountRunes

		if i != 0 {
			if input[i] == input[i-1] {
				runeArray[len(runeArray)-1].TimesAppeared++
			} else {
				character.Character = input[i]
				character.TimesAppeared++

				runeArray = append(runeArray, character)
			}
		} else {
			character.Character = input[i]
			character.TimesAppeared++

			runeArray = append(runeArray, character)
		}
	}

	return runeArray
}


func MakeString(characters []CountRunes) string {
	finalString := ""


	for i:=0; i <len(characters); i++ {
		times := strconv.Itoa(characters[i].TimesAppeared)
		runeStr := string(characters[i].Character)

		finalString = finalString + times + runeStr
	}

	return finalString
}


/* 
Your previous Plain Text content is preserved below:

// 
// in -> string -> "777733333227", "5577", "1111111", "11111111111111111"
// out-> string -> "47532217",     "2527", "71", 
//
//
 */