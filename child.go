package child

import (
	"fmt"

	parent "github.com/bstoll/import-visibility-parent"
)

func Child() {
	fmt.Println("Hello from child v2")
	parent.Parent()
}
