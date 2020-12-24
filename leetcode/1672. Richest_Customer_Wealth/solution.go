package _672__Richest_Customer_Wealth

func maximumWealth(accounts [][]int) int {

	//we need a variable to store current max wealth of any customer,
	//to be able to compare it with others
	maxWealth := 0

	//currentWeath holds the temporary wealth calculation for each customer, to be compared with richest variable
	currentWealth := 0

	//we check every account
	for _, banks := range accounts {

		//we check all the banks of one customer
		for _, bank := range banks {

			//in each account, we collect one customer money in all banks
			currentWealth += bank
		}

		//this is probably more efficient than maxWealth = max(maxWeath, currentWeath) because it only SET maxWealth
		//to new value only if maxWealth > currentWealth
		if currentWealth > maxWealth {
			maxWealth = currentWealth
		}

		//reset currentWealth
		currentWealth = 0
	}
	return maxWealth
}
