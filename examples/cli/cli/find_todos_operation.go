// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/client/todos"

	"github.com/spf13/cobra"
)

// makeOperationTodosFindTodosCmd returns a cmd to handle operation findTodos
func makeOperationTodosFindTodosCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "findTodos",
		Short: ``,
		RunE:  runOperationTodosFindTodos,
	}

	if err := registerOperationTodosFindTodosParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosFindTodos uses cmd flags to call endpoint api
func runOperationTodosFindTodos(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewFindTodosParams()
	if err, _ := retrieveOperationTodosFindTodosLimitFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTodosFindTodosSinceFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationTodosFindTodosResult(appCli.Todos.FindTodos(params, nil)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationTodosFindTodosLimitFlag(m *todos.FindTodosParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("limit") {

		var limitFlagName string
		if cmdPrefix == "" {
			limitFlagName = "limit"
		} else {
			limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
		}

		limitFlagValue, err := cmd.Flags().GetInt32(limitFlagName)
		if err != nil {
			return err, false
		}
		m.Limit = &limitFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTodosFindTodosSinceFlag(m *todos.FindTodosParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("since") {

		var sinceFlagName string
		if cmdPrefix == "" {
			sinceFlagName = "since"
		} else {
			sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
		}

		sinceFlagValue, err := cmd.Flags().GetInt64(sinceFlagName)
		if err != nil {
			return err, false
		}
		m.Since = &sinceFlagValue

	}
	return nil, retAdded
}

// printOperationTodosFindTodosResult prints output to stdout
func printOperationTodosFindTodosResult(resp0 *todos.FindTodosOK, respErr error) error {
	if respErr != nil {

		var iResp interface{} = respErr
		defaultResp, ok := iResp.(*todos.FindTodosDefault)
		if !ok {
			return respErr
		}
		if defaultResp.Payload != nil {
			msgStr, err := json.Marshal(defaultResp.Payload)
			if err != nil {
				return err
			}
			fmt.Println(string(msgStr))
			return nil
		}

		return respErr
	}

	if resp0.Payload != nil {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationTodosFindTodosParamFlags registers all flags needed to fill params
func registerOperationTodosFindTodosParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTodosFindTodosLimitParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTodosFindTodosSinceParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTodosFindTodosLimitParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	limitDescription := ``

	var limitFlagName string
	if cmdPrefix == "" {
		limitFlagName = "limit"
	} else {
		limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
	}

	var limitFlagDefault int32 = 20

	_ = cmd.PersistentFlags().Int32(limitFlagName, limitFlagDefault, limitDescription)

	return nil
}
func registerOperationTodosFindTodosSinceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sinceDescription := ``

	var sinceFlagName string
	if cmdPrefix == "" {
		sinceFlagName = "since"
	} else {
		sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
	}

	var sinceFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sinceFlagName, sinceFlagDefault, sinceDescription)

	return nil
}
