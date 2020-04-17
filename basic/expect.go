package basic

import (
	"testing"
)

func expect(expression bool, t *testing.T) {
	if !expression {
		t.Errorf("wrong expression result for %v", expression)
	}

}
