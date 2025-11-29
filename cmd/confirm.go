package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func confirmDelete(resourceType, resourceKey string, force bool) bool {
	if force {
		return true
	}

	fmt.Printf("Are you sure you want to delete %s '%s'? [y/N]: ", resourceType, resourceKey)
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	return response == "y" || response == "yes"
}
