package contextimpl

import (
	"fmt"
	"testing"
)

func TestBackgroundNotTODO(t *testing.T) {
	todo := fmt.Sprint(TODO())
	bg := fmt.Sprint(Background())

	if todo == bg {
		t.Errorf("TODO and background are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	if err := ctx.Err(); err != nil {
		t.Errorf("Err() should be nil first, got %v", err)
	}

	cancel()
	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("Error should be equal to canceled: %v", err)
	}
}
