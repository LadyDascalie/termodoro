package pomodoro

import (
	"strings"
	"testing"
	"time"
)

const (
	layout string = "Jan 01 2006 at 15:04:01"
)

func TestNewPomodoro(t *testing.T) {
	n := NewPomodoro()
	if n.Active != true {
		t.Fail()
	}
	t.Log("\n", n.Active)
}

func TestGetCurrentTime(t *testing.T) {
	ti := GetCurrentTime().Format(layout)
	t.Log("\n", ti)
}

func TestSetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
}

func TestGetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	st := GetStartTime(po)
	t.Log("\n", st)
}

func TestGetPomodoroDuration(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	time.Sleep(1 * time.Microsecond)
	pd := GetPomodoroDuration(po)
	t.Log("\n", pd)
}

func TestAddPomodoroDuration(t *testing.T) {
	n := NewPomodoro()
	n.Start = GetCurrentTime()
	AddPomodoroDuration(n)
	t.Log("\n", n.Start, n.End)
}

func TestPomodoroTimer(t *testing.T) {
	var duration time.Duration = 1
	ti := Timer(duration, time.Microsecond)
	if ti != false {
		t.Fail()
	}
}

func TestFormatDate(t *testing.T) {
	date, time := FormatDate(2015, time.January, 1, 0, 0)
	t.Log("\nDate string:", date, "\ntime.Time format:", time)
	if date != "Jan 01 2015 at 00:00:01" {
		t.Fail()
	}
}

func TestNewDefaultPomodoro(t *testing.T) {
	// Prepare a custom pomodoro
	n := NewDefaultPomodoro()
	st := n.Start.Format(layout)
	ed := n.End.Format(layout)
	ac := n.Active

	// Dates can't be empty
	// // Pomodoro must be active
	if st == "" || ed == "" {
		t.Fail()
	} else if ac != true {
		t.Fail()
	}

	t.Log("\n", n.Active, "\n", n.Start, "\n", n.End)
}

func TestFormatOutput(t *testing.T) {
	n := NewPomodoro()

	n.Active = true
	// Pass the string to the bit bucket operator (underscore) since we are not going to use it in the test
	_, n.Start = FormatDate(2015, time.January, 1, 0, 0)
	AddPomodoroDuration(n)

	expected := []string{"active", "Jan 01 2015 at 00:00:01", "Jan 01 2015 at 00:25:01"}
	got := FormatOutput(n)

	for i := 0; i < len(got); i++ {
		// INFO: godoc -http=":6060"
		// http://localhost:6060/ref/spec#Index_expressions
		if strings.Compare(expected[i], got[i]) != 0 {
			t.Fail()
			t.Log(strings.Compare(expected[i], got[i]))
		}
		t.Log("\nExpected:", expected[i], "\nReceived:", got[i])
	}

	if len(expected) != len(got) {
		t.Fail()
		t.Log("\n Expected:", len(expected), "but got:", len(got))
	}
}

// func TestTickerSince(t *testing.T) {
// 	n := NewPomodoro()
// 	n.Active = true
// 	_, n.Start = FormatDate(2015, time.January, 1, 0, 0)
// 	_, n.End = FormatDate(2015, time.January, 1, 0, 1)
// 	// po := n.End.Sub(n.Start)
// 	// c := time.Tick(1 * time.Second)
// 	tick := time.NewTicker(1 * time.Second)
// 	for now := range tick.C {
// 		t.Log(GetCurrentTime().Format(layout), now)
// 	}
//
// }
