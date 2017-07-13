package api

import (
	"fmt"
	"crypto/rand"
	"github.com/kuking/jbov/api/md"
	"errors"
	"regexp"
)

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
	if len(*uniqid)<15 || (*uniqid)[0:5] != "JBOV:" {
		return false
	}
	return true
}

func IsVolumeUniqId(uniqid *string) bool {
	if len(*uniqid)<15 || (*uniqid)[0:4] != "VOL:" {
		return false
	}
	return true
}

func isValidCname(cname *string) bool {
	re := regexp.MustCompile("^[a-z0-9_]{3,20}$") //TODO: compile just once?
	return re.MatchString(*cname)
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
			return false, errors.New("Volume uniqid does not looks valid")
		}

	}

	return true, nil
}

func Create(jbov *md.JBOV) error {

	ok, err := IsValidJBOV(jbov)
	if !ok && err != nil {
		return err
	}

	return nil
}
