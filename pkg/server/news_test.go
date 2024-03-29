/*
Copyright AppsCode Inc.

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

package server

import (
	"context"
	"fmt"
	"testing"

	gdrive "gomodules.xyz/gdrive-utils"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/youtube/v3"
)

func TestServer_NextNewsSnippet(t *testing.T) {
	t.SkipNow()

	client, err := gdrive.DefaultClient("/Users/tamal/go/src/go.bytebuilders.dev/offline-license-server", youtube.YoutubeReadonlyScope)
	if err != nil {
		t.Fatal(err)
	}

	sheetsService, err := sheets.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		t.Fatal(err)
	}

	s := &Server{
		srvSheets: sheetsService,
	}
	got, err := s.NextNewsSnippet("Kubeform")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", got)
}
