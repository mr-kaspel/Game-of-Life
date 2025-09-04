package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	symbolCell := "■ "
	dieCell := "- "
	x, y := 18, 18

	// Текущее состояние
	stateCells := make([][]int, y)
	for i := range stateCells {
		stateCells[i] = make([]int, x)
	}

	// Множество живых клеток (глайдер)
	m := map[int]map[int]int{}
	m[7] = map[int]int{9: 1}
	m[8] = map[int]int{10: 1}
	m[9] = map[int]int{8: 1, 9: 1, 10: 1}

	// Установка начального состояния
	for row, cols := range m {
		for col := range cols {
			stateCells[row][col] = 1
		}
	}

	// Направления для соседей (8 шт.)
	directions := [8][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	maxEras := 100

	for era := 1; era <= maxEras; era++ {
		fmt.Print("\033[H\033[2J") // Очистка экрана
		fmt.Printf("Эпоха: %d\n", era)

		newState := make([][]int, y)
		for i := range newState {
			newState[i] = make([]int, x)
		}
		newM := map[int]map[int]int{}

		// Собираем клетки, которые нужно проверить (живые и их соседи)
		checkSet := map[[2]int]struct{}{}
		for row, cols := range m {
			for col := range cols {
				for _, d := range directions {
					nx := col + d[0]
					ny := row + d[1]
					if nx >= 0 && nx < x && ny >= 0 && ny < y {
						checkSet[[2]int{ny, nx}] = struct{}{}
					}
				}
				// сама клетка тоже должна быть в checkSet
				checkSet[[2]int{row, col}] = struct{}{}
			}
		}

		// Применяем правила игры жизни
		for coord := range checkSet {
			row := coord[0]
			col := coord[1]

			liveNeighbors := 0
			for _, d := range directions {
				nx := col + d[0]
				ny := row + d[1]
				if nx >= 0 && nx < x && ny >= 0 && ny < y && stateCells[ny][nx] == 1 {
					liveNeighbors++
				}
			}

			if stateCells[row][col] == 1 {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					newState[row][col] = 1
					if newM[row] == nil {
						newM[row] = map[int]int{}
					}
					newM[row][col] = 1
				}
			} else if liveNeighbors == 3 {
				newState[row][col] = 1
				if newM[row] == nil {
					newM[row] = map[int]int{}
				}
				newM[row][col] = 1
			}
		}

		// Быстрое формирование вывода
		var sb strings.Builder
		for i := 0; i < y; i++ {
			row := newState[i]
			for j := 0; j < x; j++ {
				if row[j] == 1 {
					sb.WriteString(symbolCell)
				} else {
					sb.WriteString(dieCell)
				}
			}
			sb.WriteByte('\n')
		}
		fmt.Print(sb.String())

		// Переход к следующему поколению
		stateCells = newState
		m = newM

		time.Sleep(200 * time.Millisecond) // задержка для анимации
	}
}
