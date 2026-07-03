package order

import "testing"

func TestOrderSummary(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		priceRub int
		paid     bool
		want     string
	}{
		{
			name:     "paid status and paid flag",
			status:   StatusPaid,
			priceRub: 1200,
			paid:     true,
			want:     "status=paid payment=paid price_kop=120000",
		},
		{
			name:     "new status and not paid",
			status:   StatusNew,
			priceRub: 500,
			paid:     false,
			want:     "status=new payment=not_paid price_kop=50000",
		},
		{
			name:     "canceled status with paid flag",
			status:   StatusCanceled,
			priceRub: 300,
			paid:     true,
			want:     "status=canceled payment=paid price_kop=30000",
		},
		{
			name:     "unknown status",
			status:   99,
			priceRub: 10,
			paid:     false,
			want:     "status=unknown payment=not_paid price_kop=1000",
		},
		{
			name:     "paid status but not paid flag",
			status:   StatusPaid,
			priceRub: 750,
			paid:     false,
			want:     "status=paid payment=not_paid price_kop=75000",
		},
		{
			name:     "new status with paid flag",
			status:   StatusNew,
			priceRub: 100,
			paid:     true,
			want:     "status=new payment=paid price_kop=10000",
		},
		{
			name:     "zero price",
			status:   StatusNew,
			priceRub: 0,
			paid:     false,
			want:     "status=new payment=not_paid price_kop=0",
		},
		{
			name:     "negative price becomes zero",
			status:   StatusPaid,
			priceRub: -100,
			paid:     true,
			want:     "status=paid payment=paid price_kop=0",
		},
		{
			name:     "canceled and not paid",
			status:   StatusCanceled,
			priceRub: 999,
			paid:     false,
			want:     "status=canceled payment=not_paid price_kop=99900",
		},
		{
			name:     "unknown negative status",
			status:   -1,
			priceRub: 1,
			paid:     true,
			want:     "status=unknown payment=paid price_kop=100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderSummary(tt.status, tt.priceRub, tt.paid)
			if got != tt.want {
				t.Fatalf(
					"OrderSummary(%d, %d, %t)=%q, want %q",
					tt.status,
					tt.priceRub,
					tt.paid,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestStatusConstants(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{name: "StatusNew", got: StatusNew, want: 0},
		{name: "StatusPaid", got: StatusPaid, want: 1},
		{name: "StatusCanceled", got: StatusCanceled, want: 2},
		{name: "paid after new", got: StatusPaid - StatusNew, want: 1},
		{name: "canceled after paid", got: StatusCanceled - StatusPaid, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %d, want %d", tt.got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "status=paid payment=paid price_kop=120000"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
