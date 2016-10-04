# safestore

[![CircleCI](https://circleci.com/gh/blacklabcapital/safestore.svg?style=svg)](https://circleci.com/gh/blacklabcapital/safestore)



## Description

`safestore` is a Go package providing simple, useful, performant, O(1) access map-based key/value data storage implementations that offer atomic update and access methods safe for concurrent use.

Reusability is a key tenant of this package. How many times have you had to rewrite/reuse the same design patterns to support concurrent go routine access data stores? *Spoiler alert, we all know the answer. A LOT.* 

`safestore` seeks to be a general purpose "one-stop-shop" for any kind of <u>well tested</u> atomic/concurrent safe data store. No more boiler plate code for what we thought were "specialized" use cases...eventually someone else will ~~probably~~ definitely need the same thing!

*tl;dr* custom data type implementations are supported and encouraged.



## Design

The types in this package follow a traditional **OOP** design pattern, with the use of explicit *getter and setter* methods on all types. This is used to provide control over read/writes of the underlying data store, and to abstract away other private method behavior from the caller.

Atomicity is provided via use of an embedded `sync.mutex` in each type struct. Although only minor performance gains, the package conventionally avoids defering sync.Mutex Unlock() calls and instead explictly Lock() and Unlock() inline. Runtime defer calls do in fact add excess overhead. See https://github.com/golang/go/issues/14939



## Usage

`safestore `currently contains two subpackages,`primitivestore` and `seriesstore`.

The package can be and has been used in real time low-latency environments, and performs effectively even in microsecond latency environments.

The library has extensive unit tests which also double as examples for common usage. Please see the godoc for package documentation.

#### primitivestore

contains data stores for primitive or complex primitive data type values, such as `int`, `bool`, and `float`, or custom single occurence structs, among others.

#### seriesstore

contains data stores for *collection* type values , such as an array or set, which can store multiple occurences of a primitive or complex primitive data type. Unique methods for this subpackage include functions for accessing a specific index or key in the collection value, or a range of values, all of which are safe for concurrent use.



## Contributing

`master` holds the latest current stable version of safestore. Commits with a minor version are guaranteed to have no breaking API changes, only feature additions and bug fixes.

`dev` holds the latest commits and is where active development takes place. If you submit a pull request it should be against the `dev` branch.

`<major.minor>` are version branches. Tested changes from `dev` are staged for a release by merging into the appropriate version branch.