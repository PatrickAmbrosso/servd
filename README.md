# servd - Lightweight, Cross-Platform Service Manager

![GitHub Release](https://img.shields.io/github/v/release/PatrickAmbrosso/servd?color=blue)
![License](https://img.shields.io/github/license/PatrickAmbrosso/servd)
![Build](https://img.shields.io/github/actions/workflow/status/PatrickAmbrosso/servd/release.yml?label=build)

## Overview
`servd` is a lightweight **cross-platform service manager** that allows you to **install, manage, monitor, and control services** on Linux, macOS, and Windows.

**Supported Platforms:**
- **Linux** - uses  `systemd` or `init.d` as backend
- **macOS** - uses `launchd` as backend
- **Windows** - uses `sc.exe` as backend

## Features
- 🔧 **Install & Manage** services easily
- 🎯 **Cross-platform** (Windows, Linux, macOS)
- 🚨 **Monitor & Restart** services automatically
- 📜 **View Logs** for each managed service
- 🛠 **Configurable settings** via `servd config`
- 🔍 **Diagnostic checks** with `servd doctor`
- 📦 **Standalone binary** (no dependencies)

## Installation

### **Method 1 : Using Prebuilt Binaries**
Download the latest binary from [GitHub Releases](https://github.com/PatrickAmbrosso/servd/releases) and install it manually.

#### **Linux/macOS:**
```sh
curl -L https://github.com/PatrickAmbrosso/servd/releases/latest/download/servd -o /usr/local/bin/servd
chmod +x /usr/local/bin/servd
```

#### **Windows:**
```pwsh
# Windows (Run in PowerShell as Administrator):
Invoke-WebRequest -Uri "https://github.com/PatrickAmbrosso/servd/releases/latest/download/servd.exe"
```

### **Method 2 : Using the Package Manager of your choice**

#### **Homebrew (macOS):**
```sh
brew install servd
```

### **Linux (Debian/Ubuntu):**
```sh
sudo apt install servd
```

### **Linux (RPM-based):**
```sh
sudo dnf install servd
``` 

### **Linux (Arch Linux):**
```sh
sudo pacman -S servd
```

### **Linux (Alpine Linux):**
```sh
sudo apk add servd
```

### **Linux (NixOS):**
```sh
nix-env -iA nixpkgs.servd
```

#### **Scoop (Windows):**
```pwsh
scoop install servd
```

#### **Chocolatey (Windows):**
```pwsh
choco install servd
```

#### **winget (Windows):**
```pwsh
winget install -id servd.servd
```

### **Method 3 : Using Go's Package Manager**
```sh
go install github.com/PatrickAmbrosso/servd@latest
```

### **Method 4 : Building from Source**
Ensure Go 1.21+ is installed:

```sh
go version
```

Should output something like:
```sh
go version go1.21.1 linux/amd64
```

Then, clone the repository and build the binary by following the instructions below:

If you are on Linux or macOS, you can build the binary by running the following commands:

```sh
git clone https://github.com/PatrickAmbrosso/servd.git
cd servd
go build -o servd
```

If you are on Windows, you can build the binary by running the following commands:

```pwsh
git clone https://github.com/PatrickAmbrosso/servd.git
cd servd
go build -o servd.exe
```

### **Usage**

📌 General Commands
```pwsh
servd --help          # Show help menu
servd info            # Show system & platform details
servd doctor          # Run diagnostic checks
servd config --list   # View current config
```

📌 Service Management
```pwsh
servd install myservice --path /usr/bin/myapp
servd delete myservice
servd modify myservice --restart always
```

📌 Controlling Services
```pwsh
servd start myservice
servd stop myservice
servd restart myservice
servd pause myservice
servd monitor myservice
```

📌 Viewing Logs
```pwsh
servd logs myservice --tail 100
```

🔨 Build from Source
Ensure Go 1.21+ is installed:

```pwsh
git clone https://github.com/PatrickAmbrosso/servd.git
cd servd
go build -o servd
```

### **Acknowledgements**

This project wouldn't be possible without the following open-source projects:
- [cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
- [viper](https://github.com/spf13/viper) - Go configuration with fangs
- [charm](https://charm.sh/) - Supercharged terminal tooling for go (they like open source, okay!)
- [swatcher](https://github.com/patrickambrosso/swatcher) - Service watcher for Windows services

Here are some of the projects that I've used as inspiration for this project:
- [nssm](https://nssm.cc/) - A service wrapper for Windows


### **📜 License**
This project is licensed under the MIT License. See LICENSE for details.

---
