// Package content loads the site's editable content from content.yaml,
// which is embedded into the binary at build time. To change site content,
// edit content/content.yaml and rebuild (commit + push redeploys).
package content

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed content.yaml
var raw []byte

// Link is a social icon or a quick-link button.
type Link struct {
	Label string `yaml:"label"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
}

// Profile holds the hero/about identity fields.
type Profile struct {
	Name     string `yaml:"name"`
	Handle   string `yaml:"handle"`
	Location string `yaml:"location"`
	Image    string `yaml:"image"`
	Bio      string `yaml:"bio"`
}

// YouTube holds the featured video embed and channel link.
type YouTube struct {
	EmbedURL   string `yaml:"embed_url"`
	ChannelURL string `yaml:"channel_url"`
}

// Content is the full parsed site content.
type Content struct {
	Domain      string   `yaml:"domain"`
	Profile     Profile  `yaml:"profile"`
	Identities  []string `yaml:"identities"`
	YouTube     YouTube  `yaml:"youtube"`
	SocialLinks []Link   `yaml:"social_links"`
	QuickLinks  []Link   `yaml:"quick_links"`
}

// Load parses the embedded content.yaml.
func Load() (*Content, error) {
	var c Content
	if err := yaml.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("content: parse content.yaml: %w", err)
	}
	return &c, nil
}

// Site is the content parsed once at startup. A malformed content.yaml is a
// build-time bug (the file is embedded), so we fail fast rather than serve a
// broken page. The CI test suite parses it too, catching errors before deploy.
var Site = mustLoad()

func mustLoad() *Content {
	c, err := Load()
	if err != nil {
		panic(err)
	}
	return c
}
