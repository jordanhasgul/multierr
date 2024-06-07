# multierr

## Overview

`multierr` is a Go module for aggregating multiple errors into a single error. It is fully 
compatible with the [errors](https://pkg.go.dev/errors) package within the Go standard library, 
including the `errors.Is` and `errors.As` functions to provide a standardized approach for 
introspecting error values.

## Usage

### Creating an error

The `multierr.New` function is used to create a new `multierr.Error` from a list of errors.

```go
return multierr.New(
    errors.New("first error"),
    errors.New("second error"),
)
```

### Aggregating errors

The `multierr.Append` function is used to aggregate multiple errors into a single error. It has 
similar semantics to the built-in `append` function:

```go 
var errs error

err := step1()
if err != nil {
    errs = multierr.Append(errs, err)
}

err = step2()
if err != nil {
    errs = multierr.Append(errs, err)
}

return errs
```

### Checking for an error

The `errors.is` function can be used directly with a `multierr.Error` to check for a specific error:

```go
// Assume that err is a multierr.Error
err := someFunc()
if err != nil {
    if errors.Is(err, SomeError) {
        // err contains SomeError
    }
}
```

### Extracting an error

The `errors.As` function can be used directly with a `multierr.Error` to extract a specific error:

```go
// Assume that err is a multierr.Error
err := someFunc()
if err != nil {
    var someError *SomeError
    if errors.As(err, &someError) {
        // err contains SomeError and populates someError
    }
}
```

## Documentation

Documentation for `multierr` can be found [here](https://pkg.go.dev/github.com/jordanhasgul/multierr).