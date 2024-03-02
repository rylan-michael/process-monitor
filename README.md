# Process Monitor

## Overview

`process-monitor` is a simple Go program designed to monitor the lifecycle of a specific process on Unix-like operating systems. It periodically checks for the existance of a process by name and detects if the process has been restarted by observing changes in its Process ID (PID).

I set out to answer the question "How do you track when a process is stopped or restarted in userland (not kernel modules)?"

By default, `process-monitor` is set to monitor a process named "ControlCenter". This was to verify that a different script created to restart the macOS  ControlCenter when it bugs out and stops working.

## Requirements

- Go (1.11 or later recommended for module support)
- Unix-like operating system (Linux, macOS) with `pgrep` available

## Installation

1. Ensure Go is installed on your system
2. Clone this repository
3. Navigate to the project directory and build the program with `go build`.
