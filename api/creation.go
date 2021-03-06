package api

import (
	"errors"
	"os"
	"fmt"
	"path/filepath"
	"github.com/kuking/jbov/api/md"
	"io/ioutil"
	"strconv"
)

func CanCreate(jbov *md.JBOV) (bool, error) {
	if ok, _ := jbov.IsValid(); !ok {
		return false, errors.New("JBOV object should be valid")
	}

	if len(jbov.Deleted) > 0 {
		return false, errors.New("An about to be created JBOV should not have deleted files")
	}

	for cname, volume := range jbov.Volumes {
		stat, err := os.Stat(volume.LastMountPoint)
		if os.IsNotExist(err) {
			return false, errors.New(fmt.Sprintf("Volume mount point for \"%s\" does not exist: %s", cname, volume.LastMountPoint))
		}
		if !stat.IsDir() {
			return false, errors.New(fmt.Sprintf("Volume mount point for \"%s\" is not a directory: %s", cname, volume.LastMountPoint))
		}
		_ , err = os.Stat(filepath.Join(volume.LastMountPoint, md.JBOV_FNAME))
		if err == nil {
			return false, errors.New(fmt.Sprintf("Volume mount point for \"%s\" seems to be part of an existing JBOV: \"%s\" file found", cname, md.JBOV_FNAME))
		}
		_ , err = os.Stat(filepath.Join(volume.LastMountPoint, md.UNIQID_FNAME))
		if err == nil {
			return false, errors.New(fmt.Sprintf("Volume mount point for \"%s\" seems to be part of an existing JBOV: \"%s\" file found", cname, md.UNIQID_FNAME))
		}
		if volume.Deprecated {
			return false, errors.New(fmt.Sprintf("An about to be created JBOV should not start with a deprecated volume: %s", cname))
		}
	}

	return true, nil
}

func Create(jbov *md.JBOV) (bool, error) {

	if ok, err := CanCreate(jbov) ; !ok || err != nil {
		return false, err
	}

	// creates the metadata files
	jsonb, err := jbov.Marshal()
	if err != nil {
		return false, err
	}
	for _, vol := range jbov.Volumes {
		filemode64, _ := strconv.ParseUint("744", 8, 32)
		filemode := os.FileMode(filemode64)

		err := ioutil.WriteFile(filepath.Join(vol.LastMountPoint, md.JBOV_FNAME), jsonb, filemode)
		if err != nil {
			return false, err
		}

		err = ioutil.WriteFile(filepath.Join(vol.LastMountPoint, md.UNIQID_FNAME), []byte(vol.Uniqid), filemode)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}