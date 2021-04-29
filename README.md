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

Also, you can add `goimports`:

First install it: `go get golang.org/x/tools/cmd/goimports`

Then add the following to your `.vimrc`

`let g:go_fmt_command = "goimports"`

If your editor is set up to display tabs visually, you may want to disable that for Go files.

```
autocmd FileType go setlocal nolist
```

#### shortcuts

`daf` in `vim` will now allow you to cut an entire function without first
needing to select it.

#### vim-go commands

##### :GoBuild

`go build`

##### :GoRun

`go run`

##### :GoGenerate

`go generate`

##### :GoRename

Call this when your cursor is over something you'd like to rename, and it will be renamed in all of the appropriate locations.

##### :GoTest

Run this in order to test your code without leaving your `vim` session.

##### :GoTestFunc

Run this when inside a test function and only this test function will be run.  Equivalent of `go test -run MyFunctionName`.

##### :GoAlternate

Toggles file between `foo.go` and `foo_test.go`

##### :GoDef

Takes you to the location where an item is defined.  Keep running this to move further up the stack.

##### :GoDefPop

Pops the stack which you created with `:GoDef`, taking you back to the previous location(s) you called `:GoDef` from.

##### :GoDecls

See functions which are declared in the current file.

##### :GoDeclsDir

Like `:GoDecls` but finds all the functions declared in the directory of the file you're editing.

##### :GoReferrers

Find other places where a function or variable is being invoked.

##### :GoDoc

Look up docs for a function.

##### :GoInfo

Get a function's input and output parameters.

##### :GoDescribe

Like `:GoInfo`, but with more info.

##### :GoFiles

List all files in package.

##### :GoDeps

List dependencies.

##### :GoWhicherrs

Get info on what sorts of errors an `err` variable may contain.

##### :GoCallers

Find out where a function callers are.

##### :GoImpl

Generate stub methods required by an interface.

##### :GoFreevars

## Go vs Perl

In this section we'll document some commonly used Perl constructs and try to find their equivalents in Go.

### Comments

#### Perl:

```perl
# single line

=pod

Multi line

=cut
```

#### Go:

```go
// single line (C++-style)

/*
Multi-line (C-Style)
*/
```

### Print

#### Printing strings

##### Perl:

```perl
print 'hello, world';
```

##### Go:

```
package main

import "fmt"

func main () {
    fmt.Print("hello, world")
}
```

#### Formatted print statements.

##### Perl:

```perl
printf('We are open %i days per %s', 7, 'week');
```

##### Go:

```go
package main

import ( "fmt" )

func main() {
    fmt.Printf("We are open %d days per %s", 7, "week")
}

```

See [golang.org/pkg/fmt/](https://golang.org/pkg/fmt/)

#### Printing from within a test

##### Perl:

```perl
diag 'foo happens';
```

##### Go:

```go
t.Log("foo happens")
t.Logf("We are open %d days per %s", 7, "week")
```

### Variables

#### Environment Variables

##### Perl:

```perl
$ENV{FOO} = 'bar';
local $ENV{FOO} = 'bar'; # Same as above, but with local scope

print "GOPATH: $ENV{GOPATH}\n";
```

##### Go:

```go
os.Setenv("FOO", "bar")

fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
```

#### Variable Assignment

##### Perl:

```perl
my $foo = 'bar';
my $pi = 3.14;
my $no_assignment;
```

##### Go:

```go
// the following assignments are equivalent
var foo = "bar"
foo := "bar"

var pi float32 = 3.14  // explicit cast as float32
pi := float32(3.14)    // explicit cast as float32
pi := 3.14             // implicit cast as float64
pi := "3.14"           // implicit cast as string

var noAssignment string // equivalent to: noAssignment := ""
```

See [golang.org/ref/spec#Rune_literals](https://golang.org/ref/spec#Rune_literals)

#### Multiple Variables

##### Declare without explicit values

###### Perl:

```perl
my ($foo, $bar);
```

###### Go:

```go
var foo, bar int
var nothing []string // create an empty slice
```

##### Declare with explicit values

###### Perl:

```perl
my ($i, $j) = (1, 2);
```

###### Go:

```go
var i, j int = 1, 2
```

#### Double vs Single Quotes

##### Perl:

```perl
my $foo = 'bar';      # no variable interpolation
my $bar = "$foo baz"; # allow for variable interpolation
```

##### Go:

```go
foo := "本" // implicitly cast as a string
foo := '本' // implicitly cast as a rune
```

#### Multiline strings

##### Perl:

```perl
my $long_string = <<'EOF';
my multiline
string
EOF
```

Use double quotes `<<"EOF";` if you need to interpolate variables.

##### Go:

```go
longString := `
my multiline
string
`
```

#### Boolean checks (true/false)

##### Perl:

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

##### Go:

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

##### Perl:

```perl
my $foo;
if ( ! defined $foo ) {
    ...;
}
```

##### Go:

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

#### Incrementing and Decrementing Integer

##### Perl:

See [https://perldoc.perl.org/perlop.html#Auto-increment-and-Auto-decrement](https://perldoc.perl.org/perlop.html#Auto-increment-and-Auto-decrement)

```perl
$i = 0;  $j = 0;
print $i++;  # prints 0
print ++$j;  # prints 1

$i = 0;  $j = 0;
print $i--;  # prints 0
print --$j;  # prints -1
```

##### Go:

```go
counter := 1
counter++
counter--
```

#### String Concatenation

##### Perl:

```perl
my $foo = 'go';
my $bar = 'pher';

$gopher = "$foo$bar";
$gopher = $foo . $bar;
$gopher = join q{}, $foo, $bar;
$gopher = sprintf '%s%s', $foo, $bar;
```

##### Go:

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

#### Constants

##### Perl:

```perl
use Const::Fast;
const my $hello => 'Hello, world';
```

##### Go:

```go
// Create an *untyped* string constant
const hello = "Hello, world"

// Create a typed string constant
const hello string = "Hello, World"
```

Create multiple constants with one `const` declaration:

```go
const(
    hello   = "Hello, world"
    goodbye = "Goodbye!"
)
```

Constants cannot be declared using the := syntax.

### Lists

#### Create an Array

##### Perl:

```perl
my @foo = (1..3);
my $first = $foo[0];
```

##### Go:

```go
foo := [3]int{1,2,3}
first := foo[0]
```

Note that creating an empty array in Go means that it will be populated by the type's default values:

```go
var bar [5]int // creates an array of [0,0,0,0,0]
```

#### Size of an array:

##### Perl:

```perl
my $size = @array;
```

##### Go:

```Go
size := len(array)
```

#### Hashes / Structs

##### Perl:

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

##### Go:

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
    fmt.Println(v.X) // prints 4
    fmt.Printf("%+v\n", v) // prints {X:4 Y:2}

    // additional examples
    v1 := Vertex{1, 2}  // has type Vertex
    v2 := Vertex{X: 1}  // Y:0 is implicit
    v3 := Vertex{}      // X:0 and Y:0

    v.X = 0
    fmt.Printf("1: %d, 2: %d, 3: %d", v1.X, v2.X, v3.X)
}
```

#### Iterating Over a List

##### Perl:

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

##### Go:

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

##### Perl:

```perl
my @list = split ',', 'a,b,c'
```

##### Go:

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

##### Perl:

```perl
my %hash = ( key_1 => 'foo', key_2 => 'bar', );
for my $key ( keys %hash ) {
    printf( "key: %s value: %s\n", $key, $hash{$key} );
}

```

##### Go:

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

#### Checking if a Hash/Map Key Exists

##### Perl:

```perl

my %pages = ( home => 'https://metacpan.org' );
if ( exists $foo{home} ) {
    ...
}

```

##### Go:

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

#### Deleting a Hash/Map Key

##### Perl:

```perl

my %pages = ( home => 'https://metacpan.org' );
delete $pages{home};
```

##### Go:

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

##### Perl:

```perl

my %pages = ( home => 'https://metacpan.org' );
my @keys = keys %pages;
```

##### Go:

[https://play.golang.org/p/b21XYOReH6S](https://play.golang.org/p/b21XYOReH6S)

```go
package main

import (
	"fmt"
)

func main() {
	pages := make(map[string]string)
	pages["home"] = "https://metacpan.org"

	keys := []string{}
	for k := range pages {
        	keys = append(keys, k)
    	}
	fmt.Printf("%+v", keys)
}
```
#### Slices:

##### Perl:

```perl
my @array = (0..5);
my @slice = @list[2..4];
```

##### Go:

```go
array := [6]int{0,1,2,3,4,5}
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

Note also that slices in Go can use defaults for lower and upper bounds. That means that for the array of 10 integers `var a [10]int`, the following slices are equivalent:

```go
a[0:10]  // explicit lower to upper bound
a[:10]   // use default lower bound (0)
a[0:]    // use default upper bound (0)
a[:]     // use default upper and lower bounds (0 and 10)
```

Note that the lower bound is the starting point in the index (ie 0) and the _length_ of the slice is the upper bound, which is why the entire slice consists of `a[0:10` and _not_ `a[0:9]`.

See [https://tour.golang.org/moretypes/10](https://tour.golang.org/moretypes/10)

#### Appending Slices:

##### Perl:
```perl
my @array = (0..5);
my @slice = @list[2..4];
push @slice, 11, 12;
```

Note that in Perl, a slice of an array is also an array, so there's no need to make a distinction between the two.

##### Go:

```go
array := [6]int{0,1,2,3,4,5}
var slice []int = array[2:4]
slice = append(slice, 11, 12)
```

### Dumping Data Structures

#### To your terminal

##### Perl:

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

##### Go:

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

##### Perl:

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->spew( np( @list ) );
```

##### Go:

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

##### Perl:

```perl
use Data::Printer; # exports np()
use Path::Tiny qw(path);

my @list = ( 1..3 );
path('/tmp/foo.txt')->append( np( @list ) );
```

##### Go:

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

##### Perl:

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

##### Go:

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

##### Perl:

```perl
use Path::Tiny qw( path );

my $content = path('path', 'to', 'file')->slurp_utf8;
```

##### Go:

Note that in this case `content` is `[]byte`

```go
    content, err := ioutil.ReadFile(filepath.Join("path", "to", "file"))

    if err != nil {
        log.Fatal(err)
    }

    // convert byte array to string
    contentAsString := string(content[:])
```

#### Read First Line of a File

##### Perl:

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

##### Go:

`scanner.Scan()` helpfully trims newlines for us.

```go
import(
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

##### Perl:

```perl
if ( $foo > 1 ) {
    print 'bar';
}
```

##### Go:

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

##### Perl:

```perl
if ( $foo > 1 ) {
    print 'bar';
}
else {
    print 'baz';
}
```

##### Go:

```go
if foo > 1 {
     fmt.Println("bar")
} else {
     fmt.Println("baz")
}
```

#### elsif / else if

##### Perl:

```perl
if ( $foo > 1 ) {
    print 'bar';
}
elsif ( $foo < 10 ) {
    print 'baz';
}
```

##### Go:

```go
if foo > 1 {
     fmt.Println("bar")
} else if foo < 10 {
     fmt.Println("baz")
}
```

### Loops

#### For loops

##### Perl:

```perl
my $sum;
for ( my $i = 0 ; $i < 10 ; $i++ ) {
    $sum += $i;
}
```

##### Go:

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

#### While loops

##### Perl:

```perl
my $sum = 0;
my $i = 0;
while ( $i < 10 ) {
    $sum += $i++;
}
```

##### Go:

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

##### Perl:

```perl
while (1) {
}
```

##### Go:

```go
for {
}
```

#### Short-circuiting a loop iteration

Perl:

```perl
while (1) {
    ...
    next if $foo eq 'bar';
    ...
}
```

Go:

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

Perl:

```perl
while (1) {
    ...
    last if $foo eq 'bar';
    ...
}
```

Go:

```go
for {
    if foo == "bar" {
        break
    }
    // Won't get here if break is called
}

```

Note that `break` will exit the enclosing loop at the point where it is called.

### Today's Date as YYYY-MM-DD

##### Perl

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

##### Perl:

```perl
sub foo {
    print 'ok';
}
foo();
```

##### Go:

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

#### Perl:

```perl
$ perl Makefile.PL
$ make
$ make test
```

or

```perl
$ prove -l t/path/to/test.t
```

Use the `-v` flag for verbose output:

```perl
$ prove -lv t/path/to/test.t
```

#### Go:

```go
$ go test
```

Use the `-v` flag for verbose output:

```go
$ go test -v
```

If you're using `vim-go`, use `:GoTest` either from your `foo.go` or
`foo_test.go`. (Note, you can also use `:GoAlternate` to toggle between the
two files.)

To test a subset of functions:

```go
$ go test -run regexp
```

If you're using `vim-go`, move your cursor to the name of the function you'd
like to test. Running `:GoTest` here will run the function you're currently
in.

To bypass Go's test caching:

```go
$ go test -count=1
```

### Debugging

#### Printing Stack Traces

##### Perl:

```perl
use Carp qw( longmess );
print longmess();
```

##### Go:

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

#### Perl:

```perl
sleep 60;
```

#### Go:

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

#### Perl:

```perl
use Mojo::URL ();
my $url = Mojo::URL->new('https://www.google.com/search?q=schitt%27s+creek');
print $url->query->param('q'); # schitt's creek
```

#### Go:

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

#### Go:

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

```perl
print $ARGV[0], "\n" if $ARGV[0];
```

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

##### Perl:

```perl
exit(0);
```

##### Go:

```go
import("os")

os.Exit(0)
```
