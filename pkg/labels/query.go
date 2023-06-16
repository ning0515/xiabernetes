package labels

import (
	"fmt"
	"strings"
)

type Query interface {
	Matches(Labels) bool
	String() string
}

type hasTerm struct {
	label, value string
}
type notHasTerm struct {
	label, value string
}

func (h *hasTerm) Matches(labels Labels) bool {
	return labels.Get(h.label) == h.value
}

func (h *hasTerm) String() string {
	return fmt.Sprintf("%v=%v", h.label, h.value)
}

func (n *notHasTerm) Matches(labels Labels) bool {
	return labels.Get(n.label) != n.value
}

func (n *notHasTerm) String() string {
	return fmt.Sprintf("%v!=%v", n.label, n.value)
}

type andTerm []Query

func (a andTerm) String() string {
	var result []string
	for _, v := range a {
		result = append(result, v.String())
	}
	return strings.Join(result, ",")
}

func (a andTerm) Matches(labels Labels) bool {
	for _, v := range a {
		if !v.Matches(labels) {
			return false
		}
	}
	return true
}

func splitOps(part, op string) (key, value string, ok bool) {
	slice := strings.Split(part, op)
	if len(slice) == 2 {
		return slice[0], slice[1], true
	}
	return "", "", false
}
func ParseQuery(query string) Query {
	var items []Query
	parts := strings.Split(query, ",")
	for _, part := range parts {
		if part == "" {
			continue
		}
		if key, value, ok := splitOps(part, "!="); ok {
			items = append(items, &notHasTerm{label: key, value: value})
		} else if key, value, ok := splitOps(part, "=="); ok {
			items = append(items, &hasTerm{label: key, value: value})
		} else if key, value, ok := splitOps(part, "="); ok {
			items = append(items, &hasTerm{label: key, value: value})
		} else {
			fmt.Println("invalid query")
			return nil
		}
	}
	return andTerm(items)
}
