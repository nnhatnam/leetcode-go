package _431__Kids_With_the_Greatest_Number_of_Candies

func kidsWithCandies(candies []int, extraCandies int) []bool {

	//we need a variable to store the greatest number of candies that one kid has
	currentGreatest := 0

	//iterate through candies to find the greatest candies number
	for _, candy := range candies {

		if candy > currentGreatest {
			currentGreatest = candy
		}
	}

	//if a kid can have the greatest candies number with extraCandies, his/her current candies must be greater or equal to
	// currentGreatest - extraCandies.
	minimumHeigh := currentGreatest - extraCandies

	result := make([]bool, len(candies))

	for i, candy := range candies {
		result[i] = candy >= minimumHeigh
	}

	return result

}
