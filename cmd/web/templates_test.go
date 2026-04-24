package main

import (
	"testing"
	"time"

	"snippetbox.paqet.io/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	} {
		{
			name: "UTC",
			tm: time.Date(2026, 02, 1, 1, 10, 0, 0, time.UTC),
			want: "01 Feb 2026 at 01:10",
		},
		{
			name: "Empty",
			tm: time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm: time.Date(2026, 02, 1, 1, 10, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "01 Feb 2026 at 00:10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
