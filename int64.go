package operator

func ContainsInt64(l []int64, item int64) bool {
    for _, v := range l {
        if v == item {
            return true
        }
    }
    return false
}

func FilterInt64(l []int64, condition func(item int64) bool) []int64 {
    res := make([]int64, 0)
    for _, i := range res {
        if condition(i) {
            res = append(res, i)
        }
    }
    return res
}

// 去重
func DistinctInt64(l []int64) []int64 {
    return nil
}
