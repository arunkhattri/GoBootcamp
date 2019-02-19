package main

import "fmt"

/*
Unary operators have the highest precedence. As the ++ and -- operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement *p++ is the same as (*p)++.

There are five precedence levels for binary operators. Multiplication operators bind strongest, followed by addition operators, comparison operators, && (logical AND), and finally || (logical OR):

Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
Binary operators of the same precedence associate from left to right. For instance, x / y * z is the same as (x / y) * z.

+x
23 + 3*x[i]
x <= f()
^a >> b
f() || g()
x == y+1 && <-chanPtr > 0
*/
func main() {
	celsius := 37.
	fahrenheit := celsius*9/5 + 32

	fmt.Println("expression 2 + 2 * 4 / 2 results in: ", 2+2*4/2)
	fmt.Println("above is equivalent to: ", 2+((2*4)/2))
	fmt.Println("C: ", celsius, "\nF: ", fahrenheit)
}
