<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/golangci-lint-palantir"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

# golangci-lint-palantir
`golangci-lint-palantir` builds and publishes a custom build of [`golangci-lint`](https://github.com/golangci/golangci-lint)
that include the following custom linters:

* `compiles`: runs the `golangci-lint` `typecheck` check to ensure that all project code compiles
* `safelogging`: runs the [`safelogging` linter](https://github.com/palantir/safe-logging-go) to verify that unsafe values are not logged

This repository also publishes the configuration file at [pluginconfig/golangci.yml], which configures the following:

* Declares the `compiles` and `safelogging` linters to make them available for use
* Sets the enabled checks to be the golangci-lint defaults (minus staticheck), compiles, revive, and unconvert
* Configures the `revive` check to disable the check requiring package-level comments

The custom build of `golangci-lint` is built using the [`godel-distgo-asset-dist-golangci-lint`](https://github.com/palantir/godel-distgo-asset-dist-golangci-lint)
dister.
