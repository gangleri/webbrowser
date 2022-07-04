# webbrowser 
Module that provides functionality to open a URL in the platforms default web browser.

## Installation
```
go get github.com/gangleri/webbrowser
```

## Example
```golang
package main

import "github.com/gangleri/webbrowser"

func main() {
	webbrowser.Open("https://golang.org")
}

```