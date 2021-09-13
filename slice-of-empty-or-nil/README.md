## why

nil sliceとempty sliceの挙動の違いがわかっていなかったので、その記録

## what

長さ0でキャパシティ0のスライスを返す関数の書き方としては、以下の4つの書き方がある。

- `nil` を返す。
  - この場合は、nil slice。
- `[]type{}` を返す。
  - この場合は、empty slice。
- `var result []type` で宣言したのを。そのまま返す
  - この場合は、nil slice。
- `make([]type, 0, 0)` で作って返す。
  - この場合は、empty slice。

### 振る舞いその1

lengthやcapacityは、nil sliceとempty sliceとは一致し、両者ともに0である。

### 振る舞いその2

nil との等しさが成り立つのはnil sliceだけ。

振る舞いその1と合わせて、空かどうかは、長さで判定したほうがよいだろう。

### 振る舞いその3

`json.Marshal`すると、nil sliceは`"null"`になり、empty sliceは`"[]"`になる。

APIでnullableな型であることを明示するのは面倒だから、基本的にempty sliceを返したほうがいいような気がする。
