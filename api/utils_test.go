package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/kuking/jbov/api/md"
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
		assert.True(t, isValidCname(&valid), "'" + valid + "' should be a valid cname")
	}
}

func TestNotValidCname(t *testing.T) {
	invalids := []string{"", "AA", "aa", "UPPERCASE_NOT_OK", "super_long_one_is_not_valid"}
	for _, invalid := range invalids {
		assert.False(t, isValidCname(&invalid), "'" + invalid + "' should be an invalid cname")
	}
}

func TestValidCname_InvalidSymbols(t *testing.T) {
	all_invalids := "!@£$%^&*-=~`[]{}();:'\",./<>?j"
	for _, c := range all_invalids {
		st := string(c)
		assert.False(t, isValidCname(&st), "'" + st + "' should not be a valid cname")
	}
}


// integrated IsValidJBOV

func TestCanCreate_happyPath(t *testing.T) {
	jbov := givenAValidJBOV()

	ok, err := IsValidJBOV(&jbov)

	assert.True(t, ok)
	assert.NoError(t, err, "this JBOV should be valid")
}

func TestCanCreate_invalidCname(t *testing.T) {
	jbov := givenAValidJBOV()
	jbov.Cname = "INVALID"

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV cname is not valid", "invalid cname should fail")
}

func TestCanCreate_invalidJBovUniqId(t *testing.T) {
	jbov := givenAValidJBOV()
	jbov.Uniqid = "invalid"

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV UniqId does not looks valid", "invalid jbov uniq should fail")
}

func TestCanCreate_noVolumes(t *testing.T) {
	jbov := givenAValidJBOV()
	jbov.Volumes = nil

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV requires at least one volume", "jbov without volumes should fail")
}

func TestCanCreate_VolumeWithInvalidName(t *testing.T) {
	jbov := givenAValidJBOV()
	jbov.Volumes["INVALID_VOL_CNAME"] = jbov.Volumes["vol1"]

	ok, err := IsValidJBOV(&jbov)

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV volume has an invalid cname", "volume with invalid cname should fail")
}

func givenAValidJBOV() md.JBOV {
	vols := make(map[string]md.Volume)
	vols["vol1"]=md.Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol1", Deprecated: false }
	vols["vol2"]=md.Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol2", Deprecated: false }
	return md.JBOV{Cname: "valid", Uniqid: GenerateJbovUniqId(), Volumes: vols}
}