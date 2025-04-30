package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func main() {
	wordCountPtr := flag.Int("w", 5, "Number of words")
	flag.IntVar(wordCountPtr, "words", 5, "Number of words")

	charCountPtr := flag.Int("l", 28, "Number of characters")
	flag.IntVar(charCountPtr, "length", 28, "Number of characters")

	minCharsPtr := flag.Int("minc", 2, "Minimum number of characters in a word")
	maxCharsPtr := flag.Int("maxc", 11, "Maximum number of characters in a word")

	separatorPtr := flag.String("s", " ", "Separator string")
	flag.StringVar(separatorPtr, "separator", " ", "Separator string")

	verbosePtr := flag.Bool("v", false, "Enable verbose output")
	flag.BoolVar(verbosePtr, "verbose", false, "Enable verbose output")

	copyPtr := flag.Bool("c", false, "Copy to clipboard")
	flag.BoolVar(copyPtr, "copy", false, "Copy to clipboard")

	flag.Parse()

	wordCount := *wordCountPtr
	charCount := *charCountPtr
	minChars := *minCharsPtr
	maxChars := *maxCharsPtr
	separator := *separatorPtr
	verbose := *verbosePtr
	copy := *copyPtr

	var start time.Time

	if verbose {
		start = time.Now()
	}

	result, err := rollTargetSumDiceSmart(wordCount, charCount, minChars, maxChars)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		words := []string{}

		for _, roll := range result {
			word, err2 := getRandomWordFromFile(roll)
			if err2 != nil {
				fmt.Printf("Error getting word for roll %d: %v\n", roll, err)
			} else {
				words = append(words, word)
			}
		}

		password := strings.Join(words, separator)
		fmt.Println(password)

		if copy {

			err := clipboard.Init()
			if err != nil {
				fmt.Println("Copying to clipboard failed")
				if verbose {
					fmt.Println(err.Error())
				}
			}

			clipboard.Write(clipboard.FmtText, []byte(password))

			if verbose {
				fmt.Println("✅ Copied to clipboard")
			}
		}

		if verbose {
			fmt.Println("Entropy comparison:")
			comparisons := map[string]float64{
				"bspass 📟": CalculatePasswordEntropy(password),
				"123456":   CalculatePasswordEntropy("123456"),
				"password": CalculatePasswordEntropy("password"),
				"hunter2":  CalculatePasswordEntropy("hunter2"),
			}
			for key, value := range comparisons {
				fmt.Printf("  - %-10s: %.1f bits\n", key, value) // Use %-10s for alignment
			}

			elapsed := time.Since(start)
			fmt.Printf("Generation Time: %.2f ms\n", float64(elapsed.Nanoseconds())/1000000)
		}
	}
}
