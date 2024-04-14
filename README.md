# Jay

Jay aims to be the fastest production safe, serialization package written in [Go](https://go.dev) as
an alternative to
[JSON](https://pkg.go.dev/encoding/json),
[Protocol Buffers](https://pkg.go.dev/google.golang.org/protobuf), FlatBuffers,
[MessagePack](https://msgpack.org),
[Bebop](https://github.com/betwixt-labs/bebop),
and [`mus-go`](https://github.com/mus-format/mus-go).

Jay doesn't determine any types during runtime. Instead, the marshalling and unmarshalling functions are easily
generated using the [Jay commandline tool](https://github.com/speedyhoon/jay/tree/master/cmd).

## TLDR;

##### Pros:

* The fastest Go serialisation and deserialization message format.
* Encoded variables use less network overhead because no schema is added to the output when marshalling (unlike JSON).
* No custom language to learn. Jay uses Go's built in [`ast`](https://pkg.go.dev/go/ast) to find exported structs in your existing codebase.
* No hassles between client and server. Generate the code once and share _(via a Go module or copy-paste)_.
* Options to generate code optimised for:
	* Processing more requests per second _(higher CPU throughput)_ **OR**.
	* Least network bandwidth used _(10/100 networks)_.
* Doesn't introduce extra dependencies.
* Output could be compressed with `gzip`, `brotli`, [`zstd`](https://facebook.github.io/zstd/) or others.

##### Cons:

* Need to regenerate methods when Go structs are modified using the `jay` command line tool.
* Marshalled output is **not** human-readable.
* Only written for the Go language.

## Install

```shell
go install github.com/speedyhoon/jay/generate/jay
```

## Speed

Unlike JSON, Jay's marshal and unmarshal functions are generated by a command line tool.
This significantly increases execution speed during runtime by removing type reflection.

Most small structs with 10 fields can be serialized within 175 nanoseconds on old hardware _(Intel T6400 @ 2.0GHz
with GM45 GPU)_.

## Command line tool.

1. Install the Jay commandline tool.
	```shell
	go install github.com/speedyhoon/jay/generate/cmd
	```
2. Execute the command line tool and specify which `.go` file contains the exported structs.
	For example, `type Car struct` is located within `main.go`.
	```shell
	cd <<my_project_path>>
	jay
	```
3. Jay will then generate methods `MarshalJ()` and `UnmarshalJ([]byte)` in `jay.go` _(the default output file)_.
4. The new methods can then be used. For example:

```go
package main

import "fmt"

type Car struct {
	ID          uint
	Make, Model string
	Auto        bool
}

func main() {
	car := Car{ID: 42, Make: "Ford", Model: "Escort", Auto: false}
	var src []byte
	src = car.MarshalJ()
	fmt.Println(src)

	var car2 Car
	err := car2.UnmarshalJ(src)
	fmt.Printf("err: %v, %+v", err, car2)
}
```

## Message format

Exported struct fields are concatenated together in binary format, delimited by each field's length.

High-throughput mode on a 64-bit system:

```
Auto, ID,               Make,      Model         = 21 bytes
[0,   42,0,0,0,0,0,0,0, 5,F,o,r,d, 6,E,s,c,o,r,t]
```

Low-bandwidth mode: _(unfinished)_

```
Auto, ID,   Make,      Model         = 15 bytes
[0,   1,42, 5,F,o,r,d, 6,E,s,c,o,r,t]
```

## Supported types:

* `bool`
* `byte`, `[]byte`, `[]uint8`
* `float32`, `float64`
* `int`, `int8`, `int16`, `int32`, `int64`
* `rune`
* `string`
* `struct` _(Embedded structs aren't fully fuzz tested yet.)_
* `time.Time`, `time.Duration`
* `uint`, `uint8`, `uint16`, `uint32`, `uint64`

Jay also supports imported types. If the exported type is in the same package or a subpackage then `jay` will automatically find it with:

```shell
cd my_project_directory
jay .
```

To include an external type via the commandline, either:

* Specify the imported file: <br>
  `jay myFile.go ../other_project/myImportedFile.go`
* Or include that directory: <br>
  `jay myFile.go my_imported_directory`

## TODO

In order of priority:

* slices _(`[]string`)_
* arrays _(`[5]string`)_
* pointers _(`*string`)_
* Field tag options.
* Field tag documentation.
* Performance benchmarks.
* Low-bandwidth mode.
* Aliased definition types like `[]float` where `float` is defined as `type float float32`. <br>
  (`type float = float32`, `type floats = []float32` & `type floats []float32` are supported).
* maps? _(`map[string]uint`)_
* multi-dimensional arrays & slices? _(`[][]string`)_

## Done

* Specify which struct types to process in a large project directory. E.g.: ```-y Animal,settings.Config,engine.Specs```

## Not Supported

* `interface{}` requires adding reflection and/or type switches at a significant performance and complexity cost. Either
  use a different package or write your own function to convert the `interface{}` to one of the supported types
  (a `[]byte` might be the fastest option).
* Double nested pointers (`**string`).
* Generics (`any`).
* Functions.
* complex, use `[]byte` instead.

## Why

Conceived in 2011 as an interesting way to improve network communication.
It wasn't until 2023 when I was using `JSON` _(wasn't my choice by the way)_ to marshal and unmarshal data
between microcontrollers that it posed a significant bottleneck.
The aim was to process external messages within
a dozen microseconds to restore performance without upgrading the processor.

###### Name

Jay (pronounced as just `J`)  is a wordplay on [JSON](https://pkg.go.dev/encoding/json) without the `SON`, since the schema information is chopped off 🪚 and it's not human-readable.

The name `jay` also gives tribute to a 17-year-old netbook with a stuck `j` key on the keyboard. 🔁 😆 Every boot looks like:

```
jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj
jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj
jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj
```