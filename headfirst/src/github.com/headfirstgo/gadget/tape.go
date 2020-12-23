package gadget

import "fmt"

// TapePlayer Structure
type TapePlayer struct {
	Batteries string
}

// Play Song TapePlayer
func (t *TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

// Stop TapePlayer
func (t *TapePlayer) Stop() {
	fmt.Println("Stopped !")
}

// TapeRecorder Structure
type TapeRecorder struct {
	Microphones int
}

// Play Song TapeRecorder
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

// Record TapeRecorder
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Stop TapeRecorder
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped !")
}
