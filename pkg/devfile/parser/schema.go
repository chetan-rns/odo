package parser

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/openshift/odo/pkg/devfile/versions"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

// SetDevfileJSONSchema returns the JSON schema for the given devfile apiVersion
func (d *DevfileCtx) SetDevfileJSONSchema() error {

	// Check if json schema is present for the given apiVersion
	jsonSchema, err := versions.GetDevfileJSONSchema(d.apiVersion)
	if err != nil {
		return err
	}
	d.jsonSchema = jsonSchema
	return nil
}

// ValidateDevfileSchema validate JSON schema of the provided devfile
func (d *DevfileCtx) ValidateDevfileSchema() error {

	var (
		schemaLoader   = gojsonschema.NewStringLoader(d.jsonSchema)
		documentLoader = gojsonschema.NewStringLoader(string(d.rawContent))
	)

	// Validate devfile with JSON schema
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return errors.Wrapf(err, "failed to validate devfile schema")
	}

	if !result.Valid() {
		errMsg := fmt.Sprintf("invalid devfile schema. errors :\n")
		for _, desc := range result.Errors() {
			errMsg = errMsg + fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf(errMsg)
	}

	// Sucessful
	glog.V(4).Info("validated devfile schema")
	return nil
}
