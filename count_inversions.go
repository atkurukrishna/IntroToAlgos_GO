package main

import (
  "bufio"
  "fmt"
  "os"
)


func merge_count(array *[]int32, start int, mid int, mid_plus_one int, end int) (int64) {
	temp_array, loc := make([]int32, end - start + 1), 0;
	count := int64(0);
	p, q := start, mid_plus_one;
	for p <= mid && q <= end {
		if ((*array)[p] > (*array)[q]) {
			count += (int64)(mid_plus_one - p);
			temp_array[loc] = (*array)[q];
			loc += 1;
			q += 1;
		} else {
			temp_array[loc] = (*array)[p];
			loc++;
			p++;
		}
	}
	for p <= mid {
		temp_array[loc] = (*array)[p];
		p++;
		loc++;
	}
	for q <= end {
		temp_array[loc] = (*array)[q];
		q += 1;
		loc += 1;
	}
	p = start;
	loc = 0;
	for p <= end {
		(*array)[p]= temp_array[loc];
		loc += 1;
		p += 1;
	}
	return count;
}

func inversions(array *[]int32, start int, end int) (int64){
	if (start >= end) {
		return 0;
	}
	mid_point := (start + end) / 2;
	left_count := inversions(array, start, mid_point);
	//TODO(crisna): Simplify these functions using Go Slices
	right_count := inversions(array, mid_point + 1, end);
	return left_count + right_count + merge_count(array, start, mid_point, mid_point + 1, end);
}

func count_inversions(array *[]int32) {
	inversion_count := inversions(array, 0, len(*array) - 1);
	fmt.Printf("Total inversions is %v\n", inversion_count);
}

func main() {
	file, err := os.Open("count_inversions_input.txt");
	if (err != nil) {
		fmt.Printf("No input file found\n");
		return;
	}
	reader := bufio.NewReader(file);
	defer file.Close();

	input_array := make([]int32, 0);
	for {
		var buff int32;
		if _, err := fmt.Fscan(reader, &buff); err != nil {
			break;
		}
		input_array = append(input_array, buff);
	}
	count_inversions(&input_array);
}
