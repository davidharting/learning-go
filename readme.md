
# Dependnecies

For the code to run, you need `git` and `libgit2` on your system. 
Use XCode build tools to get `git`, and install `libgit2` from homebrew. 

For rpc codegen with buf's connect you need more tools on your path:

```bash
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

See [these setup docs](https://connect.build/docs/go/getting-started#install-tools) for more details.