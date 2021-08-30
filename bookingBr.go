package main

import "sort"

type Booking struct {
	name  string
	value float64 // Value in real
	time  float64 // Time in seconds
}

func (b *Booking) getValue() float64 {
	return b.value
}

func (b *Booking) getTime() float64 {
	return b.time
}

func (b *Booking) units() float64 {
	return b.getValue() / b.getTime()
}

func buildBookingOnline(names []string, values []float64, time []float64) []Booking {
	booking := []Booking{}
	for index, value := range values {
		booking = append(booking, Booking{names[index], value, time[index]})
	}
	return booking
}

func glutton(items []Booking, maxTime float64, keyString string) ([]Booking, float64) {
	var bookingCopy = make([]Booking, len(items))
	copy(bookingCopy, items)

	switch keyString {
	case "value":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return bookingCopy[i].getValue() > bookingCopy[j].getValue()
		})

	case "time":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return (1 / bookingCopy[i].getTime()) > (1 / bookingCopy[j].getValue())
		})
	case "units":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return bookingCopy[i].units() > bookingCopy[j].units()
		})
	}

	result := []Booking{}
	var totalValue float64
	var totalTime float64
	totalValue, totalTime = 0.0, 0.0
	for i := 0; i < len(bookingCopy); i++ {
		if totalTime+bookingCopy[i].getTime() <= maxTime {
			result = append(result, bookingCopy[i])
			totalTime += bookingCopy[i].getTime()
			totalValue += bookingCopy[i].getValue()
		}
	}

	return result, totalValue
}
