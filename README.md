# opensanic
Open-source SanicDB K/V cache (b+tree core benchmarks) 

This package provides benchmarks of both the binary search core and concurrent / safe database core of SanicDB. Many features of SanicDB are not implemented here. For closed source and in-development demos of SanicDB contact me! 

SanicDB is intended to allow searching through and retrieving/mutating ambiguous datasets concurrently by indexing keys. Sanic can return large sets of keys in the same time it takes to return single keys with other b+tree caches. Sanic also supports lexical keys and indexes.
