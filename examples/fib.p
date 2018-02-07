
func fib(atom n) {
    if n < 3 {
        return 1;
    }
    return fib(n-1) + fib(n-2);
}

func main() {
    atom b;
    b = fib(4);
}
