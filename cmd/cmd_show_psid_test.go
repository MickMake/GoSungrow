package cmd

import (
	"errors"
	"github.com/MickMake/GoSungrow/iSolarCloud"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowPsData_MissingDeviceType(t *testing.T) {
	// Create a new CmdShow object
	cmdShow := NewCmdShow()

	// Set up the cmds global variable for the test with a properly initialized SunGrow object
	cmds = Cmds{Api: &CmdApi{}}
	cmds.Api.SunGrow = iSolarCloud.NewSunGro("", t.TempDir())
	err := cmds.Api.SunGrow.Init()
	assert.NoError(t, err, "SunGrow Init should not fail")


	// Create a dummy cobra command
	dummyCmd := &cobra.Command{}

	// Call the function with arguments, omitting the device_type.
	// This should cause our new error to be returned.
	args := []string{"12345", "20230101", "20230102", "5"}
	err = cmdShow.funcPsData(dummyCmd, args)

	// Assert that an error occurred
	assert.Error(t, err, "funcPsData should return an error when device_type is missing")

	// Assert that the error is the specific one we expect
	assert.Equal(t, errors.New("device_type is required when specifying a date range"), err)
}
