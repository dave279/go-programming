package dog

func TestYears(t *testing.T) {
	got := Years(10)
	if got != 70 {
		t.Errorf("Years(10) = %d; want 70", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(20)
	if got != 140 {
		t.Errorf("YearsTwo(20) = %d; want 140", got)
	}
}

func BenchmarkYears(b *testing.B) {
	for b.Loop() {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for b.Loop() {
		YearsTwo(10)
	}
}
