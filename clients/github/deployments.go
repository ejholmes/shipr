package github

type DeploymentsService struct {
	client *Client
}

type Deployment struct {
	ID          *int                    `json:"id,omitempty"`
	Sha         *string                 `json:"sha,omitempty"`
	Ref         *string                 `json:"ref,omitempty"`
	Environment *string                 `json:"environment,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Payload     *map[string]interface{} `json:"payload,omitempty"`
	Repository  *Repository             `json:"repository,omitempty"`
}

type DeploymentStatus struct {
	ID         *int        `json:"id,omitempty"`
	State      *string     `json:"state,omitempty"`
	Deployment *Deployment `json:"deployment,omitempty"`
}
