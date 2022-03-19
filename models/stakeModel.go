package Models

import (
	"fmt"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *Stake) TableName() string {
	return "stake"
}

// get stake list
func GetList(take, skip int) (list []Stake) {
    list = []Stake{}

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "amount", "Status").
		From("stake").
		Where("status > 1")

	qb.OrderBy("amount").Desc().
        Limit(take).Offset(skip)

    sql := qb.String()


	num, err := o.Raw(sql).QueryRows(&list)
	if err == nil {
	    fmt.Println("error is: ", num)
	}

	return
}

func GetListNum(skip int, take int) (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("stake").
        Where("status > 1").
        Limit(take).Offset(skip)

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
}


func ReadAndCreateOrUpdate(address, amount string) (bool){
    o := orm.NewOrm()
    st := Stake{Address: address}
    if o.Read(&st) == nil {
        st.Amount = amount
        if _, err := o.Update(&st); err == nil {
            return true
        }
    } else {
        st.Amount = amount
        st.Status = 1
        o.Insert(st)
    }
    return true
}


func InsertStaker(address, amount string) {
    o := orm.NewOrm()

    st := new(Stake)
    st.Address = address
    st.Amount = amount
    st.Status = 1

    o.Insert(st)

}

func GetAddressRand(address, amount string) (num int) {
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("stake").
        Where("status > 1").
        And("amount > ?")

    sql := qb.String()


    err := o.Raw(sql, amount).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }

    return num
    
}

func ChangeStakeAmount(address, amount string) {
    o := orm.NewOrm()

    res, err := o.Raw("UPDATE stake SET amount = ? WHERE address = ?", amount, address).Exec()
    if err == nil {
        num, _ := res.RowsAffected()
        fmt.Println("mysql row affected nums: ", num)
    }

    return

}










