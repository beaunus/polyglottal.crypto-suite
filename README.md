# Pretty Sweet Crypto Suite

This project was created while I was a student at [Code Chrysalis](https://www.codechrysalis.io/).

## Polyglottal Week

As part of Code Chrysalis's immersive course, students are given one week to write an application using a language that they have never used before. Since I had received a degree in computer science, all the languages that I used in school were off limits.

Lucky for me, [Golang](https://golang.org/) was invented after my computer science course, and I had never used it prior to this assignment. I was especially attracted to Golang for the following reasons:

* It has a reputation for being very fast to write, compile, and execute.
* It was developed by rockstars.
* It has all the great features of low-level languages, with the readability of modern high-level languages.

I discovered that all the rumors were true. It was a pleasure learning about Golang. The community is vast. There are a lot of easy-to-use hands-on tutorials. If you have never used Golang, you can get up to speed in no time for this project.

## What's the Deal

This is a project to show a variety of encryption algorithms throughout history.

My intention is to include a vast amount of algorithms from https://en.wikipedia.org/wiki/History_of_cryptography.

You can see this project in action at https://pretty-sweet-crypto-suite.herokuapp.com/.

## Scripts

```bash
# Get dev tools together
go get github.com/oxequa/realize

# Get the project run tools together
go get ./...

# Develop the code with hot reloading
realize start --name="server" --run

# Run all tests [including benchmarks]
go test [-benchmem -bench .] ./...
```

## Contributing

Thank you for your interest in this project. Feel free to make a PR if you are interested in contributing.
