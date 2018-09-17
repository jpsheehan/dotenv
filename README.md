# dotenv

The `dotenv` package simplifies setting up environment variables. This provides security to your projects as you don't need to hardcode sensitive data into your applications.

## Installation
```bash
go get github.com/jpsheehan/dotenv
```

## Usage:
```go
package main

import "github.com/jpsheehan/dotenv"

func main() {
    err := dotenv.Config()
    if err != nil {
        panic(err)
    }

    // environment variables are set up!
}
```

## Sample `.env` file

```
KEY: Value
NAME: Jesse
AGE: 25
THIS LINE IS IGNORED
COUNTRY: NZ
```

## Config(filenames ...string) error
Sets the environment variables found in the array of filenames supplied. If no filenames are supplied then it defaults to `.env`. It returns an error or nil depending on if an error occurred.

## ConfigOne(filename string) error
Sets the environment variables found in the filename supplied. It returns an error or nil depending on if an error occurred.
