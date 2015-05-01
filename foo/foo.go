package foo

import "gopkg.in/errgo.v1"
import _ "gopkg.in/mgo.v2"

var Err = errgo.New("foo")
