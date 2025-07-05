# Master Server Generator (msgen) 🚀

![GitHub release](https://img.shields.io/github/release/lukeeeeeeeeeeee/master-server-generator.svg)
![GitHub issues](https://img.shields.io/github/issues/lukeeeeeeeeeeee/master-server-generator.svg)
![GitHub stars](https://img.shields.io/github/stars/lukeeeeeeeeeeee/master-server-generator.svg)

Welcome to the **Master Server Generator**! This project, known as **msgen**, is a powerful cross-platform CLI tool designed to help you create a standalone, persistent server binary. Whether you're managing a Windows or Linux environment, msgen makes server management easier and more efficient.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Cross-Platform**: Works seamlessly on both Windows and Linux.
- **Background Service**: Runs as a background service or daemon.
- **Auto-Start**: Automatically starts at boot, either as a Windows Service or a systemd unit.
- **Self-Restart**: Automatically restarts on crash or stop, ensuring your server remains available.
- **Simple CLI**: Easy to use command-line interface for quick server setup.

## Installation

To get started with msgen, you need to download the latest release. You can find the releases [here](https://github.com/lukeeeeeeeeeeee/master-server-generator/releases). Download the appropriate file for your operating system and execute it to install the tool.

### Windows Installation

1. Download the Windows executable from the [Releases](https://github.com/lukeeeeeeeeeeee/master-server-generator/releases) section.
2. Open a command prompt and navigate to the directory where you downloaded the file.
3. Run the executable to complete the installation.

### Linux Installation

1. Download the Linux binary from the [Releases](https://github.com/lukeeeeeeeeeeee/master-server-generator/releases) section.
2. Open a terminal and navigate to the download location.
3. Make the binary executable:
   ```bash
   chmod +x msgen
   ```
4. Run the binary to complete the installation.

## Usage

Using msgen is straightforward. After installation, you can generate a server binary with a simple command.

### Basic Command

To generate a new server, use the following command:

```bash
msgen create <server-name>
```

Replace `<server-name>` with your desired server name. This command will create a new server binary in the current directory.

### Options

You can customize the server generation with various options. For example:

- `--port`: Specify the port on which the server will listen.
- `--config`: Provide a path to a configuration file.
- `--daemon`: Run the server as a daemon.

### Example

To create a server named "my-server" that listens on port 8080, you would run:

```bash
msgen create my-server --port 8080
```

## Configuration

After generating your server, you may want to configure it. The generated server will create a configuration file in the same directory. You can edit this file to adjust settings like:

- **Logging Level**: Control the verbosity of logs.
- **Database Connections**: Set up connections to databases.
- **Service Settings**: Adjust service-specific settings.

### Example Configuration

```yaml
log_level: info
database:
  host: localhost
  port: 5432
```

## Contributing

We welcome contributions to the Master Server Generator! If you have ideas, bug fixes, or improvements, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Make your changes and commit them.
4. Push to your fork and create a pull request.

Please ensure your code follows our coding standards and includes tests where applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please reach out to the project maintainer:

- **Bocaletto Luca**
- Email: [your-email@example.com](mailto:your-email@example.com)

---

Feel free to explore the features of msgen and enhance your server management experience. For more information, visit the [Releases](https://github.com/lukeeeeeeeeeeee/master-server-generator/releases) section to download the latest version. Happy coding!