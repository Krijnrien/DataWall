package cassandra

import(
	"testing"

	//

	"github.com/stretchr/testify/assert"
)

/**GetDevices
 * Test if given limit is
 */
func TestGetDevicesRandomLimit(t *testing.T) {
	var locList = GetDevices(50)
	assert.Equal(t, true, len(locList) < 50, "More results were returned than given limit!")
}

/**GetDevices
 * Test if given limit is
 */
func TestGetDevicesNegativeLimit(t *testing.T) {
	// GetDevices receives uint type, limit must be positive
	var locList = GetDevices(-5)
	assert.Equal(t, true, len(locList) < 50, "More results were returned than given limit!")
}