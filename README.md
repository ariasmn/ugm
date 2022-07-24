# ugm

A TUI to consult information about UNIX users and groups.

## Installation

Install directly using Go: 
```
go install github.com/ariasmn/ugm@latest
```
or download from [releases](https://github.com/ariasmn/ugm/releases).

## Usage

To launch the tool, run the `ugm` command.

| KeyBoard      | Description                           |
| :------:      | :------------------------------------:|
|  Ctrl+c / q   | Exit                                  |
|  Tab          | Switch between user and group view    |
|    ↑          | Previous item                         |
|    ↓          | Next item                             |
|    ←          | Previous page                         |
|    →          | Next page                             |
|    /          | Enter search keywords                 |
|  enter        | Apply search                          |


## Information

`ugm` only works on UNIX based OS.

On OSX, the information reported will not be accurate. The tool relies on the `/etc/passwd` and `/etc/group` files, which are only consulted in OSX in single-user mode, and the system uses [DirectoryService](https://developer.apple.com/documentation/devicemanagement/directoryservice) to manage user and groups.

## Built with
 - [bubbletea](https://github.com/charmbracelet/bubbletea) and its ecosystem
 - [bubble-table](https://github.com/Evertras/bubble-table)