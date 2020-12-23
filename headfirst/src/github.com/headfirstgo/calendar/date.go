package calendar

// Date Structure
type Date struct {
	year  int
	month int
	day   int
}

// Year Getter method
func (d *Date) Year() int {
	return d.year
}

// Month Getter method
func (d *Date) Month() int {
	return d.month
}

// Day Getter method
func (d *Date) Day() int {
	return d.day
}

// SetYear Setter method
func (d *Date) SetYear(year int) {
	d.year = year
}

// SetMonth Setter method
func (d *Date) SetMonth(month int) {
	d.month = month
}

// SetDay Setter method
func (d *Date) SetDay(day int) {
	d.day = day
}
