package md

import (
	"encoding/json"
	"fmt"
	"regexp"
	"crypto/rand"
	"errors"
)

const JBOV_FNAME = ".jbov.metadata"
const UNIQID_FNAME = ".jbov.uniqid"
const LOCK_FNAME = ".jbov.lock"

var RE_JBOV_UNIQ = regexp.MustCompile("^JBOV:[0-9a-f]{16,64}$")
var RE_VOL_UNIQ = regexp.MustCompile("^VOL:[0-9a-f]{16,64}$")
var RE_VALID_CNAME = regexp.MustCompile("^[a-z0-9_]{3,20}$")

type JBOV struct {
	Cname          string `json:"cname"`
	Uniqid         string `json:"uniqid"`
	LastMountPoint string `json:"last-mount-point"`
	Volumes        map[string]*Volume `json:"volumes""`
	Rules          []Rule `json:"rules,omitempty"`
	Deleted        map[string]*Deleted `json:"deleted,omitempty"`
}

type Volume struct {
	Uniqid         string `json:"uniqid"`
	LastMountPoint string `json:"last-mount-point"`
	Deprecated     bool `json:"deprecated,omitempty""`
}

type Rule struct {
	Pattern        string `json:"pattern"`
	AtLeastACopyIn string `json:"at-least-a-copy-in,omitempty"`
	Ncopies        int `json:"ncopies,omitempty"`
}

type Deleted struct {
	Ts      int `json:"ts"`
	Pending []string `json:"pending"`
}

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

func IsValidCname(cname *string) bool {
	return RE_VALID_CNAME.MatchString(*cname)
}

func (jbov *JBOV) IsValid() (bool, error) {
	if !IsValidCname(&jbov.Cname) {
		return false, errors.New("JBOV cname is not valid")
	}
	if !IsJbovUniqId(&jbov.Uniqid) {
		return false, errors.New("JBOV UniqId does not looks valid")
	}
	if &jbov.Volumes == nil || len(jbov.Volumes) == 0 {
		return false, errors.New("JBOV requires at least one volume")
	}
	for cname, vol := range jbov.Volumes {
		if !IsValidCname(&cname) {
			return false, errors.New("JBOV volume has an invalid cname")
		}
		if !IsVolumeUniqId(&vol.Uniqid) {
			return false, errors.New("JBOV volume has an invalid uniqid")
		}
	}
	for _, deleted := range jbov.Deleted {
		for _, volp := range deleted.Pending {
			if _, ok := jbov.Volumes[volp]; !ok {
				return false, errors.New(fmt.Sprintf("JBOV deleted pending refers to invalid volume: %s", volp))
			}
		}
	}
	for i := 0; i < len(jbov.Rules); i++ {
		rule := jbov.Rules[i]
		if rule.AtLeastACopyIn != "" {
			if _, ok := jbov.Volumes[rule.AtLeastACopyIn]; !ok {
				return false, errors.New(fmt.Sprintf("JBOV rule at-least-a-copy-in refers to an invalid volume: %s", rule.AtLeastACopyIn))
			}
		}
	}
	return true, nil
}

func (jbov *JBOV) Marshall() ([]byte, error) {
	return json.MarshalIndent(jbov, "", "    ")
}
