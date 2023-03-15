# CasaOS-MessageBus

[![Go Reference](https://pkg.go.dev/badge/github.com/IceWhaleTech/CasaOS-MessageBus.svg)](https://pkg.go.dev/github.com/IceWhaleTech/CasaOS-MessageBus) [![Go Report Card](https://goreportcard.com/badge/github.com/IceWhaleTech/CasaOS-MessageBus)](https://goreportcard.com/report/github.com/IceWhaleTech/CasaOS-MessageBus) [![goreleaser](https://github.com/IceWhaleTech/CasaOS-MessageBus/actions/workflows/release.yml/badge.svg)](https://github.com/IceWhaleTech/CasaOS-MessageBus/actions/workflows/release.yml) [![codecov](https://codecov.io/gh/IceWhaleTech/CasaOS-MessageBus/branch/main/graph/badge.svg?token=U4S4ZSZAL9)](https://codecov.io/gh/IceWhaleTech/CasaOS-MessageBus)

Message bus accepts events and actions from various sources and delivers them to subscribers.

See [openapi.yaml](./api/message_bus/openapi.yaml) for API specification.




## publish api to npm

### edit version in package.json

### run
```bash
yarn

yarn start
```

### publish

Manual publish
```bash
yarn publish
```

Auto publish
```bash 
git push origin dev**
```