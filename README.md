# MyGoListPackage
- Goで線形リストを使用するためのパッケージです
## データの型
```Go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
```
## 関数一覧
- `listpushback`
- `listpushfront`
- `listsize`
- `listlast`
- `listclear`
- `listat`
- `listreverse`
- `listforeach`
- `listforeachif`
- `listfind`
- `listremoveif`
- `listmerge`
## Usage
### install
```bash
git clone git@github.com:JimpeiYamamoto/MyGoListPackage.git piscine
```
### 例
```Go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    link := &piscine.List{}
    piscine.ListPushBack(link, "Hello")
    piscine.ListPushBack(link, "world")
    piscine.ListPushBack(link, "!!!!")
    for link.Head != nil {
        fmt.Println(link.Head.Data)
        link.Head = link.Head.Next
    }
}
```
#### 出力
```bash
Hello
world
!!!!
```