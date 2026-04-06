package lessonpolicies

type LessonPolicies struct {
	CRUD ILessonCrudPolicy
}

func NewLessonPolicies(crud ILessonCrudPolicy) *LessonPolicies {
	return &LessonPolicies{
		CRUD: crud,
	}
}
