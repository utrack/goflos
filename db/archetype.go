package db

import (
	"path"

	"strings"

	"github.com/pkg/errors"
	"github.com/utrack/goflos/file/flini"
)

type Archetype struct {
}

func (a *Archetype) Load(flPath string) error {
	flIni, err := flini.ParseFile(path.Join(flPath, "/exe/freelancer.ini"))
	if err != nil {
		return errors.Wrap(err, "couldn't load freelancer.ini")
	}

	al := &archLoader{}
	err = al.Load(flPath, flIni)
	if err != nil {
		return errors.Wrap(err, "couldn't load archDB")
	}
	return nil
}

type archLoader struct {
}

// Load loads archetype data from INIs.
// Args are: path to Freelancer and loaded Freelancer.ini.
func (l *archLoader) Load(flPath string, flIni *flini.File) error {
	dataRelPath := flIni.Sections["Freelancer"][0].Settings["data path"].V().String(0)
	dataRelPath = strings.Replace(dataRelPath, "\\", "/", -1)
	dataPath := path.Join(flPath, "exe", dataRelPath)

	for _, sect := range flIni.Sections["Data"] {
		for _, set := range sect.Settings["WeaponModDB"] {
			iniPath := strings.ToLower(strings.Replace(set.String(0), "\\", "/", -1))
			iniPath = path.Join(dataPath, iniPath)
			iniPath = path.Clean(iniPath)
			ini, err := flini.ParseFile(iniPath)
			if err != nil {
				return errors.Wrapf(err, "WeaponModDB INI at %v", iniPath)
			}
			// TODO make use of weapon mods
			_, err = l.loadWeaponModDB(ini)
			if err != nil {
				return errors.Wrapf(err, "WeaponModDB load at %v", iniPath)
			}
		}
	}
	return nil
}

// weaponMod is the damage modifier of a weapon against types of shields.
type weaponMod struct {
	nickname string
	mods     map[string]float32
}

func (l *archLoader) loadWeaponModDB(ini *flini.File) (map[string]weaponMod, error) {
	ret := map[string]weaponMod{}
	for sectPos, sect := range ini.Sections["WeaponType"] {
		wt := weaponMod{
			nickname: sect.Settings["nickname"].V().String(0),
			mods:     map[string]float32{},
		}

		for settPos, setting := range sect.Settings["shield_mod"] {
			nick := setting.String(0)
			factor, err := setting.Float(1)
			if err != nil {
				return nil, errors.Wrapf(err, "at section %v,setting %v (nickname %v)", sectPos, settPos, nick)
			}
			wt.mods[nick] = factor
		}
		ret[wt.nickname] = wt
	}
	return ret, nil
}
