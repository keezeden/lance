package main

func contains(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}