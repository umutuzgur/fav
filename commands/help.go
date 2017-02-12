package commands

import (
	"fmt"
)

// Help prints how to use fav
func Help() {
	fmt.Print("Fav\n\n",
		"Usage:\n",
		"  fav list\n",
		"  fav save <command>\n",
		"  fav get <index>\n",
		"  fav exec <index>\n",
		"  fav delete <index>\n",
		"  fav gui",
		"  fav --help\n",
		"\nOptions:\n",
		"  --help              print this help message\n")
}
