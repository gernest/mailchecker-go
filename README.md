# mailchecker-go [![Build Status](https://travis-ci.org/gernest/mailchecker-go.svg)](https://travis-ci.org/gernest/mailchecker-go)

This is a Go(Golang) libary for detecting temporary (disposable/throwaway) emails.

This is based on [MailChecker](https://github.com/FGRibreau/mailchecker) but due to implementation details
it cannot be added to the platform list of MailChecker.


# Installation

	go get github.com/gernest/mailchecker-go
	
# How to use

```go
package main

import (
	"fmt"

	"github.com/gernest/mailchecker-go"
)

func main() {
	ok, err := mailchecker.Valid("gernest@20minutemail.co")
	if !ok {
		fmt.Println(err)
	}
}
```

# Contributing

Start with clicking the star button to make the author and his neighbors happy. Then fork the repository and submit a pull request for whatever change you want to be added to this project.

If you have any questions, just open an issue.

# Author
Geofrey Ernest

Twitter  : [@gernesti](https://twitter.com/gernesti)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.


