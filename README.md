<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/golangci-lint-palantir"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

# golangci-lint-palantir
`golangci-lint-palantir` builds and publishes a custom build of [`golangci-lint`](https://github.com/golangci/golangci-lint)
that include the following custom linters:

* `safelogging`: runs the [`safelogging` linter](https://github.com/palantir/safe-logging-go) to verify that unsafe values are not logged

This repository also publishes the configuration file at [pluginconfig/golangci.yml], which configures the following:

* Declares the `safelogging` linter as a module-based linter to make it available for use
* Sets the enabled checks to be the golangci-lint defaults (minus staticcheck), revive, and unconvert
* Configures the `revive` check to disable the check requiring package-level comments
* Adds an exclusion rule to exclude "inline: cannot inline: type parameter inference is not yet supported" errors, which
  are produced by the experimental govet inline analyzer and are not actionable

The custom build of `golangci-lint` is built using the [`godel-distgo-asset-dist-golangci-lint`](https://github.com/palantir/godel-distgo-asset-dist-golangci-lint)
dister.

The version of `golangci-lint` that is used as the base for the custom build of `golangci-lint` is specified in the
"golangci-lint-version" field of the `golangci-lint` dister configuration in [dist-plugin.yml](godel/config/dist-plugin.yml).
The godel "generate" ensures that the value of this field matches the value of the golangci-lint module specified in
[go.mod](go.mod). 
