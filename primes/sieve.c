/* Prime number Sieve. Single Threaded */

#include <stdio.h>

#define MAX_PRIME 25000
int primes[MAX_PRIME];

main() {
    int n, i, end = 0;
    for(n=2; n<MAX_PRIME; n++) {
        for(i=0; i < end; i++) {
            if(n % primes[i] == 0) goto next;
        }
        primes[end++] = n;
        printf("%d\n", n);
next:   ;
    }
}
