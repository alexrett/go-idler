# Package "IDLER" (CGO)

Golang implementation of cross platform user idle timer

Support Mac OS X, Linux and Windows platforms

MacOS dependency
```
xcode-select --install
git go gcc
```

Linux dependency
```
gcc git go libxss-dev
```

Windows dependency
```
go git gcc (tdm-gcc)
```

## How to usage

https://github.com/alexrett/idler-app
```go
i := idler.Idle{}
i.GetIdleTime() // int 
```
