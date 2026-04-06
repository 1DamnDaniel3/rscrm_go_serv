package schedulepolicies

type SchedulePolicies struct {
	CRUD IScheduleCrudPolicy
}

func NewSchedulePolicies(crud IScheduleCrudPolicy) *SchedulePolicies {
	return &SchedulePolicies{
		CRUD: crud,
	}
}
