package main

type Car struct {
	pos int
	speed int
	hour float64
}

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 || len(speed) == 0 {return 0}
	var cars []Car
	for i:=0; i<len(position); i++ {
		cars = append(cars, Car{
			pos: position[i],
			speed: speed[i],
			hour: float64((target-position[i]))/float64(speed[i]),
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})

	breakpt := make([]float64, len(cars))
	count := 1

	breakpt[len(cars)-1] = cars[len(cars)-1].hour
	for i:=len(cars)-2; i>=0; i-- {
		if cars[i].hour <= breakpt[i+1] {
			breakpt[i] = breakpt[i+1]
		} else {
			breakpt[i] = cars[i].hour
			count++
		}
	}
	// fmt.Println(cars)
	// fmt.Println(count, breakpt)
	return count
}
