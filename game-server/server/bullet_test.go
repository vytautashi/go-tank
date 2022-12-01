package server

import (
	"reflect"
	"testing"
)

// Test: bulletSizeInBytes
func TestBulletSizeInBytes_return_4(t *testing.T) {
	expect := 4
	result := bulletSizeInBytes
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: bulletSizeInBytes
func TestBulletSizeInBytes_equals_to_bullet_x_plus_y_type_size(t *testing.T) {
	var bullet Bullet

	expect := bulletSizeInBytes
	result := int(reflect.TypeOf(bullet.x).Size() + reflect.TypeOf(bullet.y).Size())
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
