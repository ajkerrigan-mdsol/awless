/*
Copyright 2017 WALLIX

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

package commands

import (
	"path/filepath"
	"testing"

	"github.com/wallix/awless/graph"
)

func TestInstanceCredentialsFromName(t *testing.T) {
	g, err := graph.NewGraphFromFile(filepath.Join("testdata", "infra.rdf"))
	if err != nil {
		t.Fatal(err)
	}

	cred, err := instanceCredentialsFromGraph(g, "inst_1", "")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := cred.IP, "1.2.3.4"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := filepath.Base(cred.KeyPath), "my-key-name"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := cred.User, ""; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	cred, err = instanceCredentialsFromGraph(g, "inst_1", "/path/toward/myinst.key")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := cred.IP, "1.2.3.4"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := cred.KeyPath, "/path/toward/myinst.key"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	cred, err = instanceCredentialsFromGraph(g, "inst_2", "/path/toward/inst2.key")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := cred.IP, "2.3.4.5"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if got, want := cred.KeyPath, "/path/toward/inst2.key"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	_, err = instanceCredentialsFromGraph(g, "inst_12", "")
	if err == nil {
		t.Fatal("expected error got none")
	}
	if _, err := instanceCredentialsFromGraph(g, "inst_3", ""); err == nil {
		t.Fatal("expected error got none")
	}
	if _, err := instanceCredentialsFromGraph(g, "inst_2", ""); err == nil {
		t.Fatal("expected error got none")
	}
}
