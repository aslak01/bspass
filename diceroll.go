package main

import (
	"fmt"
	"math/rand"
)

func rollTargetSumDiceSmart(n int, target int, minDiceVal int, maxDiceVal int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("number of dice (n) must be greater than zero")
	}
	if target < minDiceVal*n || target > maxDiceVal*n {
		return nil, fmt.Errorf("target sum %d is impossible with %d dice (range: %d - %d)", target, n, 3*n, 7*n)
	}

	rolls := make([]int, n)
	currentSum := 0

	for i := 0; i < n; i++ {
		minRemaining := (n - 1 - i) * minDiceVal
		maxRemaining := (n - 1 - i) * maxDiceVal
		lowerBound := target - currentSum - maxRemaining
		upperBound := target - currentSum - minRemaining

		if lowerBound > maxDiceVal || upperBound < minDiceVal {
			// Impossible to reach the target with the remaining dice, reroll the previous die
			if i > 0 {
				i--
				currentSum -= rolls[i]
				continue
			} else {
				// If we're at the first die and it's impossible, something is fundamentally wrong
				return nil, fmt.Errorf("impossible to reach target %d with %d dice", target, n)
			}
		}

		// If the remaining difference falls within the range of the remaining dice,
		// we can directly calculate the needed value for the current die.
		if lowerBound == upperBound && lowerBound >= 3 && lowerBound <= 7 {
			rolls[i] = lowerBound
			currentSum += lowerBound
			continue
		}

		// Otherwise, roll the current die randomly within the feasible range
		roll := rand.Intn(min(maxDiceVal, upperBound)-max(minDiceVal, lowerBound)+1) + max(minDiceVal, lowerBound)
		rolls[i] = roll
		currentSum += roll
	}

	if currentSum == target {
		return rolls, nil
	}

	// This should ideally not be reached with the smart logic, but as a fallback
	return nil, fmt.Errorf("could not find a combination of %d dice summing to %d", n, target)
}
