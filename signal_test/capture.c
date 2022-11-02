#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

const int N_SIGNALS = 4;
const int signal_val[] = { SIGINT, SIGILL, SIGTERM, SIGSEGV };
const char* signal_str[] = { "SIGINT", "SIGILL", "SIGTERM", "SIGSEGV" };

void handle_signal(int sig) {
	for(int i=0; i<N_SIGNALS; ++i) {
		if(signal_val[i] == sig) {
			printf("Received %s (%d)\n", signal_str[i], sig);
			exit(sig);
		}
	}

	printf("Received signal (%d)\n", sig);
	exit(sig);
}

int main() {
	for(int i=0; i<N_SIGNALS; ++i) {
		signal(signal_val[i], handle_signal);
	}

	while(1) {
		sleep(1);
	}

	return 0;
}
