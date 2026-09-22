// Package eventsubject implements the canonical MAP v2 domain-event subject
// contract. Domain code should use this package instead of formatting or
// splitting event subjects directly.
package eventsubject

import (
	"fmt"
	"strings"
	"unicode"

	eventv1 "github.com/Mattilsynet/mapis/gen/go/event/v1"
)

const Prefix = "event"

type Operation string

const (
	Create Operation = "create"
	Update Operation = "update"
	Delete Operation = "delete"
)

// Subject is a concrete canonical domain-event subject.
type Subject struct {
	Subsystem string
	Org       string
	Project   string
	Kind      string
	Operation Operation
	Keys      []string
}

// Build validates and renders a concrete subject.
func Build(subject Subject) (string, error) {
	parts := []string{Prefix, subject.Subsystem, subject.Org, subject.Project, subject.Kind, string(subject.Operation)}
	parts = append(parts, subject.Keys...)
	if len(subject.Keys) == 0 {
		return "", fmt.Errorf("event subject requires at least one key")
	}
	for index, part := range parts {
		if err := validateToken(part, false); err != nil {
			return "", fmt.Errorf("invalid event subject token %d: %w", index, err)
		}
	}
	if !subject.Operation.Valid() {
		return "", fmt.Errorf("invalid event operation %q", subject.Operation)
	}
	return strings.Join(parts, "."), nil
}

// Parse strictly parses a concrete canonical subject. Legacy subjects and
// wildcard filters are rejected.
func Parse(value string) (Subject, error) {
	if value == "" || strings.TrimSpace(value) != value {
		return Subject{}, fmt.Errorf("invalid event subject")
	}
	parts := strings.Split(value, ".")
	if len(parts) < 7 || parts[0] != Prefix {
		return Subject{}, fmt.Errorf("expected event.<subsystem>.<org>.<project>.<kind>.<operation>.<key...>")
	}
	for index, part := range parts {
		if err := validateToken(part, false); err != nil {
			return Subject{}, fmt.Errorf("invalid event subject token %d: %w", index, err)
		}
	}
	subject := Subject{
		Subsystem: parts[1], Org: parts[2], Project: parts[3], Kind: parts[4],
		Operation: Operation(parts[5]), Keys: append([]string(nil), parts[6:]...),
	}
	if !subject.Operation.Valid() {
		return Subject{}, fmt.Errorf("invalid event operation %q", subject.Operation)
	}
	return subject, nil
}

// FromFQRN derives ownership, scope, kind, and the first key from an FQRN.
// FQRNs remain the authoritative identity; subject keys are normalized and
// can therefore be lossy.
func FromFQRN(value string, operation eventv1.EventSpec_Operation, additionalKeys ...string) (Subject, error) {
	parts := strings.Split(strings.TrimSpace(value), "/")
	if len(parts) != 5 && len(parts) != 6 {
		return Subject{}, fmt.Errorf("invalid FQRN: expected 5 or 6 segments, got %d", len(parts))
	}
	for index, part := range parts {
		if strings.TrimSpace(part) == "" {
			return Subject{}, fmt.Errorf("invalid FQRN: empty segment %d", index)
		}
	}
	kindIndex := 3
	if len(parts) == 6 {
		kindIndex = 4
	}
	op, err := FromProtoOperation(operation)
	if err != nil {
		return Subject{}, err
	}
	keys := []string{NormalizeKey(parts[kindIndex+1])}
	for _, key := range additionalKeys {
		keys = append(keys, NormalizeKey(key))
	}
	for index, key := range keys {
		if key == "" {
			return Subject{}, fmt.Errorf("invalid FQRN: key %d normalizes to an empty token", index)
		}
	}
	subject := Subject{
		Subsystem: strings.ToLower(parts[0]),
		Org:       strings.ToLower(parts[1]),
		Project:   strings.ToLower(parts[2]),
		Kind:      strings.ToLower(parts[kindIndex]),
		Operation: op,
		Keys:      keys,
	}
	if _, err := Build(subject); err != nil {
		return Subject{}, err
	}
	return subject, nil
}

func BuildFromFQRN(value string, operation eventv1.EventSpec_Operation, additionalKeys ...string) (string, error) {
	subject, err := FromFQRN(value, operation, additionalKeys...)
	if err != nil {
		return "", err
	}
	return Build(subject)
}

func FromProtoOperation(operation eventv1.EventSpec_Operation) (Operation, error) {
	switch operation {
	case eventv1.EventSpec_OPERATION_CREATE:
		return Create, nil
	case eventv1.EventSpec_OPERATION_UPDATE:
		return Update, nil
	case eventv1.EventSpec_OPERATION_DELETE:
		return Delete, nil
	default:
		return "", fmt.Errorf("protobuf operation %s cannot be published as a domain event", operation.String())
	}
}

func (operation Operation) Valid() bool {
	return operation == Create || operation == Update || operation == Delete
}

// NormalizeKey converts an identity component to one NATS-safe token. The
// original value must remain available in Event.metadata.fqrn.
func NormalizeKey(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	var out strings.Builder
	lastDash := false
	for _, r := range value {
		valid := unicode.IsLower(r) || unicode.IsDigit(r) || r == '-' || r == '_'
		if valid {
			out.WriteRune(r)
			lastDash = r == '-'
			continue
		}
		if !lastDash {
			out.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(out.String(), "-")
}

// Filter builds a canonical wildcard filter. Empty fields mean `*`. Kind and
// operation may also be explicit. The returned filter matches any key tail.
func Filter(subsystem, org, project, kind string, operation Operation) (string, error) {
	parts := []string{Prefix, subsystem, org, project, kind, string(operation)}
	for index := 1; index < len(parts); index++ {
		if parts[index] == "" {
			parts[index] = "*"
			continue
		}
		if index == 5 && parts[index] != "*" && !operation.Valid() {
			return "", fmt.Errorf("invalid event operation %q", operation)
		}
		if err := validateToken(parts[index], true); err != nil {
			return "", err
		}
	}
	return strings.Join(parts, ".") + ".>", nil
}

func validateToken(value string, wildcard bool) error {
	if wildcard && value == "*" {
		return nil
	}
	if value == "" || strings.TrimSpace(value) != value || strings.ContainsAny(value, ".*>/ ") {
		return fmt.Errorf("invalid token %q", value)
	}
	if strings.ToLower(value) != value {
		return fmt.Errorf("token must be lowercase: %q", value)
	}
	return nil
}
