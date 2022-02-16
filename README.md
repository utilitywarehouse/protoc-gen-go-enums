# protoc-gen-go-enums

Protobuf plugin that generates go enum names with type prefix stripped.

We provide a [Buf](https://buf.build/) plugin to make generation super simple. 
Add the following to a `buf.gen.yaml` file.

```yaml
version: v1
plugins:
  # ... your other plugins
  - remote: buf.build/utilitywarehouse/plugins/go-enums:v1.0.0
    out: gen/go
    opt:
      - paths=source_relative
```
