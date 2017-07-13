package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

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

