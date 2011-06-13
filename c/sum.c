#include <stdio.h>
#include <stdlib.h>

int sum(int intArray[]) {
    int i, total;
    total = 0;
    //int len = sizeof(intArray)/sizeof(int);
    for (i = 0; i < 3; ++i) {
        total += intArray[i];
    }
    return total;
}

void main(int argc, char *argv[]) {
    int intArray[] = {1,2,3};
    printf("%d\n", sum(intArray));
}
