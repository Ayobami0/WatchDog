# 🐕 Watchdog

A lightweight and elegant file watcher that executes commands when files change. Perfect for development workflows, build processes, and automated testing.
Like a loyal dog 🐕 for your development workflows.

## 🦮 Introduction

Watchdog sits and stays by your files, watching for changes. When it spots a modification, it faithfully executes your specified commands. Perfect for live reloading, test running, and build automation.

## 🚀 Features

- 🐾 Tracks multiple files and directories simultaneously
- 🦴 Fast and lightweight (no external dependencies)
- 🐕‍🦺 Beautiful, colored console output
- 🏃‍♂️ Configurable refresh intervals
- 🔔 Smart command execution on file changes
- 🤫 Silent mode for quiet operation

## 📦 Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/watchdog.git

# Navigate to the directory
cd watchdog

# Build the binary
go build -o watchdog

# Move to a directory in your PATH (optional)
sudo mv watchdog /usr/local/bin/
```

## 🎯 Usage

```bash
watchdog [OPTIONS]... command
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-f` |  Comma-separated list of files to watch | Current directory |
| `-s` |  Turn off update messages | `false` |
| `-i` | Duration between checks | `100ms` |
| `-h` |  Show help message | - |

### Examples

1. Watch and run Go tests when files change:
```bash
watchdog go test ./...
```

2. Watch specific files for a web project:
```bash
watchdog -f main.go,templates/,static/ -i 500ms go run main.go
```

3. Quiet mode for build scripts:
```bash
watchdog -s make build
```

4. Watch multiple files and run compiler:
```bash
watchdog -f src/*.ts -i 1s tsc
```

## 🖥️ Output Format

Watchdog provides friendly, detailed output:

```
 __          __   _       _     _____
 \ \        / /  | |     | |   |  __ \
  \ \  /\  / /_ _| |_ ___| |__ | |  | | ___   __ _
   \ \/  \/ / _  | __/ __|  _ \| |  | |/ _ \ / _  |
    \  /\  / (_| | || (__| | | | |__| | (_) | (_| |
     \/  \/ \____|\__\___|_| |_|_____/ \___/ \__, |
                                              __/ |
                                             |____/

          -                                -
           --                            --
           =-=-                        -=-=
           ---==                      ==---
           -==-=-                    -=-==-
           -==--=--                --=--==-
            ++=====--            --=====++
            +++==*==--------------==*==+++
            -++=-=*--=-=------=-=-=*=-=++-
             ++=-====================-=++
             ++--=========**=========--++
             ++---=*===*======*===*=---++
             -+=-=*--=**==++==**=--*=-=+-
              --===-==+-==**==-+==-===--
              --==---=+-=-==-=-+=---==--
              -=*==----=------=----==*=-
             -=====-------==-------=====-
              ---=====--======--=====---
              ----====--======--====----
                ---==--========--==---
                 --==--+**++**+--==-- 
                ---==--+++==+++--==---
                --=======*==*=======--
                =-======------======-=
               -=-=++++=------=++++=-=-
               =====++++======++++=====
           *#*--==-=-====-==-====-=-==--*#*
            ==------==-++++++++-==------==
            -=-==**=----=++++=----=**==-=-
              ----=----------------=---- 
                 ----===-=**=-===----
                     ----====----

🐶 Watcher Started
  └── Time: 2024-11-03 15:04:05
  └── Command: go run main.go
  └── Refresh: 100ms
  └── Watching 2 file(s):
      └──  utils.go
      └──  main.go
  └── Running: ls
      │ README.md
      │ file
      │ file.go
      │ go.mod
      │ logger.go
      │ main.go
      └── ✓ Completed (took 0ms)
📁 [15:04:05] main.go
  └── Modified at: 15:04:05.123
  └── Reloading...
  └── Running: go run main.go
      │ Building...
      │ Server starting on port 8080
      └── ✓ Completed (took 1.2s)
```

## 🚨 Limitations

- Maximum of 10 watched files/directories (to prevent resource strain)
- Does not follow symbolic links
- Directory watches trigger on file deletion events
- No recursive directory watching (yet!)

## 🤝 Contributing

We'd love your help making Watchdog even better! Here's how:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/good-boy`)
3. Commit your changes (`git commit -m 'chore: Add some good boy feature'`)
4. Push to the branch (`git push origin feature/good-boy`)
5. Open a Pull Request

> [!IMPORTANT]
> See [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/)


## 📜 License

Released under MIT License - feel free to pet, play with, and modify your Watchdog!

## 🎖️ Version History

- 1.0.0
  - Initial release
  - Basic watching capabilities
  - Colored output
  - Configurable intervals

## 🐾 Future Plans

- [ ] Recursive directory watching
- [ ] Regular expression patterns
- [ ] Multiple command execution
- [ ] WebSocket notifications
- [ ] Configuration file support

---

🐕 Happy Watching! Woof! 🦴
