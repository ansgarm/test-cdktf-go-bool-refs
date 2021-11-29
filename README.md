# test-cdktf-go-bool-refs

References to boolean properties of other resources cannot be used as boolean arguments for resources.

## To reproduce run
```
yarn fail
```

## Should fail with
```
[2021-11-29T15:08:22.505] [ERROR] default - panic: unable to make an instance of unregistered interface interface {}

goroutine
[2021-11-29T15:08:22.506] [ERROR] default - 1 [running]:
github.com/aws/jsii-runtime-go/internal/kernel.(*Client).castAndSetToPtr(0xc0000b4060, 0x11d1280, 0xc000097600, 0x194, 0x11d5100, 0xc0000a5aa0, 0x15)
	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/internal/kernel/conversions.go:78 +0x1127
github.com/aws/jsii-runtime-go/internal/kernel.(*Client).CastAndSetToPtr(0xc0000b4060, 0x11c0f60, 0xc000097600, 0x11d5100, 0xc0000a5aa0)
	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/internal/kernel/conversions.go:22 +0x15a
github.com/aws/jsii-runtime-go/runtime.Get(0x120d060, 0xc0000973f0, 0x120eb38, 0x5, 0x11c0f60, 0xc000097600)
	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/runtime/runtime.go:298 +0x250
cdk.tf/go/stack/generated/random.(*jsiiProxy_String).Lower(0xc0000973f0, 0x8c37fb8, 0xc000306498)
	/Users/ansgar/projects/playground/test-cdktf-go-bool-refs/generated/random/random.go:3837 +0x7b
main.NewMyStack(0x8c30bf8, 0xc000096e60, 0x1213d70, 0x17, 0xc000096e60, 0x0)
	/Users/ansgar/projects/playground/test-cdktf-go-bool-refs/main.go:25 +0x379
main.main()
	/Users/ansgar/projects/playground/test-cdktf-go-bool-refs/main.go:34 +0x7f
```