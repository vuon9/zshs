# Zshs 🌟

[![Build Status](https://github.com/vuon9/zshs/actions/workflows/ci.yml/badge.svg)](https://github.com/vuon9/zshs/actions)

Utility functions for Zsh uses.

## Installation 📦

- Requisites: Go 1.20

Clone to local and use Go to build it:
```bash
git clone git@github.com:vuon9/zshs.git
cd zshs
go install ./cmd/zshs
```

Then it can be used anywhere by just typing `zshs`

## Features ✨

### Search a plugin command 🔍
I often found myself struggling to recall specific Zsh plugin commands. To avoid constantly referring to online cheat sheets, here is a search function:

```bash
zshs kubectl 'get pods'
zshs git 'pull --rebase'
```

## License 📜

MIT License

Copyright (c) 2024 Vuong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
