#!/usr/bin/perl

use strict;
use warnings FATAL => 'all';
use feature 'say';
use utf8;
use open qw(:std :utf8);

sub read_file {
    my ($file_name) = @_;

    open my $fh, '<', $file_name or die;
    $/ = undef;
    my $content = <$fh>;
    close $fh;

    return $content;
}

sub write_file {
    my ($file_name, $content) = @_;

    open my $fh, '>', $file_name or die;
    print $fh $content;
    close $fh;
}

sub run_go_fmt {
    my ($code) = @_;

    my $new_code;
    my $err = 0;

    write_file('a.go', $code);
    my $result = `go fmt a.go 2>&1`;

    if ($result =~ /^\s*\z/ || $result =~ '^\s*a.go\s*\z') {

        # everything is fine
        $new_code = read_file('a.go');

    } else {
        $new_code = '';
        $err = 1;
    }

    unlink 'a.go';

    return ($new_code, $err);
}

sub get_pretty_golang_code {
    my ($code) = @_;

    $code =~ s/^\s*//;
    $code =~ s/\s*\z//;
    $code = $code . "\n";

    # At first trying to fmt code as-is
    my ($pretty_code, $err) = run_go_fmt($code);
    if (not $err) {
        return $pretty_code;
    }

    # Next, trying to add 'package main'
    ($pretty_code, $err) = run_go_fmt("package main\n" . $code);
    if (not $err) {
        $pretty_code =~ s/^package main\s*//;
        return $pretty_code;
    }

    # Next, trying to put eveyrhing in main()
    ($pretty_code, $err) = run_go_fmt("package main\nfunc main() {\n" . $code . "}\n");
    if (not $err) {
        $pretty_code =~ s/
            ^
            package\smain
            \s+
            func\smain\(\)\s\{
            (.*)
            \}
        /
        $1
        /sx;

        $pretty_code =~ s/^\t//msg;

        $pretty_code =~ s/^\s*//;
        $pretty_code =~ s/\s*\z//;
        $pretty_code .= "\n";

        return $pretty_code;
    }

    # If nothing worked just returning the original trimmed code
    return $code;
}

sub main {

    die "Can't run with file a.go in the current directory" if -e 'a.go';

    my $content = read_file('README.md');
    my $new_content = '';

    my $golang_code = '';
    my $is_in_golang_block = 0;

    foreach my $line (split /\n/, $content) {

        if ($line =~ /^```\s*/ && $is_in_golang_block) {
            my $pretty_golang_code = get_pretty_golang_code($golang_code);
            $new_content .= "```go\n$pretty_golang_code```\n";
            $golang_code = '';
            $is_in_golang_block = 0;
        } elsif ($line =~ /^```go\s*/) {
            $is_in_golang_block = 1;
        } elsif ($is_in_golang_block) {
            $golang_code .= $line . "\n";
        } else {
            $new_content .= $line . "\n";
        }

    }

    write_file('README.md', $new_content);
}
main();
__END__
