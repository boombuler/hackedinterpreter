// Draw Board
function f1: var_a, var_b {
    function f1: var_a, var_b, var_c {
        if var_c == 0 {
            return 0;
        }
        var_a = var_a * 4 + 2;
        var_b = var_b * 3 + 3;

        if var_c > 999 {
            draw_text(var_a, var_b, var_c);
        }
        else {
            if var_c < 10 {
                draw_text(var_a+2, var_b, var_c);
            }
            else {
                draw_text(var_a+1, var_b, var_c);
            }
        }
    }


    var_c = 0;
    while var_c < 19 {
        draw(var_c, 1);
        draw(var_c, 14);
        if var_c < 14 {
            draw(1, var_c);
            draw(18, var_c);
        }

        var_c = var_c + 1;
    }


    var_c = 0;
    while var_c < 4 {
        var_d = 0;
        while var_d < 4 {
            f1(var_c, var_d, var_b[var_c][var_d]);
            var_d = var_d + 1;
        }
        var_c = var_c + 1;
    }
    draw_text(2, height() - 2, var_a[1])
}

// Add Random Block
function f2: var_b {
    if random(100) > 80 {
        var_a = 4
    } else {
        var_a = 2
    }

    var_c = random(4);
    var_d = random(4);
    while var_b[var_c][var_d] != 0 {
        var_c = random(4);
        var_d = random(4);
    }

    var_b[var_c][var_d] = var_a
}

// Handle Input
function f3: var_a, var_b {
    var_m = false; // Any Moved.

    if left() {
        var_d = 0;
        while var_d < 4 {
            var_c = 1;
            while var_c < 4 {
                if var_b[var_c][var_d] != 0 {
                    var_f = var_c;
                    while var_f > 0 {
                        if var_b[var_f - 1][var_d] == 0 {
                            var_b[var_f - 1][var_d] = var_b[var_f][var_d];
                            var_b[var_f][var_d] = 0;
                            var_m = true;
                        }
                        else {
                            if var_b[var_f - 1][var_d] == var_b[var_f][var_d] {
                                var_b[var_f - 1][var_d] = var_b[var_f - 1][var_d] * (0 - 2);
                                var_b[var_f][var_d] = 0;
                                var_m = true;
                            }
                            var_f = 0
                        }
                        var_f = var_f - 1;
                    }
                }
                var_c = var_c + 1;
            }
            var_d = var_d + 1;
        }
        return var_m;
    }
    if up() {
        var_d = 0;
        while var_d < 4 {
            var_c = 1;
            while var_c < 4 {
                if var_b[var_d][var_c] != 0 {
                    var_f = var_c;
                    while var_f > 0 {
                        if var_b[var_d][var_f - 1] == 0 {
                            var_b[var_d][var_f - 1] = var_b[var_d][var_f];
                            var_b[var_d][var_f] = 0;
                            var_m = true;
                        }
                        else {
                            if var_b[var_d][var_f - 1] == var_b[var_d][var_f] {
                                var_b[var_d][var_f - 1] = var_b[var_d][var_f - 1] * (0 - 2);
                                var_b[var_d][var_f] = 0;
                                var_m = true;
                            }
                            var_f = 0
                        }
                        var_f = var_f - 1;
                    }
                }
                var_c = var_c + 1;
            }
            var_d = var_d + 1;
        }
        return var_m;
    }
    if right() {
        var_d = 0;
        while var_d < 4 {
            var_c = 3;
            while var_c > (0-1) {
                if var_b[var_c][var_d] != 0 {
                    var_f = var_c;
                    while var_f < 3 {
                        if var_b[var_f + 1][var_d] == 0 {
                            var_b[var_f + 1][var_d] = var_b[var_f][var_d];
                            var_b[var_f][var_d] = 0;
                            var_m = true;
                        }
                        else {
                            if var_b[var_f + 1][var_d] == var_b[var_f][var_d] {
                                var_b[var_f + 1][var_d] = var_b[var_f + 1][var_d] * (0 - 2);
                                var_b[var_f][var_d] = 0;
                                var_m = true;
                            }
                            var_f = 4
                        }
                        var_f = var_f + 1;
                    }
                }
                var_c = var_c - 1;
            }
            var_d = var_d + 1;
        }
        return var_m;
    }
    if down() {
        var_d = 0;
        while var_d < 4 {
            var_c = 3;
            while var_c > (0-1) {
                if var_b[var_d][var_c] != 0 {
                    var_f = var_c;
                    while var_f < 3 {
                        if var_b[var_d][var_f + 1] == 0 {
                            var_b[var_d][var_f + 1] = var_b[var_d][var_f];
                            var_b[var_d][var_f] = 0;
                            var_m = true;
                        }
                        else {
                            if var_b[var_d][var_f + 1] == var_b[var_d][var_f] {
                                var_b[var_d][var_f + 1] = var_b[var_d][var_f + 1] * (0 - 2);
                                var_b[var_d][var_f] = 0;
                                var_m = true;
                            }
                            var_f = 4
                        }
                        var_f = var_f + 1;
                    }
                }
                var_c = var_c - 1;
            }
            var_d = var_d + 1;
        }
        return var_m;
    }
    return false;
}

// Check GameOver + Update Score
function f4: var_a, var_b {
    var_e = true;

    foreach var_c in var_b {
        var_d = 0;
        while var_d < 4 {
            var_f = var_c[var_d];
            if var_f < 0 {
                var_c[var_d] = 0 - var_f;
                var_a[1] = var_a[1] - var_f;
                if var_c[var_d] == 2048 {
                    var_a[2] = 1;
                }
            }

            var_d = var_d + 1;
        }
    }

    if var_a[2] == 0 {
        var_c = 0;
        while var_c < 4 {
            var_d = 0;
            while var_d < 4 {
                var_f = var_b[var_c][var_d];
                if var_f == 0 {
                    var_e = false;
                }
                if var_d > 0 {
                    if var_b[var_c][var_d - 1] == var_f {
                        var_e = false;
                    }
                }
                if var_c > 0 {
                    if var_b[var_c - 1][var_d] == var_f {
                        var_e = false;
                    }
                }
                var_d = var_d + 1;
            }
            var_c = var_c + 1;
        }

        if var_e {
            var_a[2] = 2;
        }
    }
}

function f5: var_a {
    var_c = (width() - 13) / 2;

    draw_text(var_c, 11, var_a[1]);

    draw(var_c, 4);                                                       draw(var_c+4, 4);
    draw(var_c, 5);                                                       draw(var_c+4, 5);
    draw(var_c, 6);                                                       draw(var_c+4, 6);
    draw(var_c, 7);                   draw(var_c+2, 7);                   draw(var_c+4, 7);
    draw(var_c, 8);                   draw(var_c+2, 8);                   draw(var_c+4, 8);
                    draw(var_c+1, 9);                   draw(var_c+3, 9);

    draw(var_c+6, 4); draw(var_c+7, 4); draw(var_c+8, 4);
    draw(var_c+6, 5);                   draw(var_c+8, 5);
    draw(var_c+6, 6);                   draw(var_c+8, 6);
    draw(var_c+6, 7);                   draw(var_c+8, 7);
    draw(var_c+6, 8);                   draw(var_c+8, 8);
    draw(var_c+6, 9); draw(var_c+7, 9); draw(var_c+8, 9);

    draw(var_c+10, 4);                                       draw(var_c+13, 4);
    draw(var_c+10, 5);                                       draw(var_c+13, 5);
    draw(var_c+10, 6); draw(var_c+11, 6);                    draw(var_c+13, 6);
    draw(var_c+10, 7);                    draw(var_c+12, 7); draw(var_c+13, 7);
    draw(var_c+10, 8);                                       draw(var_c+13, 8);
    draw(var_c+10, 9);                                       draw(var_c+13, 9);
}

function f6: var_a {
    var_c = (width() - 14) / 2;

    draw_text(var_c, 11, var_a[1]);

    draw(var_c, 4);
    draw(var_c, 5);
    draw(var_c, 6);
    draw(var_c, 7);
    draw(var_c, 8);
    draw(var_c, 9); draw(var_c+1, 9); draw(var_c+2, 9);

    draw(var_c+4, 4); draw(var_c+5, 4); draw(var_c+6, 4);
    draw(var_c+4, 5);                   draw(var_c+6, 5);
    draw(var_c+4, 6);                   draw(var_c+6, 6);
    draw(var_c+4, 7);                   draw(var_c+6, 7);
    draw(var_c+4, 8);                   draw(var_c+6, 8);
    draw(var_c+4, 9); draw(var_c+5, 9); draw(var_c+6, 9);

    draw(var_c+8, 4); draw(var_c+9, 4); draw(var_c+10, 4);
    draw(var_c+8, 5);
    draw(var_c+8, 6); draw(var_c+9, 6);
                      draw(var_c+9, 7); draw(var_c+10, 7);
                                        draw(var_c+10, 8);
    draw(var_c+8, 9); draw(var_c+9, 9); draw(var_c+10, 9);

    draw(var_c+12, 4); draw(var_c+13, 4); draw(var_c+14, 4);
                       draw(var_c+13, 5);
                       draw(var_c+13, 6);
                       draw(var_c+13, 7);
                       draw(var_c+13, 8);
                       draw(var_c+13, 9);
}

if var_b == 0 {
    var_b = new_list(4).map(function var_a -> new_list(4));
    var_a = [
        time() - 500, // Last Input
        0,            // Score
        0            // State 0 = in game / 1 = Won / 2 = Lost
    ];
    // Add two random blocks
    f2(var_b);
    f2(var_b);
}

if var_a[2] == 0 {
    f1(var_a, var_b);
    if (time() - var_a[0]) > 200 { // Last input at least 200ms ago?
        if f3(var_a, var_b) { // Something moved
            var_a[0] = time();
            f2(var_b);        // Add random block
            f4(var_a, var_b); // check for game over
        }
    }
} else {
    if var_a[2] == 1 {
        f5(var_a)
    } else {
        f6(var_a)
    }
    if a_btn() || b_btn() {
        var_b = 0; // restart game
    }
}

