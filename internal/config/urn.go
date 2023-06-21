package config

import (
	"fmt"
)

func BuildURN(resourceType string, id string) string {
	return fmt.Sprintf("urn:%s:%s:%s", AppConfig.URN.Namespace, resourceType, id)
}
