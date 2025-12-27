package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// merge sort rekursif
func mergeRecursive(U, V []int) []int {
    h := len(U)
    m := len(V)
    S := make([]int, h+m)

    i, j, k := 0, 0, 0
    for i < h && j < m {
        if U[i] < V[j] {
            S[k] = U[i]
            i++
        } else {
            S[k] = V[j]
            j++
        }
        k++
    }
    for i < h {
        S[k] = U[i]
        i++
        k++
    }
    for j < m {
        S[k] = V[j]
        j++
        k++
    }
    return S
}

func mergesortRecursive(S []int) []int {
    n := len(S)
    if n <= 1 {
        return S
    }
    h := n / 2
    U := make([]int, h)
    V := make([]int, n-h)
    copy(U, S[:h])
    copy(V, S[h:])
    U = mergesortRecursive(U)
    V = mergesortRecursive(V)
    return mergeRecursive(U, V)
}

// === Merge Sort Iteratif (Bottom-Up) ===
func mergeIterative(src, dst []int, left, mid, right int) {
    i, j, k := left, mid, left
    for i < mid && j < right {
        if src[i] <= src[j] {
            dst[k] = src[i]
            i++
        } else {
            dst[k] = src[j]
            j++
        }
        k++
    }
    for i < mid {
        dst[k] = src[i]
        i++
        k++
    }
    for j < right {
        dst[k] = src[j]
        j++
        k++
    }
}

func mergesortIterative(a []int) []int {
    n := len(a)
    if n <= 1 {
        return append([]int(nil), a...)
    }
    src := append([]int(nil), a...)
    dst := make([]int, n)
    width := 1
    for width < n {
        for left := 0; left < n; left += 2 * width {
            mid := left + width
            if mid > n {
                mid = n
            }
            right := left + 2*width
            if right > n {
                right = n
            }
            mergeIterative(src, dst, left, mid, right)
        }
        src, dst = dst, src
        width *= 2
    }
    return src
}

// === Helpers ===
func genData(n int, rng *rand.Rand) []int {
    a := make([]int, n)
    for i := range a {
        a[i] = rng.Intn(2_000_000_001) - 1_000_000_000
    }
    return a
}

func isSorted(a []int) bool {
    for i := 1; i < len(a); i++ {
        if a[i] < a[i-1] {
            return false
        }
    }
    return true
}

func avg(vals []float64) float64 {
    sum := 0.0
    for _, v := range vals {
        sum += v
    }
    return sum / float64(len(vals))
}

func benchmark(name string, sortFn func([]int) []int, sizes []int, trials int, seed int64) [][2]float64 {
    rng := rand.New(rand.NewSource(seed))
    results := make([][2]float64, 0, len(sizes))
    for _, n := range sizes {
        times := make([]float64, 0, trials)
        for t := 0; t < trials; t++ {
            data := genData(n, rng)	// generate data untuk dilakukan sort
            start := time.Now()
            out := sortFn(data) // algoritma sort yang dipilih untuk dilakukan pengujian
            elapsed := float64(time.Since(start).Milliseconds()) // menentukan waktu yang jenis waktu miliseconds
            if !isSorted(out) { // cek apakah output sudah terurut
                log.Fatalf("%s failed: output not sorted for n=%d", name, n)
            }
            times = append(times, elapsed) // menambahkan waktu ke dalam array
        }
        avgTime := avg(times)
        results = append(results, [2]float64{float64(n), avgTime})
    }
    return results
}

// === Main Program ===
func main() {
    sizes := []int{1000, 5000, 10000, 50000, 100000}
    trials := 5
    seed := int64(42)

    // Rekursif dulu
    resRec := benchmark("Recursive", mergesortRecursive, sizes, trials, seed)
    fmt.Println("\n=== Hasil Benchmark Merge Sort Rekursif (Top-Down) ===")
    fmt.Println("\n  n data   | avg_time (ms)")
    for _, r := range resRec {
        fmt.Printf("n=%-8d | %8.3f ms\n", int(r[0]), r[1])
    }

    // Lalu iteratif
    resIter := benchmark("Iterative", mergesortIterative, sizes, trials, seed)
    fmt.Println("\n=== Hasil Benchmark Merge Sort Iteratif ===")
    fmt.Println("\n  n data   | avg_time (ms)")

    for _, r := range resIter {
        fmt.Printf("n=%-8d | %8.3f ms\n", int(r[0]), r[1])
    }
}
