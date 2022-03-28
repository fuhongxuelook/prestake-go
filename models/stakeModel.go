package Models

import (
	"fmt"
    "time"
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

	qb.Select("id", "address", "amount","status", "created_at").
		From("stake").
		Where("status > 0")

	qb.OrderBy("amount DESC, id ASC").
        Limit(take).Offset(skip)

    sql := qb.String()


	o.Raw(sql).QueryRows(&list)

	return list
}

func GetListNum() (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("stake").
        Where("status > 0")

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
}


func ReadAndCreateOrUpdate(address, amount string) (bool){
    o := orm.NewOrm()
    
    id := GetAddressId(address);
    st := Stake{Id: id, Address: address}

    if id > 0 && o.Read(&st) == nil {
        st.Amount = amount
        if _, err := o.Update(&st); err == nil {
            return true
        }
    } else {
        st.Amount = amount
        st.Status = 1
        timestr := time.Now().Format("2006-01-02 15:04:05")
        st.CreatedAt = timestr
        o.Insert(&st)
    }
    return true
}


func InsertStaker(address, amount string) {
    o := orm.NewOrm()

    st := new(Stake)
    st.Address = address
    st.Amount = amount
    st.Status = 1
    timestr := time.Now().Format("2006-01-0215:04:05")
    st.CreatedAt = timestr

    o.Insert(st)

}

func GetAddressRand(address, amount string) (num int) {
    o := orm.NewOrm()

    id := GetAddressId(address);
    st := Stake{Id: id, Address: address}

    if id == 0 || o.Read(&st) != nil {
        return 0;
    }


    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("stake").
        Where("status > 0").
        And("amount > ?").
        Or("amount = ? and id > ?")

    sql := qb.String()

    err := o.Raw(sql, amount, amount, id).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }

    return num + 1
}



func GetAddressId(address string) (num int) {
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id").
        From("stake").
        Where("status > 0").
        And("address = ?")

    sql := qb.String()


    err := o.Raw(sql, address).QueryRow(&num)
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










