package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {

    var success float64  = 0.0 // Default for 0 speed or others
    
	if speed >= 1 && speed <= 4 {
		success = 1
    } else if speed >= 5 && speed <= 8 {
    	success = 0.9
    } else if speed >= 9 && speed <= 10 {
    	success = 0.77
    }

    return success
	//panic("SuccessRate not implemented")
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return (float64(speed * 221) * SuccessRate(speed))
    //panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
    return (int(CalculateProductionRatePerHour(speed)) / 60)
    //panic("CalculateProductionRatePerMinute not implemented")
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if v := CalculateProductionRatePerHour(speed); v >= limit {
        return limit
    } else {
    	return v
    }
    
    //panic("CalculateLimitedProductionRatePerHour not implemented")
}
