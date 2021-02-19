# Battlesnake Rules Wasm

This is a WebAssembly port of the [Battlesnake rules][rules].

The goal is to enable executing [Battlesnake][battlesnake] game rules in
environments such as web browsers, Node.js, and Electron.

## Status

This project is experimental and your milage may vary.

## Supported Rulesets

- [x] Standard
- [x] Solo
- [x] Constrictor
- [ ] Squad
- [ ] Royale

## Known Issues

- Poor memory management
- Large WASM binary size
- Use of Go's `syscall/js` which is also marked as experimental
- Lack of automated tests

Contributions welcome!

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

### Exported Functions

All functions require a single argument, which is expected to be a JSON-encoded
string. The JSON fields are interpreted as options for each function.

Likewise, the return value will be a JSON-encoded string in the case of
success, or null in the case of an error.

See `examples/index.html` for some examples of the option signature for each
available function. That example page also loads `main.wasm` file if you have
compiled it, which means you can try out those functions straight from that
page using your browser's JavaScript console.

Documentation is currently sparse, so you may need to read the source code of
`main.go` and possibly even the types defined upstream in the
[BattlesnakeOfficial/rules][rules] source code.

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
[slack]: https://battlesnake.slack.com/
