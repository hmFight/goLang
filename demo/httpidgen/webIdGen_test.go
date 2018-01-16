package httpidgen

import (
	"testing"
)

func TestGenIncrId(t *testing.T) {
	go GenIncrId()
}
