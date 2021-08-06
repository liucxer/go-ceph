package main

import (
	"fmt"

	"github.com/liucxer/go-ceph/rados"
	"github.com/liucxer/go-ceph/rbd"
)

const (
	DefaultRadosConfigFile = "/etc/ceph/ceph.conf"
	DefaultBaseImageSize   = 10 * 1024 * 1024 * 1024
	DefaultPoolName        = "test"
)

func main() {
	// connect to the cluster
	conn, _ := rados.NewConn()
	if err := conn.ReadConfigFile(DefaultRadosConfigFile); err != nil {
		fmt.Printf("Rbd read config failed: %v", err)
		return
	}
	if err := conn.Connect(); err != nil {
		fmt.Printf("Rbd connect failed: %v", err)
		return
	}

	// connect to the pool
	ioctx, err := conn.OpenIOContext(DefaultPoolName)
	if err != nil {
		fmt.Printf("Rbd open pool failed: %v", err)
		return
	}

	// create base image
	baseImageName := "hj-test"
	_, err = rbd.Create(ioctx, baseImageName, DefaultBaseImageSize, int(rbd.RbdFeatureLayering))
	if err != nil {
		fmt.Printf("Rbd create image failed: %v", err)
		return
	}

	// we should open base image first
	// if err := img.Open(); err != nil {
	// 	fmt.Printf("Rbd open image  failed: %v", err)
	// 	return
	// }

	// defer img.Close()

	// create snapshot
	// snapName := "test-snap"
	// snapshot, err := img.CreateSnapshot(snapName)
	// if err != nil {
	// 	fmt.Printf("Rbd create snapshot failed: %v", err)
	// 	return
	// }

	// protect snapshot
	// if err := snapshot.Protect(); err != nil {
	// 	fmt.Printf("Rbd create snapshot failed: %v", err)
	// 	return
	// }

	// make a clone image based on the snap shot
	// cloneImageName := "clone-test"
	// _, err = img.Clone(snapName, ioctx, cloneImageName, rbd.RbdFeatureLayering)
	// if err != nil {
	// 	fmt.Printf("Rbd clone snapshot failed: %v", err)
	// 	return
	// }

	return
}