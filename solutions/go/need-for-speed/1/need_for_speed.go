package speed

type Car struct {
    battery	int
    batteryDrain int
    speed int
    distance int
}

type Track struct {
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
    }
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var b int
    var d int
    
    if (car.battery - car.batteryDrain) >= 0 { 
        b = car.battery - car.batteryDrain 
    } else { 
        b = car.battery
    }

    if (car.battery - car.batteryDrain) >= 0 {
        d = car.distance + car.speed 
    } else { 
        d = car.distance
    }
    
	return Car{
        speed: car.speed,
        batteryDrain: car.batteryDrain,
        battery: b,
        distance: d,
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if ((car.battery / car.batteryDrain) * car.speed ) >= track.distance {
    	return true    
    } else {
        return false
    }
}
