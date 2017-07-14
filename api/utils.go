package api

import (
	"fmt"
	"crypto/rand"
	"github.com/kuking/jbov/api/md"
	"errors"
	"regexp"
)

var RE_JBOV_UNIQ = regexp.MustCompile("^JBOV:[0-9a-f]{16,64}$")
var RE_VOL_UNIQ = regexp.MustCompile("^VOL:[0-9a-f]{16,64}$")
var RE_VALID_CNAME = regexp.MustCompile("^[a-z0-9_]{3,20}$")


func generateUniqid(prefix string) string {
	buf := make([]byte, 20)
	rand.Read(buf)
	return fmt.Sprintf("%s%x", prefix, buf)
}

func GenerateVolumeUniqId() string {
	return generateUniqid("VOL:")
}

func GenerateJbovUniqId() string {
	return generateUniqid("JBOV:")
}

func IsJbovUniqId(uniqid *string) bool {
	return RE_JBOV_UNIQ.MatchString(*uniqid)
}

func IsVolumeUniqId(uniqid *string) bool {
	return RE_VOL_UNIQ.MatchString(*uniqid)
}

func isValidCname(cname *string) bool {
	return RE_VALID_CNAME.MatchString(*cname)
}

func IsValidJBOV(jbov *md.JBOV) (bool, error) {

	if !isValidCname(&jbov.Cname) {
		return false, errors.New("JBOV cname is not valid")
	}
	if !IsJbovUniqId(&jbov.Uniqid) {
		return false, errors.New("JBOV UniqId does not looks valid")
	}
	if &jbov.Volumes == nil || len(jbov.Volumes) == 0 {
		return false, errors.New("JBOV requires at least one volume")
	}
	for cname, vol := range jbov.Volumes {
		if !isValidCname(&cname) {
			return false, errors.New("JBOV volume has an invalid cname")
		}
		if !IsVolumeUniqId(&vol.Uniqid) {
			return false, errors.New("JBOV volume has an invalid uniqid")
		}
	}
	return true, nil
}


