package lessonpolicies

type LessonPolicies struct {
	CRUD         ILessonCrudPolicy
	CreatePolicy ILessonCreatePolicy
}

func NewLessonPolicies(
	crud ILessonCrudPolicy,
	CreatePolicy ILessonCreatePolicy,
) *LessonPolicies {
	return &LessonPolicies{
		CRUD:         crud,
		CreatePolicy: CreatePolicy,
	}
}
