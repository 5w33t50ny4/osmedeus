package libs

import "fmt"

const (
    // VERSION of this project
    VERSION = "v4.0.2"
    // DESC description of the tool
    DESC = "A Workflow Engine for Offensive Security"
    // BINARY name of Osmedeus
    BINARY = "osmedeus"
    // SNAPSHOT binary name of Osmedeus
    SNAPSHOT = "osm"
    // AUTHOR of this
    AUTHOR = "@j3ssiejjj"
    // DOCS private document
    DOCS = "https://docs.osmedeus.org"
    // METADATA domain for checking update
    METADATA = "https://metadata.osmedeus.org"
)

// TEMP default folder to store inputs
var TEMP = fmt.Sprintf("/tmp/%s-inputs/", SNAPSHOT)
var LDIR = fmt.Sprintf("/tmp/%s-log/", SNAPSHOT)
