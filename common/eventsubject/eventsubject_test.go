package eventsubject

import (
	"testing"

	eventv1 "github.com/Mattilsynet/mapis/gen/go/event/v1"
)

func TestBuildParse(t *testing.T) {
	want := "event.iam.acme.all.user.create.alice"
	got, err := Build(Subject{Subsystem: "iam", Org: "acme", Project: "all", Kind: "user", Operation: Create, Keys: []string{"alice"}})
	if err != nil || got != want {
		t.Fatalf("Build() = %q, %v; want %q", got, err, want)
	}
	parsed, err := Parse(got)
	if err != nil || parsed.Subsystem != "iam" || parsed.Org != "acme" || parsed.Kind != "user" || parsed.Operation != Create || parsed.Keys[0] != "alice" {
		t.Fatalf("Parse() = %#v, %v", parsed, err)
	}
}

func TestBuildFromFQRN(t *testing.T) {
	tests := []struct {
		fqrn string
		op   eventv1.EventSpec_Operation
		want string
	}{
		{"iam/acme/all/user/alice", eventv1.EventSpec_OPERATION_CREATE, "event.iam.acme.all.user.create.alice"},
		{"iam/global/all/permission/instances:create", eventv1.EventSpec_OPERATION_UPDATE, "event.iam.global.all.permission.update.instances-create"},
		{"dns/acme/frontend/zone/Example.COM", eventv1.EventSpec_OPERATION_UPDATE, "event.dns.acme.frontend.zone.update.example-com"},
		{"compute/acme/frontend/eu-west-1/compute/vm-01", eventv1.EventSpec_OPERATION_DELETE, "event.compute.acme.frontend.compute.delete.vm-01"},
	}
	for _, test := range tests {
		got, err := BuildFromFQRN(test.fqrn, test.op)
		if err != nil || got != test.want {
			t.Errorf("BuildFromFQRN(%q) = %q, %v; want %q", test.fqrn, got, err, test.want)
		}
	}
}

func TestBuildFromFQRNRejectsEmptyNormalizedKey(t *testing.T) {
	if _, err := BuildFromFQRN("iam/acme/all/user/...", eventv1.EventSpec_OPERATION_CREATE); err == nil {
		t.Fatal("expected FQRN key that normalizes to empty to be rejected")
	}
}

func TestRejectLegacyAndRead(t *testing.T) {
	if _, err := Parse("event.acme.user.change.alice"); err == nil {
		t.Fatal("legacy subject accepted")
	}
	if _, err := BuildFromFQRN("iam/acme/all/user/alice", eventv1.EventSpec_OPERATION_READ); err == nil {
		t.Fatal("READ operation accepted")
	}
}

func TestFilter(t *testing.T) {
	got, err := Filter("iam", "", "all", "rolebinding", "")
	if err != nil || got != "event.iam.*.all.rolebinding.*.>" {
		t.Fatalf("Filter() = %q, %v", got, err)
	}
}
