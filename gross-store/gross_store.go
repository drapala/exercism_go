package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// Create a Unit Map
    unitMap := map[string]int{}

    // Insert Units
    unitMap["quarter_of_a_dozen"] = 3
    unitMap["half_of_a_dozen"] = 6
    unitMap["dozen"] = 	12
    unitMap["small_gross"] = 120
    unitMap["gross"] = 144
    unitMap["great_gross"] = 1728

    // Print
    fmt.Println(unitMap)
    
    return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]

    if (exists == false ) {
        return false
    } else {
    	bill[item] = units[unit]
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given item is not in the bill
	_ , itemExists := bill[item]
    if !itemExists {
        return false
    }
    // Return false if the given unit is not in the units map
    _ , unitExists := units[unit]
    if !unitExists {
        return false
    }  

    // Next steps
	remaining := bill[item] - units[unit]
    if remaining < 0 {
        // Return false if the new quantity would be less than 0
        return false
    } else if remaining == 0 {
    	// If the new quantity is 0, completely remove the item from the bill then return true
    	delete(bill, item)
        return true
    } else {
    	// Otherwise, reduce the quantity of the item and return true
		bill[item] = remaining
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
