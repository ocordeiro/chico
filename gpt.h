#pragma once
#ifdef __cplusplus
extern "C" {
#endif

typedef void (*GoCallback)(float* text);
void goCallback(float* text);

void get_embeddings(char* text, GoCallback callback);

#ifdef __cplusplus
}  // extern "C"
#endif