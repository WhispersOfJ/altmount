package registrar

import "testing"

func TestIsBearmountDownloadClient(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{BearmountDownloadClientName, true}, // exact registered name
		{"Bearmount", true},                 // common manual name
		{"bearmount", true},                 // lowercase
		{"BearMount (SABnzbd)", true},
		{"My BearMount SAB", true},
		{"", false},
		{"qBittorrent", false},
		{"SABnzbd", false},
		{"NZBGet", false},
	}
	for _, tt := range tests {
		if got := IsBearmountDownloadClient(tt.name); got != tt.want {
			t.Errorf("IsBearmountDownloadClient(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}
