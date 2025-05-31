# webz

**webz** is a lightweight, fast, and simple website server written in Go.  
It is designed to handle multiple clients efficiently, while requiring minimal setup — just run the program and start serving.

webz aims to be as small and straightforward as possible, making it ideal for quick file hosting and static site serving.

Current version: **0.2**  
See the Roadmap section for planned and implemented features.

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
- **Setup (`--setup`)**: installs the program as a service and sets up a config file

---

## Recent updates

- Added version checking to let you know if theres a new version out
- Added a logging system
- Added password protection for the site (if user want's it)

---

## Roadmap

- [X] **v0.2** — Version checking, logging, password protection.
- [ ] **v0.3** — Static config file along with been able to use it as a service with a `--setup` mode
- [ ] **v0.4** — Provide a compiled release version and an installation script.
- [ ] **v0.5** — Ability to run as a service for long-term hosting with static configs.
- [ ] **v0.6** — Self-upgrade functionality (automatic or manual updates for the compiled binary).

---

## Notes
This project is desinged to be as easy whilst having the smallest setup time possable due to this desing philiopshy of the project   
all fetures and things that might be tradationaly be added as a .py or bash script too the project (EG: automated setup) will be built into the program and only accsesed when you specify   
for example the `--list` function won't have any accsess too the `--setup` function
