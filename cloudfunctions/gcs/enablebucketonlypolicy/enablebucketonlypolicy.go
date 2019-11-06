package enablebucketonlypolicy

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"context"
	"encoding/json"

	pb "github.com/googlecloudplatform/threat-automation/compiled/sha/protos"
	"github.com/googlecloudplatform/threat-automation/entities"
	"github.com/googlecloudplatform/threat-automation/providers/sha"
	"github.com/pkg/errors"
)

// Required contains the required values needed for this function.
type Required struct {
	BucketName, ProjectID string
}

// ReadFinding will attempt to deserialize all supported findings for this function.
func ReadFinding(b []byte) (*Required, error) {
	var finding pb.StorageScanner
	r := &Required{}
	if err := json.Unmarshal(b, &finding); err != nil {
		return nil, errors.Wrap(entities.ErrUnmarshal, err.Error())
	}
	switch finding.GetFinding().GetCategory() {
	case "BUCKET_POLICY_ONLY_DISABLED":
		r.BucketName = sha.BucketName(finding.GetFinding().GetResourceName())
		r.ProjectID = finding.GetFinding().GetSourceProperties().GetProjectId()
<<<<<<< HEAD
	default:
		return nil, entities.ErrUnsupportedFinding
=======
>>>>>>> 8a7432ce21d4d9e9221a655d9a2905020835022e
	}
	if r.BucketName == "" || r.ProjectID == "" {
		return nil, entities.ErrValueNotFound
	}
	return r, nil
}

// Execute will enable bucket only policy on buckets found within the provided folders.
func Execute(ctx context.Context, required *Required, ent *entities.Entity) error {
	resources := ent.Configuration.EnableBucketOnlyPolicy.Resources
	return ent.Resource.IfProjectWithinResources(ctx, resources, required.ProjectID, func() error {
		if err := ent.Resource.EnableBucketOnlyPolicy(ctx, required.BucketName); err != nil {
			return err
		}
		ent.Logger.Info("Bucket only policy enabled on bucket %q in project %q.", required.BucketName, required.ProjectID)
		return nil
	})
}
