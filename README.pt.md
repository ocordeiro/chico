Português | [English](README.md)

## Assistente Pessoal de Conteúdo com Inteligência Artificial

### Descrição do Projeto

Um aplicativo desktop totalmente offline, projetado para aprender com sua biblioteca pessoal de conteúdo (livros, artigos, vídeos e outros) e responder a perguntas relacionadas ao material estudado. Utilizando Inteligência Artificial (IA), o aplicativo extrai informações relevantes e fornece respostas precisas e contextualizadas, atuando como um assistente pessoal para suas necessidades de conhecimento. 
Garante privacidade e confidencialidade, pois todo o processamento de dados ocorre localmente em seu computador.

### Funcionamento

O Assistente é composto por duas partes principais:

Aprendizado de Conteúdo: O aplicativo extrai informações de arquivos de texto em sua biblioteca pessoal e emprega técnicas avançadas de Processamento de Linguagem Natural (PLN) para aprender e armazenar informações relevantes, tudo isso de maneira offline.
Consulta de Conteúdo: O aplicativo possibilita a realização de perguntas sobre o conteúdo aprendido. Ele analisa suas questões e, com o uso de algoritmos de busca e IA, fornece respostas precisas e contextualizadas, sem necessidade de conexão à internet.

### Roadmap de Recursos

- Extração de conteúdo de arquivos de texto: O aplicativo será capaz de ler e extrair informações de arquivos de texto (por exemplo, PDF, TXT, DOCX) em sua biblioteca pessoal.
- Busca contextual de conteúdo: O aplicativo implementará algoritmos avançados de busca para localizar informações relevantes no conteúdo aprendido.
- Resposta a perguntas sobre o conteúdo: O aplicativo responderá às perguntas dos usuários com base no conhecimento adquirido a partir da biblioteca pessoal de conteúdo.
- Interface gráfica: O aplicativo apresentará uma interface gráfica intuitiva e fácil de usar, desenvolvida com o framework Wails.

## Tecnologias Utilizadas

- ggml.cpp: Biblioteca Tensor para aprendizado de máquina.
- GPT2 (Word Embeddings): Utilizado para indexação baseada em contexto.
- Pinecone: Banco de dados de vetores para pesquisa contextual.
- LLaMA: IA do Facebook LLaMA para geração de respostas às perguntas dos usuários.
- Wails: Framework para desenvolvimento da interface gráfica do aplicativo.
