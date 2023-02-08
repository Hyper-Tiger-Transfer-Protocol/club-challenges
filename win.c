#include <stdio.h>
#include <stdlib.h>

#define PAYLOAD "Winner!\0"

// Code borrowed from: https://stackoverflow.com/questions/14002954/c-programming-how-to-read-the-whole-file-contents-into-a-buffer

int main() {
    char buf[] = "Winner!\n";
    FILE *f = fopen("chal", "rwb");
    fseek(f,0,SEEK_END);
    int filesize = ftell(f);
    fseek(f,0,SEEK_SET);

    char *contents = malloc(filesize + 1);
    fread(contents, filesize, 1, f);
    int locationOfString = 0;
    for(int i = 0; i < filesize; i++) {
        if (contents[i] == 'L') {
            if(contents[i+1] == 'o') {
                if (contents[i+2] == 's') {
                    if (contents[i+3] == 'e') {
                        if (contents[i+4] == 'r') {
                            if (contents[i+5] == '!') {
                                locationOfString = i;
                            }
                        }
                    }
                }
            }
        }
    }

    printf("%d\n", locationOfString);
    
    //fseek(f, locationOfString, SEEK_SET);
    fclose(f);
    f = fopen("chal", "w");
    int err = fseek(f, locationOfString, SEEK_SET);
    printf("fseek() => %d\n",err);
    err = fwrite(buf, sizeof(char), 8, f);
    printf("fwrite() => %d\n", err);

}   