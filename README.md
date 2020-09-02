你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# <i style="color: rgb(184, 73, 169)">pretty</i> Sweet ¢®¥ÞƮø Suite

This project was created while I was a student at [Code Chrysalis](https://www.codechrysalis.io/).

## Polyglottal Week

As part of Code Chrysalis's immersive course, students are given one week to write an application using a language that they have never used before. Since I had received a degree in computer science, all the languages that I used in school were off limits.

Lucky for me, [Golang](https://golang.org/) was invented after my computer science course, and I had never used it prior to this assignment. I was especially attracted to Golang for the following reasons:

- It has a reputation for being very fast to write, compile, and execute.
- It was developed by rockstars.
- It has all the great features of low-level languages, with the readability of modern high-level languages.

I discovered that all the rumors were true. It was a pleasure learning about Golang. The community is vast. There are a lot of easy-to-use hands-on tutorials. If you have never used Golang, you can get up to speed in no time for this project.

As part of the assignment, I prepared a lecture about why I chose Golang and what I learned about Golang throughout the project. You can check out the [Google Slides](https://docs.google.com/presentation/d/1gjMG6qEgGlGHaYegMGpBffWGPgNJzWmvA0NxEtB8LJ8/edit?usp=sharing).

## What's the Deal

This is a project to show a variety of encryption algorithms throughout history.

As of now, it includes the following algorithms:

- [Ceasar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher)
- [Scytale Cipher](https://en.wikipedia.org/wiki/Scytale)

My intention is to include a vast amount of algorithms from https://en.wikipedia.org/wiki/History_of_cryptography.

You can see this project in action at https://pretty-sweet-crypto-suite.herokuapp.com/.

## Getting Started

If you'd like to get this thing running, fork and/or clone the code onto your local machine and run the following scripts.

### Scripts

```bash
# Get dev tools together
# This allows hot reloading
go get github.com/oxequa/realize

# Get the project run tools together
go get ./...

# Develop the code with hot reloading [env.PORT || 8000]
realize start --name="server" --run

# Run all tests [including benchmarks]
go test [-benchmem -bench .] ./...

# Run the server on port [env.PORT || 8000]
go run server
```

## Contributing

Thank you for your interest in this project. Feel free to make a PR if you are interested in contributing.
