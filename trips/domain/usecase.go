package domain

type UseCase interface {
	AssignTrip(req *TripRequest) (*Trip, error)
	GetTrip(req *TripRequest) (*Trip, error)
	UpdateTrip(req *TripUpdateRequest) (*Trip, error)
}
