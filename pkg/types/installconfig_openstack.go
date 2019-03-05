// +build openstack

package types

import (
	"sort"

	"github.com/openshift/installer/pkg/types/openstack"
)

func init() {
	PlatformNames = append(PlatformNames, openstack.Name)
	sort.Strings(PlatformNames)
}
