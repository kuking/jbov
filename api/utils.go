package api

import (
	"fmt"
	"crypto/rand"
	"github.com/kuking/jbov/api/md"
	"strings"
	"errors"
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


func CanCreate(jbov *md.JBOV) error {

	if strings.Trim(jbov.Cname, " \t\n\\l") == "" {
		return errors.New("JBOV Cname can't be empty")
	}
	if !IsJbovUniqId(&jbov.Uniqid) {
		return errors.New("JBOV UniqId does not looks valid")
	}

	if &jbov.Volumes == nil || len(jbov.Volumes) == 0 {
		return errors.New("JBOV requires at least one Volumen")
	}


	return nil
}

func Create(jbov *md.JBOV) error {

	err := CanCreate(jbov)
	if err != nil {
		return err
	}

	return nil
}
