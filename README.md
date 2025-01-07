# mappis

Contains protobuf(protocol buffers) files and generated go types used within the plattform team

## Installation of required tools

```terminal
brew install protobuf
brew install bufbuild/buf/buf

# old?
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# new?
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

echo $PATH

# e.g.,, /Users/username/go/bin
export PATH=$PATH:YOUR_GO_BINARY_PATH 

# Check to see if it works
buf generate
```

## Development

1. CRUD your protobuf file, usually inside map-apis/<something>
2. test your changes with `buf generate mapis/<something>/v1/<someproto.proto>`
3. if generated
    3.1 `git commit -am "some smart message" && git push`
    3.2 `git tag` # look at latest tag and bump minor
    3.3 `git push --tags` # FINISHED
4. else goto 1. until 3.3 

OBS! Don't run `buf generate` (without any explicit files you've altered in your current commit) as versions can differ between developer, i.e., your small change leads to huge changes by generating all files.
