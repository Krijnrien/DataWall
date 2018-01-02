package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/**ValidPort
 * Test if given valid port is returned as true (valid)
 */
func TestValidPort(t *testing.T) {
	assert.Equal(t, true, ValidPort(25565), "Valid port returned as false (invalid)!")
}

/**ValidPort
 * Test if given invalid port is returned as false (invalid)
 */
func TestInvalidPort(t *testing.T) {
	assert.Equal(t, false, ValidPort(624563464356), "Invalid port returned as true (valid)!")
}