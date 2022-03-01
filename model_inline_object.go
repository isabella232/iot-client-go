/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
import (
	"os"
)
// InlineObject struct for InlineObject
type InlineObject struct {
	// If false, wait for the full OTA process, until it gets a result from the device
	Async bool `json:"async,omitempty"`
	// Binary expire time in minutes, default 10 mins
	ExpireInMins int32 `json:"expire_in_mins,omitempty"`
	// OTA file
	OtaFile *os.File `json:"ota_file"`
}
