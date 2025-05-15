# Codex

**Codex** is a comprehensive learning platform designed for developers who want to build applications on the **Solana Blockchain**.
It not only focuses on Solana but also provides resources to learn essential technologies such as **Rust, JavaScript, Go, and Ruby**, which are crucial for building robust Solana-based applications.

---

## 📝 Licensing

This project is offered under two different license models:

1. **Open Source:** Available for personal and educational use.
2. **Commercial License:** A special license is available for commercial use.
   Please contact us for more details: [buildwithcodex@gmail.com](mailto:buildwithcodex@gmail.com)

---

## 🚀 Quick Start

### 📥 Clone the Repository

```bash
git clone https://github.com/C-dexTeam/Codex.git --recursive
cd Codex
```

> **Note:** Be sure to use the `--recursive` flag to fetch all submodules (e.g., `Codex-Web3`, `Codex-Compiler`).

---

### ⚙️ Requirements

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [GNU Make](https://www.gnu.org/software/make/) (comes pre-installed on most Linux/macOS/WSL systems)

---

### 🧪 Run in Development Mode

```bash
make dev
```

Alternatively, you can use Docker Compose directly:

```bash
docker-compose -f ./deployment/dev.docker-compose.yaml up --build
```
---

### ❓ List All Available Make Commands

```bash
make help
```

---

## 📁 Project Structure

```
Codex/
├── backend/           # Backend services
├── frontend/          # Frontend interface
├── web3/              # Blockchain layer (submodule)
├── compiler/          # Code compiler (submodule)
├── deployment/        # Docker Compose files
├── Makefile           # Project automation commands
```
