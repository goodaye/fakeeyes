package config

import "fmt"

var ErrorPathIsNotDir = fmt.Errorf("path is not dir ")
var ErrorPathIsNotRegularFile = fmt.Errorf("path is not regular file ")
