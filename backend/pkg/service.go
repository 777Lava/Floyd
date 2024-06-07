package pkg

import (
	"fmt"
	"math"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// FloydWarshall принимает матрицу смежности графа и возвращает матрицу кратчайших путей
func (s Service) floyd(graph [][]int) [][]int {
	// Количество вершин в графе
	n := len(graph)

	// Создаем матрицу для хранения кратчайших путей, копируем входной граф
	dist := make([][]int, n)
	for i := range graph {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}
	path := ""
	// Применяем алгоритм Флойда-Уоршелла
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] < math.MaxInt32 && dist[k][j] < math.MaxInt32 && dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	fmt.Println(path)
	return dist
}

// 		1	2	3
// 1	0	1	5
// 2	0	0	1
// 3	0	0	0
