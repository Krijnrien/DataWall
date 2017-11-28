package main

import (
	"net"
	"strconv"
	"time"
)

func ValidToken(token string) bool {
	return len(token) > 0
}

func ValidPort(port int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer ln.Close()
	return err == nil
}

/**
* Timer to repeat func every given amount of time.
* @param interval in whole seconds.
* @param function name to repeat every interval tick
 */
func DoEvery(interval time.Duration, repeatFunction func(time.Time)) {
	for currentTime := range time.Tick(interval) {
		repeatFunction(currentTime)
	}
}

/**
 * The function that checks if []Device slice contains Device sample basing on Hash comparison
 * @param Slice of devices
 * @param Device to check
 * If item exists in slice, it's index is returned
 * Otherwise -1 is returned
 */
func Contains(list []Device, device Device) int {
	for index, sample := range list {
		if sample.Hash == device.Hash {
			return index
		}
	}
	return -1
}
