# Marshall Mathers

I don't know what `unmarshall`ing means

## Requirements

1. ~~Unmarshall data~~
2. ~~Change all `locked` entries to `false`~~
3. ~~Convert all `password` entries to a preferred hash~~
    - ~~explain why the hash is more secure~~
    - ~~add a random salt~~
4. ~~Marshall data back into a file~~

## Usage

This primitive implementation can be used two different ways. Either:

1. Run `make test` to satisfy the above requirements while using no salt for hashing
2. Run `ARG=anything make test` to satisfy the above requirements while using a random 128bit salt for hashing

## Assumptions

I assumed quite a few things in this simple implementations. Specific references can be found by `cat`ting for `TODO:`s in code, while a vague list follows:

1. No proper argument parsing implementation is used,
2. Many variables are preset with no ability to set in runtime,
3. Even if no salt is required at runtime, a blank one will be present in the result,
    - the original intention here was to allow for backwards compatibility: if we suppose that the output can't support a salt, it shouldn't be outputted. After all, this wasn't the case.
4. Only one hashing algorithm is currently supported - `PBKDF2`, using SHA1.

## Hash algorithm choice

`PBKDF2` was chosen due to its ability to include a salt in its hashing processes, as well as its general acceptance as a decent process in the community.
