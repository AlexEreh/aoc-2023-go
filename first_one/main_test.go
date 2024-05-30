package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		got := Reverse("amogus")
		want := "sugoma"
		require.Equal(t, want, got)
	})
	t.Run("Second test", func(t *testing.T) {
		got := Reverse("four")
		want := "ruof"
		require.Equal(t, want, got)
	})
	t.Run("Third test", func(t *testing.T) {
		got := Reverse("")
		want := ""
		require.Equal(t, want, got)
	})
	t.Run("Fourth test", func(t *testing.T) {
		got := Reverse("lol")
		want := "lol"
		require.Equal(t, want, got)
	})
}

func TestLineCalibrationNumber(t *testing.T) {
	t.Run("First test", func(t *testing.T) {
		got := LineCalibrationNumber("lol1lol")
		want := 11
		require.Equal(t, want, got)
	})
	t.Run("Second test", func(t *testing.T) {
		got := LineCalibrationNumber("ac1bg2ffer")
		want := 12
		require.Equal(t, want, got)
	})
}
