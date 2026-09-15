// Package fqrn implements the canonical MAP v2 resource-name grammar used at
// authorization boundaries.
package fqrn

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	namePattern   = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)
	idPattern     = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9-]*$`)
	regionPattern = regexp.MustCompile(`^[a-z][a-z0-9-]*-[a-z][a-z0-9-]*-[0-9]$`)
)

// Resource is a canonical non-regional or regional FQRN.
type Resource struct {
	Service string
	Org     string
	Project string
	Region  string
	Type    string
	ID      string
}

// Parse validates and parses a canonical FQRN. It deliberately accepts no
// legacy aliases or normalization: authorization identifiers must be exact.
func Parse(value string) (Resource, error) {
	if value == "" || strings.TrimSpace(value) != value || len(value) > 256 {
		return Resource{}, fmt.Errorf("invalid canonical FQRN")
	}
	parts := strings.Split(value, "/")
	if len(parts) != 5 && len(parts) != 6 {
		return Resource{}, fmt.Errorf("invalid canonical FQRN segment count: %d", len(parts))
	}
	for _, part := range parts {
		if part == "" || part == "." || part == ".." || strings.Contains(part, `\`) {
			return Resource{}, fmt.Errorf("invalid canonical FQRN segment")
		}
	}
	r := Resource{Service: parts[0], Org: parts[1], Project: parts[2]}
	if len(parts) == 5 {
		r.Type, r.ID = parts[3], parts[4]
	} else {
		r.Region, r.Type, r.ID = parts[3], parts[4], parts[5]
	}
	if !namePattern.MatchString(r.Service) || !namePattern.MatchString(r.Org) ||
		!namePattern.MatchString(r.Project) || !namePattern.MatchString(r.Type) ||
		!idPattern.MatchString(r.ID) {
		return Resource{}, fmt.Errorf("invalid canonical FQRN component")
	}
	if r.Region != "" && !regionPattern.MatchString(r.Region) {
		return Resource{}, fmt.Errorf("invalid canonical FQRN region")
	}
	if r.Service == "organization" && (r.Org != "global" || r.Project != "all" || r.Region != "" || r.Type != "organization") {
		return Resource{}, fmt.Errorf("invalid organization anchor")
	}
	if r.Service == "project" && (r.Org == "global" || r.Project != "all" || r.Region != "" || r.Type != "project" || r.ID == "all") {
		return Resource{}, fmt.Errorf("invalid project anchor")
	}
	if r.Canonical() != value {
		return Resource{}, fmt.Errorf("non-canonical FQRN")
	}
	return r, nil
}

func (r Resource) Canonical() string {
	if r.Region != "" {
		return strings.Join([]string{r.Service, r.Org, r.Project, r.Region, r.Type, r.ID}, "/")
	}
	return strings.Join([]string{r.Service, r.Org, r.Project, r.Type, r.ID}, "/")
}

func OrganizationAnchor(org string) (Resource, error) {
	return Parse("organization/global/all/organization/" + org)
}

func ProjectAnchor(org, project string) (Resource, error) {
	return Parse("project/" + org + "/all/project/" + project)
}

func (r Resource) IsOrganizationAnchor() bool {
	return r.Service == "organization" && r.Org == "global" && r.Project == "all" && r.Type == "organization"
}

func (r Resource) IsProjectAnchor() bool {
	return r.Service == "project" && r.Project == "all" && r.Type == "project"
}

// AssignmentApplies reports whether an assignment is on the exact target or
// on its canonical project or organization placement ancestor.
func AssignmentApplies(assignment, target Resource) bool {
	if assignment.Canonical() == target.Canonical() {
		return true
	}
	if assignment.IsOrganizationAnchor() {
		return target.Org != "global" && assignment.ID == target.Org
	}
	if assignment.IsProjectAnchor() {
		targetProject := target.Project
		if target.IsProjectAnchor() {
			targetProject = target.ID
		}
		return assignment.Org == target.Org && targetProject != "all" && assignment.ID == targetProject
	}
	return false
}
