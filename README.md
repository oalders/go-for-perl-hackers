
<!-- vim-markdown-toc GFM -->

* [go-for-perl-hackers](#go-for-perl-hackers)
  * [Your editor](#your-editor)
  * [Go vs Perl](#go-vs-perl)
    * [Installing binaries:](#installing-binaries)
      * [Perl](#perl)
      * [Go](#go)
    * [Constructs](#constructs)
    * [Comments](#comments)
      * [Perl](#perl-1)
      * [Go](#go-1)
    * [Print](#print)
      * [Printing strings](#printing-strings)
        * [Perl](#perl-2)
        * [Go](#go-2)
      * [Formatted print statements.](#formatted-print-statements)
        * [Perl](#perl-3)
        * [Go](#go-3)
      * [Printing from within a test](#printing-from-within-a-test)
        * [Perl](#perl-4)
        * [Go](#go-4)
    * [Variables](#variables)
      * [Environment Variables](#environment-variables)
        * [Perl](#perl-5)
        * [Go](#go-5)
      * [Variable Assignment](#variable-assignment)
        * [Perl](#perl-6)
        * [Go](#go-6)
      * [Multiple Variables](#multiple-variables)
        * [Declare without explicit values](#declare-without-explicit-values)
          * [Perl](#perl-7)
          * [Go](#go-7)
        * [Declare with explicit values](#declare-with-explicit-values)
          * [Perl](#perl-8)
          * [Go](#go-8)
      * [Double vs Single Quotes](#double-vs-single-quotes)
        * [Perl](#perl-9)
        * [Go](#go-9)
      * [Multiline strings](#multiline-strings)
        * [Perl](#perl-10)
        * [Go](#go-10)
      * [Boolean checks (true/false)](#boolean-checks-truefalse)
        * [Perl](#perl-11)
        * [Go](#go-11)
      * [Checking for (un)definedness](#checking-for-undefinedness)
        * [Perl](#perl-12)
        * [Go](#go-12)
      * [Incrementing and Decrementing Integer](#incrementing-and-decrementing-integer)
        * [Perl](#perl-13)
        * [Go](#go-13)
      * [Division with Integers](#division-with-integers)
        * [Perl](#perl-14)
        * [Go](#go-14)
      * [String Concatenation](#string-concatenation)
        * [Perl](#perl-15)
        * [Go](#go-15)
      * [String Concatenation (Existing String)](#string-concatenation-existing-string)
        * [Perl:](#perl-16)
        * [Go:](#go-16)
      * [string to byte array Conversion](#string-to-byte-array-conversion)
        * [Go](#go-17)
      * [byte array to string Conversion](#byte-array-to-string-conversion)
        * [Go](#go-18)
      * [Constants](#constants)
        * [Perl](#perl-17)
        * [Go](#go-19)
    * [Lists](#lists)
      * [Create an Array](#create-an-array)
        * [Perl](#perl-18)
        * [Go](#go-20)
      * [Size of an array:](#size-of-an-array)
        * [Perl](#perl-19)
        * [Go](#go-21)
      * [Hashes / Structs](#hashes--structs)
        * [Perl](#perl-20)
        * [Go](#go-22)
      * [Iterating Over a List](#iterating-over-a-list)
        * [Perl](#perl-21)
        * [Go](#go-23)
      * [Splitting a string](#splitting-a-string)
        * [Perl](#perl-22)
        * [Go](#go-24)
      * [Iterating Over a Hash/Map](#iterating-over-a-hashmap)
        * [Perl using keys](#perl-using-keys)
        * [Go using only primary return value](#go-using-only-primary-return-value)
        * [Perl using each](#perl-using-each)
        * [Go using both primary and secondary return values](#go-using-both-primary-and-secondary-return-values)
        * [Perl: using values](#perl-using-values)
        * [Go: ignoring primary return value, using only secondary return value](#go-ignoring-primary-return-value-using-only-secondary-return-value)
      * [Checking if a Hash/Map Key Exists](#checking-if-a-hashmap-key-exists)
        * [Perl](#perl-23)
        * [Go](#go-25)
      * [Using a Hash/Map to Track Seen Things](#using-a-hashmap-to-track-seen-things)
        * [Perl](#perl-24)
        * [Go](#go-26)
      * [Deleting a Hash/Map Key](#deleting-a-hashmap-key)
        * [Perl](#perl-25)
        * [Go](#go-27)
      * [Getting a List of Hash/Map Keys](#getting-a-list-of-hashmap-keys)
        * [Perl](#perl-26)
        * [Go](#go-28)
      * [Slices:](#slices)
        * [Perl](#perl-27)
        * [Go](#go-29)
      * [Appending Slices:](#appending-slices)
        * [Perl](#perl-28)
        * [Go](#go-30)
    * [Dumping Data Structures](#dumping-data-structures)
      * [To your terminal](#to-your-terminal)
        * [Perl](#perl-29)
        * [Go](#go-31)
      * [To disk (write)](#to-disk-write)
        * [Perl](#perl-30)
        * [Go](#go-32)
      * [To disk (append)](#to-disk-append)
        * [Perl](#perl-31)
        * [Go](#go-33)
    * [File Operations](#file-operations)
      * [Creating a directory](#creating-a-directory)
        * [Perl](#perl-32)
        * [Go](#go-34)
      * [Read an Entire File](#read-an-entire-file)
        * [Perl](#perl-33)
        * [Go](#go-35)
      * [Read First Line of a File](#read-first-line-of-a-file)
        * [Perl](#perl-34)
        * [Go](#go-36)
    * [Flow Control](#flow-control)
      * [if](#if)
        * [Perl](#perl-35)
        * [Go](#go-37)
      * [else](#else)
        * [Perl](#perl-36)
        * [Go](#go-38)
      * [elsif / else if](#elsif--else-if)
        * [Perl](#perl-37)
        * [Go](#go-39)
    * [Loops](#loops)
      * [For loops](#for-loops)
        * [Perl](#perl-38)
        * [Go](#go-40)
      * [While loops](#while-loops)
        * [Perl](#perl-39)
        * [Go](#go-41)
      * [Infinite loops](#infinite-loops)
        * [Perl](#perl-40)
        * [Go](#go-42)
      * [Short-circuiting a loop iteration](#short-circuiting-a-loop-iteration)
      * [Terminating a loop](#terminating-a-loop)
    * [Regular Expressions](#regular-expressions)
      * [Perl](#perl-41)
      * [Go](#go-43)
    * [Today's Date as YYYY-MM-DD](#todays-date-as-yyyy-mm-dd)
      * [Perl](#perl-42)
      * [Go](#go-44)
    * [Functions](#functions)
      * [Functions without signatures](#functions-without-signatures)
        * [Perl](#perl-43)
        * [Go](#go-45)
    * [Running Tests](#running-tests)
      * [Perl](#perl-44)
      * [Go](#go-46)
        * [Test Coverage](#test-coverage)
    * [Debugging](#debugging)
      * [Printing Stack Traces](#printing-stack-traces)
        * [Perl](#perl-45)
        * [Go](#go-47)
    * [Sleep](#sleep)
      * [Perl](#perl-46)
      * [Go](#go-48)
    * [Parsing URIs](#parsing-uris)
      * [Perl](#perl-47)
      * [Go](#go-49)
    * [Changing URI Query Params](#changing-uri-query-params)
      * [Go](#go-50)
    * [Command Line Scripts](#command-line-scripts)
      * [Print first argument to a script](#print-first-argument-to-a-script)
        * [Perl](#perl-48)
        * [Go](#go-51)
      * [Exiting a script](#exiting-a-script)
        * [Perl](#perl-49)
        * [Go](#go-52)

<!-- vim-markdown-toc -->
# go-for-perl-hackers

The primary aim of this cheat sheet is to help Perl programmers get up and running with Go.

## Your editor

I recommend setting up [gopls](https://pkg.go.dev/golang.org/x/tools/gopls).

## Go vs Perl

### Installing binaries:

#### Perl

```perl
cpm install -g App::perlimports
```

This will install a `perlimports` binary in your `$PATH`.

#### Go

```go
go install golang.org/x/tools/cmd/goimports@latest
```

This will install `goimports` into `$GOPATH/bin`

### Constructs

In this section we'll document some commonly used Perl constructs and try to find their equivalents in Go.

### Comments

#### Perl

```perl
# single line

=pod

Multi line

=cut
```

#### Go

```go
// single line (C++-style)

/*
Multi-line (C-Style)
*/
```

### Print

#### Printing strings

##### Perl

```perl
print 'hello, world';
```

##### Go

```go
package main

import "fmt"

func main() {
	fmt.Print("hello, world")
}
```

See "Double vs Single Quotes" for more information on why single quotes are not
used to quote strings in Go.

#### Formatted print statements.

##### Perl

```perl
printf('We are open %i days per %s', 7, 'week');
```

##### Go

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("We are open %d days per %s", 7, "week")
}
```

See [golang.org/pkg/fmt/](https://golang.org/pkg/fmt/)

#### Printing from within a test

##### Perl

```perl
diag 'foo happens';
```

##### Go

```go
t.Log("foo happens")
t.Logf("We are open %d days per %s", 7, "week")
```

### Variables

#### Environment Variables

##### Perl

```perl
$ENV{FOO} = 'bar';
local $ENV{FOO} = 'bar'; # Same as above, but with local scope

print "GOPATH: $ENV{GOPATH}\n";
```

##### Go

```go
os.Setenv("FOO", "bar")

fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
```

#### Variable Assignment

##### Perl

```perl
my $foo = 'bar';
my $pi = 3.14;
my $no_assignment;
```

##### Go

```go
// the following assignments are equivalent
var foo = "bar"
foo := "bar"

var pi float32 = 3.14 // explicit cast as float32
pi := float32(3.14)   // explicit cast as float32
pi := 3.14            // implicit cast as float64
pi := "3.14"          // implicit cast as string

var noAssignment string // equivalent to: noAssignment := ""
```

<https://go.dev/tour/basics/11>:

> When you need an integer value you should use int unless you have a specific
> reason to use a sized or unsigned integer type.

#### Multiple Variables

##### Declare without explicit values

###### Perl

```perl
my ($foo, $bar);
```

###### Go

```go
var foo, bar int
var nothing []string // create an empty slice
```

##### Declare with explicit values

###### Perl

```perl
my ($i, $j) = (1, 2);
```

###### Go

```go
var i, j int = 1, 2
```

#### Double vs Single Quotes

##### Perl

```perl
my $foo = 'bar';      # no variable interpolation
my $bar = "$foo baz"; # allow for variable interpolation
```

##### Go

```go
foo := "本" // implicitly cast as a string
foo := '本' // implicitly cast as a rune
```

Rune is an alias for `int32` and represents a Unicode code point.

For example: `fmt.Println('a')` prints 97, which is hexadecimal 61. The Unicode
for `a` is `U+0061`.

To print the Hex value of `'a'` you can also do `fmt.Println("%x", 'a')`

Go expects only a single letter inside single quotes, so this will yield a
compile time error:

```go
fmt.Printf("%s", 'hello')
```

> more than one character in rune literal

See
[golang.org/ref/spec#Rune_literals](https://golang.org/ref/spec#Rune_literals)
and <https://go.dev/tour/basics/11>

#### Multiline strings

##### Perl

```perl
my $long_string = <<'EOF';
my multiline
string
EOF
```

Use double quotes `<<"EOF";` if you need to interpolate variables.

##### Go

```go
longString := `
my multiline
string
`
```

#### Boolean checks (true/false)

##### Perl

```perl
my $success = 1;    # true
$success = 'foo';   # true
$success = 0;       # false
$success;           # (undef) false

if ($success) {
    print "This succeeded";
}

if ( !$success ) {
    print "This failed";
}

```

##### Go

```go
var success bool

success = true
success = false

if success == true {
	fmt.Println("This succeeded")
}

if success {
	fmt.Println("This succeeded")
}

if success == false {
	fmt.Println("This failed")
}

if !success {
	fmt.Println("This failed")
}
```

#### Checking for (un)definedness

##### Perl

```perl
my $foo;
if ( ! defined $foo ) {
    ...;
}
```

##### Go

```go
package main

import "fmt"

func main() {
	var myString string

	if myString == "" {
		fmt.Println("empty string")
	}

	var mySlice []int

	if mySlice == nil {
		fmt.Println("nil")
	}
	if len(mySlice) == 0 {
		fmt.Println("empty slice")
	}
}
```

[Go Playground](https://go.dev/play/p/R2rQ4F8lyCu)

#### Incrementing and Decrementing Integer

##### Perl

See [https://perldoc.perl.org/perlop.html#Auto-increment-and-Auto-decrement](https://perldoc.perl.org/perlop.html#Auto-increment-and-Auto-decrement)

```perl
$i = 0;  $j = 0;
print $i++;  # prints 0
print ++$j;  # prints 1

$i = 0;  $j = 0;
print $i--;  # prints 0
print --$j;  # prints -1
```

##### Go

```go
counter := 1
counter++
counter--
```

#### Division with Integers

##### Perl

```perl
my $amount = 18 / 20_000;
```

##### Go

```go
package main

import "fmt"

func main() {
	amount := float32(18) / float32(20000)
	fmt.Println(amount)
}
```

Note: without the `float32` conversion, `amount` will be 0.

[https://go.dev/play/p/ZIotxD2kQen](https://go.dev/play/p/ZIotxD2kQen)

#### String Concatenation

##### Perl

```perl
my $foo = 'go';
my $bar = 'pher';

my $gopher = "$foo$bar";
$gopher = $foo . $bar;
$gopher = join q{}, $foo, $bar;
$gopher = sprintf '%s%s', $foo, $bar;
```

##### Go

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var gopher string

	gopher = foo + bar
	gopher = fmt.Sprintf("%s%s", foo, bar)
}
```

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	foo := "go"
	bar := "pher"

	gopher = strings.Join([]string{"go", "pher"}, "")
	fmt.Println(gopher)
}
```

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	foo := "go"
	bar := "pher"

	buffer.WriteString(foo)
	buffer.WriteString(bar)
	fmt.Println(buffer.String())
}
```

#### String Concatenation (Existing String)

##### Perl:

```perl
my $foo = 'go';
$foo .= 'pher';
```

##### Go:

```go
package main

func main() {
	foo := "go"
	foo += "pher"
}
```

[https://go.dev/play/p/UJqGTHWCbQB](https://go.dev/play/p/UJqGTHWCbQB)

#### string to byte array Conversion

##### Go

```go
b := []byte("Each day provides its own gifts")
```

#### byte array to string Conversion

##### Go

```go
package main

import "fmt"

func main() {
	// create the byte array
	b := []byte("Each day provides its own gifts")

	// convert back to a string by passing the array
	contentAsString := string(b)
	fmt.Println(contentAsString)

	// convert back to a string by passing a slice referencing the storage of b
	contentAsString = string(b[:])
	fmt.Println(contentAsString)
}
```
[https://go.dev/play/p/DKsQJmS5yuf](https://go.dev/play/p/DKsQJmS5yuf)

#### Constants

##### Perl

```perl
use Const::Fast;
const my $hello => 'Hello, world';
```

##### Go

```go
// Create an *untyped* string constant
const hello = "Hello, world"

// Create a typed string constant
const hello string = "Hello, World"
```

Create multiple constants with one `const` declaration:

```go
const (
	hello   = "Hello, world"
	goodbye = "Goodbye!"
)
```

Constants cannot be declared using the := syntax.

### Lists

#### Create an Array

##### Perl

```perl
my @foo = (1..3);
my $first = $foo[0];
```

##### Go

```go
foo := [3]int{1, 2, 3}
first := foo[0]
```

Note that creating an empty array in Go means that it will be populated by the type's default values:

```go
var bar [5]int // creates an array of [0,0,0,0,0]
```

#### Size of an array:

##### Perl

```perl
my $size = @array;
```

##### Go

```Go
size := len(array)
```

#### Hashes / Structs

##### Perl

```perl
use Data::Printer; # exports p()

my %foo = (
    X => 1,
    Y => 2,
);

$foo{X} = 4;
print $foo{X}; # prints 4

p %foo;
# prints:
# {
#    X => 4,
#    Y => 2,
# }
```

##### Go

[https://play.golang.org/p/wyeohYSw-cf](https://play.golang.org/p/wyeohYSw-cf)

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)       // prints 4
	fmt.Printf("%+v\n", v) // prints {X:4 Y:2}

	// additional examples
	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0

	v.X = 0
	fmt.Printf("1: %d, 2: %d, 3: %d", v1.X, v2.X, v3.X)
}
```

#### Iterating Over a List

##### Perl

```perl
my @foo = ('foo', 'bar', 'baz');

for my $value ( @foo ) {
    print "$value\n";
}
```

```perl
my @foo = ('foo', 'bar', 'baz');

for (@foo) {
    print "$_\n";
}
```

```perl
// Print array index for each element
my @foo = ('foo', 'bar', 'baz');

for my $i (0..$#foo) {
    print "$i: $foo[$i]\n";
}

```

##### Go

```go
package main

import (
	"fmt"
)

func main() {
	foo := [3]int{1, 2, 3}

	for i, v := range foo {
		fmt.Printf("index: %v value: %v\n", i, v)
	}
}
```

#### Splitting a string

##### Perl

```perl
my @list = split ',', 'a,b,c'
```

##### Go

[https://play.golang.org/p/RXUdiKhZsjG](https://play.golang.org/p/RXUdiKhZsjG)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	mySlice := strings.Split("a,b,c", ",")
	fmt.Printf("%v", mySlice)
}
```

#### Iterating Over a Hash/Map

##### Perl using keys

```perl
my %hash = ( key_1 => 'foo', key_2 => 'bar', );
for my $key ( keys %hash ) {
    printf( "key: %s value: %s\n", $key, $hash{$key} );
}

```

##### Go using only primary return value

```go
package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{"key1": "foo", "key2": "bar"}

	for k := range myMap {
		fmt.Printf("key: %s value: %s\n", k, myMap[k])
	}
}
```

* [Go Playground](https://go.dev/play/p/qHcGFM1YRbF)

##### Perl using each

```perl
my %hash = ( key_1 => 'foo', key_2 => 'bar', );
for my $key, $value ( each %hash ) {
    printf( "key: %s value: %s\n", $key, $value );
}

```

##### Go using both primary and secondary return values

```go
package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{"key1": "foo", "key2": "bar"}

	for k, v := range myMap {
		fmt.Printf("key: %s value: %s\n", k, v)
	}
}
```

* [Go Playground](https://go.dev/play/p/01_1pFznIHQ)

##### Perl: using values

```perl
my %hash = ( key_1 => 'foo', key_2 => 'bar', );
for my $value ( values %hash ) {
    printf( "value: %s\n", $value );
}
```

##### Go: ignoring primary return value, using only secondary return value

```go
package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{"key1": "foo", "key2": "bar"}

	for _, v := range myMap {
		fmt.Printf("value: %s\n", v)
	}
}
```

* [Go Playground](https://go.dev/play/p/UWqh6M4dGB6)

#### Checking if a Hash/Map Key Exists

##### Perl

```perl

my %pages = ( home => 'https://metacpan.org' );
if ( exists $foo{home} ) {
    ...
}

```

##### Go

[https://stackoverflow.com/a/2050629/406224](https://stackoverflow.com/a/2050629/406224)

```go
package main

import "fmt"

func main() {
	pages := make(map[string]string)
	pages["home"] = "https://metacpan.org"
	if _, ok := pages["home"]; ok {
		fmt.Println("ok")
	}
}
```

#### Using a Hash/Map to Track Seen Things

##### Perl

```perl
my %seen;
$seen{sunrise} = 1;
if ( exists $seen{sunrise} ) {
   ...
}
```

##### Go

<https://go.dev/play/p/qJOz7_p0ptS>

```go
package main

import "fmt"

func main() {
	seen := map[string]struct{}{}
	seen["sunrise"] = struct{}{}
	if _, ok := seen["sunrise"]; ok {
		fmt.Println("I have seen the sunrise")
	}
}
```

The motivation here is that if the value has no inherent meaning, it could be
confusing to assign it something which could convey a meaning, like `true` or
`false`. There may be a small performance gain in terms of memory when using
`struct{}`, but it may not be any faster.
[https://gist.github.com/davecheney/3be245c92b61e5045f75](https://gist.github.com/davecheney/3be245c92b61e5045f75)

#### Deleting a Hash/Map Key

##### Perl

```perl

my %pages = ( home => 'https://metacpan.org' );
delete $pages{home};
```

##### Go

```go
package main

import "fmt"

func main() {
	pages := make(map[string]string)
	pages["home"] = "https://metacpan.org"
	delete(pages, "home")
}
```

#### Getting a List of Hash/Map Keys

##### Perl

```perl

my %pages = ( home => 'https://metacpan.org' );
my @keys = keys %pages;
```

##### Go

[https://go.dev/play/p/Z8OlHN7Q3GY](https://go.dev/play/p/Z8OlHN7Q3GY)

```go
package main

import (
	"fmt"
)

func main() {
	pages := make(map[string]string)
	pages["home"] = "https://metacpan.org"

	var keys []string
	for k := range pages {
		keys = append(keys, k)
	}
	fmt.Printf("%+v", keys)
}
```

#### Slices:

##### Perl

```perl
my @array = (0..5);
my @slice = @list[2..4];
```

##### Go

```go
array := [6]int{0, 1, 2, 3, 4, 5}
var slice []int = array[2:4]

var myslice []int    // create an empty slice of integers
var nothing []string // create an empty slice of strings
```

Note that arrays in Go have a fixed size, whereas slices are dynamically sized.

Also:

> A slice does not store any data, it just describes a section of an underlying array.
>
> Changing the elements of a slice modifies the corresponding elements of its underlying array.
>
> Other slices that share the same underlying array will see those changes.

See [https://tour.golang.org/moretypes/8](https://tour.golang.org/moretypes/8)

Note also that slices in Go can use defaults for lower and upper bounds. That
means that for the array of 10 integers `var a [10]int`, the following slices
are equivalent:

```go
var a [10]int

a[0:10] // explicit lower to upper bound
a[:10]  // use default lower bound (0)
a[0:]   // use default upper bound (0)
a[:]    // use default upper and lower bounds (0 and 10)
```

Note that the lower bound is the starting point in the index (ie 0) and the
_length_ of the slice is the upper bound, which is why the entire slice
consists of `a[0:10]` and _not_ `a[0:9]`.

See [https://tour.golang.org/moretypes/10](https://tour.golang.org/moretypes/10)

#### Appending Slices:

##### Perl

```perl
my @array = (0..5);
my @slice = @list[2..4];
push @slice, 11, 12;
```

Note that in Perl, a slice of an array is also an array, so there's no need to
make a distinction between the two.

##### Go

```go
array := [6]int{0, 1, 2, 3, 4, 5}
var slice []int = array[2:4]
slice = append(slice, 11, 12)
```

In Go, if you want to add two slices together, you'll need to treat the second
slice as a variadic parameter (`...`).

```go
package main

import "fmt"

func main() {
	first := []int{1, 2}
	second := []int{3, 4}
	total := append(first, second...)
	fmt.Printf("%+v", total)
}
```

See [https://go.dev/play/p/ylbpv1KDjzE](https://go.dev/play/p/ylbpv1KDjzE)

### Dumping Data Structures

#### To your terminal

##### Perl

```perl
use strict;
use warnings;
use feature qw( say );

use Data::Printer; # exports p() and np()

my %foo = ( a => 'b' );
p( %foo );

# or

say np( %foo );
```

##### Go

```go
package main

import "fmt"

func main() {
	var config struct {
		user string
		pass string
	}

	config.user = "florence"
	config.pass = "machine"

	fmt.Printf("%+v", config)

	return
}
```

Or:

```go
package main

import "github.com/sanity-io/litter"

func main() {
	var config struct {
		user string
		pass string
	}

	config.user = "florence"
	config.pass = "machine"

	litter.Dump(config)

	return
}
```

Or:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := []string{"one", "two", "three"}
	fmt.Println(prettyJSON(data))
}

func prettyJSON(target any) string {
	data, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		log.Fatalf("Cannot create pretty json from %v: %v", target, err)
	}
	return string(data)
}
```

[https://play.golang.org/p/FbSfHRNoVfP](https://play.golang.org/p/FbSfHRNoVfP)

#### To disk (write)

##### Perl

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->spew( np( @list ) );
```

##### Go

```go
package main

import (
	"log"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	list := [3]int{1, 2, 3}
	err := os.WriteFile("/tmp/foo.txt", []byte(litter.Sdump(list)), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
```

[Go Playground](https://go.dev/play/p/qjgoHlvYWfY)

#### To disk (append)

##### Perl

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->append( np( @list ) );
```

##### Go

```go
package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	list := [3]int{1, 2, 3}

	file, err := os.OpenFile(
		"/tmp/foo.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	spew.Fdump(file, list)
	file.Close()
}
```

### File Operations

#### Creating a directory

##### Perl

```perl
use Path::Tiny qw( path );
use Try::Tiny qw( catch try );

my $dir = '.my-perl-cache';
try {
    path($dir)->mkpath( { chmod => 0644 });
}
catch {
    die sprintf "Could not create dir %s because of %s", $dir, $_;
};
```

##### Go

```go
package main

import (
	"log"
	"os"
)

func main() {
	cacheDir := ".my-go-cache"
	if err := os.MkdirAll(cacheDir, 0644); err != nil {
		log.Printf("Cannot create dir %s because %v", cacheDir, err)
	}
}
```

#### Read an Entire File

##### Perl

```perl
use Path::Tiny qw( path );

my $content = path('path', 'to', 'file')->slurp_utf8;
```

##### Go

Note that in this case `dat` is `[]byte`

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	dat, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(dat))
	}
}
```

[https://go.dev/play/p/3dCR2NJuZOF](https://go.dev/play/p/3dCR2NJuZOF)

#### Read First Line of a File

##### Perl

In the Perl example, we'll chomp the line to make explicit that newlines need to be handled.

```perl
use Path::Tiny qw( path );

my $first_line;
my $fh = path('/path/to/file')->openr_utf8;

while (my $line = <$fh>) {
    $first_line = $line;
    chomp $first_line;
    last;
}

print $first_line, "\n";
```

##### Go

`scanner.Scan()` helpfully trims newlines for us.

```go
import (
	"fmt"
	"log"
)

func main() {
	file, err := os.Open("/path/to/file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstLine string
	for scanner.Scan() {
		firstLine = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(firstLine)
}
```

### Flow Control

#### if

##### Perl

```perl
if ( $foo > 1 ) {
    print 'bar';
}
```

##### Go

```go
if foo > 1 {
	fmt.Println("bar")
}

// parens are optional
if foo > 1 {
	fmt.Println("bar")
}
```

#### else

##### Perl

```perl
if ( $foo > 1 ) {
    print 'bar';
}
else {
    print 'baz';
}
```

##### Go

```go
if foo > 1 {
	fmt.Println("bar")
} else {
	fmt.Println("baz")
}
```

#### elsif / else if

##### Perl

```perl
if ( $foo > 1 ) {
    print 'bar';
}
elsif ( $foo < 10 ) {
    print 'baz';
}
```

##### Go

```go
if foo > 1 {
	fmt.Println("bar")
} else if foo < 10 {
	fmt.Println("baz")
}
```

### Loops

#### For loops

##### Perl

```perl
my $sum;
for ( my $i = 0 ; $i < 10 ; $i++ ) {
    $sum += $i;
}
```

##### Go

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
```

#### While loops

##### Perl

```perl
my $sum = 0;
my $i = 0;
while ( $i < 10 ) {
    $sum += $i++;
}
```

##### Go

```go
// The init and post statement in a Go for loop are optional.
sum := 0
i := 0
for i < 10 {
	sum += i
	i++
}
```

#### Infinite loops

##### Perl

```perl
while (1) {
}
```

##### Go

```go
for {
}
```

#### Short-circuiting a loop iteration

Perl

```perl
while (1) {
    ...
    next if $foo eq 'bar';
    ...
}
```

Go

```go
for {
	if foo == "bar" {
		continue
	}
	// Won't get here if continue is called
}
```

Note that `continue` will immediately begin the next iteration of the innermost `for` loop.

#### Terminating a loop

Perl

```perl
while (1) {
    ...
    last if $foo eq 'bar';
    ...
}
```

Go

```go
for {
	if foo == "bar" {
		break
	}
	// Won't get here if break is called
}
```

Note that `break` will exit the enclosing loop at the point where it is called.

### Regular Expressions

Match a string prefix.

#### Perl

```perl
my $str = '5.5.1';
if ( $str =~ m{\A5\.5} ) {
    print "ok\n";
}
```

#### Go

[https://go.dev/play/p/rpACvjzRt-K](https://go.dev/play/p/rpACvjzRt-K)

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\A5\.5`)
	str := "5.5.1"
	if re.MatchString(str) {
		fmt.Println("ok")
	}
}
```

[https://go.dev/play/p/-NXI8SjMjJQ](https://go.dev/play/p/-NXI8SjMjJQ)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "5.5.1"
	if strings.HasPrefix(str, "5.5") {
		fmt.Println("ok")
	}
}
```

### Today's Date as YYYY-MM-DD

#### Perl

```perl
use DateTime ();
print DateTime->now->ymd;
```

#### Go

[https://play.golang.org/p/k0ijssDDOU6](https://play.golang.org/p/k0ijssDDOU6)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02"))
}
```

### Functions

#### Functions without signatures

##### Perl

```perl
sub foo {
    print 'ok';
}
foo();
```

##### Go

```go
package main

import "fmt"

func main() {
	foo()
}
func foo() {
	fmt.Println("foo")
}
```

### Running Tests

#### Perl

```shell
$ perl Makefile.PL
$ make
$ make test
```

or

```shell
$ prove -l t/path/to/test.t
```

Use the `-v` flag for verbose output:

```shell
$ prove -lv t/path/to/test.t
```

#### Go

```shell
$ go test
```

Use the `-v` flag for verbose output:

```shell
$ go test -v
```

To test a subset of functions:

```shell
$ go test -run regexp
```

To bypass Go's test caching:

```shell
$ go test -count=1
```

##### Test Coverage

To generate a coverage report:

```shell
go test -v ./... -coverprofile cover.out && go tool cover -html cover.out -o cover.html
```

At this point you can view `cover.html` in a web browser and get a nice visual
of your test coverage.

### Debugging

#### Printing Stack Traces

##### Perl

```perl
use Carp qw( longmess );
print longmess();
```

##### Go

```go
package main

import (
	"runtime/debug"
)

func main() {
	debug.PrintStack()
}
```

### Sleep

#### Perl

```perl
sleep 60;
```

#### Go

```go
package main

import (
	"time"
)

func main() {
	time.Sleep(60 * time.Second)
}
```

### Parsing URIs

#### Perl

```perl
use Mojo::URL ();
my $url = Mojo::URL->new('https://www.google.com/search?q=schitt%27s+creek');
print $url->query->param('q'); # schitt's creek
```

#### Go

```go
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	url, err := url.Parse("https://www.google.com/search?q=schitt%27s+creek")
	if err != nil {
		log.Fatal(err)
	}
	q := url.Query()
	fmt.Println(q.Get("q")) // schitt's creek
}
```

### Changing URI Query Params

#### Go

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	url, _ := url.Parse("https://example.com")

	q := url.Query()
	q.Set("activity", "dance")
	q.Set("type", "flash")
	url.RawQuery = q.Encode()

	fmt.Println(url)
}
```

### Command Line Scripts

#### Print first argument to a script

##### Perl

```perl
print $ARGV[0], "\n" if $ARGV[0];
```

##### Go

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("%v\n", os.Args[1])
	}
}
```

#### Exiting a script

##### Perl

```perl
exit(0);
```

##### Go

```go
import("os")

os.Exit(0)
```
