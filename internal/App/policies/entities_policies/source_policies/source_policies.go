package sourcepolicies

type SourcePolicies struct {
	CRUD ISourceCrudPolicy
}

func NewSourcePolicies(crud ISourceCrudPolicy) *SourcePolicies {
	return &SourcePolicies{
		CRUD: crud,
	}
}
