package main

import (
	"math/rand"
	"errors"
	"os"
	"fmt"
)

var WRC_ERROR error = 
	errors.New("Weighted Random Choice: Length of weights must be equal to size of range!")

func weighted_random_choice(min_val int, max_val int, weights []int) int {
	if (len(weights) != (max_val - min_val)) {
		fmt.Println(WRC_ERROR)
		os.Exit(1)
	}
	var weight_sum int = 0
	for i := 0; i < len(weights); i++ {
		weight_sum += weights[i]
	}
	var randint int = rand.Intn(max_val) + min_val
	for i := 0; i < len(weights); i++ {
		if (randint < weights[i]) {
			return i + min_val
		}
		randint -= weights[i]
	}
	return -1
}
