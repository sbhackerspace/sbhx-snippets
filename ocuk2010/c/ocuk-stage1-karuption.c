#include <stdio.h>

int main(void)

{
  unsigned int n, chain, num;
  unsigned int longest_chain = 0;
  unsigned int longest_n = 0;

  for(n = 1;n < 1500000; n++) {
      num = n;
      chain = 0;

      while (num != 1) {
	  if (!(num&1)) {
	      num >>=1;
	  } else {
	    num += (num<<1) + 1;
	  }
	  chain++;
      } if (chain > longest_chain) {
	longest_chain = chain;
	longest_n = n;
      }
  }

  printf("%d", longest_n);

  return 0;
}
