// Package app.
// Здесь происходит:
// Получение и передача конфига
// Менеджмент зависимостями
// Инициализация реализаций
// Управление запуском и остановкой запущенных компонентов
//
// Поток выполнения входит из main в Run, а при необходимости
// остановить приложение - выходит отсюда же.
package app
