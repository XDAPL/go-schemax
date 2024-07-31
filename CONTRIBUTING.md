# Welcome to the go-schemax contributing guide <!-- omit in toc -->

First, welcome to the go-schemax repository. This is a rather unique Go-based package that allows interaction with LDAP schema definitions as flexible object instances rather than a line of rather esoteric text. If you are any kind of directory professional or architect, you will no doubt find this package useful.

## A Friendly Word of Warning

The subject matter demonstrated by this package, and by LDAP services in general, can be quite delicate and somewhat draining for the layperson. As such, I've tried to keep the focus of this package very narrow. There are many possible applications for this package, but I'm quite certain they should remain _outside_ of the code base as independent packages. One possible exception here is anything simple and focused enough to be stored in the `_examples` folder.

As such, those individuals wanting to craft all sorts of nifty utilities, such as a Go `template/html` demo 😉, should keep in mind that such implementations should _usually_ remain independent packages on their own. If you're uncertain, you can always ask me!

## Contributor guide

A few things should be reviewed before submitting a contribution to this repository:

 1. Read our [Code of Conduct](./CODE_OF_CONDUCT.md) to keep our community approachable and respectable.
 2. Review [RFC4512 Section 4.1](https://datatracker.ietf.org/doc/html/rfc4512#section-4.1). This package is a direct implementation of that document section as it pertains to the parsing, validation, relationships and FIFO storage of all types of schema definitions.
 3. Review the main [![Go Reference](https://pkg.go.dev/github.com/JesseCoretta/go-schemax?status.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-schemax) page, which provides the entire suite of useful documentation rendered in Go's typically slick manner 😎.
 4. Review the [Collaborating with pull requests](https://docs.github.com/en/github/collaborating-with-pull-requests) document, unless you're already familiar with its concepts ...

Once you've accomplished the above items, you're probably ready to start making contributions. For this, I thank you.
