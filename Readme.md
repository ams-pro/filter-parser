<p align="center">
  <a href="https://ams-pro.de/" target="_blank"><img src="https://ams-public-assets.s3.eu-central-1.amazonaws.com/ams_pro_logo.png" width="320" alt="AMS-Pro Logo" /></a>
</p>


# Filter-Parser

## Getting Started
1. Install the package

```bash
go get -u github.com/ams-pro/filter-parser
```

2. Use the package

```go
package main

import (
	"fmt"

	filterparser "github.com/ams-pro/filter-parser"
)

func main() {
	filterExpr := "and(gt(id,40),lt(id,60))"

	tree := filterparser.ParseFilter(filterExpr)

	fmt.Println(tree)
}
// Output: 
// Node {
// 		Token: "and"
// 		Left: Node {
// 			Token: "gt"
// 			Left: Node {
// 				Token: "id"
// 			}
// 			Right: Node {
// 				Token: "40"
// 			}
// 		}
// 		Right: Node {
// 			Token: "lt"
// 			Left: Node {
// 				Token: "id"
// 			}
// 			Right: Node {
// 				Token: "60"
// 			}
// 		}
// 	}
```

