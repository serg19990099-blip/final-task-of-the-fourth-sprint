package spentcalories

import (
  "errors"
  "fmt"
  "log"
  "strconv"
  "strings"
  "time"
)

// Основные константы, необходимые для расчетов.
const (
  lenStep                    = 0.65 // средняя длина шага.
  mInKm                      = 1000 // количество метров в километре.
  minInH                     = 60   // количество минут в часе.
  stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
  walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
  // TODO: реализовать функцию

  // Делим строку на две части по запятой
  pTrain := strings.Split(data, ",")

  if len(pTrain) != 3 {
    log.Println("")
    return 0, "", 0, errors.New("")
  }

  //Преобразовываем количество шагов в int
  stps, err := strconv.Atoi(pTrain[0])
  if err != nil {
    log.Println("")
    return 0, "", 0, errors.New("")
  }
  if stps <= 0 {
    log.Println("")
    return 0, "", 0, errors.New("")
  }

  // Парсинг времени
  tim, err := time.ParseDuration(pTrain[2])

  if err != nil {
    log.Println(err)
    return 0, "", 0, errors.New("")
  }

  if tim <= 0*time.Second {
    log.Println(err)
    return 0, "", 0, errors.New("")
  }
  if tim >= 10*time.Hour {
    log.Println(err)
    return 0, "", 0, errors.New("")
  }

  str := fmt.Sprintf("%s", pTrain[1])

  return stps, str, tim, nil

}

func distance(steps int, height float64) float64 {
  // TODO: реализовать функцию
  if steps <= 0  height <= 0 {
    log.Println("")
    return 0
  }
  stepLengt := height * stepLengthCoefficient
  distKm := (float64(steps) * stepLengt) / mInKm
  return distKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
  // TODO: реализовать функцию
  if steps <= 0  height <= 0 {
    log.Println("")
    return 0
  }
  if duration <= 0 {
    log.Println("")
    return 0
  }
  dist := distance(steps, height)
  meanSp := dist / duration.Hours()
  return meanSp
}

func TrainingInfo(data string, weight, height float64) (string, error) {
  // TODO: реализовать функцию
  stps, typ, tim, err := parseTraining(data)

  ErrUnknownType := errors.New("неизвестный тип тренировки")

  if err != nil {
    log.Println(err)
    return "", errors.New("")
  }
  if stps <= 0  tim <= 0 {
    log.Println(err)
    return "", errors.New("")
  }

  switch typ {
  case "Бег":
    runCalories, err := RunningSpentCalories(stps, weight, height, tim)
    if err != nil {
      log.Println(err)
      fmt.Println("")
    }
    dist := distance(stps, height)
    speed := meanSpeed(stps, height, tim)
    return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typ, tim.Hours(), dist, speed, runCalories), nil

  case "Ходьба":
    walkCalories, err := WalkingSpentCalories(stps, weight, height, tim)
    if err != nil {
      log.Println(err)
      fmt.Println("")
    }
    dist := distance(stps, height)
    speed := meanSpeed(stps, height, tim)
    return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typ, tim.Hours(), dist, speed, walkCalories), nil

  default:
    log.Println(ErrUnknownType)
    return "", ErrUnknownType

  }
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
  // TODO: реализовать функцию
  if steps <= 0  weight <= 0  height <= 0  duration <= 0 {
    log.Println("")
    return 0, errors.New("")
  }
  meanSp := meanSpeed(steps, height, duration)

  runSpentCalor := (weight * meanSp * duration.Minutes()) / minInH
  if runSpentCalor <= 0 {
    log.Println("")
    return 0, errors.New("")
  } else {
    log.Println("")
    return runSpentCalor, nil
  }

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
  // TODO: реализовать функцию
  if steps <= 0  weight <= 0  height <= 0 || duration <= 0 {
    log.Println("")
    return 0, errors.New("")
  }
  meanSp := meanSpeed(steps, height, duration)

  walkSpentCalor := ((weight * meanSp * duration.Minutes()) / minInH) * walkingCaloriesCoefficient
  if walkSpentCalor <= 0 {
    log.Println("")
    return 0, errors.New("")
  } else {
    return walkSpentCalor, nil
  }
}