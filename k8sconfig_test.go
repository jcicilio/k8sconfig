package main

import (
	"os"
	"testing"
)

func init(){
	configSetup()
}

func TestDefaultVariable(t *testing.T) {
	e := DEFAULTVALUE

	if k := getValue(); k != e {
		t.Errorf("default variable: got '%v' want '%v'",
			k, e)
	}
}

func TestEnvironmentVariable(t *testing.T) {
	e := ENVIRONMENTVALUE

	os.Setenv(PREFIX + "_" + NAME, e)
	if k := getValue(); k != e {
		t.Errorf("environment variable: got '%v' want '%v'",
			k, e)
	}

	os.Setenv(PREFIX + "_" + NAME, "")
}


func TestConfigurationFile(t *testing.T) {

	e := CONFIGFILEVALUE

	configReadConfigFile()

	if k := getValue(); k != e {
		t.Errorf("configuration file variable: got '%v' want '%v'",
			k, e)
	}
}