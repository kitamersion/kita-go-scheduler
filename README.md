# kita-go-scheduler

Yet another user-level cron-like application built in Go! It allows you to define and execute tasks based on cron schedules. Tasks can be configured using a YAML file.

## Installation

### Prerequisites

- Go (1.20 or later) installed on your systemm (or just use `shell.nix`)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/kitamersion/kita-go-scheduler.git
   cd kita-go-scheduler
   ```
2. Build the binary:
   ```bash
   go build -o kita-go-scheduler
   ```
3. Move the binary to a directory in your `PATH`:
   ```bash
   mv kita-go-scheduler /usr/local/bin/
   ```

## How to Use

### Running kita-go-scheduler

1. Start the application:

   ```bash
   kita-go-scheduler
   ```

   The application will look for a configuration file at `~/.config/kita-go-scheduler/config.yml`.

2. If the configuration directory or file does not exist, kita-go-scheduler will create the directory and copy an example `config.yml` file from the project root. You can modify this file to define your tasks.

### Configuration File

The configuration file is located at:

```
~/.config/kita-go-scheduler/config.yml
```

### Log File

Logs for task execution are written to:

```
~/.config/kita-go-scheduler/kita-go-scheduler.log
```

You can enable or disable logging using the `logs.enabled` property in the configuration file.

## Defining Tasks

Tasks are defined in the `config.yml` file using the following structure:

```yaml
tasks:
  - name: "Example Task 1"
    schedule: "@every 1m"
    command: "echo 'Hello, World!'"
  - name: "Backup Task"
    schedule: "0 2 * * *"
    command: "rsync -av /source /destination"

logs:
  enabled: true
```

### Task Properties

- **name**: A unique name for the task.
- **schedule**: A cron expression or predefined interval (e.g., `@every 1m`).
- **command**: The shell command to execute when the task runs.

### Examples of Cron Usage

- `@every 5s`: Run every 5 seconds.
- `@hourly`: Run once every hour.
- `0 5 * * *`: Run daily at 5:00 AM.
- `15 10 * * 1-5`: Run at 10:15 AM, Monday through Friday.

For more details on cron syntax, refer to [Cron Expressions](https://crontab.guru/).

## Troubleshooting

1. **Configuration File Not Found**: If kita-go-scheduler fails to locate `config.yml`, ensure it exists in the correct directory (`~/.config/kita-go-scheduler/`).
2. **Task Not Executing**: Verify the cron expression and the command in the `config.yml`. Check logs for any errors.

Enable logs in `config.yml`


# Disclaimer

This project is built _purely_ for learning Go, dont expect much from this project... and always thank you in advance for any feedback, issues reports and contributions!

