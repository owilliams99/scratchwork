package dockervolumes

import (
	"fmt"
	"testing"
)

func TestVolumeExpiriment(t *testing.T) {
	VolumeExpiriment()
}

func TestListVolumes(t *testing.T) {
	volumes, _ := ListVolumes()
	fmt.Printf("volumes: %v\n", volumes)
}
