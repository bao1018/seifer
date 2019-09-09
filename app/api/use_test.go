package api

import (
	"testing"

	"bitbucket.org/kyicy/seifer/config"
)

func TestSenteceToVector(t *testing.T) {
	config.Load()
	sss := SentenceToVector("sss")

	if len(sss) != 512 {
		t.Errorf("vector size is %v", len(sss))
	}
}
