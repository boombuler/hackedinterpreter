while var_a < input.length - 1 {
    var_b = var_a + 1;
    var_d = true;  // Already sorted?
    while var_b < input.length {
        if input[var_a] > input[var_b] {
            // Switch values at positions var_a and var_b
            var_c = input[var_b];
            input[var_b] = input[var_a];
            input[var_a] = var_c;

            var_b = input.length; // break the inner loop
            var_d = false;        // Mark position var_a as unsorted.
        }
        var_b++;
    }
    if var_d {
        // if var_a is already sorted continue at the next position
        var_a++;
    }
}
return input;