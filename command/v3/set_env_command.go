package v3

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v3/shared"
)

//go:generate counterfeiter . SetEnvActor

type SetEnvActor interface {
	CloudControllerAPIVersion() string
	SetEnvironmentVariableByApplicationNameAndSpace(appName string, spaceGUID string, envPair v3action.EnvironmentVariablePair) (v3action.Warnings, error)
}

type SetEnvCommand struct {
	RequiredArgs    flag.SetEnvironmentArgs `positional-args:"yes"`
	usage           interface{}             `usage:"CF_NAME set-env APP_NAME ENV_VAR_NAME ENV_VAR_VALUE"`
	relatedCommands interface{}             `related_commands:"v3-apps, env, v3-restart, set-running-environment-variable-group, set-staging-environment-variable-group, v3-stage, v3-unset-env"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       SetEnvActor
}

func (cmd *SetEnvCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, _, err := shared.NewClients(config, ui, true, "")
	if err != nil {
		return err
	}
	cmd.Actor = v3action.NewActor(ccClient, config, nil, nil)

	return nil
}

func (cmd SetEnvCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	appName := cmd.RequiredArgs.AppName
	cmd.UI.DisplayTextWithFlavor("Setting env variable {{.EnvVarName}} for app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":    appName,
		"EnvVarName": cmd.RequiredArgs.EnvironmentVariableName,
		"OrgName":    cmd.Config.TargetedOrganization().Name,
		"SpaceName":  cmd.Config.TargetedSpace().Name,
		"Username":   user.Name,
	})

	warnings, err := cmd.Actor.SetEnvironmentVariableByApplicationNameAndSpace(
		appName,
		cmd.Config.TargetedSpace().GUID,
		v3action.EnvironmentVariablePair{
			Key:   cmd.RequiredArgs.EnvironmentVariableName,
			Value: string(cmd.RequiredArgs.EnvironmentVariableValue),
		})
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	cmd.UI.DisplayText("TIP: Use 'cf v3-stage {{.AppName}}' to ensure your env variable changes take effect.", map[string]interface{}{
		"AppName": appName,
	})

	return nil
}