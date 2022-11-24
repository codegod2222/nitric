// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secret

import (
	"context"
	"fmt"
)

type SecretService interface {
	// Put - Creates a new version for a given secret
	Put(context.Context, *Secret, []byte) (*SecretPutResponse, error)
	// Access - Retrieves the value for a given secret version
	Access(context.Context, *SecretVersion) (*SecretAccessResponse, error)
}

type UnimplementedSecretPlugin struct {
	SecretService
}

var _ SecretService = (*UnimplementedSecretPlugin)(nil)

func (*UnimplementedSecretPlugin) Put(ctx context.Context, secret *Secret, value []byte) (*SecretPutResponse, error) {
	return nil, fmt.Errorf("UNIMPLEMENTED")
}

func (*UnimplementedSecretPlugin) Access(ctx context.Context, version *SecretVersion) (*SecretAccessResponse, error) {
	return nil, fmt.Errorf("UNIMPLEMENTED")
}
