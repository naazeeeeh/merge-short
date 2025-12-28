// group project
// Judul: Analisis
// anggota kelompok
// Nazeeh - 103042400055
// Agung Pratama - 103042400015
// Kallistus Wahyu Sandivan - 103042400068

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// mergeRekursif menggabungkan array U dan V
func mergeRekursif(U, V []int) []int {
	var S []int;
	var h, m, i, j, k int;
    h = len(U) // panjang array U
    m = len(V) // panjang array V
    S = make([]int, h+m) // array S dengan panjang h+m

    i, j, k = 0, 0, 0
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

func mergesortRekursif(S []int) []int {

	var n, h int;
	var U, V []int;

    // jika panjang array S hanya 1 atau kurang dari 1,
    // maka array S sudah terurut
    n = len(S)
    if n <= 1 {
        return S
    }

    // membagi array S menjadi 2 bagian
    h = n / 2
    U = make([]int, h)
    V = make([]int, n-h)

    // menyalin U dan V dengan nilai S
    copy(U, S[:h]) //fungsi GO untuk copy elemen dari array S ke array U
    copy(V, S[h:]) //copy dari array S ke array V

    // rekursif untuk melakukan merge sort pada U dan V
    U = mergesortRekursif(U)
    V = mergesortRekursif(V)

    // menggabungkan U dan V
    return mergeRekursif(U, V)
}

// mergesort iteratif
func mergeIteratif(src, dst []int, left, mid, right int) {
	var i, j, k int;
    i, j, k = left, mid, left
    
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

func mergesortIteratif(a []int) []int {

	var n, width, left, mid, right int;
	var src, dst []int;

    n = len(a)
    if n <= 1 {
        return append([]int(nil), a...)
    }
    src = append([]int(nil), a...)
    dst = make([]int, n)
    width = 1
    for width < n {
        for left = 0; left < n; left += 2 * width {
            mid = left + width
            if mid > n {
                mid = n
            }
            right = left + 2*width
            if right > n {
                right = n
            }
            mergeIteratif(src, dst, left, mid, right)
        }
        src, dst = dst, src
        width *= 2
    }
    return src
}

// generate data random tidak terurut 
// untuk dilakukan sorting
func generateData(n int, rng *rand.Rand) []int {
	var a []int;
	var i int;

    a = make([]int, n)
    for i = range a {
        a[i] = rng.Intn(2_000_000_001) - 1_000_000_000
    }
    return a
}

func isSorted(a []int) bool {
	var i int;
    for i = 1; i < len(a); i++ {
        if a[i] < a[i-1] {
            return false
        }
    }
    return true
}

// menghitung rata-rata dari waktu + trials
func avg(vals []float64) float64 {
	var sum, v float64;
    sum = 0.0
    for _, v = range vals {
        sum += v
    }
    return sum / float64(len(vals))
}

func eksperimenPerbandingan(name string, sortFn func([]int) []int, sizes []int, trials int, seed int64) [][2]float64 {
	var rng *rand.Rand;
	var n, t int;
	var data []int;

    rng = rand.New(rand.NewSource(seed))
    results := make([][2]float64, 0, len(sizes))
    for _, n = range sizes {
        times := make([]float64, 0, trials)
        for t = 0; t < trials; t++ {
            data = generateData(n, rng)	// generate data untuk dilakukan sort
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
    resRec := eksperimenPerbandingan("Rekursif", mergesortRekursif, sizes, trials, seed)
    fmt.Println("\nHasil Perbandingan Merge Sort Rekursif")
    fmt.Println("\n  n data   | avg_time (ms)")
    for _, r := range resRec {
        fmt.Printf("n=%-8d | %8.3f ms\n", int(r[0]), r[1])
    }

    // Lalu iteratif
    resIter := eksperimenPerbandingan("Iteratif", mergesortIteratif, sizes, trials, seed)
    fmt.Println("\nHasil Perbandingan Merge Sort Iteratif")
    fmt.Println("\n  n data   | avg_time (ms)")

    for _, r := range resIter {
        fmt.Printf("n=%-8d | %8.3f ms\n", int(r[0]), r[1])
    }
}
