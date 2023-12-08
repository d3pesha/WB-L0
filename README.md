### Тестовое задание: Демонстрационный сервис управления данными заказов

#### Задачи:

1. **Подготовка окружения:**
   - Развернуть локально PostgreSQL.
   - Создать собственную базу данных.
   - Настроить пользователя для взаимодействия с базой данных.

2. **Модель данных:**
   - Создать таблицы в PostgreSQL для хранения данных о заказах, опираясь на модель в формате JSON, предоставленную в задании.

3. **Разработка сервиса:**
   - Реализовать сервис для обработки данных о заказах.
   - Настроить подключение и подписку на канал в nats-streaming.
   - Записывать полученные данные в базу данных PostgreSQL.

4. **Кэширование данных:**
   - Разработать механизм кэширования в памяти для эффективного доступа к данным.
   - Сохранять полученные данные в кэше in-memory.

5. **Восстановление после сбоя:**
   - Обеспечить восстановление кэша из базы данных в случае падения сервиса.

6. **HTTP-сервер и интерфейс:**
   - Запустить HTTP-сервер для предоставления данных о заказах по их идентификатору из кэша.
   - Разработать простой веб-интерфейс для отображения данных по идентификатору заказа.

#### Рекомендации и дополнительные указания:

- Учесть статичность данных при выборе модели хранения в кэше и PostgreSQL.
- Обеспечить безопасную обработку данных из nats-streaming с использованием фильтрации и проверок.
- Разработать скрипт для тестирования работы подписки в онлайн-режиме.
- Рассмотреть стратегии обработки ошибок и восстановления для минимизации потерь данных в случае проблем с сервисом.
