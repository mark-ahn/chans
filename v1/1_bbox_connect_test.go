package chans_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/mark-ahn/chans/v1"
)

func TestConnectBytes(t *testing.T) {
	ch_in, ch_out := make(chan []byte, 1), make(chan []byte, 1)
	chans.ConnectBytes(context.TODO(), ch_in, ch_out, nil)

	data := []byte{1, 2, 3, 4, 5}
	ch_in <- data
	d := <-ch_out
	if !reflect.DeepEqual(d, data) {
		t.Errorf("expect %v, got %v", data, d)
	}

	data = []byte{0}
	ch_in <- data
	d = <-ch_out
	if !reflect.DeepEqual(d, data) {
		t.Errorf("expect %v, got %v", data, d)
	}
}
