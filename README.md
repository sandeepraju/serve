# Serve (work in progress)

__Serve__ is a simple HTTP server for serving static files.

## Installing

```
brew tap <add correct path here>
brew install serve
```

## Installing from source code

```
go get github.com/sandeepraju/serve
```

## Usage

```
NAME:
   serve - A simple HTTP server for serving static files

USAGE:
   serve [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   A simple HTTP server for serving static files.

   See https://github.com/sandeepraju/serve for more details!

EXAMPLES:
   * Serve files in the current directory

      serve

   * Serve at port 8000 on 192.168.0.3

      serve -a 192.168.0.3 -p 8000

   * Serve /tmp

      serve -d /tmp

AUTHOR:
   Sandeep Raju Prabhakar <me@sandeepraju.in>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --directory value, --dir value, -d value  The directory to serve (default: "./")
   --address value, -a value                 The IP address or hostname of the interface (default: "localhost")
   --port value, -p value                    The port to listen on (default: 8888)
   --help, -h                                show help
   --version, -v                             print the version
```


## Examples

### Serve files in the current directory

```bash
serve
2017/05/27 22:44:54 Serving ./ at http://localhost:8888/
```

### Serve at port 8000 on 192.168.0.3

```bash
serve -a 192.168.0.3 -p 8000
2017/05/27 22:45:52 Serving ./ at http://localhost:8000/
```

### Serve /tmp

```bash
serve -d /tmp
2017/05/27 22:44:54 Serving /tmp at http://localhost:8888/
```

## Setting up for development

* Make sure [dep](https://github.com/golang/dep) is installed.

```
go get -u github.com/golang/dep/cmd/dep

or

brew install dep
```

* Clone this repository

```
git clone https://github.com/sandeepraju/serve.git
```

* Install dependencies

```
dep ensure
```

* Build the project

```
go build .
```

You'll a binary with the name `serve` in the current working directory.

* Run the binary

```
./serve --help
```

## License

```
BSD 3-Clause License

Copyright (c) 2018, Sandeep Raju Prabhakar
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
