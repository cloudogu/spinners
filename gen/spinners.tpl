// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots, using data from:
// {{ .URL }}
package spinners

import (
  "io"
  "fmt"
)

type spinnerFactory func(writer io.Writer) *Spinner

var factories = map[string]spinnerFactory{ {{ range $Name, $Configuration := .Spinners }}
  "{{ $Name }}": New{{ $Name | Title }}Spinner,{{ end }}
}

// NewSpinnerByName returns the spinner with the given name
func NewSpinnerByName(name string, writer io.Writer) (*Spinner, error) {
  factory, found := factories[name]
  if ! found {
    return nil, fmt.Errorf("could not find spinner with name %s", name)
  }

  return factory(writer), nil
}

{{ range $Name, $Configuration := .Spinners }}
// New{{ $Name | Title }}Spinner creates a new {{ $Name }} spinner
func New{{ $Name | Title }}Spinner(writer io.Writer) *Spinner {
  configuration := Configuration{Interval: {{ $Configuration.Interval }}, Frames: []string{ {{ range $Configuration.Frames }}
    "{{ . | Escape }}",{{ end }}
  }}
  return NewSpinner(configuration, writer)
}
{{ end }}