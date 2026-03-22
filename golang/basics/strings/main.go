package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============ STRING OPERATIONS ============

	// TODO: Declare s := "hello"

	s := "hello"

	// TODO: Print the byte length of s using len()

	fmt.Println(len(s))
	// TODO: Print s as a slice of runes (for unicode-safe iteration)

	fmt.Println([]rune(s))

	// TODO: Print the first byte of s using s[0]

	fmt.Println(s[0])


	// TODO: Convert the first byte of s to a string and print it

	fmt.Println(string(s[0]))

	// TODO: Concatenate s with " world" and print

	s += " world"

	// TODO: Check if s contains "ell" using strings.Contains and print result

	ok := strings.Contains(s , "ell")
	if ok{
		fmt.Println("contains yes")
	}else{
		fmt.Println("contains no")
	}

	// TODO: Split "a,b,c" by "," using strings.Split and print result

	t := "a,b,c"
	res := strings.Split(t, ",")

	for _, item := range res{
		fmt.Println(item)
	}

	// TODO: Join []string{"x", "y", "z"} with "-" using strings.Join and print

	t1 := []string{"x", "y", "z"}

	res1 := strings.Join(t1,"-")

	fmt.Println(res1)


	// TODO: Check if s has prefix "he" using strings.HasPrefix and print

	ok = strings.HasPrefix(s, "he")

	if ok{
		fmt.Println("prefix yes")
	}else{
		fmt.Println("prefix no")
	}

	// TODO: Repeat "ab" 3 times using strings.Repeat and print (should be "ababab")

	res1 = strings.Repeat("ab",3)
	fmt.Println(res1)
	// ============ STRING BUILDER ============
	// Efficient string concatenation in loops

	// TODO: Create a strings.Builder called sb

	var sb strings.Builder
	// TODO: Write "hello" to sb using WriteString
	sb.WriteString("hello")
	// TODO: Write a space byte using WriteByte
	sb.WriteByte(' ')
	// TODO: Write "world" to sb using WriteString
	sb.WriteString("world")

	// TODO: Print the result using sb.String()

	fmt.Println(sb.String())


}
