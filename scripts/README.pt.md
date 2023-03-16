Português | English (README.md)

## Word Token Embeddings (WTE) para GPT-2

### Descrição

O script `gpt2_wte.py` tem como objetivo extrair os weights (pesos) de Word Token Embeddings (WTE) do modelo GPT-2 criando um modelo ggml otimizado com menos de 200MB. 
Os embeddings são representações vetoriais que são utilizados como entrada no modelo GPT-2.
Ao extrair esses embeddings de um texto, podemos armazenálos em um banco de dados vetorial para realizar semanticas.

#### Como usar
```bash
pip install transformers numpy
python gpt2_wte.py
```