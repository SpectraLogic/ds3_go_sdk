#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <dlfcn.h>

typedef struct {
  const unsigned char *data;
  int length;
} go_string;

extern go_string get() __asm__ ("ds3.ds3_c.Get");
extern void __go_init_main(void) __asm__ ("__go_init_main");

int main(int argc, char *argv[]) {
    void (*runtime_mallocinit)(void);
    void (*runtime_cpuprofinit)(void);
    void (*__go_gc_goroutine_init)(int*);
    void (*__initsig)(void);
    void (*__go_enable_gc)(void);

    printf("line 1\n");

    void *libgo;
    //TODO: better error handling
    char *libgopath = "/usr/lib/x86_64-linux-gnu/libgo.so.4.0.0";
    if (!(libgo = dlopen(libgopath, RTLD_NOW | RTLD_GLOBAL))) {
        fprintf(stderr, "%s\n", dlerror());
        return 1;
    }
    printf("line 2\n");

    *(void **)(&runtime_mallocinit) = dlsym(libgo, "runtime_mallocinit");
    *(void **)(&runtime_cpuprofinit) = dlsym(libgo, "runtime_cpuprofinit");
    *(void **)(&__go_gc_goroutine_init) = dlsym(libgo, "__go_gc_goroutine_init");
    *(void **)(&__initsig) = dlsym(libgo, "__initsig");
    *(void **)(&__go_enable_gc) = dlsym(libgo, "__go_enable_gc");

    printf("line 3\n");
    //dlclose(libgo);

    printf("line 4\n");
    //runtime_mallocinit();
    printf("line 5\n");
    //runtime_cpuprofinit();
    printf("line 6\n");
    //__go_gc_goroutine_init(&argc);
    printf("line 7\n");
    //__initsig();
    printf("line 8\n");

#if defined(HAVE_SRANDOM)
    srandom ((unsigned int) time (NULL));
#else
    srand ((unsigned int) time (NULL));
#endif

    printf("line 9\n");
    __go_init_main();
    printf("line 10\n");
    __go_enable_gc();
    printf("line 11\n");

    // use go
    //go_string foo = { "hello", 5 };
    //printf("%s\n", get().data);
    printf("line 12\n");


    return 0;
}

