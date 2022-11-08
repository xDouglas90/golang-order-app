package usecase

import (
	"database/sql"
	"testing"

	"github.com/xdouglas90/golang-order-app/internal/order/entity"
	"github.com/xdouglas90/golang-order-app/internal/order/infra/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax floar NOT NULL, final_price float NOT NULL)")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculatePrice())

	calculateFinalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
