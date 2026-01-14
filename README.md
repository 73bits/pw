# 🔐 pw - cli password manager with encryption

A minimal, secure **command-line password manager** written in Go.
It uses **AES encryption**, **key derivation**, and a clean layered architecture.

This project is built to **learn and understand** how real password managers work internally.

---

## ✨ Features

* Master password–protected vault
* AES-GCM encryption
* Password-based key derivation (PBKDF / scrypt / argon-ready)
* Secure password input (no echo)
* Simple CLI commands
* Clean separation of concerns (model / service / repo / cmd)

---

## 📁 Project Structure

```
.
├── README.md
├── cmd                     # CLI commands
│   ├── add.go
│   ├── common.go
│   ├── get.go
│   ├── helper.go
│   └── init.go
├── go.mod
├── go.sum
├── internal
│   ├── model               # Data structures
│   │   └── entry.go
│   ├── repo                # Persistence layer
│   │   ├── json.go
│   │   └── repo.go
│   └── service             # Crypto & vault logic
│       └── auth.go
├── main.go                 # CLI entry point (minimal)
└── vault.enc               # Encrypted vault file (created at runtime)
```

---

## 🔐 Security Model (High-level)

* **Master password is never stored**
* A random **salt** is generated on vault creation
* Encryption key is derived from:

  ```
  key = DeriveKey(masterPassword + salt)
  ```
* Vault data is encrypted using **AES-GCM**
* `vault.enc` contains:

  ```
  [MAGIC][SALT][ENCRYPTED_DATA]
  ```

If the wrong password is provided:

* Decryption fails
* Authentication tag validation fails
* Vault cannot be read

---

## 🚀 Usage

### Initialize vault

```bash
go run . init
```

Creates a new encrypted vault file (`vault.enc`).

---

### Add a password

```bash
go run . add -name github
```

You’ll be prompted for:

* Master password
* Entry password

---

### Get a password

```bash
go run . get -name github
```

You’ll be prompted for the master password again.

---

## 🧠 Design Principles

* **Repo layer** only stores raw `[]byte`
* **Service layer** handles encryption & vault logic
* **CLI layer** handles user input/output
* Easy to replace file storage with PostgreSQL later

---

## 🧪 Learning Goals of This Project

* Understand encryption vs hashing
* Learn proper Go project structure
* Practice separation of concerns
* Build a real, non-trivial CLI app
* Avoid common security mistakes

---

## ⚠️ Disclaimer

This project is for **learning purposes**.
Do **not** use it to store real production secrets without further hardening.

---

## 📌 Future Improvements

* `list` command
* Duplicate entry prevention
* Session-based unlock
* Vault versioning & migration
* Tests
* PostgreSQL backend
