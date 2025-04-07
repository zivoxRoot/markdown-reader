# Markdown reader

A markdown reader that output stylized markdown in your terminal !

## Features

- Automatic color support for 16 colors, 256 colors, and truecoor
- Support titles to level 6
- Bullet list and numbered list
- Quote
- Code block
- Bold, italic, strikethrough, and inline code

## Usage

```bash
md file.md
```

## Installation

1. Install script :

> Be careful when downloading scripts from the internet, always read them before ! Read the script [here](https://github.com/zivoxRoot/markdown-reader/blob/main/install.sh)

```Bash
curl -L https://raw.githubusercontent.com/zivoxRoot/markdown-reader/refs/heads/main/install.sh | bash
```

2. Install markdown-reader with go

```bash

```

3. Build it from source, you can either take the latest release or the unstable development version

```bash
git clone https://github.com/zivoxRoot/markdown-reader.git
cd markdown-reader
go build -o ./md ./cmd/main.go
sudo mv md /usr/local/bin
```

## Contributing

Contributions are always welcome!

See `CONTRIBUTING.md` for ways to get started.
