CC=gcc
GOC=gccgo
CFLAGS=-g -Wall -Werror
GOFLAGS=-fgo-prefix=ds3 -g
LDFLAGS=-g
LIBS=-lgo -ldl
OBJS=obj/main.o obj/ds3.o obj/gomain.o

ds3_client: $(OBJS)
	$(CC) $(LDFLAGS) -o bin/ds3_client $^ $(LIBS)

obj/%.o: src/%.c
	$(CC) -c $(CFLAGS) $< -o $@

obj/%.o: src/%.go
	$(GOC) -c $(GOFLAGS) $< -o $@

# Don't use the prefix for the main.
obj/gomain.o: src/gomain.go
	$(GOC) -c $< -o $@

clean:
	rm -rf bin/* obj/*

