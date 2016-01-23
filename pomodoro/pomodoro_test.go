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



func TestPomodoroTimer(t *testing.T) {
	tn := time.Now().Local()
	ti := Timer(1 * time.Second)
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
	n := NewPomodoro()
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

	_, n.Start = FormatDate(2015, time.January, 1, 0, 0)
	SetPomodoroDuration(n)

	expected := []string{"active", "Jan 01 2015 at 00:00:01.000", "Jan 01 2015 at 00:25:01.000"}
	got := FormatOutput(n)

	for i := 0; i < len(got); i++ {
		if strings.Compare(expected[i], got[i]) != 0 {
			t.Fail()
			t.Log(strings.Compare(expected[i], got[i]))
		}
		 t.Log("\nExpected:", expected[i], "\nReceived:", got[i])
	}

	// Dont test this anymore, since it's effectively useless
	if len(expected) != len(got) {
		t.Fail()
		t.Log("\n Expected:", len(expected), "but got:", len(got))
	}
}
