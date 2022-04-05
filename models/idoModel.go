package Models

import (
	"fmt"
    "time"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *Ido) TableName() string {
	return "ido"
}

// get stake list
func GetIDOList(take, skip int) (list []Ido) {
    list = []Ido{}

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "amount","status", "created_at").
		From("ido").
		Where("status > 0")

	qb.OrderBy("id ASC").
        Limit(take).Offset(skip)

    sql := qb.String()


	o.Raw(sql).QueryRows(&list)

	return list
}

func GetIDOListNum() (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("ido").
        Where("status > 0")

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
}


func IDOReadAndCreateOrUpdate(address string, amount uint64) (bool){
    o := orm.NewOrm()
    
    id := GetIDOAddressId(address);
    st := Ido{Id: id, Address: address}

    if id > 0 && o.Read(&st) == nil {
        if amount == st.Amount {
            return false;
        }
        timestr := time.Now().Format("2006-01-02 15:04:05")
        st.CreatedAt = timestr
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


func InsertIDO(address string, amount uint64) {
    o := orm.NewOrm()

    st := new(Ido)
    st.Address = address
    st.Amount = amount
    st.Status = 1
    timestr := time.Now().Format("2006-01-0215:04:05")
    st.CreatedAt = timestr

    o.Insert(st)

}


func GetIDOAddressId(address string) (num int) {
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id").
        From("ido").
        Where("status > 0").
        And("address = ?")

    sql := qb.String()


    err := o.Raw(sql, address).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }

    return num 
}

func ChangeIDOAmount(address, amount string) {
    o := orm.NewOrm()

    res, err := o.Raw("UPDATE ido SET amount = ? WHERE address = ?", amount, address).Exec()
    if err == nil {
        num, _ := res.RowsAffected()
        fmt.Println("mysql row affected nums: ", num)
    }

    return

}










