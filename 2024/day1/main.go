package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
	"strconv"
)

func integer_absolute_value(a int, b int) int {
	/* part 1's helper function */
	difference := a - b
	if difference < 0 {
		return -difference
	}
	return difference
}

func part_1() int {
	/* read text file, pair up numbers ascending,
	find difference between integer pairs and add them */
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	slice1 := []int{}
	slice2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "   ")
		
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	total_distance := 0
	for i := 0; i < len(slice1); i++ {
		total_distance += integer_absolute_value(slice1[i], slice2[i])
	}

	return total_distance
}

func part_2() int {
	/* read file, count how many times a number in left list appears in right list,
	ex: if 2 appears four times in right, the score for this is 2 * 4 = 8
	do this for each number in left list and sum it all up */
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	counter := make(map[int]int)
	slice2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "   ")
		
		num1, _ := strconv.Atoi(fields[0])
		counter[num1] = 0

		num2, _ := strconv.Atoi(fields[1])
		slice2 = append(slice2, num2)		
	}

	for i := 0; i < len(slice2); i++ {
		_, exists := counter[slice2[i]]
		if exists {
			counter[slice2[i]] += 1
		}
	}

	score := 0
	for key, value := range counter {
		score += (key * value)
	}

	return score
}


func main() {
	fmt.Println("PART 1", part_1())
	fmt.Println("PART 2", part_2())
}