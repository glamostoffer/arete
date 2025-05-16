create table public.quizz(
    id serial primary key,
    course_id bigint not null,
    title text not null,
    "description" text not null,
    difficulty text not null,
    passing_score bigint not null, -- кол-во необходимых правильных ответов для прохождения
    sequence_number bigint not null -- порядковый номер квиза в рамках курса
);

create table public.question(
    id serial primary key,
    quizz_id bigint not null references public.quizz (id),
    question text not null,
    explanation text not null
);

create table public.question_option(
    id serial primary key,
    question_id bigint not null references public.question (id),
    "option" text not null,
    is_correct boolean not null
);

create table public.user_completed_quizz(
    user_id bigint not null,
    quizz_id bigint not null references public.quizz (id),
    primary key (user_id, quizz_id)
);

create public.task(
    id serial primary key,
    course_id bigint not null,
    title text not null,
    "description" text not null,
    difficulty text not null,
    explanation text not null,
    sequence_number bigint not null -- порядковый номер задачи в рамках курса
);

create table public.prog_language(
    id serial primary key,
    "name" text not null
);

create table public.task_template(
    task_id bigint primary key references public.task (id),
    language_id int references public.prog_language (id),
    function_template text not null 
);

create public.test_case(
    id serial primary key,
    task_id bigint not null references public.task (id),
    input text not null,
    "output" text not null,
);

create table public.user_completed_task(
    user_id bigint not null,
    task_id bigint not null references public.task (id),
    time_rate numeric not null,
    memory_rate numeric not null,
    primary key (user_id, task_id)
);

create schema if not exists outbox;

create table outbox."event"(
    id serial primary key,
    "key" text not null,
    topic text not null,
    payload jsonb not null,
    idempotency_key text not null,
    created_at timestamptz not null default now(),
    locked_until timestamptz,
    attempts int not null default 0,
    error text,
    processed_at timestamptz
);


-- =========================================================================== --

-- Тесты для курса "Программирование на Go (Golang)" (course_id = 1)
INSERT INTO public.quizz (course_id, title, description, passing_score, sequence_number) VALUES
(1, 'Основы Go', 'Проверка базовых знаний синтаксиса Go', 3, 1),
(1, 'Многопоточность в Go', 'Тест по горутинам и каналам', 4, 2);

-- Вопрос 1
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(1, 'Какой синтаксис используется для объявления переменной в Go?', 
'В Go есть два основных способа объявления переменных: с указанием типа "var name type = value" и с краткой записью "name := value". Краткая запись возможна только внутри функций.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(1, 'var x int = 5', true),
(1, 'let x = 5', false),
(1, 'x = 5', false),
(1, 'const x = 5', false);

-- Вопрос 2
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(1, 'Что выведет программа: fmt.Println(len("Привет"))?', 
'Функция len() для строк возвращает количество байт, а не символов. Кириллические символы в UTF-8 занимают 2 байта, поэтому "Привет" (6 букв) займет 12 байт.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(2, '12', true),
(2, '6', false),
(2, '5', false),
(2, 'Ошибку компиляции', false);

-- Вопрос 3
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(1, 'Какой тип данных используется для хранения текста в Go?', 
'Основной тип для работы с текстом в Go - string. Тип rune представляет отдельный Unicode символ, но для строк используется string.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(3, 'string', true),
(3, 'text', false),
(3, 'rune', false),
(3, 'char', false);

-- Вопрос 4
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(1, 'Как правильно объявить функцию в Go?', 
'Функции в Go объявляются с помощью ключевого слова func, за которым следует имя функции, параметры в круглых скобках и возвращаемые типы.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(4, 'func sum(a int, b int) int {}', true),
(4, 'function sum(a, b) {}', false),
(4, 'def sum(a int, b int): int {}', false),
(4, 'sum := func(a, b) {}', false);

-- Вопрос 5
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(1, 'Что делает оператор := в Go?', 
'Оператор := используется для краткого объявления переменной с одновременной инициализацией. Тип переменной выводится автоматически из значения.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(5, 'Объявляет и инициализирует переменную с выводом типа', true),
(5, 'Сравнивает значения', false),
(5, 'Присваивает значение существующей переменной', false),
(5, 'Объявляет указатель', false);

-- Тест "Многопоточность в Go" (quizz_id = 2)
-- Вопрос 6
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(2, 'Что такое горутина в Go?', 
'Горутина - это легковесный поток выполнения, управляемый runtime Go. В отличие от потоков ОС, горутины потребляют мало памяти и переключаются быстрее.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(6, 'Легковесный поток выполнения, управляемый runtime Go', true),
(6, 'Поток операционной системы', false),
(6, 'Асинхронная функция', false),
(6, 'Метод структуры', false);

-- Вопрос 7
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(2, 'Как создать горутину в Go?', 
'Для запуска функции в отдельной горутине используется ключевое слово go перед вызовом функции. Например: go myFunction()');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(7, 'go myFunction()', true),
(7, 'goroutine.create(myFunction)', false),
(7, 'async myFunction()', false),
(7, 'thread.start(myFunction)', false);

-- Вопрос 8
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(2, 'Для чего используются каналы в Go?', 
'Каналы - это типизированные conduit, через которые горутины могут обмениваться данными. Они обеспечивают синхронизацию и безопасность при передаче данных.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(8, 'Для обмена данными между горутинами', true),
(8, 'Для хранения коллекций данных', false),
(8, 'Для работы с файлами', false),
(8, 'Для создания буфера ввода-вывода', false);

-- Вопрос 9
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(2, 'Как создать буферизированный канал?', 
'Буферизированный канал создается с указанием емкости буфера вторым аргументом в make: make(chan int, 100). Небуферизированный канал имеет нулевую емкость.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(9, 'make(chan int, 10)', true),
(9, 'make(chan int)', false),
(9, 'new(chan int, 10)', false),
(9, 'chan int{10}', false);

-- Вопрос 10
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(2, 'Что делает конструкция select с каналами?', 
'Select позволяет горутине ждать операций на нескольких каналах. Выполняется первый готовый case. Если готовы несколько, выбирается случайный.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(10, 'Ожидает операций на нескольких каналах', true),
(10, 'Выбирает случайное число', false),
(10, 'Закрывает канал', false),
(10, 'Создает новый канал', false);

-- Тесты для курса "Разработка на C#" (course_id = 2)
INSERT INTO public.quizz (course_id, title, description, passing_score, sequence_number) VALUES
(2, 'Основы C#', 'Тест по базовому синтаксису C#', 3, 1),
(2, 'ООП в C#', 'Тест по объектно-ориентированному программированию', 4, 2);

-- Вопрос 11
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(3, 'Какой символ завершает инструкции в C#?', 
'Как и в других C-подобных языках, в C# точка с запятой (;) обязательна для завершения инструкций. Это помогает компилятору определить конец выражения.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(11, '; (точка с запятой)', true),
(11, ': (двоеточие)', false),
(11, '. (точка)', false),
(11, '} (закрывающая фигурная скобка)', false);

-- Вопрос 12
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(3, 'Как объявить целочисленную переменную в C#?', 
'В C# используется стиль объявления "тип имя = значение". Для целых чисел чаще всего используют int (Int32) или long (Int64).');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(12, 'int x = 5;', true),
(12, 'var x = 5;', false),
(12, 'x := 5;', false),
(12, 'integer x = 5;', false);

-- Вопрос 13
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(3, 'Какой метод является точкой входа в программу на C#?', 
'В консольных приложениях C# точка входа - это статический метод Main в классе Program. Он может принимать аргументы командной строки или быть без параметров.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(13, 'static void Main(string[] args)', true),
(13, 'void Start()', false),
(13, 'public static int EntryPoint()', false),
(13, 'static async Task Main()', true); -- Допустимый вариант в новых версиях

-- Вопрос 14
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(3, 'Как объявить константу в C#?', 
'Константы в C# объявляются с ключевым словом const и должны быть инициализированы при объявлении. Их значение нельзя изменить после компиляции.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(14, 'const int MaxValue = 100;', true),
(14, 'readonly int MaxValue = 100;', false),
(14, 'final int MaxValue = 100;', false),
(14, 'constant int MaxValue = 100;', false);

-- Вопрос 15
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(3, 'Что выведет код: Console.WriteLine(5 / 2);?', 
'При делении целых чисел в C# результат тоже будет целым числом. Для получения дробного результата нужно использовать числа с плавающей точкой.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(15, '2', true),
(15, '2.5', false),
(15, '2.0', false),
(15, '2.5f', false);

-- Тест "ООП в C#" (quizz_id = 4)
-- Вопрос 16
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(4, 'Как объявить класс в C#?', 
'Классы в C# объявляются с ключевым словом class, за которым следует имя класса. Тело класса помещается в фигурные скобки.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(16, 'class MyClass {}', true),
(16, 'type MyClass {}', false),
(16, 'object MyClass {}', false),
(16, 'struct MyClass {}', false);

-- Вопрос 17
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(4, 'Какое ключевое слово используется для наследования класса?', 
'В C# для наследования используется символ : после имени производного класса, за которым следует имя базового класса.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(17, ': (двоеточие)', true),
(17, 'extends', false),
(17, 'inherits', false),
(17, 'super', false);

-- Вопрос 18
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(4, 'Какой модификатор доступа делает член класса доступным только внутри этого класса?', 
'private - самый строгий модификатор доступа в C#, который ограничивает видимость только содержащим классом.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(18, 'private', true),
(18, 'protected', false),
(18, 'internal', false),
(18, 'public', false);

-- Вопрос 19
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(4, 'Что такое виртуальный метод в C#?', 
'Виртуальные методы (с модификатором virtual) могут быть переопределены в производных классах с помощью override. Это основа полиморфизма в C#.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(19, 'Метод, который можно переопределить в производном классе', true),
(19, 'Метод без реализации', false),
(19, 'Статический метод', false),
(19, 'Абстрактный метод', false);

-- Вопрос 20
INSERT INTO public.question (quizz_id, question, explanation) VALUES
(4, 'Как создать экземпляр класса в C#?', 
'Для создания экземпляра класса используется оператор new, который выделяет память и вызывает конструктор класса.');

INSERT INTO public.question_option (question_id, option, is_correct) VALUES
(20, 'new MyClass()', true),
(20, 'MyClass.create()', false),
(20, 'MyClass.new()', false),
(20, 'instantiate MyClass()', false);