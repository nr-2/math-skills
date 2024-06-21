package nor

//avarge = sum /len
func Mean(data []int) float64 {
    len := float64(len(data))
    var sum float64
    for _, n := range data {
        sum += float64(n)
    }
    mean := sum / len
    return mean
}