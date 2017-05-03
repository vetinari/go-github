// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// AdminService handles communication with the admin related methods of the
// GitHub API. These API routes are normally only accessible for GitHub
// Enterprise installations.
//
// GitHub API docs: https://developer.github.com/v3/enterprise/
type AdminService service

// TeamLDAPMapping represents the mapping between a GitHub team and an LDAP group.
type TeamLDAPMapping struct {
	ID          *int    `json:"id,omitempty"`
	LDAPDN      *string `json:"ldap_dn,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	Privacy     *string `json:"privacy,omitempty"`
	Permission  *string `json:"permission,omitempty"`

	MembersURL      *string `json:"members_url,omitempty"`
	RepositoriesURL *string `json:"repositories_url,omitempty"`
}

func (m TeamLDAPMapping) String() string {
	return Stringify(m)
}

// UserLDAPMapping represents the mapping between a GitHub user and an LDAP user.
type UserLDAPMapping struct {
	ID         *int    `json:"id,omitempty"`
	LDAPDN     *string `json:"ldap_dn,omitempty"`
	Login      *string `json:"login,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitempty"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	Type       *string `json:"type,omitempty"`
	SiteAdmin  *bool   `json:"site_admin,omitempty"`

	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
}

func (m UserLDAPMapping) String() string {
	return Stringify(m)
}

// UpdateUserLDAPMapping updates the mapping between a GitHub user and an LDAP user.
//
// GitHub API docs: https://developer.github.com/v3/enterprise/ldap/#update-ldap-mapping-for-a-user
func (s *AdminService) UpdateUserLDAPMapping(ctx context.Context, user string, mapping *UserLDAPMapping) (*UserLDAPMapping, *Response, error) {
	u := fmt.Sprintf("admin/ldap/users/%v/mapping", user)
	req, err := s.client.NewRequest("PATCH", u, mapping)
	if err != nil {
		return nil, nil, err
	}

	m := new(UserLDAPMapping)
	resp, err := s.client.Do(ctx, req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// UpdateTeamLDAPMapping updates the mapping between a GitHub team and an LDAP group.
//
// GitHub API docs: https://developer.github.com/v3/enterprise/ldap/#update-ldap-mapping-for-a-team
func (s *AdminService) UpdateTeamLDAPMapping(ctx context.Context, team int, mapping *TeamLDAPMapping) (*TeamLDAPMapping, *Response, error) {
	u := fmt.Sprintf("admin/ldap/teams/%v/mapping", team)
	req, err := s.client.NewRequest("PATCH", u, mapping)
	if err != nil {
		return nil, nil, err
	}

	m := new(TeamLDAPMapping)
	resp, err := s.client.Do(ctx, req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// CreateUser creates a new user on Github Enterprise
//
// Login and Email are required attributes.
//
// https://developer.github.com/enterprise/v3/users/administration/#create-a-new-user
func (s *AdminService) CreateUser(ctx context.Context, user *User) (*User, *Response, error) {
	u := "admin/users"
	req, err := s.client.NewRequest("POST", u, user)
	if err != nil {
		return nil, nil, err
	}

	usr := new(User)
	resp, err := s.client.Do(ctx, req, usr)
	if err != nil {
		return nil, resp, err
	}
	return usr, resp, nil
}

// AdminMessage is returned from a RenameUser call
type AdminMessage struct {
	Message *string `json:"message"`
	URL     *string `json:"url"`
}

// RenameUser renames a user on Github Enterprise
//
// https://developer.github.com/enterprise/v3/users/administration/#rename-an-existing-user
func (s *AdminService) RenameUser(ctx context.Context, oldLogin, newLogin string) (*AdminMessage, *Response, error) {
	u := fmt.Sprintf("admin/users/%s", oldLogin)
	req, err := s.client.NewRequest("PATCH", u, &User{Login: &newLogin})
	if err != nil {
		return nil, nil, err
	}

	msg := new(AdminMessage)
	resp, err := s.client.Do(ctx, req, msg)
	if err != nil {
		return nil, resp, err
	}
	return msg, resp, nil
}
