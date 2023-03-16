all: clean ggml.o libgpt.so

clean:
	rm -f *.o *.so

ggml.o: ggml.c ggml.h
	$(CC) -c ggml.c -o ggml.o

libgpt.so: ggml.o
	$(CXX) -O3 -DNDEBUG -std=c++11 -fPIC -shared -o libgpt.so ggml.o gpt.cpp utils.cpp