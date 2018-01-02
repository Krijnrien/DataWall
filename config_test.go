package main

import (
	"testing"
	"regexp"

	"github.com/stretchr/testify/assert"
	"GoReqHandler/config"
)

/**
 * Test if config file exists
 */
func TestConfigFileExists(t *testing.T) {
	assert.Equal(t, true, fileExists(), "The configuration file can not be found.")
}

/**
 * Test if config ip address value has a valid IPV4 format through regex.
 */
func TestConfigIPIsValid(t *testing.T) {
	cfg := *config.Get()
	match, _ := regexp.MatchString(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, cfg.IpAddress)
	assert.Equal(t, true, match, "Ip address does not have valid")
}

/**
 * Test whether config values are set by checking if string length is longer than 0
 */
func TestConfigKeyspaceHasValue(t *testing.T) {
	cfg := *config.Get()
	assert.Equal(t, true, len(cfg.Keyspace) > 0, "Keyspace string empty")
}