package implementation_test

import (
	"testing"
	impl "tkh/code/design-patterns/creational/factory-method/go/implementation"
)

func TestGetGun(t *testing.T) {
	ak47, err := impl.GetGun("ak47")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if ak47.GetName() != "AK47 gun" {
		t.Errorf("expected AK47 gun, got %s", ak47.GetName())
	}
	if ak47.GetPower() != 4 {
		t.Errorf("expected power 4, got %d", ak47.GetPower())
	}

	musket, err := impl.GetGun("musket")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if musket.GetName() != "Musket gun" {
		t.Errorf("expected Musket gun, got %s", musket.GetName())
	}
	if musket.GetPower() != 1 {
		t.Errorf("expected power 1, got %d", musket.GetPower())
	}
}
