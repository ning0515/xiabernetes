package labels

import "strings"

type Labels interface {
	Get(label string) (value string)
}

type Set map[string]string

func (s Set) String() string {
	query := make([]string, 0, len(s))
	for key, value := range s {
		query = append(query, key+"="+value)
	}
	return strings.Join(query, ",")
}

func (s Set) Get(label string) string {
	return s[label]
}
