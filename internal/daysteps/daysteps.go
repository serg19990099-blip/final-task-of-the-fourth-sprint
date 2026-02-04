package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	// Делим строку на две части по запятой
	txt := strings.Split(data, ",")

	if len(txt) != 2 {
		log.Println("")
		return 0, 0, errors.New("")
	}

	//Преобразовываем количество шагов в int
	stps, err := strconv.Atoi(txt[0])
	if err != nil {
		log.Println(err)
		return 0, 0, errors.New("")
	}
	if stps <= 0 {
		log.Println(err)
		return 0, 0, errors.New("")
	}
	// Парсинг времени

	tim, err := time.ParseDuration(txt[1])

	if err != nil {
		log.Println(err)
		return 0, 0, errors.New("")
	}

	if tim <= 0*time.Second {
		log.Println(err)
		return 0, 0, errors.New("")
	}
	if tim >= 10*time.Hour {
		log.Println(err)
		return 0, 0, errors.New("")
	}

	return stps, tim, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	stp, tim, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if stp <= 0 || tim <= 0 {
		log.Println(err)
		return ""
	}

	// Вычисляем дистанцию в км

	dist := (float64(stp) * stepLength) / mInKm

	// Считаем калории при ходьбе

	spentCalories, errs := spentcalories.WalkingSpentCalories(stp, weight, height, tim)
	if errs != nil {
		log.Println(errs)
		return fmt.Sprintf("%s", errs)
	}

	if tim <= 0 {
		log.Println(errs)
		return fmt.Sprintf("%s", errs)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stp, dist, spentCalories)
}
