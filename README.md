# Zshs

Utility functions for Zsh uses.

## Search plugin command help
Sometimes, I forgot how to use plugin command in Zsh, or may be I just want to learn but don't want to go to Zsh cheatsheet pages. So I use CLI command to search with the keyword that I already have.

```bash
go run ./cmd/zshs/main.go kubectl 'get pods'
go run ./cmd/zshs/main.go git 'pull --rebase'
```

## Checklist

- [ ] Build CI
- [ ] Make the first release

## License