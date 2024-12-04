# TinyGo Starter

A starter template for TinyGo projects with built-in support for Raspberry Pi Pico/Pico W development and Gitpod integration.

## Features

- 🚀 Pre-configured development environment with Gitpod and VS Code
- 🔧 Automated tool installation and configuration
- 📦 Board Support Package (BSP) for Raspberry Pi Pico and Pico W
- 🛠️ Command-line utilities package
- 📝 Easy-to-use scripts for building and flashing

## Project Structure

```
.
├── cmd/                    # Command directory for your programs
│   └── template/          # Template program
├── pkg/                   # Reusable packages
│   ├── bsp/              # Board Support Package
│   │   ├── bsp.go        # Common BSP interface
│   │   ├── bsp_pico.go   # Pico-specific implementation
│   │   └── bsp_picow.go  # Pico W-specific implementation
│   └── cmdline/          # Command-line utilities
├── scripts/              # Utility scripts
│   ├── build.sh         # Build script
│   ├── flash.sh         # Flashing script for your local machine
│   └── add-target.sh    # New target creation script
```

## Getting Started

1. Open this repository using Gitpod Flex. Head over to https://app.gitpod.io and sign in with your GitHub account.
2. Wait for the automated setup to complete:
   - Go tools installation
   - VS Code Go configuration
   - Build targets configuration

## Available Automations

| Command | Description |
|---------|-------------|
| `newTarget` | Creates a new target program with basic structure |
| `updateGoRoot` | Updates Go root path for VS Code |
| `installTools` | Installs essential Go development tools |
| `refreshBuildTargets` | Refreshes build configurations for all targets |

## Creating a New Program

1. Run the `newTarget` automation or use the script directly:
```bash
./scripts/add-target.sh your_program_name
```

2. Navigate to `cmd/your_program_name`
3. Modify the template code to implement your program

## Building and Flashing

### Building
```bash
./scripts/build.sh your_program_name
```

### Flashing to Device
Download this flash script to your local machine and run it:
```bash
./scripts/flash.sh your_program_name
```

## BSP (Board Support Package)

The BSP package provides a hardware abstraction layer for supported boards:
- Raspberry Pi Pico
- Raspberry Pi Pico W

Feel free to extend this bsp package to support other boards or drivers as you see fit.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the terms found in the LICENSE file.