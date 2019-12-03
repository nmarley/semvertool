# semvertool

> Utility to validate semantic version strings and related operations

## Table of Contents
- [Install](#install)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

```sh
go get -u github.com/nmarley/semvertool/cmd/semvertool
```

## Usage

Example validating SemVer:

```sh
# Example of a long-ass version string, silently exits with RC=0
> semvertool 3.4.0-dev.4+buildmeta1234

# Performs basic SemVer validation, useful for CI scripts
> semvertool malarkey
malarkey is not valid SemVer

# Has quiet mode, useful for CI scripts
> semvertool 3.4.0 --quiet;echo $?
0

> semvertool malarkey --quiet;echo $?
1
```

Also prints version permutations:

```sh
# This is useful for Docker tagging
> semvertool 3.4.0 --show-permutations
3 3.4 3.4.0

# Can print prerelease...
> semvertool 1.2.3-beta.4+meta123 --prerelease
beta.4

# ... and also just prerelease 'head', e.g. the first part before the dot:
> semvertool 1.2.3-beta.4+meta123 --prerelease-head
beta

# Print permutations with prerelease 'head':
> semvertool 3.2.4-alpha --show-permutations
3-alpha 3.2-alpha 3.2.4-alpha
```

### Permutations

The main reason the `-show-permutations` flag exists is to use for Docker tagging, e.g. version `1.2.3` should return the strings `1.2.3`, `1.2`, and `1`, as they are commonly applied to Docker images. Tags like like `latest` are not included in this tool and should be considered separately, as this only deals with version numbers.

For versions like `3.4.0-dev.4+buildmeta1234`, it probably doesn't make sense to have tags a tag w/only major version + extra info like `3-dev.4+buildmeta1234`. But for a normal version string like `1.2.3` then it might make more sense. It's up to the caller whether or not to use this flag.

## Contributing

Feel free to dive in! [Open an issue](https://github.com/nmarley/semvertool/issues/new) or submit PRs.

## License

[ISC](LICENSE)
