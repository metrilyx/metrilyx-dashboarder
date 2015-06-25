package v3

import (
	"testing"
)

func Test_NewPodPanel(t *testing.T) {
	nc := NewPodPanel()
	t.Logf("%s\n", nc.Id)
	nc = NewPodPanel("test_Id")
	t.Logf("%s\n", nc.Id)
	nc = NewPodPanel("test_Name", "Test ID")
	t.Logf("%s\n", nc.Id)
}
