package services

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct{}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
