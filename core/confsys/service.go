/*
 * === This file is part of ALICE O² ===
 *
 * Copyright 2019 CERN and copyright holders of ALICE O².
 * Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

package confsys

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/AliceO2Group/Control/common/logger"
	"github.com/AliceO2Group/Control/configuration"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = logger.New(logrus.StandardLogger(), "confsys")

var (
	once sync.Once
	instance *Service
)

func Instance() *Service {
	once.Do(func() {
		var err error
		configUri := viper.GetString("globalConfigurationUri")
		instance, err = newService(configUri)
		if err != nil {
			log.WithField("globalConfigurationUri", configUri).Fatal("bad configuration URI")
		}
	})
	return instance
}



type Service struct {
	src configuration.Source
}

/* Expected structure:
/o2/aliecs/
{
	run_number: 47102,

	settings: {
		log_level: "DEBUG"
	},
	global_vars: {
		n_flps: 42
	},
	repositories: {
		base: {
			url: "https://gitlab.cern.ch/foo/bar",
			user: "gituser",
			pass: "gitpass"
		},
		extra: {...}
	}
}

*/
func formatKey(key string) (consulKey string) {
	// Trim leading slashes
	consulKey = strings.TrimLeft(key, "/")
	return
}

func newService(uri string) (svc *Service, err error) {
	var src configuration.Source
	src, err = configuration.NewSource(uri)
	return &Service{src: src}, err
}

func (s *Service) NewDefaultRepo(defaultRepo string) error {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Put("o2/control/default_repo", defaultRepo)
	} else {
		data := []byte(defaultRepo)
		return ioutil.WriteFile(filepath.Join(s.GetReposPath(),"default_repo"), data, 0644)
	}
}

func (s *Service) GetDefaultRepo() (defaultRepo string, err error) {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Get("o2/control/default_repo")
	} else {
		var defaultRepoData []byte
		defaultRepoData, err = ioutil.ReadFile(filepath.Join(s.GetReposPath(),"default_repo"))
		if err != nil {
			return
		}
		defaultRepo = strings.TrimSuffix(string(defaultRepoData), "\n")
		return
	}
}

func (s *Service) NewDefaultRevision(defaultRevision string) error {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Put("o2/control/default_revision", defaultRevision)
	} else {
		data := []byte(defaultRevision)
		return ioutil.WriteFile(filepath.Join(s.GetReposPath(),"default_revision"), data, 0644)
	}
}

func (s *Service) GetDefaultRevision() (defaultRevision string, err error) {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Get("o2/control/default_revision")
	} else {
		var defaultRevisionData []byte
		defaultRevisionData, err = ioutil.ReadFile(filepath.Join(s.GetReposPath(),"default_revision"))
		if err != nil {
			return
		}
		defaultRevision = strings.TrimSuffix(string(defaultRevisionData), "\n")
		return
	}
}

func (s *Service) GetRepoDefaultRevisions() (map[string]string, error) {
	var defaultRevisions map[string]string
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		data, err := cSrc.Get("o2/control/default_revisions")
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(data), &defaultRevisions)
		if err != nil {
			return nil, err
		}
	} else {
		defaultRevisionData, err := ioutil.ReadFile(filepath.Join(s.GetReposPath(),"default_revisions.json"))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(defaultRevisionData, &defaultRevisions)
	}
	return defaultRevisions, nil
}

func (s *Service) SetRepoDefaultRevisions(defaultRevisions map[string]string) error {
	data, err := json.MarshalIndent(defaultRevisions, "", "    ")
	if err != nil {
		return err
	}

	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		err = cSrc.Put("o2/control/default_revisions", string(data))
	} else {
		err = ioutil.WriteFile(filepath.Join(s.GetReposPath(),"default_revisions.json"), data, 0644)
	}
	return err
}

func (s *Service) NewRunNumber() (runNumber uint32, err error) {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.GetNextUInt32("o2/control/run_number")
	} else {
		// Unsafe check-and-set, only for file backend
		var rnf string
		rnf = filepath.Join(viper.GetString("coreWorkingDir"), "runcounter.txt")
		if _, err = os.Stat(rnf); os.IsNotExist(err) {
			err = ioutil.WriteFile(rnf, []byte("0"), 0644)
			if err != nil {
				return
			}
		}
		var raw []byte
		raw, err = ioutil.ReadFile(rnf)
		if err != nil {
			return
		}
		var rn64 uint64
		rn64, err = strconv.ParseUint(string(raw[:]), 10, 32)
		if err != nil {
			return
		}
		runNumber = uint32(rn64)
		runNumber++
		raw = []byte(strconv.FormatUint(uint64(runNumber), 10))
		err = ioutil.WriteFile(rnf, raw, 0)
		return
	}
}

func (s *Service) GetROSource() configuration.ROSource {
	return s.src
}

// maybe this one shouldn't exist at all, because vars should get inserted
// response: but not all of them! some vars will likely only get parsed at deployment time i.e. right
//    before pushing TaskInfos
func (s *Service) GetVars() map[string]string {
	//FIXME: implement
	return nil
}

// Or maybe even "RefreshConfig" which will refresh all the things that happen to be runtime-refreshable
func (s *Service) RefreshRepositories() {
	panic("not implemented yet")
}

// Returns a YAML file OR even a structure made of Roles or Nodes with:
// import() functions already computed and resolved
// vars inserted (todo: figure out which vars get parsed when and where)
func (s *Service) GenerateWorkflowDescriptor(wfPath string, vars map[string]string /*vars from cli/gui*/) string {
	panic("not implemented yet")
}

// Persist Mesos Framework ID by saving to Consul, or to a local file.
func (s *Service) NewMesosFID(fidValue string) error {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Put("o2/aliecs/mesos_fid", fidValue)
	} else {
		data := []byte(fidValue)
		return ioutil.WriteFile(filepath.Join(viper.GetString("coreWorkingDir"), "mesos_fid.txt"), data, 0644)
	}
}

// Retrieve Mesos Framework ID from Consul, or local file.
func (s *Service) GetMesosFID() (fidValue string, err error) {
	if cSrc, ok := s.src.(*configuration.ConsulSource); ok {
		return cSrc.Get("o2/aliecs/mesos_fid")
	} else {
		var byteFidValue []byte
		byteFidValue, err = ioutil.ReadFile(filepath.Join(viper.GetString("coreWorkingDir"), "mesos_fid.txt"))
		if err != nil {
			return
		}
		fidValue = strings.TrimSuffix(string(byteFidValue), "/n")
		return
	}
}

func (s *Service) GetReposPath() string {
	return filepath.Join(viper.GetString("coreWorkingDir"), "repos")
}