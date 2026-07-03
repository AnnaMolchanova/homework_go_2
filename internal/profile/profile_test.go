package profile

import "testing"

func TestBuildUserCard(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		age      int
		active   bool
		want     string
	}{
		{
			testName: "adult active with spaces and uppercase",
			name:     "  MARIA  ",
			age:      20,
			active:   true,
			want:     "name=Maria age=20 group=adult status=active",
		},
		{
			testName: "minor inactive lowercase",
			name:     "ivan",
			age:      17,
			active:   false,
			want:     "name=Ivan age=17 group=minor status=inactive",
		},
		{
			testName: "adult boundary mixed case",
			name:     "  aLiCe",
			age:      18,
			active:   false,
			want:     "name=Alice age=18 group=adult status=inactive",
		},
		{
			testName: "minor active",
			name:     "peter",
			age:      10,
			active:   true,
			want:     "name=Peter age=10 group=minor status=active",
		},
		{
			testName: "empty name",
			name:     "",
			age:      30,
			active:   true,
			want:     "name=Unknown age=30 group=adult status=active",
		},
		{
			testName: "spaces only name",
			name:     "     ",
			age:      30,
			active:   false,
			want:     "name=Unknown age=30 group=adult status=inactive",
		},
		{
			testName: "negative age is minor",
			name:     "Max",
			age:      -1,
			active:   false,
			want:     "name=Max age=-1 group=minor status=inactive",
		},
		{
			testName: "already normalized",
			name:     "Maria",
			age:      25,
			active:   true,
			want:     "name=Maria age=25 group=adult status=active",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got := BuildUserCard(tt.name, tt.age, tt.active)
			if got != tt.want {
				t.Fatalf("BuildUserCard(%q, %d, %t)=%q, want %q", tt.name, tt.age, tt.active, got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "name=Maria age=20 group=adult status=active"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
