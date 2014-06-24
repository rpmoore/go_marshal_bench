Go Marshal Bench
----------------

This contains benchmarks comparing several different encodings all written in Go

### Encodings

Currently the following encodings are being compared:

* JSON - Uses the built in JSON encoding library
* BSON - Uses [gobson](http://labix.org/gobson)
* Protobuf - Uses [goprotobuf](https://code.google.com/p/goprotobuf/)

### Planned

* Message Pack
* Capnproto

### Running

To run the code, make sure that protbuf is installed.  Once it's install run `./build.sh`.  This will create the go protobuf files and build an executable.  To run the benchmarks run``go_marshal_bench`

To read the output, the `Internal` label means that it is using the non protobuf generated struct.  I wanted to see if there would be a performance different between serialzing the protobuf generated struct vs a non protobuf generated struct.

Here is an example of the output:

```
Starting Benchmark
Setting up test data with 100000 entries
Setting up test data with 100000 entries
Internal Json:	444.386804ms
Json:		450.493275ms
Internal Bson:	413.80553ms
Bson:		514.284076ms
ProtoBuf:	188.630753ms
Setting up test data with 1000000 entries
Setting up test data with 1000000 entries
Internal Json:	4.444901537s
Json:		4.605246395s
Internal Bson:	4.105297768s
Bson:		5.214341554s
ProtoBuf:	1.847902519s
Setting up test data with 5000000 entries
Setting up test data with 5000000 entries
Internal Json:	22.745075218s
Json:		22.662297948s
Internal Bson:	20.410903404s
Bson:		26.646342526s
ProtoBuf:	9.557486452s
Finished Benchmark
```
