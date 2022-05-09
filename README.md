# :memo: changelog-generator-default

[![CI](https://github.com/ted-vo/changelog-generator-default/workflows/CI/badge.svg?branch=main)](https://github.com/ted-vo/changelog-generator-default/actions?query=workflow%3ACI+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/ted-vo/changelog-generator-default)](https://goreportcard.com/report/github.com/ted-vo/changelog-generator-default)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ted-vo/changelog-generator-default)](https://pkg.go.dev/github.com/ted-vo/changelog-generator-default)

The default changelog generator for [semantic-release](https://github.com/ted-vo/semantic-release).

## Output of the changelog

The changelog generator will order the types of commits in the changelog in the following order:

- Breaking Changes
- Feature
- Bug Fixes
- Reverts
- Performance Improvements
- Documentation
- Tests
- Code Refactoring
- Styles
- Chores
- Build
- CI

## Emoji changelogs

In order to use emoji changelogs including a prefixed emoji, you need to provide the following config when calling semantic-relase: `--changelog-generator-opt "emojis=true"`. Or add the config within your `.semrelrc` file.

[Example Change Log](./examples/GENERATED_CHANGELOG.md)

## Licence

The [MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright © 2020 [Christoph Witzko](https://twitter.com/christophwitzko)
