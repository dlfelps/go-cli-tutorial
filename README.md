# GoCliTeacher

An interactive command-line tool for learning how to build CLI applications in Go.

## Features

- **Interactive Tutorials**: Step-by-step guides with examples
- **Hands-on Exercises**: Practice your skills with guided exercises
- **Progress Tracking**: Keep track of completed tutorials and exercises
- **Comprehensive Topics**: From basics to best practices

## Installation

### Using Go

```bash
go install github.com/yourusername/gocli-teacher@latest
```

### From Binary Releases

Download the latest binary release for your platform from the [Releases page](https://github.com/yourusername/gocli-teacher/releases).

## Getting Started

After installation, run the tool with no arguments to see the welcome message:

```bash
gocli-teacher
```

Start with the basics tutorial:

```bash
gocli-teacher tutorial basics
```

## Available Tutorials

- **basics**: Learn the basic structure of a CLI app in Go
- **flags**: Understand how to use command-line flags
- **commands**: Learn about subcommands and command hierarchy
- **interactive**: Build interactive CLI features
- **best-practices**: Understand CLI design best practices

## Available Exercises

- **simple-cli**: Create a basic CLI tool
- **flag-exercise**: Practice using command-line flags
- **command-exercise**: Create a CLI tool with subcommands
- **interactive**: Build an interactive CLI application

## Tracking Your Progress

View your progress through tutorials and exercises:

```bash
gocli-teacher progress
```

View recent activity:

```bash
gocli-teacher progress --recent
```

Reset your progress (if needed):

```bash
gocli-teacher progress --reset
```

## Project Structure

- `cmd/`: Command definitions
- `tutorials/`: Tutorial content
- `exercises/`: Hands-on exercises
- `utils/`: Utility functions
- `progress/`: Progress tracking system

## Development

### Prerequisites

- Go 1.18 or later

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gocli-teacher.git
   cd gocli-teacher
   ```

2. Build the application:
   ```bash
   go build
   ```

3. Run the application:
   ```bash
   ./gocli-teacher
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI library for Go
- [Go](https://golang.org/) - The Go programming language