EosForce API library for Go
=========================

[![GoDoc](https://godoc.org/github.com/eosforce/goeosforce?status.svg)](https://godoc.org/github.com/eosforce/goeosforce)

This library provides simple access to data structures (binary packing
and JSON interface) and API calls to an EosForce RPC server, running
remotely or locally.  It provides wallet functionalities (KeyBag), or
can sign transaction through the `keosd` wallet. It also knows about
the P2P protocol on port 9876.

As of before the June launch, this library is pretty much in
flux. Don't expect stability, as we're moving alongside the main
`eosio` codebase, which changes very fast.


The Library is fork from https://github.com/eosforce/goeosforce

License
-------

MIT
