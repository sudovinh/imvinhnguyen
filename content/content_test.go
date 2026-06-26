package content

import "testing"

// TestLoad ensures the embedded content.yaml parses and has the fields the
// page relies on. This runs in CI, so a malformed or incomplete edit to
// content.yaml fails the build before it can deploy a broken page.
func TestLoad(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatalf("content.yaml failed to parse: %v", err)
	}

	if c.Domain == "" {
		t.Error("domain is empty")
	}
	if c.Profile.Name == "" {
		t.Error("profile.name is empty")
	}
	if c.Profile.Handle == "" {
		t.Error("profile.handle is empty")
	}
	if c.Profile.Bio == "" {
		t.Error("profile.bio is empty")
	}

	if len(c.SocialLinks) == 0 {
		t.Error("no social_links defined")
	}
	for i, l := range c.SocialLinks {
		if l.Label == "" || l.URL == "" || l.Icon == "" {
			t.Errorf("social_links[%d] missing label, url, or icon: %+v", i, l)
		}
	}

	for i, l := range c.QuickLinks {
		if l.Label == "" || l.URL == "" {
			t.Errorf("quick_links[%d] missing label or url: %+v", i, l)
		}
	}

	if len(c.Identities) == 0 {
		t.Error("no identities defined for the animated line")
	}
	for i, id := range c.Identities {
		if id == "" {
			t.Errorf("identities[%d] is empty", i)
		}
	}
}
