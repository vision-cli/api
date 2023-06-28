package v1

// A plugin must accept a json serialized PluginRequest from stdin and return a json serialized PluginResponse
// specific to the command in the request. The two commands in api are 'usage' and 'run', and hence a plugin
// must return a json serialized PluginUsageResponse or PluginRunResponse

const (
	CommandUsage = "usage"
	CommandRun   = "run"
)

type PluginRequest struct {
	Command      string // v1 plugins must support 'usage' and 'run'
	Args         string // blank if the command is 'usage', otherwise the sub command e.g. create, delete, etc if the command is 'run'
	Placeholders string // blank if the command is 'usage', otherwise the placeholders structure built from the config file and flags
}

// Commands that a plugin must support:
// - usage
//   - args: none
//   - placeholders: none
//   - response: Unmarshalled into a PluginUsageResponse, showing the usage of the plugin
//	   all fields are mandatory. An error is thrown if a field is missing, however any additional fields returned are ignored
//     subcommands are a comma separated list of subcommands that the plugin supports
//
// - run
//   - args: the sub command e.g. create, delete, specific to the plugin
//   - placeholders: the placeholders structure
//   - response: Unmarshalled into a PluginRunResponse, showing the result of the plugin execution
//	   if a result is returned, the error is ignored. If an error is returned, the result is ignored
//	   if both are returned then an error is thrown

type PluginUsageResponse struct {
	Version     string
	Use         string
	Short       string
	Long        string
	Example     string
	Subcommands []string
}

type PluginRunResponse struct {
	Result string
	Error  string
}
