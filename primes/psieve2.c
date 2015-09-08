/* Prime number sieve. Multithreaded using pthreads and condition vars */

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#define MAX_PRIME 25000
#define QUEUE_SIZE 4

struct proc {
    struct proc *link;
    int buf[QUEUE_SIZE];
    int N;
    pthread_mutex_t mux;
    pthread_cond_t ok_to_write, ok_to_read;
    pthread_t thread;
} 
*proc_init(struct proc *parent, void *(*routine)(void*))
{
    struct proc *cb = malloc(sizeof(struct proc));
    assert(cb);
    cb->link = parent;
    cb->N = 0;
    pthread_mutex_init(&cb->mux, NULL);
    pthread_cond_init(&cb->ok_to_read, NULL);
    pthread_cond_init(&cb->ok_to_write, NULL);
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
	pthread_mutex_lock(&cb->mux);
	while(cb->N == QUEUE_SIZE) {
	    pthread_cond_wait(&cb->ok_to_write, &cb->mux);
	}
	cb->buf[i] = x++;
	++cb->N;
	pthread_mutex_unlock(&cb->mux);
	i = (i+1)%QUEUE_SIZE;
	pthread_cond_signal(&cb->ok_to_read);
    } while(1);
}

void *outputter(void *arg)
{
    struct proc *cb = arg;
    int i = 0;
    do {
	pthread_mutex_lock(&cb->link->mux);
	while(cb->link->N == 0)  {
	    pthread_cond_wait(&cb->link->ok_to_read, &cb->link->mux);
	}
        int n = cb->link->buf[i];
	--cb->link->N;
	pthread_mutex_unlock(&cb->link->mux);
	pthread_cond_signal(&cb->link->ok_to_write);
	i = (i+1)%QUEUE_SIZE;
	printf("%d\n", n);
    } while(1);
}

void *sifter(void *arg)
{
    struct proc *cb = arg;
    struct proc *next = NULL;

    int i = 0, j = 0;

    pthread_mutex_lock(&cb->link->mux);
    while(cb->link->N == 0)  {
	pthread_cond_wait(&cb->link->ok_to_read, &cb->link->mux);
    }
    int p = cb->link->buf[0];
    --cb->link->N;
    pthread_mutex_unlock(&cb->link->mux);
    pthread_cond_signal(&cb->link->ok_to_write);
    i = (i+1)%QUEUE_SIZE;

    if(p >= MAX_PRIME) exit(1);
    printf("%d\n", p);

    do {
	int n;
	pthread_mutex_lock(&cb->link->mux);
	while(cb->link->N == 0)  {
	    pthread_cond_wait(&cb->link->ok_to_read, &cb->link->mux);
	}
	n = cb->link->buf[i];
	--cb->link->N;
	pthread_mutex_unlock(&cb->link->mux);
	pthread_cond_signal(&cb->link->ok_to_write);
	if(n % p != 0) {
	    pthread_mutex_lock(&cb->mux);
	    while(cb->N == QUEUE_SIZE)  {
		pthread_cond_wait(&cb->ok_to_write, &cb->mux);
	    }
	    cb->buf[j] = n;
	    ++cb->N;
	    pthread_mutex_unlock(&cb->mux);
	    pthread_cond_signal(&cb->ok_to_read);
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
