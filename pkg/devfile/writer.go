package devfile

import (
	"encoding/json"
	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// WriteJsonDevfile creates a devfile.json file
func (d *DevfileObj) WriteJsonDevfile() error {

	// Encode data into JSON format
	jsonData, err := json.MarshalIndent(d.Data, "", "  ")
	if err != nil {
		return errors.Wrapf(err, "failed to marshal devfile object into json")
	}

	// Write to devfile.json
	fs := d.Ctx.GetFs()
	err = fs.WriteFile(OutputDevfileJsonPath, jsonData, 0644)
	if err != nil {
		return errors.Wrapf(err, "failed to create devfile json file")
	}

	// Successful
	glog.V(4).Infof("devfile json created at: '%s'", OutputDevfileJsonPath)
	return nil
}

// WriteYamlDevfile creates a devfile.yaml file
func (d *DevfileObj) WriteYamlDevfile() error {

	// Encode data into YAML format
	yamlData, err := yaml.Marshal(d.Data)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal devfile object into yaml")
	}

	// Write to devfile.yaml
	fs := d.Ctx.GetFs()
	err = fs.WriteFile(OutputDevfileYamlPath, yamlData, 0644)
	if err != nil {
		return errors.Wrapf(err, "failed to create devfile yaml file")
	}

	// Successful
	glog.V(4).Infof("devfile yaml created at: '%s'", OutputDevfileYamlPath)
	return nil
}
