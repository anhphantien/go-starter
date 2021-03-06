package enums

type UserRole struct {
	ADMIN string
	USER  string
}

type UserStatus struct {
	NOT_ACTIVATED string
	ACTIVE        string
	IS_DISABLED   string
}

var User = struct {
	Role   UserRole
	Status UserStatus
}{
	Role: UserRole{
		ADMIN: "ADMIN",
		USER:  "USER",
	},
	Status: UserStatus{
		NOT_ACTIVATED: "NOT_ACTIVATED",
		ACTIVE:        "ACTIVE",
		IS_DISABLED:   "IS_DISABLED",
	},
}
