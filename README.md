# gojq

A minimal CLI tool like `jq` that pretty-prints JSON from a stream. If the input isn't valid JSON, it just prints it as-is. Useful for working with structured logs like `zerolog`.

## 🔧 Features

- ✅ Pretty-prints JSON from stdin
- ✅ Gracefully falls back to plain output for non-JSON lines
- 🎨 Optional colored JSON output (default)
- 📦 Supports flags for compact output or no color

## 🧪 Example Usage

    your-app | gojq

Output:

    {
      "level": "info",
      "msg": "Order created",
      "time": "2025-03-21T13:12:17Z"
    }

## 🚀 Install

### From source:

    go install github.com/trythrow/gojq@latest

Make sure `$GOBIN` is in your `$PATH`.

### Or run directly (for testing):

    echo '{"hello":"world"}' | go run .

## 🏁 Flags

| Flag          | Description                     |
|---------------|---------------------------------|
| `--no-color`  | Disable color output            |
| `--compact`   | Output compact JSON (no indent) |

### Examples

    your-app | gojq --no-color
    your-app | gojq --compact
    your-app | gojq --no-color --compact

## 🔗 License

MIT © Ahmad Said
