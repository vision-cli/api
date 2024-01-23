package v1

type Info struct {
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
}

type Version struct {
	SemVer string `json:"sem_ver"`
}

type Init struct {
	Config  any  `json:"config"`
	Success bool `json:"success"`
}

type Add struct {
	Config  any  `json:"config"`
	Success bool `json:"success"`
}

type Generate struct {
	Success bool `json:"success"`
}
