# chai-lang.dev

This is the source for the [Chai programming language](https://github.com/ComedicChimera/chai) website.  
You can find it hosted here *insert link*.

## Dependencies:

This website is built with Django and Svelte.js.

Download and Install:

- Python 3.9 or later
- Node.js (and NPM)

## Building and Running

All actions in this respository can be easily run from the Powershell script `cacao.ps1`. 

Install Required Dependencies:

    cacao setup

Build Static Website Content:

    cacao build

Run Development Server:

    cacao rundev

## Repository Layout

| Directory | Purpose |
| --------- | ------- |
| `content` | Static, markdown content for the website (docs, etc.) |
| `chaisite` | The main Django project |
| `chaisite/whirlsite` | Primary project configuration and common files |
| `chaisite/home` | Home page and primary subpages |
| `chaisite/docs` | Documentation |
| `chaisite/packages` | Package documentation/index |

## Guide Layout

TBD


