package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashReference(t *testing.T) {
	assert.False(t, HashReference("hola"))
	assert.False(t, HashReference("ho$la"))
	assert.False(t, HashReference("hol{a}"))
	assert.False(t, HashReference("h$}{ola"))
	values := []string{
		"mongodb://${parent.database.hostname}/db",
		"mongodb://${parent.database}/db",
		"mongodb://${parent}/db",
		"mongodb://${parent.database.hostname}",
	}
	for _, value := range values {
		assert.True(t, HashReference(value))
		rs := re.FindStringSubmatch(value)
		fmt.Println(value)
		for _, r := range rs {
			fmt.Printf(" - %s\n", r)
		}
	}
}
