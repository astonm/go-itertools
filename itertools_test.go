package itertools

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toSlice[T any](s iter.Seq[T]) []T {
	maxLen := 20
	next, _ := iter.Pull(s)
	out := make([]T, 0)
	for {
		if len(out) > maxLen {
			panic("iterator ran too long")
		}

		v, ok := next()
		if !ok {
			break
		}
		out = append(out, v)
	}
	return out
}

func assertSequenceMatch[V any](t *testing.T, gotSeq iter.Seq[V], want []V) {
	got := toSlice(gotSeq)
	assert.Equal(t, want, got)
}

func TestNewSeq(t *testing.T) {
	assertSequenceMatch(t, NewSeq(1, 2, 3), []int{1, 2, 3})
}

func TestEnumerate(t *testing.T) {
	keys := make([]int, 0)
	vals := make([]int, 0)
	for k, v := range Enumerate(NewSeq(0, 1, 2, 3)) {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	assert.Equal(t, keys, vals)
}

func TestTake(t *testing.T) {
	assertSequenceMatch(t, Take(NewSeq(1, 2, 3), 0), []int{})
	assertSequenceMatch(t, Take(NewSeq(1, 2, 3, 4, 5, 6), 3), []int{1, 2, 3})
}

func TestChain(t *testing.T) {
	assertSequenceMatch(t,
		Chain(NewSeq(1, 2, 3), NewSeq(4, 5, 6)),
		[]int{1, 2, 3, 4, 5, 6},
	)
}

func TestCount(t *testing.T) {
	assertSequenceMatch(t, Take(Count(), 3), []int{0, 1, 2})
}

func TestCycle(t *testing.T) {
	assertSequenceMatch(t, Take(Cycle(NewSeq(1, 2, 3)), 5), []int{1, 2, 3, 1, 2})
}

func TestRepeat(t *testing.T) {
	assertSequenceMatch(t, Repeat("a", 5), []string{"a", "a", "a", "a", "a"})
}

func TestAccumulate(t *testing.T) {
	runningSums := Accumulate(NewSeq(1, 2, 3), func(x, y int) int { return x + y }, 0)
	assertSequenceMatch(t, runningSums, []int{1, 3, 6})

	runningProducts := Accumulate(NewSeq(1, 2, 3), func(x, y int) int { return x * y }, 1)
	assertSequenceMatch(t, runningProducts, []int{1, 2, 6})
}

func TestBatched(t *testing.T) {
	assertSequenceMatch(t,
		Batched(NewSeq(1, 2, 3, 4, 5, 6, 7, 8), 3),
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}},
	)
}

func TestCombinations(t *testing.T) {
	assertSequenceMatch(t,
		Combinations([]string{"A", "B", "C", "D"}, 2),
		[][]string{{"A", "B"}, {"A", "C"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"C", "D"}},
	)
	assertSequenceMatch(t,
		Combinations([]int{0, 1, 2, 3}, 3),
		[][]int{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3}},
	)
}

func TestCombinationsWithReplacement(t *testing.T) {
	assertSequenceMatch(t,
		CombinationsWithReplacement([]string{"A", "B", "C"}, 2),
		[][]string{{"A", "A"}, {"A", "B"}, {"A", "C"}, {"B", "B"}, {"B", "C"}, {"C", "C"}},
	)
}

func TestCompress(t *testing.T) {
	assertSequenceMatch(t,
		Compress(NewSeq([]byte("ABCDEFG")...), []int{1, 0, 1, 0, 1, 1}),
		[]byte("ACEF"),
	)

	assertSequenceMatch(t,
		Compress(Count(), []bool{false, true, false, true, false, true}),
		[]int{1, 3, 5},
	)
}

func TestDropWhile(t *testing.T) {
	assertSequenceMatch(t,
		DropWhile(func(x int) bool { return x < 5 }, NewSeq(1, 4, 6, 3, 8)),
		[]int{6, 3, 8},
	)
}

func TestFilterFalse(t *testing.T) {
	assertSequenceMatch(t,
		FilterFalse(func(x int) bool { return x < 5 }, NewSeq(1, 4, 6, 3, 8)),
		[]int{6, 8},
	)
}

func TestGroupBy(t *testing.T) {
	want := map[string][]string{
		"A": []string{"A", "A", "A", "A"},
		"B": []string{"B", "B", "B"},
		"C": []string{"C", "C"},
		"D": []string{"D"},
	}

	for k, g := range GroupBy(NewSeq("A", "A", "A", "A", "B", "B", "B", "C", "C", "D")) {
		assertSequenceMatch(t, g, want[k])
	}
}

func TestSlice(t *testing.T) {
	assertSequenceMatch(t,
		Slice(NewSeq([]byte("ABCDEFG")...), 2, 4),
		[]byte("CD"),
	)

	assertSequenceMatch(t,
		Slice(NewSeq([]byte("ABCDEFG")...), 2, -1),
		[]byte("CDEFG"),
	)
}

func TestPairwise(t *testing.T) {
	want := []string{"AB", "BC", "CD", "DE", "EF", "FG"}

	var i int
	for x, y := range Pairwise(NewSeq([]byte("ABCDEFG")...)) {
		assert.Equal(t, x, want[i][0])
		assert.Equal(t, y, want[i][1])
		i++
	}
}

func TestPermutations(t *testing.T) {
	assertSequenceMatch(t,
		Permutations([]string{"A", "B", "C", "D"}, 2),
		[][]string{
			{"A", "B"},
			{"A", "C"},
			{"A", "D"},
			{"B", "A"},
			{"B", "C"},
			{"B", "D"},
			{"C", "A"},
			{"C", "B"},
			{"C", "D"},
			{"D", "A"},
			{"D", "B"},
			{"D", "C"},
		},
	)

	assertSequenceMatch(t,
		Permutations([]int{0, 1, 2}, 3),
		[][]int{
			{0, 1, 2},
			{0, 2, 1},
			{1, 0, 2},
			{1, 2, 0},
			{2, 0, 1},
			{2, 1, 0},
		},
	)
}

func TestProduct(t *testing.T) {
	assertSequenceMatch(t,
		Product([]byte("ABCD"), []byte("xy")),
		[][]byte{
			{'A', 'x'},
			{'A', 'y'},
			{'B', 'x'},
			{'B', 'y'},
			{'C', 'x'},
			{'C', 'y'},
			{'D', 'x'},
			{'D', 'y'},
		},
	)
}

func TestProductRepeat(t *testing.T) {
	assertSequenceMatch(t,
		ProductRepeat([]int{0, 1}, 3),
		[][]int{
			{0, 0, 0},
			{0, 0, 1},
			{0, 1, 0},
			{0, 1, 1},
			{1, 0, 0},
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 1},
		},
	)
}

func TestMap(t *testing.T) {
	assertSequenceMatch(t,
		Map(func(x int) byte { return byte('0' + x) }, NewSeq(0, 1, 2)),
		[]byte{'0', '1', '2'},
	)
}

func TestTakeWhile(t *testing.T) {
	assertSequenceMatch(t,
		TakeWhile(func(x int) bool { return x < 5 }, NewSeq(1, 4, 6, 3, 8)),
		[]int{1, 4},
	)
}

func TestTee(t *testing.T) {
	vals := []int{2, 4, 6, 8}
	seqs := Tee(OfSlice(vals), 3)

	s0, stop0 := iter.Pull(seqs[0])
	s1, stop1 := iter.Pull(seqs[1])
	s2, stop2 := iter.Pull(seqs[2])

	defer stop0()
	defer stop1()
	defer stop2()

	for _, v := range vals {
		var x int
		var ok bool

		x, ok = s0()
		assert.True(t, ok)
		assert.Equal(t, x, v)

		x, ok = s1()
		assert.True(t, ok)
		assert.Equal(t, x, v)

		x, ok = s2()
		assert.True(t, ok)
		assert.Equal(t, x, v)
	}
}

func TestZip(t *testing.T) {
	chrs := OfSlice([]byte("2468"))
	nums := OfSlice([]int{2, 4, 6, 8})

	for c, n := range Zip(chrs, nums) {
		assert.Equal(t, c, byte('0'+n))
	}
}

func TestPullZip3(t *testing.T) {
	chrs := OfSlice([]byte("2468"))
	nums := OfSlice([]int{2, 4, 6, 8})
	flts := OfSlice([]float32{2.0, 4.0, 6.0, 8.0})

	next, stop := PullZip3(chrs, nums, flts)
	defer stop()

	for {
		c, n, f, ok := next()
		if !ok {
			break
		}

		assert.Equal(t, c, byte('0'+n))
		assert.Equal(t, f, float32(n))
	}
}

func TestPullZip4(t *testing.T) {
	mat := []iter.Seq[int]{
		NewSeq(0, 4, 8, 12),
		NewSeq(1, 5, 9, 13),
		NewSeq(2, 6, 10, 14),
		NewSeq(3, 7, 11, 15),
	}

	next, stop := PullZip4(mat[0], mat[1], mat[2], mat[3])
	defer stop()

	for {
		a, b, c, d, ok := next()
		if !ok {
			break
		}

		assert.Equal(t, []int{a, b, c, d}, []int{a, a + 1, a + 2, a + 3})
	}
}
