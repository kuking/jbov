package api

import (
	"testing"
	"os"
	"path/filepath"
	"github.com/kuking/jbov/api/md"
	"github.com/stretchr/testify/assert"
)

// CanCreateJBOV

func TestCanCreateJBOV_ShouldFailWithInvalidJBOV(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Uniqid = "invalid"

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "JBOV object should be valid")
}

func TestCanCreateJBOV_AllVolumeMountPointsMustExist(t *testing.T) {
	jbov := givenValidJBOV()
	givenMountPointsExist(&jbov)
	defer cleanupMountPoints(&jbov)

	ok, err := CanCreateJBOV(&jbov)

	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestCanCreateJBOV_failsWhenVolumeMountPointDoNotExist(t *testing.T) {
	jbov := givenValidJBOV()
	delete(jbov.Volumes, "vol2")
	jbov.Volumes["vol1"].LastMountPoint = "/this-surely-does-not-exists/"

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "Volume mount point for \"vol1\" does not exist: /this-surely-does-not-exists/")
}

func TestCanCreateJBOV_failsWhenVolumeMountPointIsAFile(t *testing.T) {
	jbov := givenValidJBOV()
	delete(jbov.Volumes, "vol2")
	jbov.Volumes["vol1"].LastMountPoint = "/etc/passwd"

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "Volume mount point for \"vol1\" is not a directory: /etc/passwd")
}

func TestCanCreateJBOV_failsWhenThereIsAMetadataFile(t *testing.T) {
	jbov := givenValidJBOV()
	givenMountPointsExist(&jbov)
	defer cleanupMountPoints(&jbov)
	f, _ := os.Create( filepath.Join(jbov.Volumes["vol1"].LastMountPoint, md.JBOV_FNAME))
	f.Close()

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "Volume mount point for \"vol1\" seems to be part of an existing JBOV: \".jbov.metadata\" file found")
}

func TestCanCreateJBOV_failsWhenThereIsAUniqIDFile(t *testing.T) {
	jbov := givenValidJBOV()
	givenMountPointsExist(&jbov)
	defer cleanupMountPoints(&jbov)
	f, _ := os.Create( filepath.Join(jbov.Volumes["vol1"].LastMountPoint, md.UNIQID_FNAME))
	f.Close()

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "Volume mount point for \"vol1\" seems to be part of an existing JBOV: \".jbov.uniqid\" file found")
}

func TestCanCreateJBOV_ShouldNotHaveDeletedFiles(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Deleted = make(map[string]*md.Deleted)
	jbov.Deleted["a/path"] = &md.Deleted{Ts: 123}

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "An about to be created JBOV should not have deleted files")
}

func TestCanCreateJBOV_shouldNotStartWithDeprecatedVolumes(t *testing.T) {
	jbov := givenValidJBOV()
	givenMountPointsExist(&jbov)
	defer cleanupMountPoints(&jbov)
	jbov.Volumes["vol1"].Deprecated = true

	ok, err := CanCreateJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualError(t, err, "An about to be created JBOV should not start with a deprecated volume: vol1")
}

