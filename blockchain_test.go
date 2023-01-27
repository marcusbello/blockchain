package learn_bitcoin

import (
	"crypto/sha256"
	"reflect"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	t.Run("create new Blockchain", func(t *testing.T) {
		tx := []string{"hello", "world"}
		prevHash := []byte("")
		b := NewBlock(tx, prevHash)
		if b == nil {
			t.Errorf("NewBlock() = nil")
		}
	})
	t.Run("validate first transaction", func(t *testing.T) {
		tx := []string{"hello", "world"}
		prevHash := []byte("")
		b := NewBlock(tx, prevHash)
		wantedTx := []byte("hello")
		got := []byte(b.transactions[0])
		if !reflect.DeepEqual(wantedTx, got) {
			t.Errorf("First Tx = %v, want %v", got, wantedTx)
		}
	})
}

func TestNewHash(t *testing.T) {
	t.Run("create new hash", func(t *testing.T) {
		tx := []string{"hello", "world"}
		prevHash := []byte("")
		currentTime := time.Now()
		h := NewHash(currentTime, tx, prevHash)
		if h == nil {
			t.Errorf("NewHash() = nil")
		}
	})
	t.Run("confirm Tx hashing", func(t *testing.T) {
		txs := []string{"hello", "world"}
		var prevHash = []byte("")
		currentTime := time.Now()
		prevHashTime := append(prevHash, currentTime.String()...)
		for _, tx := range txs {
			prevHashTime = append(prevHashTime, string(tx)...)
		}
		want := sha256.Sum256(prevHashTime)
		got := NewHash(currentTime, txs, prevHash)
		if !reflect.DeepEqual(got, want[:]) {
			t.Errorf("Hash() = %v, want %v", got, want)
		}
	})
}
