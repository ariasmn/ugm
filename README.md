ugm
======

<p>
    <a href="https://github.com/ariasmn/ugm/releases"><img src="https://img.shields.io/github/v/release/ariasmn/ugm" alt="Latest Release"></a>
    <a href="https://goreportcard.com/report/github.com/ariasmn/ugm"><img src="https://goreportcard.com/badge/ariasmn/ugm" alt="Go ReportCard"></a>
</p>

A TUI to view information about UNIX users and groups.

![How it works](https://user-images.githubusercontent.com/33121576/180660203-9c9f3801-5298-4ebc-b23a-e25e35582974.gif)

## Installation

Install directly using Go: 
```
go install github.com/ariasmn/ugm@latest
```
or download from [releases](https://github.com/ariasmn/ugm/releases).

## Usage

To launch the tool, run the `ugm` command.

| KeyBoard              | Description                           |
| :------:              | :------------------------------------:|
|  Ctrl+c / q / Esc     | Exit                                  |
|  Tab                  | Switch between user and group view    |
|    ↑ / k              | Previous item                         |
|    ↓ / j              | Next item                             |
|    ← / h              | Previous page                         |
|    → / l              | Next page                             |
|    /                  | Enter search keywords                 |
|  Enter                | Apply search                          |


## Notes

`ugm` currently supports Linux, FreeBSD, OpenBSD, and NetBSD only. Although macOS is UNIX-based, it is not supported: the account parser and release targets exclude Darwin. macOS uses [Directory Services](https://developer.apple.com/documentation/devicemanagement/directoryservice) to manage users and groups, so `/etc/passwd` and `/etc/group` are not an authoritative account source there.

## Built with
 - [bubbletea](https://github.com/charmbracelet/bubbletea) and its ecosystem
 - [bubble-table](https://github.com/Evertras/bubble-table)