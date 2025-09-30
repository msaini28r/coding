1. Case sensitive language, variables should be defined when using go also there are types like String, Bool, Integer ( unit8, unit64, int8, int 64, uintpr ), float32, float 64, Complex

2. Advance types cover Array, Slices, Maps, Structs, Pointers

3. When we use const, the Capital letter make its public otherwise private

```

const Login string = "Mohit"  // Exported (public)
const password string = "secret"  // Unexported (private)

```

4. The best way to declare any varible, is using := because it will decide the type itself. This is called walrus operator

```
Name := "Mohit"
Number := 4
Value := 5.67
```