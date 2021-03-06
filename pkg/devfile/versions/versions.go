package versions

import (
	"reflect"

	v100 "github.com/openshift/odo/pkg/devfile/versions/1.0.0"
)

// SupportedApiVersions stores the supported devfile API versions
type supportedApiVersion string

// Supported devfile API versions in odo
const (
	apiVersion100 supportedApiVersion = "1.0.0"
)

// List of supported devfile API versions
var supportedApiVersionsList = []supportedApiVersion{apiVersion100}

// ------------- Init functions ------------- //

// apiVersionToDevfileStruct maps supported devfile API versions to their corresponding devfile structs
var apiVersionToDevfileStruct map[supportedApiVersion]reflect.Type

// Initializes a map of supported devfile api versions and devfile structs
func init() {
	apiVersionToDevfileStruct = make(map[supportedApiVersion]reflect.Type)
	apiVersionToDevfileStruct[apiVersion100] = reflect.TypeOf(v100.Devfile100{})
}

// Map to store mappings between supported devfile API versions and respective devfile JSON schemas
var devfileApiVersionToJSONSchema map[supportedApiVersion]string

// init initializes a map of supported devfile apiVersions with it's respective devfile JSON schema
func init() {
	devfileApiVersionToJSONSchema = make(map[supportedApiVersion]string)
	devfileApiVersionToJSONSchema[apiVersion100] = v100.JsonSchema100
}
