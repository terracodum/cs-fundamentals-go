# cs-fundamentals-go

Закрытие CS-базы (структуры данных, алгосы, сложность) руками на Go, перед собесами (Avito, backend intern).

Формат: NeetCode 150 как основной трек + пара структур данных с нуля (hash table, linked list) там, где Go не даёт этого бесплатно.

## Структура репо

```
cs-fundamentals-go/
├── README.md
├── from-scratch/           # структуры данных с нуля, без stdlib
│   ├── linked-list/        # односвязный список
│   ├── doubly-linked-list/
│   ├── stack-queue/
│   ├── dynamic-array/
│   ├── hashtable/
│   ├── bst/
│   ├── heap/
│   ├── graph/
│   ├── trie/
│   └── dsu/
├── linked-list/            # neetcode задачи по теме
├── stack-queue/
├── trees/
├── hashmap/
├── graphs/
├── dp/
└── misc/                   # sorting, two-pointers, binary search и что не влезло в темы выше
```

Каждая задача — своя папка: `topic/00X-task-name/main.go` + `main_test.go`.

## Порядок прохождения

### 1. Структуры данных (руками, без библиотек)

- [ ] `from-scratch/linked-list` — односвязный список (insert/delete/reverse/find middle/detect cycle)
- [ ] `from-scratch/doubly-linked-list` — двусвязный список
- [ ] `from-scratch/stack-queue` — стек и очередь (поверх своего списка или slice)
- [ ] `from-scratch/dynamic-array` — динамический массив (рост capacity, как в Go slice)
- [ ] `from-scratch/hashtable` — хэш-таблица с нуля (свой hash function + collision resolution)
- [ ] `from-scratch/bst` — BST: insert/delete/search, обходы in-order/pre-order/post-order, BFS по уровням
- [ ] `from-scratch/heap` — куча: min-heap и max-heap на массиве
- [ ] `from-scratch/graph` — граф: adjacency list/matrix, BFS, DFS
- [ ] `from-scratch/trie` — префиксное дерево
- [ ] `from-scratch/dsu` — Disjoint Set Union (Union-Find)

### 2. NeetCode 150 по темам

- [ ] `linked-list` — NeetCode задачи по теме
- [ ] `stack-queue`
- [ ] `trees` — BST, обходы, BFS/DFS
- [ ] `hashmap`
- [ ] `graphs` — BFS/DFS, Dijkstra, topological sort
- [ ] `dp` — если останется время

## Правила

- Без гуглинга паттерна решения. Если завис на 20+ минут — смотрю разбор, но переписываю сам, не копирую.
- Каждая задача — тест, не просто main с print.
- Раз в неделю — ретро: что реально было gap'ом, а что уже знал.

## Полезные ссылки

- [NeetCode 150](https://neetcode.io/practice)
- Чеклист базы: см. Notion → CS Fundamentals Checklist