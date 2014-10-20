foreach var_b in input {
    if var_b == "(" {
        var_a++
    } else {
        var_a--
    }
    if var_a < 0 {
        return false
    }
}
return var_a == 0