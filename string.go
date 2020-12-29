package operator

func ZipStr(keys, values []string) map[string]string {
    res := make(map[string]string, 0)
    for i, item := range values {
        res[keys[i]] = item
    }
    return res
}

func GetOrElseStr(m map[string]string, key, defaultVal string) string {
    v, ok := m[key]
    if ok {
        return v
    }
    return defaultVal
}

func GetOrElsePartialStr(defaultVal string) func(m map[string]string, key string) string {
    return func(m map[string]string, key string) string {
        return GetOrElseStr(m, key, defaultVal)
    }
}
