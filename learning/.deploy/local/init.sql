create table public.course_category (
    id serial primary key,
    name text not null
);

create table public.course (
    id serial primary key,
    title text not null,
    "description" text not null,
    duration text not null,
    difficulty text not null,
    category_id bigint not null references public.course_category (id),
    image_url text not null
);

create table public.lesson (
    id serial primary key,
    course_id bigint not null references public.course (id),
    title text not null,
    "description" text not null,
    duration text not null
);

create table public.lesson_content (
    lesson_id bigint primary key not null references public.lesson (id),
    content text not null
);

create table public.user_course (
    user_id bigint not null,
    course_id bigint not null references public.course (id),
    primary key (user_id, course_id)
);

-- =========================================================================== --

insert into public.course_category (
    name
) values 
(
    'programming'
),
(
    'web'
),
(
    'gamedev'   
),
(
    'algorithms'
),
(
    'databases'
);


insert into public.course (
    title,
    "description",
    duration,
    difficulty,
    category_id,
    image_url
) values 
(
    'Программирование на Go (Golang)',
    'Изучите язык программирования Go: многопоточность, эффективность и простота синтаксиса.',
    '6 недель',
    'Средний',
    1,
    'https://images.unsplash.com/photo-1617791160536-598cf32026fb?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
),
(
    'Разработка на C#',
    'Полный курс по C#: от основ до продвинутых концепций, включая .NET и разработку приложений.',
    '7 недель',
    'Средний',
    1,
    'https://images.unsplash.com/photo-1573164713988-8665fc963095?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
),
(
    'Go для веб-разработки',
    'Создание высоконагруженных веб-приложений с использованием Golang и современных фреймворков.',
    '5 недель',
    'Продвинутый',
    2,
    'https://images.unsplash.com/photo-1499951360447-b19be8fe80f5?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
),
(
    'C# и Unity для разработки игр',
    'Основы разработки игр на Unity с использованием языка программирования C#.',
    '8 недель',
    'Средний',
    3,
    'https://images.unsplash.com/photo-1551103782-8ab07afd45c1?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
),
(
    'Алгоритмы и структуры данных',
    'Углубленное изучение алгоритмов и структур данных для подготовки к техническим собеседованиям.',
    '6 недель',
    'Продвинутый',
    4,
    'https://images.unsplash.com/photo-1558494949-ef010cbdcc31?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
),
(
    'Основы баз данных',
    'Изучение SQL и NoSQL баз данных, проектирование и оптимизация запросов.',
    '5 недель',
    'Средний',
    5,
    'https://images.unsplash.com/photo-1460925895917-afdab827c52f?ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80'
);

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(1, 'Введение в Go', 'Основные концепции и синтаксис языка Go', '20 мин'),
(1, 'Многопоточность', 'Горутины и каналы для параллельного программирования', '30 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(1, '<h2>Введение в Go</h2>
<p>Go (Golang) — современный язык программирования от Google.</p>
<h3>Особенности:</h3>
<ul>
<li>Простой синтаксис</li>
<li>Быстрая компиляция</li>
<li>Встроенная поддержка многопоточности</li>
</ul>
<pre><code>package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}</code></pre>'),
(2, '<h2>Многопоточность в Go</h2>
<p>Горутины — легковесные потоки в Go.</p>
<pre><code>package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}</code></pre>');

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(2, 'Основы C#', 'Синтаксис и базовые конструкции', '25 мин'),
(2, 'ООП в C#', 'Классы, объекты и наследование', '30 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(3, '<h2>Основы C#</h2>
<p>C# — объектно-ориентированный язык от Microsoft.</p>
<pre><code>using System;

class Program {
    static void Main() {
        Console.WriteLine("Привет, мир!");
        int x = 5;
        double y = 3.14;
        Console.WriteLine($"x = {x}, y = {y}");
    }
}</code></pre>'),
(4, '<h2>ООП в C#</h2>
<p>Пример класса и наследования:</p>
<pre><code>public class Animal {
    public string Name { get; set; }
    
    public virtual void MakeSound() {
        Console.WriteLine("Some sound");
    }
}

public class Dog : Animal {
    public override void MakeSound() {
        Console.WriteLine("Гав!");
    }
}</code></pre>');

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(3, 'Веб-сервер на Go', 'Создание сервера с net/http', '30 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(5, '<h2>Веб-сервер на Go</h2>
<p>Простейший HTTP-сервер:</p>
<pre><code>package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет, мир!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}</code></pre>');

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(4, 'Скриптинг в Unity', 'Основы написания скриптов на C#', '35 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(6, '<h2>Скриптинг в Unity</h2>
<p>Пример простого скрипта для движения объекта:</p>
<pre><code>using UnityEngine;

public class PlayerController : MonoBehaviour {
    public float speed = 5.0f;
    
    void Update() {
        float moveX = Input.GetAxis("Horizontal");
        float moveZ = Input.GetAxis("Vertical");
        Vector3 movement = new Vector3(moveX, 0.0f, moveZ);
        transform.Translate(movement * speed * Time.deltaTime);
    }
}</code></pre>');

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(5, 'Быстрая сортировка', 'Реализация алгоритма быстрой сортировки', '40 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(7, '<h2>Быстрая сортировка</h2>
<p>Реализация на C#:</p>
<pre><code>public static void QuickSort(int[] array, int left, int right) {
    if (left < right) {
        int pivot = Partition(array, left, right);
        QuickSort(array, left, pivot - 1);
        QuickSort(array, pivot + 1, right);
    }
}

private static int Partition(int[] array, int left, int right) {
    int pivot = array[right];
    int i = left - 1;
    
    for (int j = left; j < right; j++) {
        if (array[j] <= pivot) {
            i++;
            Swap(ref array[i], ref array[j]);
        }
    }
    Swap(ref array[i + 1], ref array[right]);
    return i + 1;
}</code></pre>');

INSERT INTO public.lesson (course_id, title, description, duration) VALUES
(6, 'SQL запросы', 'Основные команды SELECT, INSERT, UPDATE', '30 мин');

INSERT INTO public.lesson_content (lesson_id, content) VALUES
(8, '<h2>Основы SQL</h2>
<p>Примеры базовых SQL-запросов:</p>
<pre><code>-- Создание таблицы
CREATE TABLE Users (
    id INT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100)
);');