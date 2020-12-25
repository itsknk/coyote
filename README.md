# coyote
[![Go Reference](https://pkg.go.dev/badge/github.com/itsknk/coyote.svg)](https://pkg.go.dev/github.com/itsknk/coyote)
[![Build Status](https://travis-ci.com/itsknk/coyote.svg?branch=master)](https://travis-ci.com/itsknk/coyote)

Downloading files made easy.

coyote is a library written in go that helps in downloading files in an elegant manner.

## Install
If you don't have go installed, you can download it [here](https://golang.org/doc/install).
```
$ go get github.com/itsknk/coyote
``` 
After installing you can import it just like everyother library.
```
import (
    "github.com/itsknk/coyote"
)
```
## Usage
```
package main
import (
	"github.com/itsknk/coyote"
)
func main() {
	url := []string{
		"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
		"https://www.telegraph.co.uk/content/dam/film/Entourage/entouragetv-xlarge.jpg",
	}

	filename := []string{
		"entourage_one.png",
		"entourage_two.png",
	}
	downloader := coyote.Downloader{}
	downloader.Coyote(url, filename)
}
```

## Contributing
- Fork it and then do the changes or else download the zip file, test to make sure nothing is going sideways.
- Make a pull request with a detailed explanation. 

## License
[MIT](https://github.com/itsknk/statusHTTP/blob/master/LICENSE)