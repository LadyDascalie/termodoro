package pomodoro

import (
	"fmt"
	"time"
)

const (
	// PomodoroLength defines the default length of a pomodoro
	// in the future this value will be sourced via a .pomodororc file
	PomodoroLength time.Duration = 1

	// DefaultDuration calculates the default duration of a Pomodoro
	DefaultDuration = PomodoroLength * time.Second

	// Layout represents the default time layout format to use for time functions
	Layout string = "Jan 01 2006 at 15:04:01.000"
)

func main() {}

// Pomodoro defines the blueprint for a pomodoro
type Pomodoro struct {
	Start    time.Time
	Active   bool
	End      time.Time
	Elapsed  time.Time
	Duration time.Duration
}

// NewPomodoro creates a pomodoro object in memory
func NewPomodoro() *Pomodoro {
	return &Pomodoro{Active: true, Duration: 0}
}

// NewDefaultPomodoro sets the default values for a new pomodoro
func NewDefaultPomodoro() (n *Pomodoro) {
	n = NewPomodoro()
	SetStartTime(n, DefaultDuration)
	//SetPomodoroDuration(n)
	// StartTimer(n)
	return n
}

// GetCurrentTime is an exported wrapper for the time.Now() function
func GetCurrentTime() (t time.Time) {
	t = time.Now()
	return
}

// GetPomodoroDuration calculates how much time has passed since the pomodoro started
func GetPomodoroDuration(p *Pomodoro) (t time.Duration) {
	t = time.Since(p.Start)
	return
}

// SetPomodoroDuration sets the length of the pomodoro
func SetPomodoroDuration(p *Pomodoro) {
	p.End = p.Start.Add(PomodoroLength * time.Minute)
}

// SetStartTime sets the starting time of the pomodoro
// later will be used to also set pomodoros in advance
func SetStartTime(p *Pomodoro, duration time.Duration) {
	//p.Start = GetCurrentTime()
	fmt.Println(Timer())
}

// GetStartTime gets the time at which the pomodoro started
func GetStartTime(p *Pomodoro) (st time.Time) {
	st = p.Start
	return
}

// Timer counts down the time until active pomodoro ends
func Timer() (ending time.Time) {
	fmt.Println(time.Now().Local())
	timer := time.NewTimer(DefaultDuration)
	<-timer.C
	ending = time.Now().Local()
	return
}

// FormatDate is a wrapper function that allows easier setting of dates
func FormatDate(year int, month time.Month, day, hour, min int) (formatedDate string, date time.Time) {
	date = time.Date(year, month, day, hour, min, 0, 0, time.Local)
	formatedDate = date.Format(Layout)
	return
}

// FormatOutput returns the pomodoros as a slice of strings
func FormatOutput(p *Pomodoro) (output []string) {
	st := p.Start.Format(Layout)
	ed := p.End.Format(Layout)
	ac := p.Active
	var state string
	if ac == true {
		state = "active"
	}

	output = []string{state, st, ed}
	return
}
