package usecase

type usecase struct{}

func Init() UsecaseItf {
	return &usecase{}
}
