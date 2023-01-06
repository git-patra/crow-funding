package user

import (
	"CrowFundingV2/src/helper"
	"CrowFundingV2/src/modules/user"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE EXECUTE ...")

	m.Run()

	// after
	fmt.Println("AFTER EXECUTE ...")
}

func TestFindByID(t *testing.T) {
	db, _ := helper.GetDBConnection()
	userRepo := user.NewRepository(db)
	result, _ := userRepo.FindById(3)

	if runtime.GOOS == "windows" {
		t.Skip("Cannot run in windows.")
	}

	// assert akan memanggil Fail() setelah pengecekan
	assert.Equal(t, 1, result.ID, "User not found!")

	//if result.ID == 0 {
	//	// Error akan print log dan memanggil Fail()
	//	t.Error("User not found!")
	//}
	//
	// kalo Fail() akan lanjut ekskusi code dibawah
	//fmt.Println("Cek Fail() Print")
}

func TestFindByEmail(t *testing.T) {
	db, _ := helper.GetDBConnection()
	userRepo := user.NewRepository(db)
	result, _ := userRepo.FindByEmail("patraaa@gmail.com")

	// require akan memanggil FailNow() setelah pengecekan
	require.Equal(t, 1, result.ID, "User not found!")

	//if result.ID == 0 {
	//	// Fatal akan print log dan memanggil FailNow()
	//	t.Fatal("User not found!")
	//}
	//
	// kalo FailNow() tidak aga di eksuksi perintah dibawah nya
	//fmt.Println("Cek FailNow() Print")
}
