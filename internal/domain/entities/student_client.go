package entities

type StudentClient struct {
	ID        int64
	StudentID int64
	ClientID  int64
	IsPayer   bool
	Relation  string
	SchoolID  string
}
