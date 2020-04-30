package pattern

import (
	"fmt"
	"testing"
)

//Person 定义一个本身的人
type Person struct {
	Name         string
	WalletAssets int //每个人都有钱包
}

//Tenant 租客 继承Person
type Tenant struct {
	Person
	furniture string
}

//Landlord 房东，也继承Person，要收房租
type Landlord struct {
	Person
	RentAccout int //房东的租金账户
}

type Mediator struct {
	Person
	tenant   interface{}
	landlord interface{}
}

func TestClassCompose(t *testing.T) {

	med := &Mediator{Person: Person{"mediator", 1001}}

	landlord := &Landlord{Person: Person{"landlord", 2000}, RentAccout: 500}

	tenant := &Tenant{Person{"tenant", 500}, "desk"}

	fmt.Println("mediator:", med)
	fmt.Println("landlord Name:", landlord.Name)
	fmt.Println("tenant furniture:", tenant.furniture)
	fmt.Println("----------------------------")
	fmt.Println("mediator:", med)
	fmt.Println("landlord:", landlord)
	fmt.Println("tenant:", tenant)
}
