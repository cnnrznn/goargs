# goargs

A simple command line parser.

### Inclusion

### Module structure

### Algorithm

A command line provided by os.Args can look like the following:
```
[C:\Users\cnnrz\Documents\src\goargs\echo.exe 1 two three -flag -flagwithargs 1 2 3]
```

This is a `[]string` parsed by whitespace.

The algorithm is roughly the following:

1. Strip executable path from `os.Args`.
2. We are left with an array of strings. In the above example, `[1 two thre...]`.
3. Separate the slice into sub-slices where each sub-slice begins with a flag token: `-flagName` or `--flag`.
4. Sub slices look like `[--flag argOne argTwo argN]`.
5. For each sub-slice, populate the relevant entry in `Vals` struct.