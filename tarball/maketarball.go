package tarball

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//WriteToTar tars any found volumes
func WriteToTar() {
	volumes, _ := ListVolumes()
	for _, vol := range volumes {
		dir, err := os.Open(vol.Mountpoint)
		if err != nil {
			log.Fatal(err)
		}
		defer dir.Close()

		files, err := dir.Readdir(0)
		if err != nil {
			log.Fatal(err)
		}

		destfile := vol.Name + ".tar"

		//
		tarfile, err := os.Create(destfile)
		defer tarfile.Close()

		var fileWriter io.WriteCloser = tarfile

		tarfileWriter := tar.NewWriter(fileWriter)
		defer tarfileWriter.Close()

		for _, fileInfo := range files {

			if fileInfo.IsDir() {
				continue
			}

			file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			header := new(tar.Header)
			header.Name = file.Name()
			header.Size = fileInfo.Size()
			header.Mode = int64(fileInfo.Mode())
			header.ModTime = fileInfo.ModTime()

			err = tarfileWriter.WriteHeader(header)

			_, err = io.Copy(tarfileWriter, file)
		}
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
