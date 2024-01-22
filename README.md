# ![logo](./images/vision-logo.svg "Vision") &nbsp; Vision Plugin Api

Vision plugins must implement one of the versions of the api. Vision will
be backward compatible so if the vision-cli tool is on v2, it will continue
to support v1. This does not mean the api spec is backward compatible, only
that the tool will detect that a plugin or project is on a previous version
and will execute the appropriate code.

Plugins must return the message structs which will be serialized into json
during communication between the vision-cli and the plugins.
Standard vision plugins including project, service, graphql, gateway, etc also implement this interface
