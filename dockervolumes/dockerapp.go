package dockervolumes

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//ListVolumes works like docker volume ls
func ListVolumes() (volumeList []*types.Volume, err error) {
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		return volumeList, err
	}

	emptyArgs := filters.Args{}
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	if err != nil {
		return volumeList, err
	}
	volumeList = volumeListBody.Volumes

	//print to took identical to docker volume ls
	fmt.Println("DRIVER\t\tVOLUME NAME")
	for _, vol := range volumeList {
		fmt.Printf("%v\t\t%v\n", vol.Driver, vol.Name)
	}
	return volumeList, err
}

//VolumeExpiriment gives more detail on volumes, similar to docker volume inspect
func VolumeExpiriment() {
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	emptyArgs := filters.Args{}
	//
	volumeListBody, err := cli.VolumeList(context.Background(), emptyArgs)
	fmt.Printf("volume list body content \n%v\n", volumeListBody)
	fmt.Print("Volumes:\n")
	v := volumeListBody.Volumes

	// TODO add filter on driver
	for _, vol := range v {
		fmt.Printf(" Name: %s\n MountPath: %s\n Driver: %s\n\n", vol.Name, vol.Mountpoint, vol.Driver)
	}
}
