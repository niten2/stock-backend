# Тестовое задание

Разработать прототип системы складского учета,
в которой фиксируется
- учет,
- прибытие,
- убытие товаров на складе.

Характеристики груза

## Каждый тип груза хранимый на складе имеет следующие характеристики:

- Название
- Вес одного ящика (в килограммах)
- № Ангара где будет производится хранение

## Роли пользователей и функционал

В системе должно быть представлено 3 роли пользователей:

### Поставщик товара.
- Вносит информацию о новых поставках грузов (без возможности дальнейшего редактирования\удаления),
- Обязательно указывает типА поставляемого груза и количества ящиков (в штуках).

### Оптовик.
- Видит информацию о текущем остатке (кол-во ящиков) по каждому типу груза.
- Производить забор необходимого количества ящиков груза любого типа со склада, уменьшая тем самым его доступный остаток.

### Администратор системы. Имеет возможность использовать следующий функционал:
- Видеть полную историю всех операций поставки и забора грузов с указанием точной даты и времени.
- Видеть информацию о текущих остатках.
- Создавать новые типы грузов, которые затем могут использоваться остальными пользователями системы.
- Видеть общую массу всех грузов, хранимых в каждом ангаре

## Пример

Имеется категория грузов:

Телефон Vertu
- Масса 10 кг
- Ангар №5.

Поставщик «Скоровоз» вносит на склад следующий товар:

- Категория груза «Телефон Vertu»
- Количество: 17 ящиков

Оптовик «Селлер» берет со склада:
- Категория «Телефон Vertu»
- Количество: 10 ящиков

После этого Администратор авторизуется в системе и видит следующую информацию:

![image](https://github.com/niten2/test_tasks/blob/master/stock/image.png)
