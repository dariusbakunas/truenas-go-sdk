# TrueNAS Go SDK Development

To generate go SDK from current openapi spec use this command:

```bash
$ openapi-generator-cli generate -i cfg/openapi.yaml -c cfg/config.yaml -o . -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
```