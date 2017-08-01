package atoms

import (
	"os"
	"testing"

	"github.com/missionMeteora/journaler"
)

const (
	testErrInvalidValueFmt = "Invalid value, expected %v and received %v"
	testErrInvalidTypeFmt  = "Invalid type, expected %T and received %T"
	testErrInvalidSwapFmt  = "Swapped successfully when should have failed"
)

func TestMain(m *testing.M) {
	journaler.Notification("Running atoms test suite")
	sc := m.Run()
	os.Exit(sc)
}
