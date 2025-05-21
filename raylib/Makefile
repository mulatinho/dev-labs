CC=gcc
CFLAGS=-Wall -Wextra
TEST_CFLAGS=-Wall -g -ggdb -DMLT_TESTING=1 -I/usr/include/raylib -I/usr/include
LDLIBS=-lraylib -lm -ldl -lpthread -lGL -lrt -lX11

all: main 

clean:
	rm -fv src/*~ src/*.o src/*.swp src/raylib0[0-9] src/3d src/main
	rm -fv tests/*~ tests/*.o tests/*.swp tests/*_test

image:
	docker build -t main:latest .

push-image:
	docker push mulatinho/main:latest

main:
	gcc $(CFLAGS) -o src/main src/main.c $(LDLIBS)
	gcc $(CFLAGS) -o src/3d src/3d.c $(LDLIBS)
	gcc $(CFLAGS) -o src/raylib01 src/raylib01.c $(LDLIBS)
	gcc $(CFLAGS) -o src/raylib02 src/raylib02.c $(LDLIBS)
	gcc $(CFLAGS) -o src/raylib03 src/raylib03.c $(LDLIBS)

main-debug:
	gcc $(TEST_CFLAGS) -o src/main src/main.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o src/3d src/3d.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o src/raylib01 src/raylib01.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o src/raylib02 src/raylib02.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o src/raylib03 src/raylib03.c $(LDLIBS)

test:
	gcc $(TEST_CFLAGS) -o src/raylib02 src/raylib02.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o src/main.o -c src/main.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o tests/main_test.o -c tests/main_test.c $(LDLIBS)
	gcc $(TEST_CFLAGS) -o tests/main_test src/main.o tests/main_test.o $(LDLIBS) 


format:
	clang-format --verbose -i $(SRC) $(T_SRC)

lint:
	clang-tidy src/*.c tests/*.c -- -std=c99 $(HEADERS)

install:
	install -m 0755 src/main /usr/local/bin/main
	install -m 0644 doc/main.1 /usr/local/share/man/man1/main.1

