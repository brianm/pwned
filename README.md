# pwned 

CLI tool to check passwords against a bloom filter derived from
[';--have i been pwned?](https://haveibeenpwned.com/) . The bloom 
filter is tuned for a 0.000001% false positive rate.

Passwords may be passed on the command line or on stdin.
Each argument on the command line will be treated as a
candidate, or each line from stdin will be treated as a
candidate.

Candidate passwords which match (which have been pwned)
will be printed to stdout.

# Building

Use the makefile to build, it appends embeds the precompiled bloom filter
in the binary. If you want to rebuild the bloom filter you can do so
by [downloading the password list](https://haveibeenpwned.com/Passwords) and
placing it in the project directory, then running `make bloom`. You can adjust
the makefile to change bloom filter tuning params.

# License

This work is [Apache-2.0](https://opensource.org/licenses/Apache-2.0).

Depends on [DCSO/bloom](https://github.com/DCSO/bloom) which is 
[BSD-3-Clause](https://github.com/DCSO/bloom/blob/master/LICENSE)
