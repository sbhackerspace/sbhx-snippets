#include <stdio.h>
#include <stdlib.h>

int add(int x, int y) {
  return x + y;
}

int main(int argc, char *argv[]) {
  int a = add(10, 20);
  printf("%d\n", a);

  return 0;
}
