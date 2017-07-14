package md

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
		assert.True(t, IsValidCname(&valid), "'"+valid+"' should be a valid cname")
	}
}

func TestNotValidCname(t *testing.T) {
	invalids := []string{"", "AA", "aa", "UPPERCASE_NOT_OK", "super_long_one_is_not_valid"}
	for _, invalid := range invalids {
		assert.False(t, IsValidCname(&invalid), "'"+invalid+"' should be an invalid cname")
	}
}

func TestValidCname_InvalidSymbols(t *testing.T) {
	all_invalids := "!@£$%^&*-=~`[]{}();:'\",./<>?j"
	for _, c := range all_invalids {
		st := string(c)
		assert.False(t, IsValidCname(&st), "'"+st+"' should not be a valid cname")
	}
}

// serialise

func TestMarshalJBOV(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Uniqid = "JBOV:0000000000000000000000000000000000000000"
	jbov.Volumes["vol1"].Uniqid = "VOL:1111111111111111111111111111111111111111"
	jbov.Volumes["vol2"].Uniqid = "VOL:2222222222222222222222222222222222222222"

	content, err := jbov.Marshall()

	assert.NoError(t, err)
	expected := `{
		"cname": "valid",
		"uniqid": "JBOV:0000000000000000000000000000000000000000",
		"last-mount-point": "",
		"volumes": {
			"vol1": {
				"uniqid": "VOL:1111111111111111111111111111111111111111",
				"last-mount-point": "/mnt/vol1"
			},
			"vol2": {
				"uniqid": "VOL:2222222222222222222222222222222222222222",
				"last-mount-point": "/mnt/vol2"
			}
		}
	}`
	assert.JSONEq(t, expected, string(content))
}

func TestUnmarshalJBOV(t *testing.T) {
	// TODO
}

func TestRoundTripMarshaller(t *testing.T) {
	// TODO
}

//  IsValid

func TestIsValid_happyPath(t *testing.T) {
	jbov := givenValidJBOV()

	ok, err := jbov.IsValid()

	assert.True(t, ok)
	assert.NoError(t, err, "this JBOV should be valid")
}

func TestIsValidJBOV_invalidCname(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Cname = "INVALID"

	ok, err := jbov.IsValid()

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV cname is not valid", "invalid cname should fail")
}

func TestIsValidJBOV_invalidJBovUniqId(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Uniqid = "invalid"

	ok, err := jbov.IsValid()

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV UniqId does not looks valid", "invalid jbov uniq should fail")
}

func TestIsValidJBOV_noVolumes(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes = nil

	ok, err := jbov.IsValid()

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV requires at least one volume", "jbov without volumes should fail")
}

func TestIsValidJBOV_VolumeWithInvalidName(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes["INVALID_VOL_CNAME"] = jbov.Volumes["vol1"]

	ok, err := jbov.IsValid()

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV volume has an invalid cname", "volume with invalid cname should fail")
}

func TestIsValidJBOV_VolumeWithInvalidUniqid(t *testing.T) {
	jbov := givenValidJBOV()
	jbov.Volumes["vol1"].Uniqid = GenerateJbovUniqId()

	ok, err := jbov.IsValid()

	assert.False(t, ok)
	assert.EqualErrorf(t, err, "JBOV volume has an invalid uniqid", "volume with invalid uniqid should fail")
}



// utils


func givenValidJBOV() JBOV {
	vols := make(map[string]*Volume)
	vols["vol1"] = &Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol1", Deprecated: false }
	vols["vol2"] = &Volume{Uniqid: GenerateVolumeUniqId(), LastMountPoint: "/mnt/vol2", Deprecated: false }
	return JBOV{Cname: "valid", Uniqid: GenerateJbovUniqId(), Volumes: vols}
}
