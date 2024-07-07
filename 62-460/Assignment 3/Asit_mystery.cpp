#include <stdio.h>

/**
 * Adds 5 to the value pointed to by the parameter `k`.
 * Then, it returns the result of multiplying the updated value by 3 and subtracting 1.
 *
 * @param k A pointer to an integer value.
 * @return The result of the calculation: 3 * (*k + 5) - 1.
 */
int mystery(int* k) {
    *k += 5;
    return 3 * (*k) - 1;
}

/**
 * The main function is the entry point of the program.
 * It calculates the sum1 and sum2 using the mystery function and prints the results.
 *
 * @return 0 indicating successful execution of the program.
 */
int main() {
    int i = 20, j = 20, sum1, sum2;
    sum1 = (i / 2) + mystery(&i);
    sum2 = mystery(&j) + (j / 2);
    
    printf("sum1: %d\n", sum1);
    printf("sum2: %d\n", sum2);

    return 0;
}
