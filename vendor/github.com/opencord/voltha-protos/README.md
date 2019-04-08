# voltha-protos

Protobuf files used by [VOLTHA](https://wiki.opencord.org/display/CORD/VOLTHA).

Currently this is used to generate both Golang and Python protobufs and gRPC stubs.

Protobuf definition files are located in `protos/voltha_protos` directory. This directory hierarchy and import scheme is required to allow the python code generated by the gRPC compiler to [have the correct import paths](https://github.com/grpc/grpc/issues/9575#issuecomment-293934506).

> NOTE: The `protos/google/api` directory has files copied from the [Google APIs](https://github.com/googleapis/googleapis), and is only included for initial compilation of the VOLTHA protobuf files - these API's should be installed independently via either the python [googleapis-common-protos](https://pypi.org/project/googleapis-common-protos/)
> package or the golang [go-genproto](https://github.com/google/go-genproto) repo.



## Go Environment

Get the voltha-protos repository:
```
mkdir -p ~/source
cd ~/source
git clone https://gerrit.opencord.org/voltha-protos
cd voltha-protos
```

### Setting up the Go environment

After installing Go on a Mac or Linux environment, the GOPATH environment variable needs be set.  These instructions assume its ~/go.
Create a symbolic link in the $GOPATH/src tree to the voltha-go repository:

```sh
mkdir -p $GOPATH/src/github.com/opencord
ln -s ~/source/voltha-protos $GOPATH/src/github.com/opencord/voltha-protos
```



## Go Dependencies

### Install Dependencies

Checkout and go install correct version of protoc-gen-go.  make install build and runs go install
```sh
cd $GOPATH/src/github.com/opencord/voltha-protos
go get -d github.com/golang/protobuf/
cd $GOPATH/src/github.com/golang/protobuf
git checkout v1.3.1
make install
```

## Building locally

Install libprotoc 3.7.0 manually or use the Makefile target install it.  Then build the python and golang stubs
```sh
cd $GOPATH/src/github.com/opencord/voltha-protos
make install-protoc
make build
```

use dist/*.tar.gz for local python imports
use go/ for local go imports



## Using voltha-protos in your project

### Python

Installation from Pypi:

`pip install voltha-protos`

or from a local build:

`pip install ~/source/voltha-protos/dist/*.tar.gz`

To use it within your code (for example)

`from voltha_protos import voltha_pb2`



### Go

```sh
go get github.com/opencord/voltha-protos
cd $GOPATH/github.com/opencord/voltha-protos
make build
````
protos should be importable from github.com/opencord/voltha-protos/go/packagename


To use the libraries, import protos with the root path github.com/opencord/voltha-protos/go/



## Testing

`make test` will run tests for all languages.
