# What is this about?

I am learning `Go` and the best way to learn is to write.

This file keeps links to other file and resouces.

## Proto
- Step 1: Install protoc (proto compiler) if you did not doing so, run `brew install protoc`; then you can use `protoc` in your command line
- Step 2: Download the package & make it available inside the project by running `go get github.com/golang/protobuf/proto`; then you generated golang file can use google golang packages
- Step 3: Decide where to keep your proto. It is recommend to keep all your proto file in `./proto`. 
- Step 4: Config your proto files.

Within your proto, there is a `package` option. This is not go package but the path from the `./proto` to your proto files.

For `option go_package`, please use the full go path to where you want your generated struct resides. e.g: `github.com/project-name/parent/child`. This should match with your usage of the Struct later on.

- Step 5: Generate the `.pb.go` file using `protoc`

Run `protoc --go_out=../ proto/*.proto` in the root of project.

Note: let `go_out` the parent of the project folder help generate the `.pb.go`file in correct folder.



# Resouces

- [Where I take code from](https://gobyexample.com)
- [Thinking in Go](ThinkingInGo.md)

