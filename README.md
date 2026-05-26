# mapis

Contains protobuf(protocol buffers) files and generated go types used within the platform team

## Installation of required tools

`protoc-gen-go-lite` is a protoc plugin. That means your system must have the protobuf compiler (`protoc`) installed and available in `PATH`, in addition to the Go plugin binary itself.

Example system prerequisites:

- macOS: `brew install protobuf bufbuild/buf/buf`
- Debian/Ubuntu: `sudo apt-get install -y protobuf-compiler`

Then install the Go-based generators:

```terminal
brew install protobuf
brew install bufbuild/buf/buf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/aperturerobotics/protobuf-go-lite/cmd/protoc-gen-go-lite@latest

echo $PATH

# e.g.,, /Users/username/go/bin
export PATH=$PATH:YOUR_GO_BINARY_PATH 

# Check to see if it works
protoc --version
which protoc-gen-go-lite
buf generate apis/iam/v1/serviceuser_pat.proto
```

## Generation model

Mapis now generates protobuf-go-lite output directly into `gen/go`.

- generator: `protoc-gen-go-lite`
- output path: `gen/go`
- enabled features: `marshal+unmarshal+size+equal+clone`

## Development

1. CRUD your protobuf file, usually inside map-apis/<something>
2. test your changes with `buf generate apis/<something>/v1/<someproto.proto>`
3. if generated
    3.1 `git commit -am "some smart message" && git push`
    3.2 `git tag` # look at latest tag and bump minor
    3.3 `git push --tags` # FINISHED
4. else goto 1. until 3.3 

OBS! Prefer generating explicit files you changed (for example `buf generate apis/iam/v1/serviceuser_pat.proto`) to avoid broad generated diffs.
