// Node Info Cli App
#include <sys/utsname.h>
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

int main()
{
    // define utsname struct
    struct utsname info;
    // call uname() and check call
    if (uname(&info) < 0)
    {
        // error if call fails
        perror("uname");
        fprintf(stderr, "Error in uname : %d\n", errno);
        exit(EXIT_FAILURE);
    } else {
        // Print Node Info
        printf("Node Info\n----------\n");
        printf("node : %s\n", info.nodename);
        printf("system : %s\n", info.sysname);
        printf("version : %s\n", info.version);
        printf("release : %s\n", info.release);
        printf("arch : %s\n", info.machine);
        exit(EXIT_SUCCESS);
    }
}
