## GPT-2 Word Token Embeddings (WTE) Extractor

### Description
O script `gpt2_wte.py` aims to extract the weights of Word Token Embeddings (WTE) from the GPT-2 model, creating an optimized ggml model with less than 200MB.
The embeddings are vector representations that are used as input in the GPT-2 model.
By extracting these embeddings from a text, we can store them in a vector database for performing semantics.

### How to use
```bash
pip install transformers numpy
python gpt2_wte.py```