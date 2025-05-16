package chaos

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CreateEntry(title string) error {
	safeTitle := strings.ReplaceAll(strings.ToLower(title), " ", "-")
	filename := fmt.Sprintf("entries/%s_%s.md", time.Now().Format("2006-01-02"), safeTitle)

	content := fmt.Sprintf(
		`# %s
		**Date:** %s

	## Hyphthesis

	...

	## Experiments

	... 

	## Logs

	...
	
	## Notes

	...

	## Mutation Plans

	...`, title, time.Now().Format(time.RFC1123))
	// prem file mode means permission on this file
	// 0644 6 -  r+w 4-- r

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
