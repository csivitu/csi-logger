package schemas

type ResourceCreateSchema struct {
	Name      string `json:"name" validate:"required,max=50"`
	HostedURL string `json:"hosted_url" validate:"required"`
}

type ResourceUpdateSchema struct {
	HostedURL string `json:"hosted_url" validate:"required"`
}