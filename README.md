English | [Português](readme.pt.md)

Personal Content Assistant with Artificial Intelligence

Project Description
A fully offline desktop application designed to learn from your personal library of content (books, articles, videos, and more) and answer questions related to the material studied. Using Artificial Intelligence (AI), the application extracts relevant information and provides precise and contextualized answers, acting as a personal assistant for your knowledge needs. It ensures privacy and confidentiality, as all data processing occurs locally on your computer.

Operation
The Assistant is composed of two main parts:

Content Learning: The application extracts information from text files in your personal library and employs advanced Natural Language Processing (NLP) techniques to learn and store relevant information, all offline.
Content Query: The application allows for asking questions about the learned content. It analyzes your questions and, using search algorithms and AI, provides precise and contextualized answers, without the need for an internet connection.

Feature Roadmap
Extraction of content from text files: The application will be able to read and extract information from text files (e.g., PDF, TXT, DOCX) in your personal library.
Contextual content search: The application will implement advanced search algorithms to locate relevant information in the learned content.
Answering questions about content: The application will answer user questions based on the knowledge acquired from the personal content library.
Graphical User Interface: The application will present an intuitive and easy-to-use graphical interface, developed with the Wails framework.

Technologies Used

ggml.cpp: Tensor Library for machine learning.
GPT2 (Word Embeddings): Used for context-based indexing.
Pinecone: Vector database for contextual search.
LLaMA: Facebook LLaMA AI for generating answers to user questions.
Wails: Framework for developing the application's graphical interface.