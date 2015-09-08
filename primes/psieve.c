/* Prime number sieve. Multithreaded using pthread and semaphores. */

#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#define MAX_PRIME 25000
#define QUEUE_SIZE 4

struct proc {
    struct proc *link;
    int buf[QUEUE_SIZE];
    sem_t avail, N;
    pthread_t thread;
} 
*proc_init(struct proc *parent, void *(*routine)(void*))
{
    struct proc *cb = malloc(sizeof(struct proc));
    assert(cb);
    cb->link = parent;
    sem_init(&cb->avail, 0, QUEUE_SIZE);
    sem_init(&cb->N,     0, 0);
    if(routine)
	pthread_create(&cb->thread, NULL, routine, cb);
    return cb;
}

void *enumerator(void *arg)
{
    struct proc *cb = arg;
    int x = 2;
    int i = 0;
    do {
	sem_wait(&cb->avail);
	cb->buf[i] = x++;
	sem_post(&cb->N);
	i = (i+1)%QUEUE_SIZE;
    } while(1);
}

void *outputter(void *arg)
{
    struct proc *cb = arg;
    int i = 0;
    do {
	sem_wait(&cb->link->N);
        int n = cb->link->buf[i];
	sem_post(&cb->link->avail);
	printf("%d\n", n);
	i = (i+1)%QUEUE_SIZE;
    } while(1);
}

void *sifter(void *arg)
{
    struct proc *cb = arg;
    struct proc *next = NULL;

    int i = 0, j = 0;

    sem_wait(&cb->link->N);
    int p = cb->link->buf[0];
    sem_post(&cb->link->avail);
    i = (i+1)%QUEUE_SIZE;

    if(p >= MAX_PRIME) exit(1);
    printf("%d\n", p);

    do {
	int n;
	sem_wait(&cb->link->N);
	n = cb->link->buf[i];
	sem_post(&cb->link->avail);
	if(n % p != 0) {
	    sem_wait(&cb->avail);
	    cb->buf[j] = n;
	    sem_post(&cb->N);
	    j = (j+1)%QUEUE_SIZE;
	    if(!next)
		next = proc_init(cb, sifter);
	}
	i = (i+1)%QUEUE_SIZE;
    } while(1);
}

int main(void) 
{
    struct proc *main = proc_init(NULL, NULL);
    proc_init(main, sifter);
    enumerator(main);
}
