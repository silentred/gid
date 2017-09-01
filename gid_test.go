package gid

import "testing"

func TestGID(t *testing.T) {
	id := Get()
	t.Log(id)
}
