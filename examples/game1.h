function f1: var_a, var_b {
    var_i = 0;
    while var_i < 4 {
        draw(var_b[0]+var_i, var_b[1]);
        draw(var_b[0]+var_i, var_b[1]+2);
        var_i = var_i + 1;
    }

    if var_a > 999 {
        draw_text(var_b[0], var_b[1]+1, var_a);
    }
    else {
        if var_a < 10 {
            draw_text(var_b[0]+2, var_b[1]+1, var_a);
        }
        else {
            draw_text(var_b[0]+1, var_b[1]+1, var_a);
        }
    }
}






f1(1, [1, 1])