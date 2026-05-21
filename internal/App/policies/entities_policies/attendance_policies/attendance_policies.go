package attendancepolicies

type AttendancePolicies struct {
	CRUD         IAttendanceCrudPolicy
	CreatePolicy IAttendanceCreatePolicy
	UpdatePolicy IAttendanceUpdatePolicy
}

func NewAttendancePolicies(
	crud IAttendanceCrudPolicy,
	CreatePolicy IAttendanceCreatePolicy,
	UpdatePolicy IAttendanceUpdatePolicy,
) *AttendancePolicies {
	return &AttendancePolicies{
		CRUD:         crud,
		CreatePolicy: CreatePolicy,
		UpdatePolicy: UpdatePolicy,
	}
}
