# TrueNAS Go SDK Development

To generate go SDK from current openapi spec use this command:

```bash
$ openapi-generator-cli generate -i cfg/openapi.yaml -c cfg/config.yaml -o . -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
```

Workaround for TrueNAS content negotiation issue:

```bash
patch client.go < client.patch
```

## Code quality

* Install git pre-commit: [https://pre-commit.com/](https://pre-commit.com/) and run `pre-commit install` in project root
* Install commitizen and make commits with `git cz`:

```bash
npm install -g commitizen cz-conventional-changelog
echo '{ "path": "cz-conventional-changelog" }' > ~/.czrc
```