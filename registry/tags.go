package registry

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (registry *Registry) Tags(repository string) (tags []string, err error) {
	url := "/v2/%s/tags/list"

	var response tagsResponse
	for {
		registry.Logf("registry.tags url=%s repository=%s", url, repository)
		url = registry.url(url, registry)
		url, err = registry.getPaginatedJson(url, &response)
		switch err {
		case ErrNoMorePages:
			tags = append(tags, response.Tags...)
			return tags, nil
		case nil:
			tags = append(tags, response.Tags...)
			continue
		default:
			return nil, err
		}
	}
}
