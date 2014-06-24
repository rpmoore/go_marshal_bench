Go Marshal Bench
----------------

This contains benchmarks comparing several different encodings all written in Go

### Encodings

Currently the following encodings are being compared:

* JSON - Uses the built in JSON encoding library
* BSON - Uses [gobson](http://labix.org/gobson)
* Protobuf - Uses [goprotobuf](https://code.google.com/p/goprotobuf/)

### Planned

* Capnproto

### Running

To run the code, make sure that protbuf is installed.  Once it's install run `./build.sh`.  This will create the go protobuf files and build an executable.  To run the benchmarks run``go_marshal_bench`

To read the output, the `Internal` label means that it is using the non protobuf generated struct.  I wanted to see if there would be a performance different between serialzing the protobuf generated struct vs a non protobuf generated struct.
