# Contributing

Thank you for your interest in contributing to ChainCraft!

All work on the code base should be motivated by our project
requirements. If you would like to work on a feature or improvement, please
document your changes clearly in commit messages and code comments.

All new contributions should be well-documented and tested. The
contribution helps solve problems and allows for
easy maintenance. Once the changes are implemented, ensure they follow
the established patterns and coding standards.

When implementing large structural
changes to the code base, these changes should be documented in the form of an
[Architectural Decision Record (ADR)](./docs/adr/) if applicable. The ADR will help
maintain coherence in the larger context. If you are not comfortable with writing an
ADR, you can document your changes in code comments and commit messages.

> How to pick a number for the ADR?

Find the largest existing ADR number and bump it by 1.

Each contribution should be aimed at creating maintainable code which aligns with the project's goals:

- Contributors should ensure their changes are well-tested and documented.
- Maintainers should have the necessary context to understand and review contributions.

## Commit Naming

Commits should be titled as following:

```txt
pkg: Concise title of commit
```

**Example:**

```txt
service/header: Remove race in core_listener
```

## Branching Model and Release

The main development branch is `main`.

Every release is maintained in a release branch named `vX.Y`. On each respective release branch, we tag the releases
vX.Y.0, vX.Y.1 and so forth.

Note all commits should be clear and well-documented. This keeps the commit history clean and makes it
easy to understand what changes were introduced.

### Development Procedure

The latest state of development is on `main`, which must never fail `make test`. *Never* force push `main`.

To begin contributing, create a development branch from `main`.

Make changes, and ensure your changes are well-tested and documented. Also, `git
rebase` on top of the latest `main` to keep your branch current.

#### Committing Changes

Before committing your changes:

- Ensure your branch is up-to-date with a recent `main`
- Run `make test` to ensure that all tests pass
- Ensure your code follows the project's coding standards
- Write clear commit messages following the established format

### Git Commit Style

We follow the [Go style guide on commit messages](https://tip.golang.org/doc/contribute.html#commit_messages). Write concise commits that start with the package name and have a description that finishes the sentence "This change modifies ChainCraft to...". For example,

```sh
cmd/debug: execute p.Signal only when p is not nil

[potentially longer description in the body]
```

It is recommended to prepend the type of change the commit is making to the commit message as documented [here](https://www.conventionalcommits.org/en/v1.0.0/).

```txt
feat(service/header): Title of commit.
```

Each commit should be clear and well-documented. Be sure to write descriptive commit messages that explain the purpose and scope of your changes.

## Testing

### Unit tests

Unit tests are located in `_test.go` files as directed by [the Go testing
package](https://golang.org/pkg/testing/). If you're adding or removing a
function, please check there's a `TestType_Method` test for it.

Run: `make test`

## Protobuf

If your PR modifies `*.proto` files, you will need to regenerate protobuf files with `make pb-gen`. Note this command assumes you have installed [protoc](https://grpc.io/docs/protoc-installation/).
