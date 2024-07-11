package repositories

import (
	pb "authentication_service/genproto/authentication_service"
	"authentication_service/pkg"
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
)

func Connect() *sql.DB {
	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal("Connectionda xatolik?")
	}

	return db
}

func TestLogin(t *testing.T) {
	db := Connect()
	testlogin := &pb.LoginRequest{
		Email:    "kimdirovkimdir@gmail.com",
		Password: "kimdirov",
	}

	resp, err := NewUserRepository(db).Login(testlogin)
	if err != nil {
		t.Fatalf("Loginni test qilishda xatolik: %v", err)
	}

	if !resp.Success {
		t.Fatal("Login muvaffaqiyatsiz bo'ldi")
	}

	testlogin.Password = "kgsfgdsgfsbvewogf"
	resp, err = NewUserRepository(db).Login(testlogin)
	if err == nil {
		t.Fatalf("Xato parol bilan login ishladi bu kutulgan xato: %v", err)
	}

	if resp.Success {
		t.Fatal("Xato parol nilan login ishladi?")
	}

	_, err = db.Exec("delete from users where email=$1", testlogin.Email)
	if err != nil {
		t.Fatal("Soxta userni o'chirishda xatolik?")
	}
}

func TestGetProfileById(t *testing.T) {
	db := Connect()
	user := pb.UserIdRequest{
		Id: "5a9c15ca-9b54-4b41-9e2e-a5dcc23d4254",
	}

	profile, err := NewUserRepository(db).GetProfileById(&user)
	if err != nil {
		t.Fatal("Profile ni olishda xatolik?")
	}

	if profile.Profile == nil {
		t.Fatal("Profile topilmadi?")
	}

	kutilganism := "kimdir"
	kutilganemaail := "kimdirovkimdir@gmail.com"

	if profile.Profile.Name != kutilganism {
		t.Fatalf("Ism mos emas. Kutilgan ism: %s,Olingan ism: %s", kutilganism, profile.Profile.Name)
	}

	if profile.Profile.Email != kutilganemaail {
		t.Fatalf("Email mos emas. Kutilgan email: %s,Olingan email: %s", kutilganemaail, profile.Profile.Email)
	}
}

func TestRegister(t *testing.T) {
	db := Connect()
	id := uuid.NewString()
	profile := &pb.Profile{
		Id:       id,
		Name:     "kimdir",
		Email:    "kimdirovkimdir@gmail.com",
		Password: "kimdirov",
		Role:     "role",
	}

	req := pb.RegisterRequest{
		Profile: profile,
	}

	resp, err := NewUserRepository(db).Register(&req)
	if err != nil {
		t.Fatalf("Ro'yxatdan o'tishda xatolik: %v", err)
	}

	if resp.Profile == nil {
		t.Fatalf("Ro'yxatdan o'tib bo'lmadi?")
	}

	if resp.Profile.Name != profile.Name {
		t.Errorf("Ism mos kelmadi. Kutilgan ism: %s, Olingan ism: %s", profile.Name, resp.Profile.Name)
	}

	if resp.Profile.Email != profile.Email {
		t.Errorf("Email mos kelmadi. Kutilgan email: %s, Olingan email: %s", profile.Email, resp.Profile.Email)
	}

	if resp.Profile.Password != profile.Password {
		t.Errorf("Password mos emas. Kutilgan password: %v,Olingan password: %v", profile.Password, resp.Profile.Password)
	}

	_, err = db.Exec("delete from users where id=$1", id)
	if err != nil {
		t.Fatalf("Sinov uchun yaratilgan userni o'chirishda xatolik: %v", err)
	}
}
