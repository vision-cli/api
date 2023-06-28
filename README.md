# ![logo](./images/vision-logo.svg "Vision") &nbsp; Vision Plugin Api

Vision plugins must implement one of the versions of the api. Vision will
be backward compatible so if the vision-cli tool is on v2, it will continue
to support v1.

Plugins must handle the message structs which will be serialized into json
duing communication between the vision-cli and the plugins.
Standard vision plugins including project, service, graphql, gateway, etc also implement this interface
