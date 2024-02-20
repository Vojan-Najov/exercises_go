/*
Конфигурация

Обычно, реальные проекты используют для конфигурации отдельный пакет.
Напишите пакет config, который будет содержать в себе структуру go

type Config struct{
Width int
Height int
}

и функцию New() *Config, для установки значений по умолчанию
(Width = 100 и Height = 50).

Примечания
Функцию main создавать не надо.

*/

package config

type Config struct {
  Width  int
  Height int
}

func New() *Config {
  return &Config{Width: 100, Height: 50}
}
