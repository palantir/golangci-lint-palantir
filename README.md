<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/golangci-lint-palantir"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

# golangci-lint-palantir
`golangci-lint-palantir` builds and publishes a custom build of [`golangci-lint`](https://github.com/golangci/golangci-lint)
that include the following custom linters:

* `compiles`: runs the `golangci-lint` `typecheck` check to ensure that all project code compiles

This repository also publishes the configuration file at [pluginconfig/golangci.yml], which defines default
configuration that enables the `compiles` linter and defines the default set of enabled linters.

The custom build of `golangci-lint` is built using the [`godel-distgo-asset-dist-golangci-lint`](https://github.com/palantir/godel-distgo-asset-dist-golangci-lint)
dister.
