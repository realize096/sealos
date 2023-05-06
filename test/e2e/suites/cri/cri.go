/*
Copyright 2023 cuisongliu@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cri

import (
	"fmt"

	"github.com/labring/sealos/test/e2e/testhelper/cmd"
	"github.com/labring/sealos/test/e2e/testhelper/settings"
)

type Interface interface {
	Pull(name string) error
	ImageList() error
	ValidateImage(name string) error
}

type fakeCRIClient struct {
	SealosCmd *cmd.SealosCmd
}

func NewCRIClient() Interface {
	return &fakeCRIClient{
		SealosCmd: cmd.NewSealosCmd(settings.E2EConfig.SealosBinPath, &cmd.LocalCmd{}),
	}
}

func (f *fakeCRIClient) Pull(name string) error {
	if f.SealosCmd.CriBinPath == "" {
		if err := f.SealosCmd.SetCriBinPath(); err != nil {
			return err
		}
	}
	return f.SealosCmd.CRIImagePull(name)
}
func (f *fakeCRIClient) ImageList() error {
	if f.SealosCmd.CriBinPath == "" {
		if err := f.SealosCmd.SetCriBinPath(); err != nil {
			return err
		}
	}
	_, err := f.SealosCmd.CRIImageList(true)
	return err
}
func (f *fakeCRIClient) ValidateImage(name string) error {
	data, err := f.SealosCmd.CRIImageList(false)
	if err != nil {
		return err
	}
	for _, v := range data.Images {
		for _, tag := range v.RepoTags {
			if tag == name {
				return nil
			}
		}
	}
	return fmt.Errorf("image %s not found", name)
}