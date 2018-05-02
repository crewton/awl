# awl

Awl is a Golang AWS SDK wrapper for lazy coders. It doesn't cover everything,
or even most things. In fact, the vast majority of the AWS SDK is not available
through Awl.


## Philosophy

Awl is all about being lazy. It requires very minimal setup, and lazily
provisions API clients on demand. This probably greatly increases the risk of
runtime errors in complex applications. Awl is not meant for these purposes.

Awl also forces you to use assumed roles to do any work. It doesn't help you set
up your initial credentials (it will pick things up from the standard places
via the SDK of course), and it requires that you give it a role to assume before
it will do anything for you. If you are not a fan of that approach, Awl may not
be the library for you.


## Status

Awl is in the very earliest stages of development. You should probably avoid
it.


## Who is Awl for?

Awl is intended to be for lazy coders who aren't writing ultra-high reliability
programs, but instead just want to get something done and not spend too much
time futzing with the nits of the AWS SDK. Seriously, though, this is not the
library to use if you are deploying high-stakes non-trivial production level
services.

However, given its extremely early stage of development, you probably still
shouldn't use it.


## Testing

Run `go test`.


## License

This software is public domain. No rights are reserved. See LICENSE for more
information.
