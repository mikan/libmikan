// Copyright 2016 mikan.

package crypt

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	data, err := Encrypt("test", "password")
	if err != nil {
		t.Errorf("Error occurred in encrypt.")
	}
	plain, err := Decrypt(data, "password")
	if err != nil {
		t.Errorf("Error occurred in decrypt.")
	}
	if plain != "test" {
		t.Errorf("Not match:")
	}
}
