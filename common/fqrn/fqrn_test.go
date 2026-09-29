package fqrn

import "testing"

func TestAssignmentApplies(t *testing.T) {
	target, _ := Parse("secret/acme/frontend/secret/db")
	tests := []struct {
		assignment string
		want       bool
	}{
		{"secret/acme/frontend/secret/db", true},
		{"project/acme/all/project/frontend", true},
		{"organization/global/all/organization/acme", true},
		{"project/acme/all/project/backend", false},
		{"organization/global/all/organization/other", false},
		{"secret/acme/frontend/secret/other", false},
	}
	for _, tt := range tests {
		assignment, err := Parse(tt.assignment)
		if err != nil {
			t.Fatalf("Parse(%q): %v", tt.assignment, err)
		}
		if got := AssignmentApplies(assignment, target); got != tt.want {
			t.Errorf("AssignmentApplies(%q) = %v, want %v", tt.assignment, got, tt.want)
		}
	}
}

func TestParseRejectsAliasesAndMalformedResources(t *testing.T) {
	for _, value := range []string{
		"orgs/acme",
		"secret.acme.frontend.secret.db",
		"secret/acme//secret/db",
		"secret/acme/frontend/secret/../db",
		" project/acme/all/project/frontend",
	} {
		if _, err := Parse(value); err == nil {
			t.Errorf("Parse(%q) unexpectedly succeeded", value)
		}
	}
}

func TestParseAcceptsProjectServiceActivationResource(t *testing.T) {
	value := "project/acme/frontend/projectserviceactivation/dns"
	resource, err := Parse(value)
	if err != nil {
		t.Fatalf("Parse(%q): %v", value, err)
	}
	if resource.Service != "project" || resource.Org != "acme" || resource.Project != "frontend" || resource.Type != "projectserviceactivation" || resource.ID != "dns" {
		t.Fatalf("resource = %+v", resource)
	}
	if resource.Canonical() != value {
		t.Fatalf("Canonical() = %q, want %q", resource.Canonical(), value)
	}
}
