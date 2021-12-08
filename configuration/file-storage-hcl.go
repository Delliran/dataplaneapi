// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package configuration

import (
	"io/ioutil"
	"strings"

	"github.com/haproxytech/dataplaneapi/log"
	"github.com/hashicorp/hcl"
	"github.com/jinzhu/copier"
	"github.com/rodaine/hclencoder"
)

type StorageHCL struct {
	filename string
	cfg      *StorageDataplaneAPIConfiguration
}

func (s *StorageHCL) Load(filename string) error {
	s.filename = filename
	cfg := &StorageDataplaneAPIConfiguration{}
	var hclFile []byte
	var err error
	hclFile, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = hcl.Decode(&cfg, string(hclFile))
	if err != nil {
		return err
	}
	s.cfg = cfg
	return nil
}

func (s *StorageHCL) Get() *StorageDataplaneAPIConfiguration {
	if s.cfg == nil {
		s.cfg = &StorageDataplaneAPIConfiguration{}
	}
	return s.cfg
}

func (s *StorageHCL) Set(cfg *StorageDataplaneAPIConfiguration) {
	s.cfg = cfg
}

func (s *StorageHCL) SaveAs(filename string) error {
	var err error
	var hcl []byte
	var localCopy StorageDataplaneAPIConfiguration
	err = copier.Copy(&localCopy, &s.cfg)
	if err != nil {
		return err
	}
	// check if we have cluster log targets in config file
	if s.cfg.Cluster != nil && len(s.cfg.Cluster.ClusterLogTargets) > 0 {
		// since this can contain " character, escape it
		for index, value := range localCopy.Cluster.ClusterLogTargets {
			localCopy.Cluster.ClusterLogTargets[index].LogFormat = strings.Replace(value.LogFormat, `"`, `\"`, -1)
		}
	}
	// check if we have cluster log targets in config file
	if localCopy.Cluster != nil && len(localCopy.Cluster.ClusterLogTargets) > 0 {
		// since this can contain " character, escape it
		for index, value := range localCopy.Cluster.ClusterLogTargets {
			localCopy.Cluster.ClusterLogTargets[index].LogFormat = strings.Replace(value.LogFormat, `"`, `\"`, -1)
		}
	}
	if localCopy.LogTargets != nil && len(*localCopy.LogTargets) > 0 {
		var logTargets []log.Target
		for _, value := range *localCopy.LogTargets {
			value.LogFormat = strings.Replace(value.LogFormat, `"`, `\"`, -1)
			value.ACLFormat = strings.Replace(value.ACLFormat, `"`, `\"`, -1)
			logTargets = append(logTargets, value)
		}
		localCopy.LogTargets = (*log.Targets)(&logTargets)
	}
	if localCopy.Log != nil {
		if localCopy.Log.ACLFormat != nil {
			aclF := strings.Replace(*localCopy.Log.ACLFormat, `"`, `\"`, -1)
			localCopy.Log.ACLFormat = &aclF
		}
		if localCopy.Log.LogFormat != nil {
			logF := strings.Replace(*localCopy.Log.LogFormat, `"`, `\"`, -1)
			localCopy.Log.LogFormat = &logF
		}
	}

	hcl, err = hclencoder.Encode(localCopy)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, hcl, 0644) //nolint:gosec
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageHCL) Save() error {
	return s.SaveAs(s.filename)
}
