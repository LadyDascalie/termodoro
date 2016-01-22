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
	SetStartTime(po, DefaultDuration)
}

func TestGetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po, DefaultDuration)
	st := GetStartTime(po)
	t.Log("\n", st)
}

func TestGetPomodoroDuration(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po, DefaultDuration)
	time.Sleep(1 * time.Microsecond)
	pd := GetPomodoroDuration(po)
	t.Log("\n", pd)
}

func TestSetPomodoroDuration(t *testing.T) {
	n := NewPomodoro()
	n.Start = GetCurrentTime()
	SetPomodoroDuration(n)
	t.Log("\n", n.Start, n.End)
}

func TestPomodoroTimer(t *testing.T) {
	tn := time.Now().Local()
	ti := Timer()
	if ti.Format(layout) != tn.Add(1*time.Second).Format(layout) {
		t.Fail()
		t.Log("\nExpected:", tn, "\nReceived:", ti)
	}
	t.Log("Succesufully added one second")
}

func TestFormatDate(t *testing.T) {
	date, time := FormatDate(2015, time.January, 1, 0, 0)
	t.Log("\nDate string:", date, "\ntime.Time format:", time)
	if date != "Jan 01 2015 at 00:00:01.000" {
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
	SetPomodoroDuration(n)

	expected := []string{"active", "Jan 01 2015 at 00:00:01.000", "Jan 01 2015 at 00:01:01.000"}
	got := FormatOutput(n)

	for i := 0; i < len(got); i++ {
		// INFO: godoc -http=":6060"
		// http://localhost:6060/ref/spec#Index_expressions
		if strings.Compare(expected[i], got[i]) != 0 {
			t.Fail()
			t.Log(strings.Compare(expected[i], got[i]))
		}
		// t.Log("\nExpected:", expected[i], "\nReceived:", got[i])
	}

	// Dont test this anymore, since it's effectively useless
	if len(expected) != len(got) {
		t.Fail()
		t.Log("\n Expected:", len(expected), "but got:", len(got))
	}
}
