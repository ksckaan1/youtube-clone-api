package golang_jwt

import (
	"testing"
)

func TestJWT_Generate(t *testing.T) {
	j := New("verysecretkey", 24)

	if _, err := j.Generate(1); err != nil {
		t.Error(err)
	}
}

func TestJWT_Parse(t *testing.T) {
	j := New("verysecretkey", 24)

	var (
		generatedToken string
		err            error
	)

	t.Run("GenerateToken", func(t *testing.T) {
		if generatedToken, err = j.Generate(1); err != nil {
			t.Error(err)
		}
	})

	t.Run("ParseToken", func(t *testing.T) {
		if _, err = j.Parse(generatedToken); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("MalformedToken", func(t *testing.T) {
			var fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJDb21wYW55MSIsInN1YiI6IkF1dGhvcml6YXRpb24iLCJleHAiOjE2NzgzOTEyMDksImp0aSI6IjU1NSJ9.eo8F0RwmrlKZRqTRoEBE1YbOs8Dmy7F43eY1LHUblwos"

			if _, err = j.Parse(fakeToken); err == nil {
				t.Error("expected: token signature is invalid: signature is invalid")
			}
		})
	})
}
