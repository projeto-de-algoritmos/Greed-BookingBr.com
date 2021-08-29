package main

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
