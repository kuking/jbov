package api

import (
	"testing"
	"log"
)

func assert(t bool) {
	if !t {
		log.Panic("assertion fail")
	}
}

func TestGeneratedJbovUniqIdValidates(t *testing.T) {
	uniqid := GenerateJbovUniqId()
	assert(IsJbovUniqId(&uniqid))
}

func TestGeneratedVolUniqIdValidates(t *testing.T) {
	uniqid := GenerateVolumeUniqId()
	assert(IsVolumeUniqId(&uniqid))
}

func TestUniqIdsCrossed(t *testing.T) {
	vol := GenerateVolumeUniqId()
	jbov := GenerateJbovUniqId()
	assert(vol != jbov)
	assert(IsVolumeUniqId(&vol))
	assert(!IsVolumeUniqId(&jbov))
	assert(IsJbovUniqId(&jbov))
	assert(!IsJbovUniqId(&vol))
}
