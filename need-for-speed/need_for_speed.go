package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
        speed: speed,
        batteryDrain: batteryDrain,
        battery: 100,
        distance: 0,
    }
    return car
    //panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct{
    distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
    }
    return track
    //panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// If there is not enough battery to drive one more time,
	// the car will not move.
	if car.battery - car.batteryDrain < 0 {
        return car
    } else {
    	car.distance = car.distance + car.speed
        car.battery = car.battery - car.batteryDrain
        return car
    }
    //panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// How many times can I call drive?
    var canDrive int  = car.battery / car.batteryDrain
	
    if (canDrive * car.speed) >= track.distance{
        return true
    } else {
    	return false
    }
    // panic("Please implement the CanFinish function")
}
