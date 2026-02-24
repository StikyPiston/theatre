---
title: Theatre Documentation
author: StikyPiston
---

# Welcome to Theatre!

**Theatre** is a terminal-based markdown *presentation tool*, written in Go!

Slides are separated with `---`

---

# Installation

## using Nix (Recommended)

To install, add the input to your flake,

```nix
inputs = {
  theatre.url = "github:stikypiston/theatre"
}
```

and install it with

```nix
environment.systemPackages = [
  inputs.theatre.packages.${pkgs.stdenv.hostPlatform.system}.theatre
]
```

## using Go

To install, simply run:

```shell
go install github.com/stikypiston/theatre@latest
```

Make sure that `~/go/bin` is in your PATH

## from Source

### Prerequisites

- A working `go` installation (1.25.x)

### Steps

#### Clone the repo

```shell
git clone https://github.com/stikypiston/theatre
```

#### Build the project

```shell
go build
```

---

# Getting Started

To start, create a simple markdown file.

Then, at the top, add the *metadata*, as such:

```yaml
title: Your Presentation's Title
author: Your Name
```

After that, create slide.

```markdown
# Hello, world!

This is a slide.
```

Then, end the slide with three dashes (`---`) to separate it.

Add as many slides as you want!
