# Filx

Filx is a command-line utility for password protecting files. Run `filx help` for info.

Filx uses 256-bit AES keys to encrypt files, generated via the given password.

## Install

1. Clone this repository
2. Run `go install` to install to your `$GOBIN`

## Reference

Encrypting a file (use -d to delete the input file)

```bash
filx enc ./myfile.dat [-d]
```

Decrypting a file (use -d to delete the input file)

```bash
filx dec ./myfile.dat.enc [-d]
```

## Contributing

I just made this one day while I was bored, feel free to open a PR.
