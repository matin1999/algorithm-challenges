package main

import (
	"fmt"
	"math"
	"sort"
)

type DataPoint struct {
	X     float64
	Y     float64
	Label string
}

func EuclideanDistance(p1, p2 DataPoint) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

func KNN(dataPoints []DataPoint, queryPoint DataPoint, k int) string {
	type DistanceLabel struct {
		Distance float64
		Label    string
	}

	var distances []DistanceLabel
	for _, point := range dataPoints {
		distance := EuclideanDistance(point, queryPoint)
		distances = append(distances, DistanceLabel{Distance: distance, Label: point.Label})
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	labelCount := make(map[string]int)
	for i := 0; i < k && i < len(distances); i++ {
		labelCount[distances[i].Label]++
	}

	var majorityLabel string
	maxCount := 0
	for label, count := range labelCount {
		if count > maxCount {
			maxCount = count
			majorityLabel = label
		}
	}

	return majorityLabel
}

func main() {
	dataPoints := []DataPoint{
		{X: 1, Y: 2, Label: "A"},
		{X: 2, Y: 3, Label: "A"},
		{X: 3, Y: 3, Label: "B"},
		{X: 5, Y: 4, Label: "B"},
		{X: 6, Y: 1, Label: "A"},
		{X: 7, Y: 5, Label: "B"},
	}

	queryPoint := DataPoint{X: 4, Y: 4}
	k := 3

	result := KNN(dataPoints, queryPoint, k)
	fmt.Printf("The predicted class for the query point (%.1f, %.1f) is: %s\n", queryPoint.X, queryPoint.Y, result)
}
