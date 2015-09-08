/* Prime number sieve. Multiprocess, using pipes and forks. */

#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#define MAX_PRIME 25000

void enumerator(int out)
{
    int x = 2;
    do {
	write(out, &x, sizeof x);
	++x;
    } while(1);
}

void outputter(int in)
{
    do {
        int n;
	read(in, &n, sizeof n);
        if(n >= MAX_PRIME) exit(0);
	printf("%d\n", n);
    } while(1);
}

void sifter(int in, int out)
{
    pid_t pid = 0;
    int p;
    read(in, &p, sizeof p);
    write(out, &p, sizeof p);

    int sub[2];
    pipe(sub);

    do {
	int n;
	read(in, &n, sizeof n);
	if(n % p != 0) {
            if(!pid && (pid=fork()) == 0) {
                sifter(sub[0], out);
            }
            write(sub[1], &n, sizeof n);
	}
    } while(1);
}

int main(void) 
{
    int io[4];
    pipe(io);
    pipe(io+2);
    if(fork() == 0) {
        enumerator(io[1]);
    } else if(fork() == 0) {
       sifter(io[0],io[3]);
    }
    outputter(io[2]);
}
