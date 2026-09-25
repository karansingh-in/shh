# shh

A local, encrypted personal vault for journal entries and files — CLI-based, zero cloud, zero telemetry. Your data never leaves your machine.

## Why

Most password managers and note apps either phone home to a server or trust you to trust them. `shh` doesn't touch the network at all. Everything is encrypted with a key derived from your own master password and stored as a single file on disk.

## Features

- **Encrypted journal vault** — add, view, update, delete, and list text entries, all encrypted at rest
- **Multiline entries** — write real journal entries, not one-liners
- **Arbitrary file encryption** — encrypt/decrypt any file (`shh encrypt photo.jpg`, `shh decrypt photo.jpg.shh`)
- **Real cryptography, not hand-waved**:
  - **Argon2id** for password-based key derivation (memory-hard, resists GPU brute-forcing)
  - **AES-256-GCM** for authenticated encryption (tamper detection built in — wrong password fails loudly, not silently)
  - Fresh random salt per vault, fresh random nonce per encryption operation

## Installation

### Build from source
Requires [Go 1.21+](https://go.dev/dl/).

```bash
git clone https://github.com/<your-username>/shh.git
cd shh
go build -o shh
```

On Windows, this produces `shh.exe`. On macOS/Linux, `shh`.

### Cross-compiling
Building for a different OS than the one you're on:
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o shh.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o shh

# Linux
GOOS=linux GOARCH=amd64 go build -o shh
```

## Usage
for windows
```bash
./shh.exe add                    # add a new journal entry
./shh.exe get <name>              # view an entry
./shh.exe list                    # list all entry names
./shh.exe update <name>           # edit an entry's body
./shh.exe delete <name>           # delete an entry

./shh.exe encrypt <path>          # encrypt any file → <path>.shh
./shh.exe decrypt <path>.shh      # decrypt it back
```

Every command prompts for your master password (hidden input, not echoed to the terminal). Your vault lives at `~/.shh/vault`.

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