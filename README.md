# Sample Elvish plugin

Elvish plugins are Elvish modules written in Go. They can be imported like
modules written in Elvish.

Plugins are only supported on platforms supported by Go's
[plugin](https://pkg.go.dev/plugin) package.

## Writing a plugin

To create an Elvish plugin, follow these steps:

1.  Initialize a new Go module (replace `github.com/elves/sample-plugin` with
    your repo path):

    ```sh
    go mod init github.com/elves/sample-plugin
    ```

2.  Add Go source files to form a `main` package. The package should export a
    variable named `Ns` that represents an Elvish namespace.

    See `main.go` in this repository for an example.

    In general, you can reference the various [builtin
    modules](https://src.elv.sh/pkg/eval/mods) for how to build the namespace
    and how Go functions are exported to Elvish.

3.  Add the latest commit of Elvish as a dependency:

    ```sh
    go get src.elv.sh@master
    ```

    This step is required for now, because the last released version of Elvish
    does not contain plugin support yet.

4.  Get the `go.mod` and `go.sum` files into a buildable state:

    ```sh
    go mod tidy
    ```

## Building the plugin

With the source files ready, build the plugin with the following steps:

1.  Build Elvish with plugin support, at exactly the same version used by the
    plugin:

    ```sh
    CGO_ENABLED=1 go install src.elv.sh/cmd/elvish@$(go list -f '{{.Version}}' -m src.elv.sh)
    ```

    Elvish version of the command:

    ```sh
    E:CGO_ENABLED=1 go install src.elv.sh/cmd/elvish@(go list -f '{{.Version}}' -m src.elv.sh)
    ```

2.  Build the plugin:

    ```sh
    go build -buildmode=plugin
    ```

    This will create a new `${name}.so` in the current directory, `${name}`
    being the same as the directory name.

Alternatively, if you have a clone of the Elvish repository, you can also use
these steps:

1.  Build Elvish with plugin support:

    ```sh
    cd /path/to/elvish/repository
    git checkout version-encoded-in-plugin-go.mod
    make get ELVISH_PLUGIN_SUPPORT=1
    ```

2.  Build the plugin:

    ```sh
    go build -buildmode=plugin -trimpath
    ```

    The `-trimpath` flag is needed to be consistent with the Elvish binary built
    in step 1.

After building the plugin, scripts may now use the `.so` file as a module:

```sh
elvish sample-script.elv
# Outputs "bar"
```

Whenever you rebuild Elvish or the plugin, you have to make sure that their
versions are synchronized, and built with the same version of Go compiler.
