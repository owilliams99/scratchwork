package tarball

import (
	"fmt"
	"testing"
)

func TestWriteToTar(t *testing.T) {
	err := WriteToTar()
	if err != nil {
		fmt.Println(err)
	}
}
