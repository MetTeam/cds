package sdk

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

// StaticFiles define a files needed to be save for serving static files
type StaticFiles struct {
	ID           int64     `json:"id" db:"id" cli:"id"`
	Name         string    `json:"name" db:"name" cli:"name"`
	NodeRunID    int64     `json:"workflow_node_run_id" db:"workflow_node_run_id"`
	NodeJobRunID int64     `json:"workflow_node_run_job_id,omitempty" db:"-"`
	EntryPoint   string    `json:"entrypoint" db:"entrypoint"`
	PublicURL    string    `json:"public_url" db:"public_url" cli:"public_url"`
	Created      time.Time `json:"created" db:"created" cli:"created"`

	TempURL   string `json:"temp_url,omitempty" db:"-"`
	SecretKey string `json:"secret_key,omitempty" db:"-"`
}

//GetName returns the name the artifact
func (c *StaticFiles) GetName() string {
	return base64.RawURLEncoding.EncodeToString([]byte(c.Name))
}

//GetPath returns the path of the artifact
func (c *StaticFiles) GetPath() string {
	container := base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%d-%s", c.NodeJobRunID, url.PathEscape(c.Name))))
	return container
}
