/**
 * The AsitPatel_mystery class contains a method to calculate a mystery value based on an input integer.
 */
public class Asit_mystery {

    /**
     * Calculates the mystery value based on the given integer.
     *
     * @param k the input integer
     * @return the mystery value calculated using the input integer
     */
    public static int mystery(int k) {
        k += 5;
        return 3 * k - 1;
    }

    /**
     * The main method is the entry point of the program.
     * It performs some calculations using the mystery method and prints the results.
     *
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        int i = 20, j = 20, sum1, sum2;
        sum1 = (i / 2) + mystery(i);
        sum2 = mystery(j) + (j / 2);

        System.out.println("sum1: " + sum1);
        System.out.println("sum2: " + sum2);
    }
}