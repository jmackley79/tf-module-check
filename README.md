# Terraform Module Version Checker
This is a simple command-line program written in Go that scans Terraform files in the current directory to look for module references and versions. It then checks the latest release on GitHub for each referenced module and alerts the user if a newer version is available.

## Installation
To use this program, you need to have Go installed on your system. Once Go is installed, you can download and build the program using the following commands:
```
$ go get github.com/jmackley79/tf-module-check
$ cd $GOPATH/src/github.com/jmackley79/tf-module-check
$ go build
```
This will create an executable file called tf-module-check in the current directory.

## Usage
To use the program, simply run the executable in the directory containing your Terraform files:
```
$ ./tf-module-check
```
The program will scan all .tf files in the current directory and its subdirectories and check for module references and versions. If a newer version of a referenced module is available on GitHub, the program will print an alert message.

## Limitations
This program only works with Terraform files that use the standard module block syntax, and may not work with alternative syntaxes or configurations.
The program only checks for updates on GitHub, and may not detect newer versions available from other sources.
The program only checks for updates for the latest release of a module, and does not support checking for specific versions or releases.
