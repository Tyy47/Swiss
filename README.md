<h1 align="center">Swiss</h1>
<p align="center">A cli army knife of tools for developers</p>
<hr>

[Documentation](./DOCS.md) 

## What is Swiss?
Swiss is a multitool in the terminal. It was built on the premise of being reliable and transportable, a tool that you can carry everywhere with no issues.

## What does Swiss provide?
Swiss provides reliability to a developer who just wants things to work. We all love to tinker, but sometimes we want something to just work out of the box with no issues. Swiss is meant to be a tool that provides everything a developer might need with no issues.

## Install

### Binary:
Prepackaged binaries are available to download via the releases section. After downloading one of the executables, place it in a PATH'd location to access it via the command line.

### Go:
Running ``go install`` inside of the Swiss directory will install the program in a PATH'd location via Go.

### Updates via Swiss
If an update is available for Swiss. Run `swiss -u` or `swiss update` to update Swiss. It will make an install directory locally and clone the Swiss repo into that folder then cd into it. It will then prompt you for the Go install route or the local/bin install, which is only available for Linux. After install, make sure either Go's install path is added to your system PATH, for the Linux only option, make sure ~/.local/bin is added to your PATH.

## FAQ

### Why Go?
Go is a simple language that doesn't get in the way of developing programs. I needed a language for easy system calls and no compiler annoyance. Go is most likely the best choice in that regard.

### Why are there no packages outside of the standard library?
I want Swiss to be "my own". Meaning, I want to create everything from scratch. Another perk of this, it lowers the risk of supply chain attacks through libraries as well as breakage from outside packages. I want this program to be expanded upon indefinitely without the reliance of others for critical systems. If there is any problems in the future with compiling the program it's because of something I broke or the standard library got changed in some way.

### There is a bug in Swiss, how do I report it?
Create an issue report stating what command you we're running, what you expected to happen, and what actually happened.

### How can I contribute to Swiss?
I'd like this project to be kept as my own but I'm open to suggestions of improvements or other tools that could be added. Feel free to open a issue report with the "Suggestion" post flag and I'll take a look!
