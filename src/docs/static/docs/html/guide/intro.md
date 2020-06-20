# Introduction

Welcome to the Whirlwind Guide.  This guide is intended to go over and help
familiarize any programmer new to Whirlwind with its fundamental concepts,
ideas, syntax, and semantics.  It is by no means comprehensive but should serve
as a good introduction.  It will cover all major language topics although mostly
at a high level.  This guide will also briefly cover installing the Whirlwind
toolchain as well as setting up a very basic starter project without the use of
[Blend](/docs/blend-intro).  You may find other resources such as the [CLI
Reference](/docs/cli-reference) or the [Style Guide](/docs/style-guide).

Note that it is often helpful to reinforce concepts with a little bit of
practice.  For this reason, some may find it helpful to work through the
[Whirlwind Koans](/docs/koans) in tandem with the guide.

## Installation

To install Whirlwind, first visit the [install page](/install) to find a
suitable installer (or installation "strategy") for your platform.  If no such
version exists, you can try building from the compiler source located
[here](https://github.com/ComedicChimera/Whirlwind).  Note that if no installer
exists for your platform, Whirlwind may not *currently* have a version supporting
your platform yet (we are still a very new language and implementing support for
all platforms takes time).  However, a satisfactory version should exist for the
vast majority of people (**TODO: update as necessary** Windows, OSX, several
popular linux distros).  

Regardless, one you have gone through the installation process, you should be
able to type the command `whirl` into the command-line and get a meaningful
result back.

## Setting Up

After you install, the first thing you want to do is pick an editor to program
in.  There are a number of editors for which a satisfactory Whirlwind plugin
exists.  If you are new to programming or don't have a preferred editor, I
recommend VSCode (link below) as it was the editor used to develop the language
and as such is the editor with the most support.  It is also fairly friendly to
newer programmers (who don't want a terminal editor of course).

Here is a table of all the known supported editors and their corresponding plugins:

| Editor | Plugin |
| ------ | ------ |
| [VSCode](https://code.visualstudio.com/) | *insert link here* |

Once you have found a satisfactory editor and have configured it for Whirlwind,
create an empty directory to store your first project and open your editor in
the directory along with a terminal ready to go (either in-editor or along the
side).  Finally, create a blank file in that directory with the `.wrl`
extension. This denotes that the file is a Whirlwind file.  With that, you are
all set and ready to go.
