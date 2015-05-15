# go-initstruct

Finally no mess with struct initialization!  
It may be a bit non-idiomatic but anyways... I like it :)

NOTE: doesn't work with private fields!

```
   type TestStruct struct {
	A  int            `init:"3"`
	B  string         `init:"test"`
	C  float32        `init:"3.14"`
	D  *TestStruct2   `init:"yes"`
	E  TestStruct3    `init:"yes"`
    F  map[string]int `init:"yes"`
    G  []int          `init:"yes"`
    H  chan int       `init:"10"`   // 10 is capacity
}

```

It is so easy now to auto-init it!  
Also, it is __recursive__, so D and E are also initialized!

```
	import ( . "github.com/rshmelev/go-initstruct" )

	...
        
	s := &TestStruct{}
    
    // simple way to replace zero-values with something more useful
	InitZeroFieldsRecursively(s)
    
    // reset struct!
    ResetAllFields(s)
    
```

...or, use a bit more complex but flexible `StructInitializer`

I'm not sure about using it in production, however it seems to be quite stable! 

Author: rshmelev@gmail.com
