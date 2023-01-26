# protoc-gen-go-enums

Protobuf plugin that generates go enum names with type prefix stripped.

Using Buf generation:

```yaml
version: v1
plugins:
  - name: uwentity
    out: gen/go
    opt:
      - paths=source_relative
```
