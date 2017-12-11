package main

// we consolidate all generation here.
// normally the generate commands would live "close" to the respective code,
// but stringer must go after the go-bindata methods...

//go:generate go-bindata -ignore "\\.DS_Store" -pkg main -o generated_bindata.go bindata/...

//REMOVE_TO_ENABLE go:generate stringer -type=EXAMPLE
