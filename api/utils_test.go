package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/kuking/jbov/api/md"
	"os"
	"io/ioutil"
)

// uniq ids creation and validation

func TestGeneratedJbovUniqIdValidates(t *testing.T) {
	uniqid := GenerateJbovUniqId()
	assert.True(t, IsJbovUniqId(&uniqid), "GenerateJbovUniqId and IsJbovUniqId does not seem to agree")
}

func TestGeneratedVolUniqIdValidates(t *testing.T) {
	uniqid := GenerateVolumeUniqId()
	assert.True(t, IsVolumeUniqId(&uniqid), "GenerateVolumeUniqId and IsVolumeUniqId does not seem to agree")
}

func TestJbovUniqIdDoesNotValidatesAsVolId(t *testing.T) {
	jbovUniqId := GenerateJbovUniqId()
	assert.False(t, IsVolumeUniqId(&jbovUniqId), "GenerateJbovUniqId should not be validated by IsVolumeUniqId")
}

func TestGenerateVolumeUniqIdDoesNotValidateAsJBovId(t *testing.T) {
	jbovUniqId := GenerateVolumeUniqId()
	assert.False(t, IsJbovUniqId(&jbovUniqId), "GenerateVolumeUniqId should not be validated by IsJbovUniqId")
}

// cname validation

func TestValidCname(t *testing.T) {
	valids := []string{"valid", "valvalval", "long_valid", "lowercase_is_ok", "vol1", "vol", "numbers_ok_111"}
	for _, valid := range valids {
		assert.True(t, isValidCname(&valid), "'"+valid+"' should be a valid cname")
	}
}

func TestNotValidCname(t *testing.T) {
	invalids := []string{"", "AA", "aa", "UPPERCASE_NOT_OK", "super_long_one_is_not_valid"}
	for _, invalid := range invalids {
		assert.False(t, isValidCname(&invalid), "'"+invalid+"' should be an invalid cname")
	}
}

func TestValidCname_InvalidSymbols(t *testing.T) {
	all_invalids := "!@£$%^&*-=~`[]{}();:'\",./<>?j"
	for _, c := range all_invalids {
		st := string(c)
		assert.False(t, isValidCname(&st), "'"+st+"' should not be a valid cname")
	}
}

//  IsValidJBOV

func TestIsValidJBOV_happyPath(t *testing.T) {
	jbov := givenValidJBOV()

	ok, err := IsValidJBOV(&jbov)

	assert.True(t, ok)
	assert.NoError(t, err, "this JBOV should be valid")
}

func TestIsValidJBOV_invalidCname(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Cname = "INVALID"

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV cname is not valid", "invalid cname should fail")
}

func TestIsValidJBOV_invalidJBovUniqId(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Uniqid = "invalid"

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV UniqId does not looks valid", "invalid jbov uniq should fail")
}

func TestIsValidJBOV_noVolumes(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes = nil

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV requires at least one volume", "jbov without volumes should fail")
}

func TestIsValidJBOV_VolumeWithInvalidName(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes["INVALID_VOL_CNAME"] = jbov.Volumes["vol1"]

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV volume has an invalid cname", "volume with invalid cname should fail")
}

func TestIsValidJBOV_VolumeWithInvalidUniqid(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes["vol1"].Uniqid = GenerateJbovUniqId()

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV volume has an invalid uniqid", "volume with invalid uniqid should fail")
}

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


// utility

func givenValidJBOV() md.JBOV {
	vols := make(map[string]*md.Volume)
	vols["vol1"] = &md.Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol1", Deprecated: false }
	vols["vol2"] = &md.Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol2", Deprecated: false }
	return md.JBOV{Cname: "valid", Uniqid: GenerateJbovUniqId(), Volumes: vols}
}

func givenMountPointsExist(jbov *md.JBOV) {
	for _, vol := range jbov.Volumes {
		dir, _ := ioutil.TempDir(os.TempDir(), "")
		vol.LastMountPoint = dir
	}
}

func cleanupMountPoints(jbov *md.JBOV) {
	for _, vol := range jbov.Volumes {
		os.RemoveAll(vol.LastMountPoint)
	}

}
