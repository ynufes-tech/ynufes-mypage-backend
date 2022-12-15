package entity

type UserDashboard struct {
	Grants []string `firestore:"dashboard-grants"`
}
