package nl.servicehouse.demo;

public class BinomialDistribution {

    public static void main(String[] args) {
        combinatorial(27,9);
    }

    private static long combinatorial(long m, long r) {
        return (m >= r) ? factorial(m) / (factorial(r) * factorial(m - r)) : 1;
    }

    private static long factorial(long x) {
        return (x >= 1) ? x * factorial(x - 1) : 1;
    }

}
