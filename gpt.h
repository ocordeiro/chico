#pragma once
#ifdef __cplusplus
extern "C" {
#endif

typedef void (*GoCallback)(char* text);
void goCallback(char* text);

void get_embeddings(char* text, GoCallback callback);

#ifdef __cplusplus
}  // extern "C"
#endif