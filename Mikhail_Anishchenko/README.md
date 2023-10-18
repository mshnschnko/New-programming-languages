## Лабораторные по дисциплине Новые языки программирования

### Описание заданий
1. [Рекурсия. Erlang](Erlang_Recursion)
    - Постановка задачи: обход двоичного несбалансированного дерева в ширину.
    - Формат входных данных: в файле `tree.txt` задается дерево рядом натуральных чисел по принципу: $i$ - узел, $2i+1$ - левый сын, $2i+2$ - правый сын, где $i$ - индекс элемента ($i\geq0$). Отсутствие сына обозначается числом $-1$ в соответствующей ячейке.
    - Формат выходных данных: вывод содержимого дерева в порядке обхода в консоль.
2. [Многопоточность. Go](Go_multithreading)
    - Постановка задачи: Параллельное чтение нескольких файлов и вычисление определителей матриц. Вычисление определителя производится с помощью разложения матрицы по первой строке и доволнительных миноров матрицы.
    - Формат входных данных: путь до директории с текстовыми файлами, содержащими квадратные матрицы с вещественными числами. По умолчанию задана директория matrices. Если матрица не является квадратной, выводится соответствующее сообщение об ошибке и программа завершается.
    - Формат выходных данных: вывод в консоль строки в формате `Determinant of matrix in  <filepath>  =  <det>`.
3. [Алгоритм. Swift](Swift_Algorithm)
    - Постановка задачи: поиск образца в тексте по алгоритму Рабина-Карпа с использованием полиномиальной хэш-функции.
    - Формат входных данных: в консоль с клавиатуры вводятся образец для поиска (в первой строке), путь до файла с текстом (во второй строке).
    - Формат выходных данных: в консоль выводится массив индексов вхождений образца в исходный текст. Индекс соответствует началу подстроки со вхождением.
4. [Асинхронность. Rust](Rust_Async)
5. [Ввод/вывод. Lua](Lua_IOStream)
    - Постановка задачи: введенную в консоль с клавиатуры строку разделить по пробелам, создать текстовый файл с названием, равным первому слову в введенной строке (если такой файл существует, то перезаписать его), вносить в созданный файл только те слова из оставшихся, длина которых не меньше длины первого слова.
    - Формат входных данных: в консоль с клавиатуры вводится строка со словами, разделенными пробелами.
    - Формат выходных данных: текстовый файл со словами введенной строки, длина которых не меньше длины первого слова.