package internal

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	tmr := NewTimer()
	if tmr == nil {
		t.Fatal("NewTimer() returned nil")
	}
}

func TestTimerElapsed(t *testing.T) {
	tmr := NewTimer()

	// Sleep briefly
	time.Sleep(10 * time.Millisecond)

	elapsed := tmr.Elapsed()
	if elapsed < 10*time.Millisecond {
		t.Errorf("Elapsed() = %v, want >= 10ms", elapsed)
	}
}

func TestTimerRemaining(t *testing.T) {
	tmr := NewTimer()

	remaining := tmr.Remaining()
	if remaining <= 0 || remaining > Timeout {
		t.Errorf("Remaining() = %v, want > 0 and <= %v", remaining, Timeout)
	}
}

func TestTimerIsTimedOut(t *testing.T) {
	tmr := NewTimer()

	if tmr.IsTimedOut() {
		t.Error("IsTimedOut() = true immediately after creation")
	}
}

func TestTimerIsTTY(t *testing.T) {
	tmr := NewTimer()

	// Just verify it doesn't panic and returns a boolean
	_ = tmr.IsTTY()
}

func TestTimerShowProgress_NoTTY(t *testing.T) {
	tmr := NewTimer()
	tmr.isTTY = false

	// Should not panic when not a TTY
	tmr.ShowProgress()
}

func TestTimerClearProgress_NoTTY(t *testing.T) {
	tmr := NewTimer()
	tmr.isTTY = false

	// Should not panic when not a TTY
	tmr.ClearProgress()
}

func TestTimerShowCompletion_NoTTY(t *testing.T) {
	tmr := NewTimer()
	tmr.isTTY = false

	// Should not panic when not a TTY
	tmr.ShowCompletion(1 * time.Second)
}

// TestTimeoutConstant validates spec requirement of 5-minute timeout.
func TestTimeoutConstant(t *testing.T) {
	if Timeout != 5*time.Minute {
		t.Errorf("Timeout = %v, want 5 minutes", Timeout)
	}
}

func TestProgressWidthConstant(t *testing.T) {
	if ProgressWidth != 20 {
		t.Errorf("ProgressWidth = %d, want 20", ProgressWidth)
	}
}
