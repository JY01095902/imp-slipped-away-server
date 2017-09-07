package services

import (
	"fmt"
	"time"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	entities "imp-slipped-away-server/entities"
)

func SaveUser(user *entities.User) error {
	fmt.Println()
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "users"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("open_id", user.OpenId)

	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("nick_name", user.NickName)
	putRowChange.AddColumn("gender", user.Gender)
	putRowChange.AddColumn("phone_number", user.PhoneNumber)
	putRowChange.AddColumn("city", user.City)
	putRowChange.AddColumn("province", user.Province)
	putRowChange.AddColumn("country", user.Country)
	putRowChange.AddColumn("avatar_url", user.AvatarUrl)
	putRowChange.AddColumn("source", user.Source)
	putRowChange.AddColumn("created_at", time.Now().UTC().Format("2006-01-02 15:04:05"))
	putRowChange.AddColumn("deleted_at", "")
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	_, err := client.PutRow(putRowRequest)

	if err != nil {
		fmt.Println("SaveUser failed with error:", err)
	} else {
		fmt.Println("SaveUser finished")
	}

	return err
}