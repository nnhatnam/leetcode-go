package _603_Design_Parking_System

type ParkingSystem struct {
	max      []int
	occupied []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		max:      []int{big, medium, small},
		occupied: make([]int, 3),
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.occupied[carType-1] >= this.max[carType-1] {
		return false
	} else {
		this.occupied[carType-1]++
		return true
	}
}

