package matchers

import "fmt"

// Matcher is a type that matches expected against actuals.
type Matcher interface {
	Match(actual interface{}) (resultValue interface{}, err error)
}

// NotMatcher accepts a matcher and will succeed if the specified matcher fails.
type NotMatcher struct {
	child Matcher
}

// Not returns a NotMatcher with the specified child matcher.
func Not(child Matcher) NotMatcher {
	return NotMatcher{
		child: child,
	}
}

func (m NotMatcher) Match(actual interface{}) (interface{}, error) {
	v, err := m.child.Match(actual)
	if err == nil {
		return nil, fmt.Errorf("match %#v", m.child)
	}

	return v, nil
}
