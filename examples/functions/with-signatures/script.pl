#!/usr/bin/env perl

# Prints the lesser of two values.

use v5.20;
use feature qw(signatures);
no warnings qw(experimental::signatures);

sub min ( $x, $y ) {
    return $x < $y ? $x : $y;
}

say min( 3, 4 );
