package auth

import (
	"testing"

	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/andresidrim/cesupa-hospital/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB abre um DB SQLite em memória e faz AutoMigrate de models.User
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.User{})
	assert.NoError(t, err)
	return db
}

// TestServiceRegister cobre registro bem-sucedido e duplicidade de CPF
func TestServiceRegister(t *testing.T) {
	db := setupTestDB(t)
	svc := NewService(db)

	tests := []struct {
		name      string
		input     models.User
		preInsert bool
		wantErr   bool
	}{
		{name: "successful register", input: models.User{Name: "Alice", CPF: "12345678900", Password: "secret", Role: "admin"}, wantErr: false},
		{name: "duplicate cpf", input: models.User{Name: "Bob", CPF: "12345678900", Password: "pwd123", Role: "doctor"}, preInsert: true, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// limpa tabela
			db.Exec("DELETE FROM users")

			if tt.preInsert {
				// insere user inicial para duplicidade
				assert.NoError(t, svc.Register(&models.User{Name: "Init", CPF: tt.input.CPF, Password: "x", Role: "admin"}))
			}

			// armazena senha antes de hash
			raw := tt.input.Password

			err := svc.Register(&tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			// verifica hash
			var saved models.User
			r := db.Where("cpf = ?", tt.input.CPF).First(&saved)
			assert.NoError(t, r.Error)
			assert.NotEqual(t, raw, saved.Password)
			// valida hash
			err = utils.CheckPassword(saved.Password, raw)
			assert.NoError(t, err)
		})
	}
}

// TestServiceLogin cobre erros e sucesso de login via JWT
func TestServiceLogin(t *testing.T) {
	db := setupTestDB(t)
	svc := NewService(db)

	// cria usuário para login
	plain := "mypassword"
	user := models.User{Name: "Carol", CPF: "55544433322", Password: plain, Role: "doctor"}
	assert.NoError(t, svc.Register(&user))

	tests := []struct {
		name        string
		cpf         string
		password    string
		wantErr     bool
		wantParseOK bool
	}{
		{name: "user_not_found", cpf: "nouser", password: "any", wantErr: true},
		{name: "incorrect_password", cpf: user.CPF, password: "wrongpass", wantErr: true},
		{name: "success", cpf: user.CPF, password: plain, wantErr: false, wantParseOK: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := svc.Login(tt.cpf, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
			if tt.wantParseOK {
				uid, perr := utils.ParseJWT(token)
				assert.NoError(t, perr)
				assert.Equal(t, user.ID, uid)
			}
		})
	}
}
