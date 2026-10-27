package user

import (
	osuser "os/user"
	"strings"
	"testing"
)

func TestRenderGroupTableSkipsUnresolvedMemberships(t *testing.T) {
	got := renderGroupTable([]*osuser.Group{
		nil,
		{Gid: "1000", Name: "staff"},
	})

	if !strings.Contains(got, "staff") {
		t.Fatalf("expected resolved group to remain visible, got %q", got)
	}
}

func TestRenderGroupTableWithNoResolvedMemberships(t *testing.T) {
	if got := renderGroupTable([]*osuser.Group{nil}); got != "No resolvable group details available" {
		t.Fatalf("expected an honest unavailable-details message, got %q", got)
	}
}
