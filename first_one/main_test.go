package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		got := Reverse("amogus")
		want := "sugoma"
		require.Equal(t, got, want)
	})
	t.Run("Second test", func(t *testing.T) {
		got := Reverse("four")
		want := "ruof"
		require.Equal(t, got, want)
	})
	t.Run("Third test", func(t *testing.T) {
		got := Reverse("")
		want := ""
		require.Equal(t, got, want)
	})
	t.Run("Fourth test", func(t *testing.T) {
		got := Reverse("lol")
		want := "lol"
		require.Equal(t, got, want)
	})
}
