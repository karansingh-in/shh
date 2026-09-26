# v3il

A local, encrypted personal vault for journal entries and files — CLI-based, zero cloud, zero telemetry. Your data never leaves your machine.

## Why

Most password managers and note apps either phone home to a server or trust you to trust them. `v3il` doesn't touch the network at all. Everything is encrypted with a key derived from your own master password and stored as a single file on disk.

## Features

- **Encrypted journal vault** — add, view, update, delete, and list text entries, all encrypted at rest
- **Multiline entries** — write real journal entries, not one-liners
- **Arbitrary file encryption** — encrypt/decrypt any file (`v3il encrypt photo.jpg`, `v3il decrypt photo.jpg.v3il`)
- **Real cryptography, not hand-waved**:
  - **Argon2id** for password-based key derivation (memory-hard, resists GPU brute-forcing)
  - **AES-256-GCM** for authenticated encryption (tamper detection built in — wrong password fails loudly, not silently)
  - Fresh random salt per vault, fresh random nonce per encryption operation

## Installation

### Build from source
Requires [Go 1.21+](https://go.dev/dl/).

```bash
git clone https://github.com/karansingh-in/v3il.git
cd v3il
go build -o v3il
```

On Windows, this produces `v3il.exe`. On macOS/Linux, `v3il`.

### Cross-compiling
Building for a different OS than the one you're on:
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o v3il.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o v3il

# Linux
GOOS=linux GOARCH=amd64 go build -o v3il
```

## Usage
for windows
```bash
./v3il.exe add                    # add a new journal entry
./v3il.exe get <name>              # view an entry
./v3il.exe list                    # list all entry names
./v3il.exe update <name>           # edit an entry's body
./v3il.exe delete <name>           # delete an entry

./v3il.exe encrypt <path>          # encrypt any file → <path>.v3il
./v3il.exe decrypt <path>.v3il      # decrypt it back
```

Every command prompts for your master password (hidden input, not echoed to the terminal). Your vault lives at `~/.v3il/vault`.

## Security model — what this protects against, and what it doesn't

**Protects against:**
- Someone getting a copy of your vault file (stolen laptop, backup leak, etc.) — without your master password, it's unrecoverable ciphertext
- Tampering — AES-GCM's authentication means a modified file fails to decrypt rather than silently returning corrupted data

**Does NOT protect against:**
- A keylogger or malware running on your machine while you type your password
- Someone with access to your unlocked session/memory while the vault is decrypted
- A weak master password — Argon2id slows down brute-forcing, it doesn't make a weak password strong

This is a personal tool built to actually understand the cryptography involved, not an audited production security product. Use accordingly.

## Roadmap

- [ ] TUI (terminal UI) — interactive navigation instead of one-shot CLI commands
- [ ] Secure file sharing between trusted devices

## License

MIT