package raptor_test

import (
	"testing"

	"github.com/raptorbox/raptor-sdk-go"
)

//
// func TestMain(m *testing.M) {
// 	log.SetLevel(log.DebugLevel)
// 	os.Exit(m.Run())
// }

func TestRaptor(t *testing.T) {
	raptor.New("http://raptor.local")
}
