result := someFunc(1, "two", map[string]float64{})

// equal to

throwawayMap := make(map[string]float64)
result := someFunc(1, "two", throwawayMap)
