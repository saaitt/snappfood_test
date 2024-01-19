package usecase

import (
	"fmt"
	"github.com/saaitt/snappfood_test/agents/domain"
)

func (uc *useCase) GetAssigned() (*domain.AgentHistory, error) {
	var message []byte
	fmt.Println("hi")
	message = <-uc.consumer.MsgCh
	fmt.Println(string(message))
	//var t domain.Trip
	//_, err := uc.repo.FindByFilter(&t, "order_id", req.OrderID)
	//if err != nil && err != repository.ErrItemNotFound {
	//	return nil, err
	//}
	//if !t.Empty() {
	//	return nil, domain.ErrorOrderIsAlreadyAssigned
	//}
	//
	//res, err := uc.repo.Create(&domain.Trip{
	//	OrderID:   req.OrderID,
	//	Status:    domain.StatusAssigned,
	//	CreatedAt: time.Now().UTC(),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//trip := res.(*domain.Trip)
	//return trip, nil
	return nil, nil
}
