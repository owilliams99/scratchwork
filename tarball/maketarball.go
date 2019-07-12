package tarball

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//WriteToTar takes in list of volumes and saves info to tarfile in archive
func WriteToTar() {

	//read
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	volumes, _ := ListVolumes()

	for _, vol := range volumes {
		volName := vol.Name + ".txt"

		hdr := &tar.Header{
			Name: volName,
			Size: int64(len(vol.Mountpoint)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(vol.Mountpoint)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	//write
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}

//ListVolumes is straight copy from docker.go
func ListVolumes() (volumeList []*types.Volume, err error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		return volumeList, err
	}

	emptyArgs := filters.Args{}
	//
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	if err != nil {
		return volumeList, err
	}
	volumeList = volumeListBody.Volumes
	return volumeList, err
}
