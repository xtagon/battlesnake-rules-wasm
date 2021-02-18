# Battlesnake Rules Wasm
========================

This is a WebAssembly port of the [Battlesnake rules](rules).

The goal is to enable executing [Battlesnake](battlesnake) game rules in
environments such as web browsers, Node.js, and Electron.

## Status

This project is experimental and your milage may vary.

## Known Issues

- Poor Error handling (such as panics if arguments are passed incorrectly)
- Poor memory management
- Large WASM binary size
- Use of Go's `syscall/js` which is also marked as experimental
- Undocumented API for using exported functions
- Lack of automated tests

Contributions welcome!

## Supported Rulesets

Presently only the "Standard" ruleset is supported.

## How to Use

### Prerequisites

- Go 1.15
- GNU Make (Optional)

### Compiling

If GNU Make is available:

```shell
make
```

If GNU Make is not available, read `Makefile` to know how to build without it.

### Loading WebAssembly

After compiling, copy the files in `/dist` to your project, including
`main.wasm` and `wasm_exec.js`, so that you can instantiate it using
WebAssembly. See `examples/index.html` for a basic example.

After it is loaded, a `window.BattlesnakeRules` or `global.BattlesnakeRules`
will be set which is an object containing the exported functions. Typically the
global will be scoped to `window` in browsers, or `global` in environments such
as Node.js and Electron.

The function signatures are not yet documented, so you will need to understand
the exported functions in `main.go` in order to use them. Note that it is
currently required to stringify and parse JSON arguments and return values on
the consuming end, as JSON arguments are passed as JSON-encoded strings to the
WASM functions, not as JS objects or Go structs.

## Open Invite

If you have any questions, or just wish to geek out and chat about Battlesnake
feel free to reach out!

If you're using this library for a project, I'd love to hear about it.

You can reach me at [xtagon@gmail.com](mailto:xtagon@gmail.com), or catch me in
[Battlesnake Slack][slack] (username: `@xtagon`).

## License

This project is released under the terms of the [MIT License](LICENSE.txt).

[battlesnake]: https://play.battlesnake.com/
[rules]: https://github.com/BattlesnakeOfficial/rules
