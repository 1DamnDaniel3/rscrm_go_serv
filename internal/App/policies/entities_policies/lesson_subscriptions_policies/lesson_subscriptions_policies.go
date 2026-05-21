package lessonsubscriptionspolicies

type LessonSubscriptionsPolicies struct {
	CRUD ILessonSubscriptionsCrudPolicy
}

func NewLessonSubscriptionsPolicies(crud ILessonSubscriptionsCrudPolicy) *LessonSubscriptionsPolicies {
	return &LessonSubscriptionsPolicies{
		CRUD: crud,
	}
}
