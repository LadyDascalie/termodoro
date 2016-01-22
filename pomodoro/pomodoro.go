package pomodoro

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/0xAX/notificator"
)

const (
	// PomodoroLength defines the default length of a pomodoro
	// in the future this value will be sourced via a .pomodororc file
	PomodoroLength time.Duration = 25

	// DefaultDuration calculates the default duration of a Pomodoro
	DefaultDuration = PomodoroLength * time.Minute

	// Layout represents the default time layout format to use for time functions
	Layout string = "Jan 01 2006 at 15:04:01.000"
)

var notify *notificator.Notificator

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
	return &Pomodoro{Active: true, Duration: DefaultDuration}
}

// NewDefaultPomodoro sets the default values for a new pomodoro
func NewDefaultPomodoro() (n *Pomodoro) {
	n = NewPomodoro()
	ending := Timer()
	if ending.Format(Layout) != "" {
		n.Active = true
	}

	if !n.Active {
		plsNotify()
	}

	return
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

func checkOsVersion() (os string) {
	os = runtime.GOOS
	return os
}

func plsNotify() {
	os := checkOsVersion()
	if os == "darwin" {
		cmd := exec.Command("bash", "-c", `osascript -e  'display  notification  "Your Pomodoro has ended!"  with  title  "Pomodoro"'`)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error sending notification")
		}
	} else if os == "linux" {
		notify = notificator.New(notificator.Options{
			AppName: "Termodoro",
		})
		notify.Push("Termordoro", "Your pomodoro has ended!", "", notificator.UR_NORMAL)
	} else if os == "window" {
		fmt.Println("Notifications aren't supported on your platform.")
	}

}
