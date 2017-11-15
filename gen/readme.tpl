# Spinners

Go implementation of the [sindresorhus/cli-spinners](https://github.com/sindresorhus/cli-spinners).

## Usage

```go
# create new dots spinner
spinner := spinners.NewDotsSpinner(os.Stdout)

# start spinner in a separate go routing
spinner.Start("waiting for something long ...")

# do something
...

# stop spinner and reset the line
spinner.Stop()

# or get a spinner by its name
spinner, err := spinners.NewSpinnerByName("dots", os.Stdout)
if err != nil {
    log.Fatal(err)
}
```

## Build

The spinners.go can be regenerated using the `go generate` command. The command will load the latest spinners.json from
master branch of [sindresorhus/cli-spinners](https://github.com/sindresorhus/cli-spinners) and regenerate the
spinners.go and the README.md file.

## Available Spinners
| Name | Interval | Frames |
| ---- | -------- | ------ |{{ range $Name, $Configuration := .Spinners }}
| {{ $Name }} | {{ $Configuration.Interval }} | {{ range $Configuration.Frames }}{{ . }} {{ end }} |{{ end }}