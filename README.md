# webz

**webz** is a lightweight, fast, and simple website server written in Go.  
It is designed to handle multiple clients efficiently, while requiring minimal setup — just run the program and start serving.

webz aims to be as small and straightforward as possible, making it ideal for quick file hosting and static site serving.

Current version: **0.2**  
See the [Roadmap](##Roadmap) section for planned and implemented features.

---

## Installation

You can install **webz** either by downloading a pre-compiled binary from the [releases](https://github.com/sataniccoder/webz/releases) page or by building from source.

### Build from source

```bash
git clone https://github.com/sataniccoder/webz.git
cd webz
go build webz.go
```

This will produce the `webz` executable binary.

---

## Features

1. **Simple file hosting** — easily share files and folders over HTTP using command-line arguments.
2. **Live editing** — make changes to your files, and webz will serve updates automatically without needing a restart.
3. **Auto-update** — webz can check for updates and update itself while preserving your current configuration.

---

## Modes

- **Default mode:** Serves a primary HTML file. Requires an `index.html` file in the directory where webz runs or in the specified folder.
- **List mode (`--list`):** Allows clients to browse files and directories via a web interface, view files in the browser, and download them directly.

---

## Recent updates

- Added `--list` mode for directory browsing.
- Added `--silent` option to suppress output logs.

---

## Roadmap

- [X] **v0.2** — Version checking, logging, password protection.
- [ ] **v0.3** — HTTPS support with self-signed certificates and static configuration files.
- [ ] **v0.4** — Provide a compiled release version and an installation script.
- [ ] **v0.5** — Ability to run as a service for long-term hosting with static configs.
- [ ] **v0.6** — Self-upgrade functionality (automatic or manual updates for the compiled binary).

---

## Notes

webz is designed to remain small and fast. If the project grows significantly, there may be two versions in the future:

- A **slim** version for minimal resource usage.
- An **expanded** version with additional features.

---

If you have any questions, feature requests, or issues, feel free to open an issue or reach out!
