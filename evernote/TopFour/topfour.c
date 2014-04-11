#include <stdio.h>

#define MAX	1000000

struct element {
	int v;
	int i;
};

int main() {
	int n, i;
	int num[MAX];
	struct element m, m1, m2, m3;

	scanf("%d\n", &n);
	for(i = 0; i < n; i++)
		scanf("%d\n", &num[i]); 

	m.i = 0; m.v = num[0];
	m1.i = m2.i = m3.i = -1;

	for(i = 1; i < n; i++) {
		if(num[i] > m.v) {
			m3 = m2;
			m2 = m1;
			m1 = m;
			m.i = i;
			m.v = num[i];
			continue;
		}

		if(m1.i == -1 || num[i] > m1.v) {
			m3 = m2;
			m2 = m1;
			m1.i = i;
			m1.v = num[i];
			continue;
		}

		if(m2.i == -1 || num[i] > m2.v) {
			m3 = m2;
			m2.i = i;
			m2.v = num[i];
			continue;
		}

		if(m3.i == -1 || num[i] > m3.v) {
			m3.i = i;
			m3.v = num[i];
		}
	}

	if(m.i != -1)
		printf("%d\n", m.v);

	if(m1.i != -1)
		printf("%d\n", m1.v);

	if(m2.i != -1)
		printf("%d\n", m2.v);

	if(m3.i != -1)
		printf("%d\n", m3.v);
	return 0;
}

