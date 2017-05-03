// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAdminService_UpdateUserLDAPMapping(t *testing.T) {
	setup()
	defer teardown()

	input := &UserLDAPMapping{
		LDAPDN: String("uid=asdf,ou=users,dc=github,dc=com"),
	}

	mux.HandleFunc("/admin/ldap/users/u/mapping", func(w http.ResponseWriter, r *http.Request) {
		v := new(UserLDAPMapping)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"id":1,"ldap_dn":"uid=asdf,ou=users,dc=github,dc=com"}`)
	})

	mapping, _, err := client.Admin.UpdateUserLDAPMapping(context.Background(), "u", input)
	if err != nil {
		t.Errorf("Admin.UpdateUserLDAPMapping returned error: %v", err)
	}

	want := &UserLDAPMapping{
		ID:     Int(1),
		LDAPDN: String("uid=asdf,ou=users,dc=github,dc=com"),
	}
	if !reflect.DeepEqual(mapping, want) {
		t.Errorf("Admin.UpdateUserLDAPMapping returned %+v, want %+v", mapping, want)
	}
}

func TestAdminService_UpdateTeamLDAPMapping(t *testing.T) {
	setup()
	defer teardown()

	input := &TeamLDAPMapping{
		LDAPDN: String("cn=Enterprise Ops,ou=teams,dc=github,dc=com"),
	}

	mux.HandleFunc("/admin/ldap/teams/1/mapping", func(w http.ResponseWriter, r *http.Request) {
		v := new(TeamLDAPMapping)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"id":1,"ldap_dn":"cn=Enterprise Ops,ou=teams,dc=github,dc=com"}`)
	})

	mapping, _, err := client.Admin.UpdateTeamLDAPMapping(context.Background(), 1, input)
	if err != nil {
		t.Errorf("Admin.UpdateTeamLDAPMapping returned error: %v", err)
	}

	want := &TeamLDAPMapping{
		ID:     Int(1),
		LDAPDN: String("cn=Enterprise Ops,ou=teams,dc=github,dc=com"),
	}
	if !reflect.DeepEqual(mapping, want) {
		t.Errorf("Admin.UpdateTeamLDAPMapping returned %+v, want %+v", mapping, want)
	}
}

func TestAdminService_CreateUser(t *testing.T) {
	setup()
	defer teardown()

	input := &User{
		Login: String("octocat"),
		Email: String("octocat@github.com"),
	}

	mux.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{ "login": "octocat", "id": 1, "avatar_url": "https://github.com/images/error/octocat_happy.gif","gravatar_id": "", "url": "https://api.github.com/users/octocat", "html_url": "https://github.com/octocat","followers_url": "https://api.github.com/users/octocat/followers","following_url": "https://api.github.com/users/octocat/following{/other_user}","gists_url": "https://api.github.com/users/octocat/gists{/gist_id}","starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}","subscriptions_url": "https://api.github.com/users/octocat/subscriptions","organizations_url": "https://api.github.com/users/octocat/orgs","repos_url": "https://api.github.com/users/octocat/repos","events_url": "https://api.github.com/users/octocat/events{/privacy}","received_events_url": "https://api.github.com/users/octocat/received_events","type": "User","site_admin": false}`)
	})
	mapping, _, err := client.Admin.CreateUser(context.Background(), input)
	if err != nil {
		t.Errorf("Admin.CreateUser returned error: %v", err)
	}

	want := &User{
		Login:             String("octocat"),
		ID:                Int(1),
		AvatarURL:         String("https://github.com/images/error/octocat_happy.gif"),
		HTMLURL:           String("https://github.com/octocat"),
		GravatarID:        String(""),
		Type:              String("User"),
		SiteAdmin:         Bool(false),
		URL:               String("https://api.github.com/users/octocat"),
		EventsURL:         String("https://api.github.com/users/octocat/events{/privacy}"),
		FollowingURL:      String("https://api.github.com/users/octocat/following{/other_user}"),
		FollowersURL:      String("https://api.github.com/users/octocat/followers"),
		GistsURL:          String("https://api.github.com/users/octocat/gists{/gist_id}"),
		OrganizationsURL:  String("https://api.github.com/users/octocat/orgs"),
		ReceivedEventsURL: String("https://api.github.com/users/octocat/received_events"),
		ReposURL:          String("https://api.github.com/users/octocat/repos"),
		StarredURL:        String("https://api.github.com/users/octocat/starred{/owner}{/repo}"),
		SubscriptionsURL:  String("https://api.github.com/users/octocat/subscriptions"),
	}
	if !reflect.DeepEqual(mapping, want) {
		t.Errorf("Admin.CreateUser returned %+v, want %+v", mapping, want)
	}
}

func TestAdminService_RenameUser(t *testing.T) {
	setup()
	defer teardown()

	oldLogin := "octocat"
	newLogin := "monalisa"
	input := &User{Login: &newLogin}

	mux.HandleFunc("/admin/users/octocat", func(w http.ResponseWriter, r *http.Request) {
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"message": "Job queued to rename user. It may take a few minutes to complete.", "url": "https://api.github.com/user/1"}`)
	})
	msg, _, err := client.Admin.RenameUser(context.Background(), oldLogin, newLogin)
	if err != nil {
		t.Errorf("Admin.RenameUser returned error: %v", err)
	}

	want := &AdminMessage{
		Message: String("Job queued to rename user. It may take a few minutes to complete."),
		URL:     String("https://api.github.com/user/1"),
	}
	if !reflect.DeepEqual(msg, want) {
		t.Errorf("Admin.RenameUser returned %+v, want %+v", msg, want)
	}
}
