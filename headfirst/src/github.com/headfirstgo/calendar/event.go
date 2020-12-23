package calendar

// Event Structure
type Event struct {
	title string
	Date
}

// Title Getter Method
func (e *Event) Title() string {
	return e.title
}

// SetTitle Setter Method
func (e *Event) SetTitle(title string) {
	e.title = title
}
