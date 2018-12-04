# go-for-perl-hackers

The primary aim of this cheat sheet is to help Perl programmers get up and running with Go.

## Your editor

### vim

Consider adding the `vim-go` plugin to your `.vimrc`

If you're using [vim-plug](https://github.com/junegunn/vim-plug), that would
look something like:

```
Plug 'fatih/vim-go'
```

If you're using Pathogen, that would look something like:

```
runtime bundle/vim-pathogen/autoload/pathogen.vim
" Bundle: tpope/vim-pathogen
call pathogen#infect()

" Bundle: https://github.com/fatih/vim-go.git
```

Now, open `vim` after installing `vim-go` and enter `:GoInstallBinaries`

Now you can take advantage of `:GoRename` when refactoring code and `:GoTest`
in order to test your code without leaving your `vim` session.

Also, you can add `goimports`:

First install it: `go get golang.org/x/tools/cmd/goimports`

Then add the following to your `.vimrc`

`let g:go_fmt_command = "goimports"`

Lastly, if your editor is set up to display tabs, you may want to disable that for Go files.

```
autocmd FileType go setlocal nolist
```

#### shortcuts

`daf` in `vim` will now allow you to cut an entire function without first
needing to select it.

## Go vs Perl

In this section we'll document some commonly used Perl constructs and try to find their equivalents in Go.

### Comments

Perl:

```perl
# single line

=pod

Multi line

=cut
```

Go:

```go
// single line (C++-style)

/*
Multi-line (C-Style)
*/
```

### Print

#### Printing strings

Perl:

```perl
print 'hello, world';
```

Go:

```
package main

import "fmt"

func main () {
    fmt.Println("hello, world")
}
```

#### Formatted print statements.

Perl:

```perl
printf('We are open %i days per %s', 7, 'week');
```

Go:

```go
package main

import ( "fmt" )

func main() {
    fmt.Printf("We are open %d days per %s", 7, "week")
}

```

See [golang.org/pkg/fmt/](https://golang.org/pkg/fmt/)

#### Printing from within a test

Perl:

```perl
diag 'foo happens';
```

Go:

```go
t.Log("foo happens")
t.Logf("We are open %d days per %s", 7, "week")
```

### Variables

#### Variable Assignment

Perl:

```perl
my $foo = 'bar';
my $pi = 3.14;
```

Go:

```go
// the following assignments are equivalent
var foo = "bar"
foo := "bar"

var pi float32 = 3.14  // explicit cast as float32
pi := float32(3.14)    // explicit cast as float32
pi := 3.14             // implicit cast as float64
pi := "3.14"           // implicit cast as string
```

##### Instantiate multiple variables at once

Perl:

```perl
my ($one, $two, $three);
```

Go:

```go
var one, two, three string
```

##### Double vs Single Quotes

Perl:

```perl
my $foo = 'bar'; // no variable interpolation
my $bar = "$foo baz"; // allow for variable interpolation
```

Go:

```go
foo := "本" // implicitly cast as a string
foo := '本' // implicitly cast as a rune
```

See [golang.org/ref/spec#Rune_literals](https://golang.org/ref/spec#Rune_literals)

##### Multiple Variables

##### Declare without explicit values

Perl:

```perl
my ($foo, $bar);
```

Go:

```go
var foo, bar int
```

##### Declare with explicit values

Perl:

```perl
my ($i, $j) = (1, 2);
```

Go:

```go
var i, j int = 1, 2
```

#### Checking for (un)definedness

Perl:

```perl
my $foo;
if ( ! defined $foo ) {
    ...;
}
```

Go:

```go
var myString string

if myString == "" {
    fmt.Println("Empty")
}


var mySlice []int

if mySlice == nil {
    fmt.Println("nil")
}
```

#### String Concatenation

Perl:

```perl
my $foo = 'go';
my $bar = 'pher';

$gopher = "$foo$bar";
$gopher = $foo . $bar;
$gopher = join q{}, $foo, $bar;
$gopher = sprintf '%s%s', $foo, $bar;
```

Go:

```go
foo := "go"
bar := "pher"

gopher := foo + bar
gopher := fmt.Sprintf("%s%s", foo, bar)
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

### Constants

Perl:

```perl
use Const::Fast;
const my $hello => 'Hello, world';
```

Go:

```go
// Create an *untyped* string constant
const hello = "Hello, world"

// Create a typed string constant
const hello string = "Hello, World"
```

Constants cannot be declared using the := syntax.

### Arrays

#### Create an Array

Perl:

```perl
my @foo = (1..3);
my $first = $foo[0];
```

Go:

```go
foo := [3]int{1,2,3}
first := foo[0]
```

#### Size of an array:

Perl:

```perl
my $size = @array;
```

Go

```Go
size := len(array)
```

### Hashes / Structs

Perl:

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

delete $foo{X};
```

Go:

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
    fmt.Println(v.X) // prints 4
    fmt.Printf("%+v\n", v) // prints {X:4 Y:2}

    // additional examples
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0

    delete(v, "X")
}
```

See [tour.golang.org/moretypes/5](https://tour.golang.org/moretypes/5)

### Dumping Data Structures

#### To your terminal

Perl:

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

Go:

```go
package main

import "fmt"

func main() {
    var config struct {
        user    string
        pass    string
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

import "github.com/davecgh/go-spew/spew"

func main() {
    var config struct {
        user    string
        pass    string
    }

    config.user = "florence"
    config.pass = "machine"

    spew.Dump(config)

    return
}
```

#### To disk (write)

Perl:

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->spew( np( @list ) );
```

Go:

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
        os.O_CREATE|os.O_WRONLY,
        0666,
    )

    if err != nil {
        log.Fatal(err)
    }

    spew.Fdump(file, list)
    file.Close()
}
```

#### To disk (append)

Perl:

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->append( np( @list ) );
```

Go:

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

### Flow Control

#### if

Perl:

```perl
if ( $foo > 1 ) {
    print 'bar';
}
```

Go:

```go
if ( foo > 1 ) {
    fmt.Println("bar")
}

// parens are optional
if foo > 1 {
    fmt.Println("bar")
}
```

#### else

Perl:

```perl
if ( $foo > 1 ) {
    print 'bar';
}
else {
    print 'baz';
}
```

Go:

```go
if foo > 1 {
     fmt.Println("bar")
} else {
     fmt.Println("baz")
}
```

### Loops

#### For loops

Perl:

```perl
my $sum;
for ( my $i = 0 ; $i < 10 ; $i++ ) {
    $sum += $i;
}
```

Go:

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

#### While loops

Perl:

```perl
my $sum = 0;
my $i = 0;
while ( $i < 10 ) {
    $sum += $i++;
}
```

Go:

```go
// The init and post statement in a Go for loop are optional.
sum := 0
i := 0
for i < 10 {
    sum += i
    i += 1
}
```

#### Infinite loops

Perl:

```perl
while (1) {
}
```

Go:

```go
for {
}
```

### Running Tests

Perl:

```perl
$ perl Makefile.PL
$ make
$ make test
```

or

```perl
$ prove -l t/path/to/test.t
```

Go:

```go
$ go test
```

If you're using `vim-go`, use `:GoTest` either from your `foo.go` or
`foo_test.go`.  (Note, you can also use `:GoAlternate` to toggle between the
two files.)

To test a subset of functions:

```go
$ go test -run regexp
```

If you're using `vim-go`, move your cursor to the name of the function you'd
like to test.  Running `:GoTest` here will run the function you're currently
in.
