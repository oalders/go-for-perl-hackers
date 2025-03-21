#!/usr/bin/env perl

# This was translated from main.go by Copilot.
use strict;
use warnings;

# Node class
package Node;

sub new {
    my ( $class, $payload ) = @_;
    my $self = {
        Payload => $payload,
        Next    => undef,
    };
    bless $self, $class;
    return $self;
}

# LinkedList class
package LinkedList;

sub new {
    my ($class) = @_;
    my $self = {
        Head => undef,
    };
    bless $self, $class;
    return $self;
}

sub add {
    my ( $self, $data ) = @_;
    my $new_node = Node->new($data);
    if ( !defined $self->{Head} ) {
        $self->{Head} = $new_node;
        return;
    }
    my $current = $self->{Head};
    while ( defined $current->{Next} ) {
        $current = $current->{Next};
    }
    $current->{Next} = $new_node;
}

sub print_list {
    my ($self) = @_;
    my $current = $self->{Head};
    while ( defined $current ) {
        print $current->{Payload}, " -> ";
        $current = $current->{Next};
    }
    print "nil\n";
}

# Main script
package main;
my $ll = LinkedList->new();
$ll->add(10);
$ll->add(20);
$ll->add(30);
$ll->print_list();    # Output: 10 -> 20 -> 30 -> nil
