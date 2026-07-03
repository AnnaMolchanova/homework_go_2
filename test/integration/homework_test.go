package integration

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCommandOutput(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{path: "./cmd/01_integer", want: "add=22 div=3 mod=2 pages=11 even=true clamp=100"},
		{path: "./cmd/02_bases_bytes", want: "same=true byte=255 bin=11111111 hex=ff packed=258"},
		{path: "./cmd/03_float", want: "discount=850.00 tax=239.88 avg=12.5 rounded=12.35"},
		{path: "./cmd/04_boolean", want: "enter=true rest=true access=true leap=true withdraw=true"},
		{path: "./cmd/05_text", want: "bytes=12 runes=6 first=П changed=рот email=maria@example.com"},
		{path: "./cmd/06_constants", want: "app=go-homework-2 attempts=3 status=paid role=admin final=true"},
		{path: "./cmd/07_conversion", want: "double=84 kop=120000 sum=30 user=1001"},
		{path: "./cmd/08_pointers", want: "afterInc=11 value=22 nil=5"},
		{path: "./cmd/09_calculator", want: "add=15 div=3"},
		{path: "./cmd/10_profile", want: "name=Maria age=20 group=adult status=active"},
		{path: "./cmd/11_order", want: "status=paid payment=paid price_kop=120000"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			cmd := exec.Command("go", "run", tt.path)
			cmd.Dir = "../.."
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go run %s failed: %v\n%s", tt.path, err, out)
			}

			got := strings.TrimSpace(string(out))
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
