package version

import (
	"fmt"

	"github.com/acquia/hub/git"
)

var Version = "2.5.1"

func FullVersion() (string, error) {
	gitVersion, err := git.Version()
	if err != nil {
		gitVersion = "git version (unavailable)"
	}
	return fmt.Sprintf("%s\nhub version %s", gitVersion, Version), err
}
