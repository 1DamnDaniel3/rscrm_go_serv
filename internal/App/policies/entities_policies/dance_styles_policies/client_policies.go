package dancestylespolicies

type DanceStylesPolicies struct {
	CRUD IDanceStylesCrudPolicy
}

func NewDanceStylesPolicies(crud IDanceStylesCrudPolicy) *DanceStylesPolicies {
	return &DanceStylesPolicies{
		CRUD: crud,
	}
}
