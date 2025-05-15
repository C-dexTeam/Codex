# Codex
This repository is a comprehensive learning platform designed for developers who want to build on the Solana Blockchain. It not only focuses on Solana but also provides resources to learn essential technologies like Rust, JavaScript, Go, and Ruby, which are crucial for creating Solana-based applications.

## Lisanslama  
This project is offered with two different license models: 
1. **Open Source:** It is available for personal and educational use.  
2. **Commercial License:** A special license is available for those who wish to use the software for commercial purposes. Contact [buildwithcodex@gmail.com](mailto:buildwithcodex@gmail.com) for more information.


## 🚀 Quick Start

### 📥 Clone the Repository

```bash
git clone --recurse-submodules https://github.com/C-dexTeam/Codex.git
cd Codex

Use --recurse-submodules to fetch all submodules (e.g. Codex-Web3, Codex-Compiler).

```
⚙️ Requirements
Docker

Docker Compose

GNU Make (default in Linux/macOS/WSL)



🧪 Run in Development Mode
bash
make dev

🛑 Stop and Remove Containers
bash
make down

🛠️ Build Containers (Optional)
bash
make dev-build

❓ Show Available Commands
bash
make help

📁 Project Structure
```
Codex/
├── backend/           # Backend services
├── frontend/          # Frontend interface
├── web3/              # Blockchain layer (submodule)
├── compiler/          # Code compiler (submodule)
├── deployment/        # Docker Compose files
├── Makefile           # Project automation commands
```
