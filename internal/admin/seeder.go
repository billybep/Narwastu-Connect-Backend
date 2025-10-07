package admin

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&Admin{}).Count(&count)
	if count > 0 {
		return
	}

	seedAdmins := []Admin{
		{Name: "billybep", PhoneNumber: "6285146313141", Email: "sysadmin@narwastu.com", Role: RoleSystemAdmin},
		{Name: "Narwastu", PhoneNumber: "087077778987", Email: "admin@narwastu.com", Role: RoleAdmin},
		{Name: "Yemima", PhoneNumber: "62890789699", Email: "yemima@narwastu.com", Role: RoleAdmin},
		{Name: "Christy", PhoneNumber: "6285341695155", Email: "christy@narwastu.com", Role: RoleAdmin},
		{Name: "Member", PhoneNumber: "0833333333", Email: "member@narwastu.com", Role: RoleMember},
		{Name: "Guest", PhoneNumber: "0844444444", Email: "guest@narwastu.com", Role: RoleGuest},
	}

	for i := range seedAdmins {
		hash, _ := bcrypt.GenerateFromPassword([]byte("passWord#123"), bcrypt.DefaultCost)
		seedAdmins[i].Password = string(hash)
		if err := db.Create(&seedAdmins[i]).Error; err != nil {
			log.Println("failed to seed admin:", err)
		}
	}

	log.Println("✅ Admin seeding done")
}
