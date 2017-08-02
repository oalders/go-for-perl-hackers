# go-for-perl-hackers

The primary aim of this cheat sheet is to help Perl programmers get up and running with Go.

## Your editor

### vim

Consider adding the `vim-go` plugin to your `.vimrc`

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

##### Initialize without values

Perl:

```perl
my ($foo, $bar);
```

Go:

```go
var foo, bar int
```

##### Initialize with values

Perl:

```perl
my ($i, $j) = (1, 2);
```

Go:

```go
var i, j int = 1, 2
```

#### Checking for definedness

Perl:

```perl
my $foo;
if ( !defined $foo ) {
    ...;
}
```

Go:
```go
var foo
if foo == nil {
}
```

#### String Concatenation

Perl:

```perl
my $gopher = 'go' . 'pher';
```

Go:

```go
gopher := "go"+"pher"
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
my @array = (1..3);
```

Go:

```go
array := [3]int{1,2,3}
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
}
```

See [tour.golang.org/moretypes/5](https://tour.golang.org/moretypes/5)

### Dumping Data Structures

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
		pass	string
    }
	
	config.user = "florence"
	config.pass = "machine"
	
	fmt.Printf("%+v", config)
	
	return
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
while ( $sum < 10 ) {
    $sum += $i;
}
```

Go:

```go
// The init and post statement in a Go for loop are optional.
sum := 0
for sum < 10 {
	sum += sum
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
