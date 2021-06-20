# Sample Elvish plugin

Elvish plugins are Elvish modules written in Go. They can be imported like
modules written in Elvish.

Plugins are only supported on platforms supported by Go's
[plugin](https://pkg.go.dev/plugin) package.

To create an Elvish plugin, first follow these steps to create the source tree:

1.  Initialize a new Go module (replace `github.com/elves/sample-plugin` with
    your repo path):

    ```sh
    go mod init github.com/elves/sample-plugin
    ```

2.  Add Go source files to form a `main` package. The package should export a
    variable named `Ns` that represents an Elvish namespace.

    You also need to add a `main` function, but it will not be executed, so you
    can leave it empty.

    See `main.go` in this repository for an example.

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

Now you have the source files ready, build and use the plugin with the following
steps:

1.  Build the plugin:

    ```sh
    go build -buildmode=plugin
    ```

    This will create a new `${name}.so` in the current directory, `${name}`
    being the same as the directory name.

2.  Build Elvish with plugin support, at exactly the same version used by the
    plugin:

    ```sh
    CGO_ENABLED=1 go install src.elv.sh/cmd/elvish@$(go list -f '{{.Version}}' -m src.elv.sh)
    ```

    Elvish version of the command:

    ```sh
    E:CGO_ENABLED=1 go install src.elv.sh/cmd/elvish@(go list -f '{{.Version}}' -m src.elv.sh)
    ```

3.  Scripts may now use the `.so` file as a module:

    ```sh
    elvish sample-script.elv
    # Outputs "bar"
    ```

Whenever you rebuild Elvish or the plugin, you have to make sure that their
versions are synchronized, and built with the same version of Go compiler.
