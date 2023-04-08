package bcrypt

import "testing"

func TestBcrypt_Generate(t *testing.T) {
	secret := "verysecretkey"

	bc := New(secret, 10)

	if _, err := bc.Generate("asdf1234"); err != nil {
		t.Error(err)
	}
}

func TestBcrypt_Compare(t *testing.T) {
	secret := "verysecretkey"

	bc := New(secret, 10)

	var (
		hashed    string
		plainPass = "asdf1234"
		err       error
	)

	if hashed, err = bc.Generate(plainPass); err != nil {
		t.Error(err)
	}

	t.Run("Compare", func(t *testing.T) {
		if err = bc.Compare(hashed, plainPass); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("WrongPassword", func(t *testing.T) {
			if err = bc.Compare(hashed, "abc"); err == nil {
				t.Error("expected: hashedPassword is not the hash of the given password")
			}
		})
	})
}
