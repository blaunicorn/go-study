# learning gin videos

[video address]https://www.bilibili.com/video/BV1CE411H7bQ

### first day

```
mkdir ginessential
cd ginessential
cd . > main.go
```

edit helllowrd

```
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

```

```
go run main.go
```

```
go mod init github.com/oceanlearn.teach/ginessential
```

download gin dependency

```
go get -u github.com/gin-gonic/gin
```
