GO-NBT
======

Simple to use the NBT format parser and generator. Code might be better but we got what we have.

## Example usage

```go
package main

import (
	"os"
	"compress/gzip"
	"github.com/bohdan4ik/go-nbt"
)

func main() {
	// Read example
	f, err := os.Open("nbt_file.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	r, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// r must implement io.Reader interface and call to
	// io.Read must provide uncompressed data
	rootCompound, rootCompoundName, err := nbt.ReadFrom(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compound:", rootCompoundName)
	fmt.Printf("%#v\n", rootCompound)


	// Write example
	f, err = os.Create("other_nbt_file.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	w, err := gzip.NewWriterLevel(f, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = nbt.WriteTo(w, rootCompound, rootCompoundName)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Close()
}
```

## TODO
- Add (more) documentation
