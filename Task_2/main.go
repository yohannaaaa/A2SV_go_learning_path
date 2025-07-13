package main

import (
	"fmt"
	"strings"
)

func wordCount(word string) map[string] int {
	new := strings.Split(word, " ")
	counter := map[string] int {}
	for _ , val:= range new{
		counter[string(val)] += 1
	}
	return counter

}

func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	new := ""
	for _,val := range word {
		if (val >= 'a' && val <= 'z') || (val >= '0' && val <= '9'){
			new +=string(val)

		}
	}
	for i:= 0; i < len(new)/2; i++{
		if new[i] != new[len(new) - 1 - i]{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println("Word Frequency:")
	fmt.Println(wordCount("Go is great, and Go is fast!")) 

	fmt.Println("Palindrome:")
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("racecar"))                        
	fmt.Println(isPalindrome("yohanna")) 

} 