#pragma once
#ifdef __cplusplus
extern "C" {
#endif

typedef void (*GoCallback)(float* embeddings, int size);
void goCallback(float* embeddings, int size);

void get_embeddings(char* text, GoCallback callback);

#ifdef __cplusplus
}  // extern "C"
#endif